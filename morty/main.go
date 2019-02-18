package main

import (
	. ".."
	"flag"
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
	"os"
	"time"
)

func main() {
	default_listen_addr := os.Getenv("MORTY_ADDRESS")
	if default_listen_addr == "" {
		default_listen_addr = "127.0.0.1:3000"
	}
	default_key := os.Getenv("MORTY_KEY")
	listen := flag.String("listen", default_listen_addr, "Listen address")
	key := flag.String("key", default_key, "HMAC url validation key (hexadecimal encoded) - leave blank to disable validation")
	ipv6 := flag.Bool("ipv6", false, "Allow IPv6 HTTP requests")
	version := flag.Bool("version", false, "Show version")
	requestTimeout := flag.Uint("timeout", 2, "Request timeout")
	flag.Parse()

	if *version {
		fmt.Println(VERSION)
		return
	}

	if *ipv6 {
		CLIENT.Dial = fasthttp.DialDualStack
	}

	p := &Proxy{RequestTimeout: time.Duration(*requestTimeout) * time.Second}

	if *key != "" {
		p.Key = []byte(*key)
	}

	log.Println("listening on", *listen)

	if err := fasthttp.ListenAndServe(*listen, p.RequestHandler); err != nil {
		log.Fatal("Error in ListenAndServe:", err)
	}
}
