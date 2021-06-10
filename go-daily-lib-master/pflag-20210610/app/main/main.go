package pflag

import flag "github.com/spf13/pflag"

import (
	"fmt"
	"strings"
)

// 定义命令行参数对应的变量
var cliName = flag.StringP("name", "n", "nick", "Input Your Name")
var cliAge = flag.IntP("age", "a",22, "Input Your Age")
var cliGender = flag.StringP("gender", "g","male", "Input Your Gender")
var cliOK = flag.BoolP("ok", "o", false, "Input Are You OK")
var cliDes = flag.StringP("des-detail", "d", "", "Input Description")
var cliOldFlag = flag.StringP("badflag", "b", "just for test", "Input badflag")

func wordSepNormalizeFunc(f *flag.FlagSet, name string) flag.NormalizedName {
	from := []string{"-", "_"}
	to := "."
	for _, sep := range from {
		name = strings.Replace(name, sep, to, -1)
	}
	return flag.NormalizedName(name)
}

func main() {
	// 设置标准化参数名称的函数
	flag.CommandLine.SetNormalizeFunc(wordSepNormalizeFunc)

	//pflag 包支持通过简便的方式为参数设置默认值之外的值，实现方式为设置参数的 NoOptDefVal 属性
	// 为 age 参数设置 NoOptDefVal
	// 比如-a不带参数则为25
	flag.Lookup("age").NoOptDefVal = "25"

	// 把 badflag 参数标记为即将废弃的，请用户使用 des-detail 参数
	flag.CommandLine.MarkDeprecated("badflag", "please use --des-detail instead")

	// 把 badflag 参数的 shorthand 标记为即将废弃的，请用户使用 des-detail 的 shorthand 参数
	flag.CommandLine.MarkShorthandDeprecated("badflag", "please use -d instead")

	// 在帮助文档中隐藏参数 gender
	flag.CommandLine.MarkHidden("badflag")

	// 把用户传递的命令行参数解析为对应变量的值
	flag.Parse()

	fmt.Println("name=", *cliName)
	fmt.Println("age=", *cliAge)
	fmt.Println("gender=", *cliGender)
	fmt.Println("ok=", *cliOK)
	fmt.Println("des=", *cliDes)
}

// //帮助文档中没有显示 badflag 的信息。其实在把参数标记为废弃时，同时也会设置隐藏参数。
//cts@cts-pc:~/test$ ./test
//name= nick
//age= 22
//gender= male
//ok= false
//des=
//cts@cts-pc:~/test$

//cts@cts-pc:~/test$ ./test -b=sssss
//Flag shorthand -b has been deprecated, please use -d instead
//Flag --badflag has been deprecated, please use --des-detail instead
//name= nick
//age= 22
//gender= male
//ok= false
//des=

//cts@cts-pc:~/test$ ./test --age=369
//name= nick
//age= 369
//gender= male
//ok= false
//des=

//cts@cts-pc:~/test$ ./test -a=369
//name= nick
//age= 369
//gender= male
//ok= false
//des=

//cts@cts-pc:~/test$ ./test -a
//name= nick
//age= 25
//gender= male
//ok= false
//des=