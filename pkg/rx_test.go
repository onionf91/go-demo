package pkg

import (
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"log"
	"syscall"
	"testing"
)

func logMsg(format string, v ...interface{}) {
	msg := fmt.Sprintf(format, v...)
	log.Printf("[%v] %q", syscall.Gettid(), msg)
}

func TestJust(t *testing.T) {
	msg := "Hello, World!"
	observable := rxgo.Just(msg)()
	ch := observable.Observe()
	item := <-ch

	logMsg("value %v", item.V)

	if item.V != msg {
		t.Errorf("got %v, wanted %v", item.V, msg)
	}
}
