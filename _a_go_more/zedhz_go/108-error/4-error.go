package main

import (
	"fmt"
	"github.com/pkg/errors"
)

// 包装错误

func  wrap1(err error)  error{
	if err != nil {
		return fmt.Errorf("something failed: %v", err)
	}
	return nil
}

// 另外，在 Go 语言的开发者中，更为普遍的做法是将错误包装在另一个错误中，同时保留原始内容：

type authorizationError struct {
	operation string
	err error   // original error
}

func (e *authorizationError) Error() string {
	return fmt.Sprintf("authorization failed during %s: %v", e.operation, e.err)
}

//todo 当然，更好的方式是通过一种标准的访问方法，这样，我们最好使用一个接口，比如 causer接口中实现 Cause() 方法来暴露原始错误，以供进一步检查：

type causer interface {
	Cause() error
}

func (e *authorizationError) Cause() error {
	return e.err
}

// 这里有个好消息是，这样的代码不必再写了，有一个第三方的错误库，对于这个库，我无论到哪儿都能看到它的存在，
// 所以，这个基本上来说就是事实上的标准了。代码示例如下：
type MyError int

func ( a MyError) Error() string {
	fmt.Println(a.Error())
	return "nil"
}
func funcErrorWrap(err error) error{

	//错误包装
	if err != nil {
		return errors.Wrap(err, "read failed")
	}
	// Cause接口
	switch err := errors.Cause(err).(type) {
	case *MyError:
		fmt.Println(err)
		// handle specifically
	default:
		// unknown error
	}
	return nil
}