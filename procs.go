package common

import (
	"fmt"
	"runtime"
)

// PrintCPU :
func PrintCPU() {
	fmt.Println("\nCPU INFO")
	fmt.Println("CPUs: ", runtime.NumCPU())
	fmt.Println("MAXPROCS: ", runtime.GOMAXPROCS(0))
	fmt.Println("GoRoot: ", runtime.GOROOT())
	fmt.Println("Archive: ", runtime.GOOS)
	fmt.Println("---------------------")

	// runtime.GOMAXPROCS(4)
	// runtime.GOMAXPROCS(runtime.NumCPU())
}
