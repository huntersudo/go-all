
https://jishuin.proginn.com/p/763bfbd6476a

### install 
``` 
go get github.com/onsi/ginkgo/ginkgo
go get github.com/onsi/gomega/...
```
### 创建套件

我们通过 ginkgo bootstrap 命令，来初始化一个 Ginkgo 测试套件。

``` 
ginkgo bootstrap
```

此时在 gopher.go 同级目录中，生成了 gopher_suite_test.go 文件，内容如下

``` 
package gopher_test

import (
 "testing"

 . "github.com/onsi/ginkgo"
 . "github.com/onsi/gomega"
)

func TestGopher(t *testing.T) {
 RegisterFailHandler(Fail)
 RunSpecs(t, "Gopher Suite")
}
```
此时，我们就可以运行测试套件了，通过命令 go test 或 ginkgo 均可。

``` 
$ go test
Running Suite: Gohper Suite
===========================
Random Seed: 1642684664
Will run 0 of 0 specs


Ran 0 of 0 Specs in 0.006 seconds
SUCCESS! -- 0 Passed | 0 Failed | 0 Pending | 0 Skipped
PASS
ok      gexp/go_other/bdd_test/gohper   0.246s

```

当然，空测试套件没有什么价值，我们需要在此套件下编写测试（Spec）用例。

我们可以在 gopher_suite_test.go 中编写测试，但是推荐分离到独立的文件中，特别是包中有多个需要被测试的源文件的情况下。

### 创建 Spec 

执行 ginkgo generate gopher  可以生成一个 gopher_test.go 测试文件。
``` 
ginkgo generate gopher
```

此时测试文件中的内容如下

``` 
package gopher_test

import (
 . "github.com/onsi/ginkgo"
)

var _ = Describe("Gopher", func() {

})
```
### 编写 Spec

我们基于此测试文件撰写实际的测试用例,详见文件gopher_test.go。          

可以看到，BDD 风格的测试案例在代码中就被描述地非常清晰。由于我们的测试用例与预期相符，执行 go test 执行测试套件会校验通过。            

读者可自行更改数据致测试不通过，你会看到 Ginkgo 将打印出堆栈与错误描述性信息。         


TDD 和 BDD 是敏捷开发中常被提到的方法论。与TDD相比，BDD 通过编写 行为和规范 来驱动软件开发。这些行为和规范在代码中体现于更 ”繁琐“ 的描述信息。          

关于 BDD 的本质，有另外一种表达方式：BDD 帮助开发人员设计软件，TDD 帮助开发人员测试软件。   

Ginkgo 是 Go 语言中非常优秀的 BDD 框架，它通过 DSL 语法（Describe/Context/It）有效地帮助开发者组织与编排测试用例。
本文只是展示了 Ginkgo 非常简单的用例，权当是抛砖引玉。


读者在使用 Ginkgo 过程中，需要理解它的执行生命周期， 重点包括  It、Context、Describe、BeforeEach、AfterEach、JustBeforeEach、
BeforeSuite、AfterSuite、By、Fail  这些模块的执行顺序与语义逻辑。

Ginkgo 有很多的功能本文并未涉及，例如异步测试、基准测试、持续集成等强大的支持。
其仓库位于 https://github.com/onsi/ginkgo ，同时提供了英文版与中文版使用文档，读者可以借此了解更多 Ginkgo 信息。

### todo 

最后，K8s 项目中也使用了 Ginkgo 框架，用于编写其端到端 (End to End，E2E) 测试用例，值得借鉴学习。

### 附录


