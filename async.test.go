package async

import (
	"fmt"
	"testing"
	"time"

	"github.com/sjisntsuperman/async"
)

func DoneAsync() int {
	fmt.Println("Warming up ...")
	time.Sleep(3 * time.Second)
	fmt.Println("Done ...")
	return 1
}

func TestAsync(t *testing.T) {
	fmt.Println("Let's start ...")
	future := async.Executor(func() interface{} {
		return DoneAsync()
	})
	if future == nil {
		t.Fatalf("testing error: %s", future)
	}
	fmt.Println("Done is running ...")
	val := future.Await()
	fmt.Println(val)
}