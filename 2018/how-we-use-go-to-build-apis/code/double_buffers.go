package main

import (
	"runtime/debug"
	"fmt"
	"runtime"
)

func init() {
	debug.SetGCPercent(-1)
}

type A struct {
	B int
	_ int
}

// START OMIT
func main() {
	profile(func(){
		lists := make([][]*A, 2) // HL
		lists[0] = make([]*A, 1000) // HL
		lists[1] = make([]*A, 1000) // HL

		for x := 0; x < 5; x++ {
			for y := 0; y < 1000; y++ {
				a := lists[0][y] // HL
				if a == nil {
					a = &A{}
				}
				a.B = x*y
				lists[0][y] = a  // HL
			}

			lists[0], lists[1] = lists[1], lists[0] // HL
		}
	})
}
// END OMIT

func profile(fn func()){
	o := new(runtime.MemStats)
	runtime.ReadMemStats(o)
	fn()
	n := new(runtime.MemStats)
	runtime.ReadMemStats(n)
	fmt.Printf("objects %v alloc %v", n.HeapObjects-o.HeapObjects, n.HeapAlloc - o.HeapAlloc)
}