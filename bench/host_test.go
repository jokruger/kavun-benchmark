// Host info prelude. Runs once before any benchmark and prints machine metadata in `key: value` form. Captured by `tee`
// into the raw results file and rendered by cmd/report at the top of the Markdown report.
//
// This package contains no benchmarks itself — each engine lives in its own sub-package under `bench/` so that
// engines with conflicting init-time global state (e.g. duplicate gob registrations between tengo and its kavun fork)
// can coexist as separate test binaries. `go test ./bench/...` runs this package first (alphabetically), so its
// host_* prelude is the one captured at the top of the raw results file.
//
// Standard `go test -bench` output already includes goos/goarch/pkg/cpu; here we add NumCPU, GOMAXPROCS, Go version,
// and (best-effort) total system memory.
package bench

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	printHostInfo(os.Stdout)
	os.Exit(m.Run())
}

func printHostInfo(w *os.File) {
	fmt.Fprintf(w, "host_numcpu: %d\n", runtime.NumCPU())
	fmt.Fprintf(w, "host_gomaxprocs: %d\n", runtime.GOMAXPROCS(0))
	fmt.Fprintf(w, "host_goversion: %s\n", runtime.Version())
	fmt.Fprintf(w, "host_totalmem: %s\n", totalMem())
	fmt.Fprintf(w, "host_timestamp: %s\n", time.Now().UTC().Format(time.RFC3339))
}

func strOr(s, fallback string) string {
	if s == "" {
		return fallback
	}
	return s
}

// totalMem returns total system RAM as a human-readable string. Best-effort across darwin/linux; returns "unknown"
// elsewhere.
func totalMem() string {
	switch runtime.GOOS {
	case "darwin":
		out, err := exec.Command("sysctl", "-n", "hw.memsize").Output()
		if err == nil {
			if n, err := strconv.ParseUint(strings.TrimSpace(string(out)), 10, 64); err == nil {
				return humanBytes(n)
			}
		}
	case "linux":
		data, err := os.ReadFile("/proc/meminfo")
		if err == nil {
			for line := range strings.SplitSeq(string(data), "\n") {
				if strings.HasPrefix(line, "MemTotal:") {
					fields := strings.Fields(line)
					if len(fields) >= 2 {
						if n, err := strconv.ParseUint(fields[1], 10, 64); err == nil {
							return humanBytes(n * 1024)
						}
					}
				}
			}
		}
	}
	return "unknown"
}

func humanBytes(n uint64) string {
	const (
		KB = 1 << 10
		MB = 1 << 20
		GB = 1 << 30
	)
	switch {
	case n >= GB:
		return fmt.Sprintf("%.2f GB", float64(n)/float64(GB))
	case n >= MB:
		return fmt.Sprintf("%.2f MB", float64(n)/float64(MB))
	case n >= KB:
		return fmt.Sprintf("%.2f KB", float64(n)/float64(KB))
	}
	return fmt.Sprintf("%d B", n)
}
