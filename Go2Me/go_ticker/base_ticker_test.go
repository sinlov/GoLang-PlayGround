package go_ticker

import (
	"sync"
	"testing"
	"time"
)

func Test_BaseTickerSecond(t *testing.T) {

	tickerSecond, err := BaseTickerSecond(3)
	if err != nil {
		t.Fatalf("BaseTickerSecond err %v", err)
	}
	var wg sync.WaitGroup
	wg.Add(1)

	cnt := 3

	go func(tic *time.Ticker) {
		defer wg.Done()
		for {
			<-tic.C
			cnt--
			if cnt < 0 {
				tic.Stop()
				t.Logf("tick stop [t:%v]", time.Now().Format("2006-01-02 15:04:05"))
				return
			}

			t.Logf("tick in [t:%v]", time.Now().Format("2006-01-02 15:04:05"))

		}
	}(tickerSecond)
	wg.Wait()
}
