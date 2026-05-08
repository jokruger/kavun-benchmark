// Command report parses raw `go test -bench` output and produces a Markdown report with per-task tables
// (sorted by ns/op) and a summary table.
//
// The canonical list of tasks comes from the task package, so the report can detect when an engine skipped a task and
// apply a penalty.
//
// Penalty model for missing tasks:
//
//   - ratio = max(ratio observed on this task) * -penalty   (default 2.0)
//   - rank  = (number of engines that ran this task) + 1
//
// Both values feed into the geomean / avg-rank summary as if they were real measurements, so "didn't implement it" is
// worse than "implemented it badly" without being catastrophic.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/jokruger/kavun-benchmark/task"
)

type sample struct {
	nsPerOp     float64
	bytesPerOp  float64
	allocsPerOp float64
}

type aggregate struct {
	engine  string
	task    string
	samples []sample
}

func (a *aggregate) mean() sample {
	if len(a.samples) == 0 {
		return sample{}
	}
	var s sample
	for _, x := range a.samples {
		s.nsPerOp += x.nsPerOp
		s.bytesPerOp += x.bytesPerOp
		s.allocsPerOp += x.allocsPerOp
	}
	n := float64(len(a.samples))
	s.nsPerOp /= n
	s.bytesPerOp /= n
	s.allocsPerOp /= n
	return s
}

var benchLine = regexp.MustCompile(`^Benchmark\S+\s+\d+\s+`)

// Captures: 1=engine (whatever follows "Benchmark"), 2=task path (everything between first '/' and the optional -CPUS
// suffix).
var nameRE = regexp.MustCompile(`^Benchmark([A-Za-z0-9_]+)/(.+?)(?:-\d+)?$`)

// preludeRE matches `key: value` lines that appear before any benchmark line — both Go's standard goos/goarch/pkg/cpu
// header and our own host_* additions.
var preludeRE = regexp.MustCompile(`^([a-z_][a-z0-9_]*):\s+(.+)$`)

// parseFile returns engine -> task -> aggregate, plus the order engines were first seen, plus the host/prelude
// key-value map.
func parseFile(path string) (map[string]map[string]*aggregate, []string, map[string]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, nil, nil, err
	}
	defer f.Close()

	out := map[string]map[string]*aggregate{}
	engineOrder := []string{}
	engineSeen := map[string]bool{}
	prelude := map[string]string{}
	sawBench := false

	sc := bufio.NewScanner(f)
	sc.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if !sawBench {
			if m := preludeRE.FindStringSubmatch(line); m != nil {
				// Don't clobber if the same key appears multiple times
				if _, exists := prelude[m[1]]; !exists {
					prelude[m[1]] = m[2]
				}
			}
		}
		if !benchLine.MatchString(line) {
			continue
		}
		sawBench = true
		fields := strings.Fields(line)
		if len(fields) < 3 {
			continue
		}
		m := nameRE.FindStringSubmatch(fields[0])
		if m == nil {
			continue
		}
		engineName := strings.ToLower(m[1])
		taskName := m[2]

		s := sample{}
		for i := 2; i+1 < len(fields); i += 2 {
			val, err := strconv.ParseFloat(fields[i], 64)
			if err != nil {
				continue
			}
			switch fields[i+1] {
			case "ns/op":
				s.nsPerOp = val
			case "B/op":
				s.bytesPerOp = val
			case "allocs/op":
				s.allocsPerOp = val
			}
		}
		if s.nsPerOp == 0 {
			continue
		}

		if _, ok := out[engineName]; !ok {
			out[engineName] = map[string]*aggregate{}
		}
		if _, ok := out[engineName][taskName]; !ok {
			out[engineName][taskName] = &aggregate{engine: engineName, task: taskName}
		}
		out[engineName][taskName].samples = append(out[engineName][taskName].samples, s)

		if !engineSeen[engineName] {
			engineSeen[engineName] = true
			engineOrder = append(engineOrder, engineName)
		}
	}
	return out, engineOrder, prelude, sc.Err()
}

// row is one line in a per-task table.
type row struct {
	engine   string
	mean     sample
	missing  bool    // true => engine did not run this task; mean/ratio are penalty values
	rank     int     // 1-based by ns/op asc; missing engines get N+1
	ratio    float64 // ns/op vs best on this task
	memRatio float64 // B/op vs best on this task (1.0 if both zero)
}

