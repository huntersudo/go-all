package main

import (
	"flag"
	"fmt"
)

//定义⼀些全局变量存储选项的值，如这⾥的 intflag/boolflag/stringflag ；
//在 init ⽅法中使⽤ flag.TypeVar ⽅法定义选项，这⾥的 Type 可以为基本类型 Int/Uint/Float64/Bool ，还可以是时间 间隔 time.Duration 。定义时传⼊变量的地址、选项名、默认值和帮助信息；
//在 main ⽅法中调⽤ flag.Parse 从 os.Args[1:] 中解析选项。因为 os.Args[0] 为可执⾏程序路径，会被剔除。
var (
	// intflag int
	// boolflag bool
	// stringflag string

	intflag    *int
	boolflag   *bool
	stringflag *string
)

func init() {
	// flag.IntVar(&intflag, "intflag", 0, "int flag value")
	// flag.BoolVar(&boolflag, "boolflag", false, "bool flag value")
	// flag.StringVar(&stringflag, "stringflag", "default", "string flag value")

	intflag = flag.Int("intflag", 0, "int flag value")
	boolflag = flag.Bool("boolflag", false, "bool flag value")
	stringflag = flag.String("stringflag", "default", "string flag value")
}

func main() {
	flag.Parse()

	// fmt.Println("int flag:", intflag)
	// fmt.Println("bool flag:", boolflag)
	// fmt.Println("string flag:", stringflag)

	fmt.Println("int flag:", *intflag)
	fmt.Println("bool flag:", *boolflag)
	fmt.Println("string flag:", *stringflag)

	fmt.Println(flag.Args())
	fmt.Println("Non-Flag Argument Count:", flag.NArg())
	for i := 0; i < flag.NArg(); i++ {
		fmt.Printf("Argument %d: %s\n", i, flag.Arg(i))
	}

	fmt.Println("Flag Count:", flag.NFlag())
}
