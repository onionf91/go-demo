package pkg

import (
	"log"
	"runtime"
	"syscall"
	"testing"
)

func TestGettid(t *testing.T) {
	runtime.GOMAXPROCS(runtime.NumCPU())

	num := 20
	sem := make(chan int, num)

	for i := 0; i < num; i++ {
		go func() {
			log.Println(syscall.Gettid())

			sem <- 0
		}()
	}

	for i := 0; i < num; i++ {
		<-sem
	}
}