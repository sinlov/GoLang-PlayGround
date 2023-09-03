package go_timer

import (
	"context"
	"time"
)

type HandlerFunc func(ctx *context.Context)

// d more as time.Second * time.Duration(3000)
func SetInterval(d time.Duration, ctx *context.Context, function ...HandlerFunc) {
	ticker := time.NewTicker(d)
	go func() {
		for _ = range ticker.C {
			for _, eachFunction := range function {
				go eachFunction(ctx)
			}
		}
	}()
}
