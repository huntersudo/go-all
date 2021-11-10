package main

import (
	"fmt"
	"testing"
)
/**
不要启动 gopath

$ go test puzzlers/src/puzzlers/article20/q0
ok      puzzlers/src/puzzlers/article20/q0      0.375s

 被缓存
$ go test puzzlers/src/puzzlers/article20/q0
ok      puzzlers/src/puzzlers/article20/q0      (cached)

go clean -cache

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

/**
  === RUN   TestHello
      demo52_test.go:34: The expected greeting is "Hello, Robert!".
  --- PASS: TestHello (0.00s)
  PASS
 */
}

/**
对于功能测试函数来说，其名称必须以Test为前缀，并且参数列表中只应有一个*testing.T类型的参数声明。
！！测试函数名称不合法
 */
func testIntroduce(t *testing.T) { // 请注意这个测试函数的名称。
	intro := introduce()
	expected := "Welcome to my Golang column."
	if intro != expected {
		t.Errorf("The actual introduce %q is not the expected.",
			intro)
	}
	t.Logf("The expected introduce is %q.\n", expected)
}
