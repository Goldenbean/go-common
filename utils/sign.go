package utils

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// SignalWatch :
func SignalWatch() {

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		sig := <-sigs

		fmt.Printf("caught sig: %+v \n", sig)
		fmt.Println("Wait for 1 second to finish processing")
		time.Sleep(1 * time.Second)
		os.Exit(0)
	}()
}
