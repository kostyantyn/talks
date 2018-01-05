package main

import (
	"runtime/debug"
	"runtime"
	"fmt"
)

func init() {
	debug.SetGCPercent(-1)
}

// START OMIT
type A struct { B int; _ int }

func main() {
	profile(func(){
		for i := 0; i < 10; i++ {
			list := make([]*A, 1000) // HL
			for i := 0; i < 1000; i++ {
				list[i] = &A{B: i} // HL
			}
		}
	})
}

func profile(fn func()){
	o := new(runtime.MemStats)
	runtime.ReadMemStats(o)
	fn()
	n := new(runtime.MemStats)
	runtime.ReadMemStats(n)
	fmt.Printf("objects %v alloc %v", n.HeapObjects-o.HeapObjects, n.HeapAlloc - o.HeapAlloc)
}
// END OMIT