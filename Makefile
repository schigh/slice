.PHONY: base
base:
	@{ \
		cd internal/base; \
		go run base.go; \
	}

.PHONY: slicify
slicify:
	go build -o "${GOBIN}/slicify" ./cmd/slicify/slicify.go

.PHONY: test
test:
	go test -race -v -cover ./...

.PHONY: benchmarks
benchmarks:
	@{ \
  		touch ./_data/benchmark.txt; \
  		echo "-------------------------------------------------" >> ./_data/benchmark.txt; \
  		echo "$(shell date)" >> ./_data/benchmark.txt; \
  		echo "running int slice benchmark"; \
  		go test -bench BenchmarkInt_ -benchmem >> ./_data/benchmark.txt; \
  		echo "running int8 slice benchmark"; \
  		go test -bench BenchmarkInt8_ -benchmem >> ./_data/benchmark.txt; \
  		echo "running int16 slice benchmark"; \
  		go test -bench BenchmarkInt16_ -benchmem >> ./_data/benchmark.txt; \
  		echo "running int32 slice benchmark"; \
  		go test -bench BenchmarkInt32_ -benchmem >> ./_data/benchmark.txt; \
  		echo "running int64 slice benchmark"; \
  		go test -bench BenchmarkInt64_ -benchmem >> ./_data/benchmark.txt; \
  		echo "running uint slice benchmark"; \
  		go test -bench BenchmarkUInt_ -benchmem >> ./_data/benchmark.txt; \
  		echo "running uint8 slice benchmark"; \
  		go test -bench BenchmarkUInt8_ -benchmem >> ./_data/benchmark.txt; \
  		echo "running uint16 slice benchmark"; \
  		go test -bench BenchmarkUInt16_ -benchmem >> ./_data/benchmark.txt; \
  		echo "running uint32 slice benchmark"; \
  		go test -bench BenchmarkUInt32_ -benchmem >> ./_data/benchmark.txt; \
  		echo "running uint64 slice benchmark"; \
  		go test -bench BenchmarkUInt64_ -benchmem >> ./_data/benchmark.txt; \
  		echo "running float32 slice benchmark"; \
  		go test -bench BenchmarkFloat32_ -benchmem >> ./_data/benchmark.txt; \
  		echo "running uint float64 benchmark"; \
  		go test -bench BenchmarkFLoat64_ -benchmem >> ./_data/benchmark.txt; \
  		echo "running string slice benchmark"; \
  		go test -bench BenchmarkString_ -benchmem >> ./_data/benchmark.txt; \
	}
