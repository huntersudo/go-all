// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 21.

// Server3 is an "echo" server that displays request parameters.
package main

import (
	"fmt"
	"log"
	"net/http"
)

//GET /?q=query HTTP/1.1
//Header["Sec-Fetch-Dest"] = ["document"]
//Header["Connection"] = ["keep-alive"]
//Header["User-Agent"] = ["Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.96 Safari/537.36"]
//Header["Accept"] = ["text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"]
//Header["Sec-Fetch-Site"] = ["none"]
//Header["Sec-Fetch-Mode"] = ["navigate"]
//Header["Sec-Fetch-User"] = ["?1"]
//Header["Upgrade-Insecure-Requests"] = ["1"]
//Header["Accept-Encoding"] = ["gzip, deflate, br"]
//Header["Accept-Language"] = ["zh-CN,zh;q=0.9"]
//Host = "localhost:8000"
//RemoteAddr = "127.0.0.1:52777"
//Form["q"] = ["query"]

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!+handler
// handler echoes the HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

//!-handler