// rankTask builds the per-task table including penalty rows for any canonical engine that didn't measure this task.
func rankTask(taskName string, byEngine map[string]map[string]*aggregate, allEngines []string, penalty float64) []row {
	rows := make([]row, 0, len(allEngines))
	present := []row{}
	missingEngines := []string{}

	for _, eng := range allEngines {
		if a, ok := byEngine[eng][taskName]; ok {
			present = append(present, row{engine: eng, mean: a.mean()})
		} else {
			missingEngines = append(missingEngines, eng)
		}
	}
	sort.Slice(present, func(i, j int) bool { return present[i].mean.nsPerOp < present[j].mean.nsPerOp })

	if len(present) == 0 {
		// Nothing to compare; emit synthetic rows so the report shows them as missing without dividing by zero.
		for _, eng := range missingEngines {
			rows = append(rows, row{engine: eng, missing: true, ratio: penalty, memRatio: penalty, rank: 1})
		}
		return rows
	}

	bestNs := present[0].mean.nsPerOp
	worstNs := present[len(present)-1].mean.nsPerOp
	bestMem := math.Inf(1)
	worstMem := 0.0
	for _, r := range present {
		if r.mean.bytesPerOp > 0 && r.mean.bytesPerOp < bestMem {
			bestMem = r.mean.bytesPerOp
		}
		if r.mean.bytesPerOp > worstMem {
			worstMem = r.mean.bytesPerOp
		}
	}
	if math.IsInf(bestMem, 1) {
		bestMem = 1
	}

	for i := range present {
		present[i].rank = i + 1
		if bestNs > 0 {
			present[i].ratio = present[i].mean.nsPerOp / bestNs
		} else {
			present[i].ratio = 1
		}
		if present[i].mean.bytesPerOp == 0 || bestMem == 0 {
			present[i].memRatio = 1
		} else {
			present[i].memRatio = present[i].mean.bytesPerOp / bestMem
		}
	}
	rows = append(rows, present...)

	// Penalty rows: ratio = (worst observed ratio) * penalty, rank = N+1.
	worstRatio := worstNs / bestNs
	worstMemRatio := 1.0
	if bestMem > 0 && worstMem > 0 {
		worstMemRatio = worstMem / bestMem
	}
	for _, eng := range missingEngines {
		rows = append(rows, row{
			engine:   eng,
			missing:  true,
			rank:     len(present) + 1,
			ratio:    worstRatio * penalty,
			memRatio: worstMemRatio * penalty,
		})
	}
	return rows
}

type summary struct {
	engine     string
	cpuGeomean float64
	memGeomean float64
	avgRank    float64
	worst      float64
	wins       int
	tasksRun   int
	missing    int
}

func buildSummary(perTask map[string][]row, taskOrder, allEngines []string) []summary {
	type acc struct {
		cpuLogSum, memLogSum, rankSum, worst float64
		wins, count, missing                 int
	}
	accs := map[string]*acc{}
	for _, eng := range allEngines {
		accs[eng] = &acc{}
	}
	for _, t := range taskOrder {
		for _, r := range perTask[t] {
			a := accs[r.engine]
			if a == nil {
				continue
			}
			a.cpuLogSum += math.Log(math.Max(r.ratio, 1e-12))
			a.memLogSum += math.Log(math.Max(r.memRatio, 1e-12))
			a.rankSum += float64(r.rank)
			if r.ratio > a.worst {
				a.worst = r.ratio
			}
			if r.rank == 1 && !r.missing {
				a.wins++
			}
			a.count++
			if r.missing {
				a.missing++
			}
		}
	}
	out := make([]summary, 0, len(accs))
	for _, eng := range allEngines {
		a := accs[eng]
		if a.count == 0 {
			continue
		}
		n := float64(a.count)
		out = append(out, summary{
			engine:     eng,
			cpuGeomean: math.Exp(a.cpuLogSum / n),
			memGeomean: math.Exp(a.memLogSum / n),
			avgRank:    a.rankSum / n,
			worst:      a.worst,
			wins:       a.wins,
			tasksRun:   a.count - a.missing,
			missing:    a.missing,
		})
	}
	sort.Slice(out, func(i, j int) bool { return out[i].cpuGeomean < out[j].cpuGeomean })
	return out
}

// renderHost emits a Markdown "Host" section summarizing machine info. Both the standard Go bench header keys
// (goos, goarch, pkg, cpu) and our host_* prelude keys are recognized; unknown keys are ignored.
func renderHost(out *os.File, host map[string]string) {
	if len(host) == 0 {
		return
	}
	// (display label, key) — order is the report order.
	rows := [][2]string{
		{"OS", "goos"},
		{"Arch", "goarch"},
		{"CPU", "cpu"},
		{"Logical CPUs", "host_numcpu"},
		{"GOMAXPROCS", "host_gomaxprocs"},
		{"Total memory", "host_totalmem"},
		{"Go version", "host_goversion"},
		{"Package", "pkg"},
		{"Run at (UTC)", "host_timestamp"},
	}
	any := false
	for _, r := range rows {
		if _, ok := host[r[1]]; ok {
			any = true
			break
		}
	}
	if !any {
		return
	}
	fmt.Fprintln(out, "## Host")
	fmt.Fprintln(out)
	fmt.Fprintln(out, "| Field | Value |")
	fmt.Fprintln(out, "|-------|-------|")
	for _, r := range rows {
		if v, ok := host[r[1]]; ok {
			fmt.Fprintf(out, "| %s | %s |\n", r[0], v)
		}
	}
	fmt.Fprintln(out)
}

func formatNs(v float64) string {
	switch {
	case v >= 1e9:
		return fmt.Sprintf("%.2f s", v/1e9)
	case v >= 1e6:
		return fmt.Sprintf("%.2f ms", v/1e6)
	case v >= 1e3:
		return fmt.Sprintf("%.2f µs", v/1e3)
	default:
		return fmt.Sprintf("%.0f ns", v)
	}
}

