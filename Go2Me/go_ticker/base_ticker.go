package go_ticker

import (
	"fmt"
	"time"
)

func BaseTickerSecond(second int) (*time.Ticker, error) {
	if second < 1 {
		return nil, fmt.Errorf("second is less than 1")
	}
	ticker := time.NewTicker(time.Second * time.Duration(second))
	return ticker, nil
}
