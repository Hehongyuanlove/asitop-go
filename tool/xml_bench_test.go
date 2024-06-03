package tool

import (
	"asitop-go/model"
	"testing"
)

// 比较 map => struct 与 map => json => struct

// [0526] 可能是因为 结构体并不完整的原因  MapToPcInfo 要好很多
// go test -benchmem  asitop-go/shell -bench . -benchmem
// goos: darwin
// goarch: arm64
// pkg: asitop-go/shell
// Benchmark_MapToPcInfo-8           197736              6074 ns/op            5695 B/op         64 allocs/op
// Benchmark_MapJsonPcInfo-8            267           4457239 ns/op         1586896 B/op      26674 allocs/op

// ADD GPU DESK
// Benchmark_MapToPcInfo-8            26275             45612 ns/op           40231 B/op        664 allocs/op
// Benchmark_MapJsonPcInfo-8            266           4478226 ns/op         1590514 B/op      26721 allocs/op

// Add  GPU DESK Network
// Benchmark_MapToPcInfo-8            23109             50564 ns/op           43488 B/op        730 allocs/op
// Benchmark_MapJsonPcInfo-8            266           4483141 ns/op         1592601 B/op      26729 allocs/op

// Add  GPU DESK Network Processor
// Benchmark_MapToPcInfo-8            17898             67953 ns/op           57651 B/op        925 allocs/op
// Benchmark_MapJsonPcInfo-8            261           4490205 ns/op         1599679 B/op      26744 allocs/op

func Benchmark_MapToPcInfo(b *testing.B) {

	v := map[string]interface{}{}
	ElpanXml("../simple/all.xml", &v)

	// 重置计数器
	b.ResetTimer()

	// 执行阶段
	for i := 0; i < b.N; i++ {
		// 每次迭代前重置 info，以避免潜在的累积效应
		info := model.PcInfo{}
		err := MapToPcInfo(&v, &info)
		if err != nil {
			b.Fatal("MapToPcInfo failed:", err)
		}
	}
}

func Benchmark_MapJsonPcInfo(b *testing.B) {

	v := map[string]interface{}{}
	ElpanXml("../simple/all.xml", &v)

	// 重置计数器
	b.ResetTimer()

	// 执行阶段
	for i := 0; i < b.N; i++ {
		// 每次迭代前重置 info，以避免潜在的累积效应
		info := model.PcInfo{}
		err := MapJsonPcInfo(&v, &info)
		if err != nil {
			b.Fatal("MapJsonPcInfo failed:", err)
		}
	}
}
