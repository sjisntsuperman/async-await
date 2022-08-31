package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/sjisntsuperman/async/async"
)

func Sleep(delay int) int {
	fmt.Println("waiting")
	time.Sleep(time.Duration(3) * time.Second)
	fmt.Println("finished")
	return 2
}

func TestAsync(t *testing.T) {
	future := async.Executor(func() interface{} {
		return Sleep(3)
	})
	val := future.Await()
	if val == nil {
		t.Fatalf("error: %s", future)
	}
	fmt.Println(val)
}