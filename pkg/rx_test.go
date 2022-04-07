package pkg

import (
	"context"
	"fmt"
	"github.com/reactivex/rxgo/v2"
	"log"
	"math/rand"
	"syscall"
	"testing"
	"time"
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

//func TestConcurrentProcess(t *testing.T) {
//	observable := rxgo.Range(1, 4).FlatMap(func(i rxgo.Item) rxgo.Observable {
//		logMsg("proc %v", i.V)
//		return rxgo.Range(1, 4).Map(func(_ context.Context, v interface{}) (interface{}, error) {
//			logMsg("task %v emit %v", i.V, v)
//			//time.Sleep(time.Millisecond * time.Duration(rand.Intn(200)))
//			time.Sleep(time.Second * time.Duration(rand.Intn(5)))
//			return fmt.Sprintf("task %v emit %v", i.V, v), nil
//		}, rxgo.WithPool(4))
//	}, rxgo.WithBufferedChannel(4)).DoOnNext(func(s interface{}) {
//		logMsg("~~~ %v", s)
//	})
//	<-observable
//}

func TestConcurrentProcess(t *testing.T) {
	observable := rxgo.Range(1, 16).Map(func(_ context.Context, v interface{}) (interface{}, error) {
		logMsg("task emit %v", v)
		//time.Sleep(time.Millisecond * time.Duration(rand.Intn(200)))
		time.Sleep(time.Second * time.Duration(rand.Intn(5)))
		return fmt.Sprintf("task emit %v", v), nil
	}, rxgo.WithPool(4)).DoOnNext(func(s interface{}) {
		logMsg("~~~ %v", s)
	})
	<-observable
}
