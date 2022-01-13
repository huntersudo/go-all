package concurrency

import (
	"testing"
	"time"
)

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
//goos: windows
//goarch: amd64
//pkg: github.com/quii/learn-go-with-tests/concurrency/v1
//BenchmarkCheckWebsites-12    	       1	2075924900 ns/op
//PASS
