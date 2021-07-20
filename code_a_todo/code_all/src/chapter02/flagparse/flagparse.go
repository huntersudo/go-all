package main

// 导入系统包
import (
	"flag"
	"fmt"
)

// 定义命令行参数
// args: 名称【注册的指针】、默认值、说明
var mode = flag.String("mode", "", "process mode")

func main() {

	// 解析命令行参数
	flag.Parse()

	// 输出命令行参数
	fmt.Println(*mode)
}
