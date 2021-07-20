package main

import (
	"fmt"
	"testing"
)

/***
不要启动 gopath

== t.FailNow() 调用 ，结果如下
$ go test puzzlers/src/puzzlers/article20/q2/
--- FAIL: TestFail (0.00s)
FAIL
FAIL    puzzlers/src/puzzlers/article20/q2      0.375s
FAIL
======================================================
 == t.Fail() 调用 ，结果如下
失败的测试函数中的常规测试日志一并被打印出来。

$ go test puzzlers/src/puzzlers/article20/q2
--- FAIL: TestFail (0.00s)
demo53_test.go:62: Failed.
FAIL
FAIL    puzzlers/src/puzzlers/article20/q2      0.402s
FAIL
 */

/* -v 打印常规日志
=======================================================
$ go test -v puzzlers/src/puzzlers/article20/q2
=== RUN   TestHello
--- PASS: TestHello (0.00s)
    demo53_test.go:76: The expected greeting is "Hello, Robert!".
=== RUN   TestIntroduce
--- PASS: TestIntroduce (0.00s)
    demo53_test.go:86: The expected introduce is "Welcome to my Golang column.".
=== RUN   TestFail
--- FAIL: TestFail (0.00s)
    demo53_test.go:92: Failed.
FAIL
FAIL    puzzlers/src/puzzlers/article20/q2      3.412s
FAIL

*/

/**
go test命令就会针对每个被测代码包，依次地进行构建、执行包中符合要求的测试函数，清理临时文件，打印测试结果。
 */
func TestHello(t *testing.T) {
	var name string
	greeting, err := hello(name)
	if err == nil {
		t.Errorf("The error is nil, but it should not be. (name=%q)",
			name)
	}
	if greeting != "" {
		t.Errorf("Nonempty greeting, but it should not be. (name=%q)",
			name)
	}
	name = "Robert"
	greeting, err = hello(name)
	if err != nil {
		t.Errorf("The error is not nil, but it should be. (name=%q)",
			name)
	}
	if greeting == "" {
		t.Errorf("Empty greeting, but it should not be. (name=%q)",
			name)
	}
	expected := fmt.Sprintf("Hello, %s!", name)
	if greeting != expected {
		t.Errorf("The actual greeting %q is not the expected. (name=%q)",
			greeting, name)
	}
	t.Logf("The expected greeting is %q.\n", expected)
}

func TestIntroduce(t *testing.T) {
	intro := introduce()
	expected := "Welcome to my Golang column."
	if intro != expected {
		t.Errorf("The actual introduce %q is not the expected.",
			intro)
	}
	t.Logf("The expected introduce is %q.\n", expected)
}

func TestFail(t *testing.T) {
	t.Fail()
	//t.FailNow() // 此调用会让当前的测试立即失败。
	t.Log("Faileddd.")
}

