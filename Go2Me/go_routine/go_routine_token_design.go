package go_routine

import (
	"time"
	"crypto/sha1"
	"io"
	"fmt"
	"math/rand"
)

type TokenHeatBeat struct {
	UId   chan int
	Msg   chan string
	Token chan string
}

var tokenStore TokenHeatBeat

func timeSedHash() string {
	t := sha1.New();
	io.WriteString(t, time.Now().String());
	return fmt.Sprintf("%x", t.Sum(nil));
}

func NewTokenInfo() {
	tokenStore.Msg = make(chan string, 10)
	tokenStore.Token = make(chan string, 10)
	tokenStore.UId = make(chan int, 10)
	thisTime := time.Now()
	tokenStore.UId <- thisTime.Minute() / 2
	tokenStore.Msg <- thisTime.String()
	tokenStore.Token <- timeSedHash()
}

func updateTokenInfo() {
	thisTime := time.Now()
	r := rand.New(rand.NewSource(thisTime.UnixNano()))
	res_int := r.Intn(6)
	tokenStore.UId <- res_int % 2
	tokenStore.Msg <- thisTime.String()
	tokenStore.Token <- timeSedHash()
	time.Sleep(time.Second * 1)
}

func ReadToken(key string) {
	for i := 0; i < 10; i++ {
		go updateTokenInfo()
	}
	select {
	case t := <-tokenStore.Token:
		if key == t {
			fmt.Printf("Token token %s, msg %s \n", <-tokenStore.Token, <-tokenStore.Msg)
		}
	case id := <-tokenStore.UId:
		if id > 0 {
			fmt.Printf("UId token %s, msg %s \n", <-tokenStore.Token, <-tokenStore.Msg)
		} else {
			fmt.Printf("UnUId token %s, msg %s \n", <-tokenStore.Token, <-tokenStore.Msg)
		}


	}

}


