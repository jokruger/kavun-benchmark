.PHONY: bench report all clean tidy

RESULTS := results/raw.txt
REPORT  := results/REPORT.md
COUNT   ?= 5
BENCHTIME ?= 1s
PENALTY ?= 2.0

all: bench report

bench:
	@mkdir -p results
	go test -bench=. -benchmem -run=^$$ -count=$(COUNT) -benchtime=$(BENCHTIME) ./bench/... | tee $(RESULTS)

report:
	go run ./cmd/report -in $(RESULTS) -out $(REPORT) -penalty $(PENALTY)

clean:
	rm -rf results

tidy:
	go mod tidy
