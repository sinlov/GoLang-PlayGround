package go_error

import (
	"fmt"
	"errors"
)

type equalError struct {
	Num int
}

//error Error()
func (e equalError) Error() string {
	return fmt.Sprintf("Now Number is %d , must less than 10", e.Num)
}

func Equal(n int) (int, error) {
	if n > 55 {
		return -1, errors.New("not gt 10") //new error struct
	}
	return n, nil
}
