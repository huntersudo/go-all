package main

import (
	"encoding/binary"
	"io"
)

//说到 Go 语言的 if err !=nil 的代码了，这样的代码的确是能让人写到吐
//
func parse2(r io.Reader) (*Point, error) {
	var p Point
	var err error
	read := func(data interface{}) {
		if err != nil {
			return
		}
		err = binary.Read(r, binary.BigEndian, data)
	}

	read(&p.Longitude)
	read(&p.Latitude)
	read(&p.Distance)
	read(&p.ElevationGain)
	read(&p.ElevationLoss)

	if err != nil {
		return &p, err
	}
	return &p, nil
}

// 这样大量的 if err!=nil 处理得很干净了，但是会带来一个问题，那就是有一个 err 变量和一个内部的函数，感觉不是很干净。

/*func scan()  {
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		token := scanner.Text()
		// process token
	}
	if err := scanner.Err(); err != nil {
		// process the error
	}

}*/
// ，退出循环后有一个 scanner.Err() 的检查，看来使用了结构体的方式。模仿它


type Reader struct {
	r   io.Reader
	err error
}

func (r *Reader) read(data interface{}) {
	if r.err == nil {
		r.err = binary.Read(r.r, binary.BigEndian, data)
	}
}

func parse3(input io.Reader) (*Point, error) {
	var p Point
	r := Reader{r: input}

	r.read(&p.Longitude)
	r.read(&p.Latitude)
	r.read(&p.Distance)
	r.read(&p.ElevationGain)
	r.read(&p.ElevationLoss)

	if r.err != nil {
		return nil, r.err
	}

	return &p, nil
}
//todo 有了刚刚的这个技术，我们的“流式接口 Fluent Interface”也就很容易处理了。如下所示：



