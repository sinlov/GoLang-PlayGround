package main

import (
	"flag"
	"fmt"
	"github.com/pkg/profile"
	"github.com/sinlov/GoLang-PlayGround/Go2Me/go_profile/httpdebug"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	hookProfile()
	flag.Parse()
	rand.Seed(time.Now().UTC().UnixNano())
	http.HandleFunc("/", HelloHandler)
	http.HandleFunc("/debug/vars_base", httpdebug.DebugHandler)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

//nolint:golint,unused
func a() {
	for i := 0; i < 10000; i++ {
		fmt.Print(".")
	}
	fmt.Println()
}

func hookProfile() {
	defer profile.Start(
		profile.MutexProfile,
		profile.CPUProfile,
		profile.GoroutineProfile,
		profile.MemProfile,
		profile.MemProfileAllocs,
		profile.MemProfileHeap,
		profile.MemProfileAllocs,
		//profile.BlockProfile,
		//profile.NoShutdownHook,
		profile.ProfilePath("."),
	).Stop()
}
