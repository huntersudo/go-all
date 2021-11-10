package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

func main() {
	// 示例1。
	builder := new(strings.Builder)
	_ = interface{}(builder).(io.Writer)
	_ = interface{}(builder).(io.ByteWriter)
	_ = interface{}(builder).(fmt.Stringer)

	// 示例2。
	reader := strings.NewReader("")
	_ = interface{}(reader).(io.Reader)
	_ = interface{}(reader).(io.ReaderAt)
	_ = interface{}(reader).(io.ByteReader)
	_ = interface{}(reader).(io.RuneReader)
	_ = interface{}(reader).(io.Seeker)
	_ = interface{}(reader).(io.ByteScanner)
	_ = interface{}(reader).(io.RuneScanner)
	_ = interface{}(reader).(io.WriterTo)

	// 示例3。
	buffer := bytes.NewBuffer([]byte{})
	_ = interface{}(buffer).(io.Reader)
	_ = interface{}(buffer).(io.ByteReader)
	_ = interface{}(buffer).(io.RuneReader)
	_ = interface{}(buffer).(io.ByteScanner)
	_ = interface{}(buffer).(io.RuneScanner)
	_ = interface{}(buffer).(io.WriterTo)

	_ = interface{}(buffer).(io.Writer)
	_ = interface{}(buffer).(io.ByteWriter)
	_ = interface{}(buffer).(io.ReaderFrom)

	_ = interface{}(buffer).(fmt.Stringer)

	// 示例4。
	src := strings.NewReader(
		"CopyN copies n bytes (or until an error) from src to dst. " +
			"It returns the number of bytes copied and " +
			"the earliest error encountered while copying.")
	dst := new(strings.Builder)
	// 变量src和dst的类型分别是strings.Reader和strings.Builder，
	// 但是当它们被传到io.CopyN函数的时候，就已经分别被包装成了io.Reader类型和io.Writer类型的值。
	// io.CopyN函数也根本不会去在意，它们的实际类型到底是什么。
	written, err := io.CopyN(dst, src, 58)
	// io.CopyN函数的作者通过面向接口编程，极大地拓展了它的适用范围和应用场景
	// 换个角度看，正因为strings.Reader类型和strings.Builder类型都实现了不少接口，所以它们的值才能够被使用在更广阔的场景中。
	if err != nil {
		fmt.Printf("error: %v\n", err)
	} else {
		fmt.Printf("Written(%d): %q\n", written, dst.String())
	}
}
