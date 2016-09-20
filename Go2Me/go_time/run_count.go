package go_time

import (
	"time"
	"strings"
	"fmt"
)

func StringsJoinSliceStringTime(slice []string) interface{} {
	start_time := time.Now()
	str := strings.Join(slice, " ")
	end_time := time.Since(start_time).Nanoseconds()
	fmt.Println("FastTimeTest res: ", str)
	return end_time
}