func formatBytes(v float64) string {
	switch {
	case v >= 1<<30:
		return fmt.Sprintf("%.2f GB", v/(1<<30))
	case v >= 1<<20:
		return fmt.Sprintf("%.2f MB", v/(1<<20))
	case v >= 1<<10:
		return fmt.Sprintf("%.2f KB", v/(1<<10))
	default:
		return fmt.Sprintf("%.0f B", v)
	}
}

func render(out *os.File, perTask map[string][]row, taskOrder []string, sum []summary, penalty float64, host map[string]string) {
	fmt.Fprintln(out, "# Kavun Benchmark Report")
	fmt.Fprintln(out)
	fmt.Fprintln(out, "Comparison of embeddable scripting engines for Go, focused on repeated execution of pre-compiled scripts.")
	fmt.Fprintln(out, "Compilation cost is excluded from measurements.")
	fmt.Fprintln(out)

	renderHost(out, host)

	fmt.Fprintln(out, "## Summary")
	fmt.Fprintln(out)
	fmt.Fprintln(out, "Primary ranking metric: each engine's time on a task divided by the fastest engine's time on that task.")
	fmt.Fprintln(out, "Lower is better; 1.00× means \"fastest on every task\".")
	fmt.Fprintln(out)
	fmt.Fprintf(out, "Missing-task penalty: ratio = (worst observed ratio on the task) × %.2f, rank = last + 1.\n", penalty)
	fmt.Fprintln(out)
	fmt.Fprintln(out, "| Rank | Engine | CPU geomean | Avg rank | Worst ratio | Wins | Mem geomean | Tasks run | Missing |")
	fmt.Fprintln(out, "|------|--------|-------------|----------|-------------|------|-------------|-----------|---------|")
	for i, s := range sum {
		fmt.Fprintf(out, "| %d | %s | %.2f× | %.2f | %.2f× | %d | %.2f× | %d | %d |\n",
			i+1, s.engine, s.cpuGeomean, s.avgRank, s.worst, s.wins, s.memGeomean, s.tasksRun, s.missing)
	}
	fmt.Fprintln(out)

	fmt.Fprintln(out, "## Per-task results")
	fmt.Fprintln(out)
	for _, taskName := range taskOrder {
		rows := perTask[taskName]
		t, _ := task.Lookup(taskName)
		fmt.Fprintf(out, "### %s\n\n", taskName)
		if t.Description != "" {
			fmt.Fprintf(out, "_%s_\n\n", t.Description)
		}
		fmt.Fprintln(out, "| Rank | Engine | ns/op | B/op | allocs/op | vs best |")
		fmt.Fprintln(out, "|------|--------|-------|------|-----------|---------|")
		for _, r := range rows {
			if r.missing {
				fmt.Fprintf(out, "| %d | %s | _missing_ | _missing_ | _missing_ | %.2f× (penalty) |\n",
					r.rank, r.engine, r.ratio)
			} else {
				fmt.Fprintf(out, "| %d | %s | %s | %s | %.0f | %.2f× |\n",
					r.rank, r.engine, formatNs(r.mean.nsPerOp),
					formatBytes(r.mean.bytesPerOp), r.mean.allocsPerOp, r.ratio)
			}
		}
		fmt.Fprintln(out)
	}
}

func main() {
	in := flag.String("in", "results/raw.txt", "raw `go test -bench` output")
	outPath := flag.String("out", "results/REPORT.md", "Markdown report path")
	penalty := flag.Float64("penalty", 2.0, "missing-task ratio penalty multiplier (>=1.0)")
	flag.Parse()
	if *penalty < 1.0 {
		fmt.Fprintln(os.Stderr, "penalty must be >= 1.0")
		os.Exit(1)
	}

	byEngine, engineOrder, host, err := parseFile(*in)
	if err != nil {
		fmt.Fprintln(os.Stderr, "parse:", err)
		os.Exit(1)
	}
	if len(byEngine) == 0 {
		fmt.Fprintln(os.Stderr, "no benchmark lines found in", *in)
		os.Exit(1)
	}

	taskOrder := task.Names()

	perTask := map[string][]row{}
	for _, t := range taskOrder {
		perTask[t] = rankTask(t, byEngine, engineOrder, *penalty)
	}
	sum := buildSummary(perTask, taskOrder, engineOrder)

	if err := os.MkdirAll(dir(*outPath), 0o755); err != nil {
		fmt.Fprintln(os.Stderr, "mkdir:", err)
		os.Exit(1)
	}
	f, err := os.Create(*outPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "create:", err)
		os.Exit(1)
	}
	defer f.Close()
	render(f, perTask, taskOrder, sum, *penalty, host)
	fmt.Println("wrote", *outPath)
}

func dir(p string) string {
	for i := len(p) - 1; i >= 0; i-- {
		if p[i] == '/' {
			return p[:i]
		}
	}
	return "."
}
