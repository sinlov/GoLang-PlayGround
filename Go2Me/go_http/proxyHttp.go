package go_http

import (
	"net/url"
	"net/http/httputil"
	"net/http"
	"log"
	"flag"
)

type handle struct {
	reverseProxy string
}

type ProxyFlag struct {
	Listen  string `flag:"l"`
	Reverse string `flag:"r"`
}

func (this *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	remote, err := url.Parse(this.reverseProxy)
	if err != nil {
		log.Fatalln(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	r.Host = remote.Host
	proxy.ServeHTTP(w, r)
	log.Println(r.RemoteAddr + " " + r.Method + " " + r.URL.String() + " " + r.Proto + " " + r.UserAgent())
}

func Bind(listen string, reverse string) {
	bind := flag.String("l", listen, "listen on ip:port")
	remote := flag.String("r", reverse, "reverse proxy addr")
	flag.Parse()
	log.Printf("Listening on %s, forwarding to %s", *bind, *remote)
	h := &handle{reverseProxy: *remote}
	err := http.ListenAndServe(*bind, h)
	if err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}
