package worker_pool

import (
	"context"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleWorkPool(t *testing.T) {
	// mock SimpleWorkPool
	t.Logf("~> mock SimpleWorkPool")
	tasks := make([]TaskFunc, 0, 100)
	// do SimpleWorkPool
	t.Logf("~> do SimpleWorkPool")
	for i := 0; i < 100; i++ {
		func(val int) {
			tasks = append(tasks, func(ctx context.Context) error {
				t.Logf("task id: %v", val)
				if val == 51 {
					return errors.New(fmt.Sprintf("task %v , err %v", val, "missing"))
				}
				return nil
			})
		}(i)
	}
	// verify SimpleWorkPool
	err := ExecuteAll(0, tasks...)
	if err == nil {
		//fmt.Println("missing error")
		t.Fatal("missing error")
	} else {
		t.Logf("some task err %v", err.Error())
		assert.Equal(t, "task 51 , err missing", err.Error())
	}
	//assert.Equal(t, "", "")
}
