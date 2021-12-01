// Copyright 2015 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package errors implements functions to manipulate errors.
package errors // import "github.com/chai2010/errors"

import (
	"errors"
	"fmt"
	"strings"
)

var (
	_ Error        = (*_Error)(nil)
	_ fmt.Stringer = (*_Error)(nil)
)

// Error 其中 Error 为接口类型，是 error 接口类型的扩展
// 用于给错误增加调用栈信息，同时支持错误 的多级嵌套包装，支持错误码格式。
type Error interface {
	Caller() []CallerInfo
	Wraped() []error
	Code() int
	error

	private()
}

type _Error struct {
	XCode   int          `json:"Code"`
	XError  error        `json:"Error"`
	XCaller []CallerInfo `json:"Caller,omitempty"`
	XWraped []error      `json:"Wraped,omitempty"`
}

type CallerInfo struct {
	FuncName string
	FileName string
	FileLine int
}

// New 辅助函数
// New 用于构建新的错误类型，和标准库中 errors.New 功能类似，但是增加了出错时的函数调用栈 信息
func New(msg string) error {
	return &_Error{
		XCaller: Caller(2),
		XError:  errors.New(msg),
	}
}

func NewFrom(err error) error {
	if err == nil {
		return nil
	}
	if e, ok := err.(Error); ok {
		return e
	}
	return &_Error{
		XCaller: Caller(2),
		XError:  err,
	}
}

func Newf(format string, args ...interface{}) error {
	return &_Error{
		XCaller: Caller(2),
		XError:  fmt.Errorf(format, args...),
	}
}

// NewWithCode 辅助函数
// NewWithCode 则是构造一个带 错误码的错误，同时也包含出错时的函数调用栈信息
func NewWithCode(code int, msg string) error {
	return &_Error{
		XCaller: Caller(2),
		XError:  errors.New(msg),
		XCode:   code,
	}
}

func NewWithCodef(code int, format string, args ...interface{}) error {
	return &_Error{
		XCaller: Caller(2),
		XError:  fmt.Errorf(format, args...),
		XCode:   code,
	}
}

func MustFromJson(json string) error {
	p, err := newErrorStructFromJson(json)
	if err != nil {
		panic(err)
	}
	return p.ToStdError()
}

// FromJson 辅助函数
// 用于从JSON字符串编码的错误中恢复错误对象
func FromJson(json string) (Error, error) {
	p, err := newErrorStructFromJson(json)
	if err != nil {
		return nil, &_Error{
			XCaller: Caller(1), // skip == 1
			XWraped: []error{err},
			XError:  errors.New(fmt.Sprintf("errors.FromJson: jsonDecode failed: %v!", err)),
		}
	}

	return p.ToErrorInterface(), nil
}

func ToJson(err error) string {
	if p, ok := (err).(*_Error); ok {
		return p.String()
	}
	p := &_Error{XError: err}
	return p.String()
}

// Wrap 辅助函数
// Wrap 和 WrapWithCode 则是错误二次包装 函数，用于将底层的错误包装为新的错误，但是保留的原始的底层错误信息。
// 这里返回的错误对象都可 以直接调用 json.Marshal 将错误编码为JSON字符串
func Wrap(err error, msg string) error {
	p := &_Error{
		XCaller: Caller(2),
		XWraped: []error{err},
		XError:  errors.New(fmt.Sprintf("%s -> {%v}", msg, err)),
	}
	if e, ok := err.(Error); ok {
		p.XWraped = append(p.XWraped, e.Wraped()...)
	}
	return p
}

func Wrapf(err error, format string, args ...interface{}) error {
	p := &_Error{
		XCaller: Caller(2),
		XWraped: []error{err},
		XError:  errors.New(fmt.Sprintf("%s -> {%v}", fmt.Sprintf(format, args...), err)),
	}
	if e, ok := err.(Error); ok {
		p.XWraped = append(p.XWraped, e.Wraped()...)
	}
	return p
}

// WrapWithCode 辅助函数
func WrapWithCode(code int, err error, msg string) error {
	p := &_Error{
		XCaller: Caller(2),
		XWraped: []error{err},
		XError:  errors.New(fmt.Sprintf("%s -> {%v}", msg, err)),
		XCode:   code,
	}
	if e, ok := err.(Error); ok {
		p.XWraped = append(p.XWraped, e.Wraped()...)
	}
	return p
}

func WrapWithCodef(code int, err error, format string, args ...interface{}) error {
	p := &_Error{
		XCaller: Caller(2),
		XWraped: []error{err},
		XError:  errors.New(fmt.Sprintf("%s -> {%v}", fmt.Sprintf(format, args...), err)),
		XCode:   code,
	}
	if e, ok := err.(Error); ok {
		p.XWraped = append(p.XWraped, e.Wraped()...)
	}
	return p
}

func Caller(skip int) []CallerInfo {
	var infos []CallerInfo
	for ; ; skip++ {
		name, file, line, ok := callerInfo(skip + 1)
		if !ok {
			return infos
		}
		if strings.HasPrefix(name, "runtime.") {
			return infos
		}
		infos = append(infos, CallerInfo{
			FuncName: name,
			FileName: file,
			FileLine: line,
		})
	}
	panic("unreached!")
}

func (p *_Error) Caller() []CallerInfo {
	return p.XCaller
}

func (p *_Error) Wraped() []error {
	return p.XWraped
}

func (p *_Error) Error() string {
	return p.XError.Error()
}

func (p *_Error) Code() int {
	return p.XCode
}

func (p *_Error) String() string {
	return jsonEncodeString(p)
}

func (p *_Error) MarshalJSON() ([]byte, error) {
	return jsonEncode(newErrorStruct(p)), nil
}

func (p *_Error) UnmarshalJSON(data []byte) error {
	px, err := newErrorStructFromJson(string(data))
	if err != nil {
		return err
	}
	if px != nil {
		*p = *px.ToError()
	} else {
		*p = _Error{}
	}
	return nil
}

func (p *_Error) private() {
	panic("unreached!")
}
