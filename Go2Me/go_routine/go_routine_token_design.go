package go_routine

import (
	"time"
	"crypto/sha1"
	"io"
	"fmt"
	"math/rand"
)

type TokenHeatBeat struct {
	// Uid behavior
	UId   chan int
	// message behavior
	Msg   chan string
	// token check behavior
	Token chan string
}

var tokenStore TokenHeatBeat

func timeSedHash() string {
	t := sha1.New();
	io.WriteString(t, time.Now().String());
	return fmt.Sprintf("%x", t.Sum(nil));
}

// init token info
func NewTokenInfo() {
	tokenStore.Msg = make(chan string, 10)
	tokenStore.Token = make(chan string, 10)
	tokenStore.UId = make(chan int, 10)
	thisTime := time.Now()
	tokenStore.UId <- thisTime.Minute() % 2
	tokenStore.Msg <- thisTime.String()
	tokenStore.Token <- timeSedHash()
}

// you can set your update token logic
func updateTokenInfo() {
	thisTime := time.Now()
	r := rand.New(rand.NewSource(thisTime.UnixNano()))
	res_int := r.Intn(6)
	tokenStore.UId <- res_int % 2
	tokenStore.Msg <- thisTime.String()
	tokenStore.Token <- timeSedHash()
	time.Sleep(time.Second * 1)
}

// ReadToken this is routine way update, reactive info by different behavior
func ReadToken(key string) {
	for i := 0; i < 10; i++ {
		go updateTokenInfo()
	}
	select {
	case t := <-tokenStore.Token:
	// behavior token response
		if key == t {
			fmt.Printf("Token token %s, msg %s \n", <-tokenStore.Token, <-tokenStore.Msg)
		}
	case id := <-tokenStore.UId:
	// behavior UId check response
		if id > 0 {
			fmt.Printf("UId token %s, msg %s \n", <-tokenStore.Token, <-tokenStore.Msg)
		} else {
			fmt.Printf("UnUId token %s, msg %s \n", <-tokenStore.Token, <-tokenStore.Msg)
		}


	}

}


