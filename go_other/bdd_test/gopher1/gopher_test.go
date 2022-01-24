package gopher1_test

import (
	. "github.com/onsi/ginkgo"
	 "github.com/onsi/gomega"

	gopher "bdd_test/gopher1"
)

func mockInputData() ([]gopher.Gopher, error) {
	inputData := []gopher.Gopher{
		{Name: "菜刀",
			Gender: "男",
			Age:    18,
		},
		{
			Name:   "小西瓜",
			Gender: "女",
			Age:    19,
		},
		{
			Name:   "机器铃砍菜刀",
			Gender: "男",
			Age:    17,
		},
		{
			Name:   "小菜刀",
			Gender: "男",
			Age:    20,
		},
	}
	return inputData, nil
}

// todo
//  您可以使⽤ Describe 块来描述代码的各个⾏为， Context 块在不同情况下执⾏这些⾏为。
//  使⽤ BeforeEach 块在多个测试⽤例中去除重复的步骤以及共享通⽤的设置：
//   BeforeEach 在每个 spec 之前运⾏，从⽽确保每个 spec 都具有状态的原始副本。使⽤闭包变量共享公共状态（在本例中 为 var book Book ）。您还可以在 AfterEach 块中执⾏清理操作
//    当嵌套 Describe 和 Context 块时， It 执⾏时，围绕 It 的所有容器节点的 BeforeEach 块，从最外层到最内层运⾏。
//   通常，容器块中的唯⼀代码应该是 It 块或 BeforeEach / JustBeforeEach / JustAfterEach / AfterEach 块 或闭包变量声明。
//  在容器块中进⾏断⾔通常是错误的。 在容器块中初始化闭包变量也是错误的。
//  如果你的⼀个 It 改变了这个变量，后期 It 将会收到改变后的值。这是⼀个测试污染的案例，很难追查。始终在 BeforeEach 块中 初始化变量。
//   ===========
//  我们的顶级 BeforeEach 使⽤有效的 JSON 创建了⼀个新的 book ,但是 较低级别的 Context 使⽤⽆效的JSON创建的 book 执⾏。
//  JustBeforeEach 块保证在所有 BeforeEach 块运⾏之后，并且在 It 块运⾏之前运⾏。我们可以使⽤这个特性来清除 Book spec：

var _ = Describe("Gopher", func() {
	BeforeEach(func() {
		By("当测试不通过时，我会在这里打印一个消息 【BeforeEach】")
	})

	inputData, err := mockInputData()

	Describe("校验输入数据", func() {

		Context("当获取数据没有错误发生时", func() {
			It("它应该是接收数据成功了的", func() {
				gomega.Expect(err).Should(gomega.BeNil())
			})
		})

		// TODO 这里出错
		Context("当获取的数据校验失败时", func() {
			It("当数据校验返回错误为：名字太短，不能小于3 时", func() {
				gomega.Expect(gopher.Validate(inputData[0])).Should(gomega.MatchError("名字太短，不能小于33"))
			})

			It("当数据校验返回错误为：只要男的 时", func() {
				gomega.Expect(gopher.Validate(inputData[1])).Should(gomega.MatchError("只要男的"))
			})

			It("当数据校验返回错误为：岁数太小，不能小于18 时", func() {
				gomega.Expect(gopher.Validate(inputData[2])).Should(gomega.MatchError("岁数太小，不能小于18"))
			})
		})

		Context("当获取的数据校验成功时", func() {
			It("通过了数据校验", func() {
				gomega.Expect(gopher.Validate(inputData[3])).Should(gomega.BeNil())
			})
		})
	})

	AfterEach(func() {
		By("当测试不通过时，我会在这里打印一个消息 【AfterEach】")
	})
})
/**

+?[1mSTEP?[0m: 当测试不通过时，我会在这里打印一个消息 【BeforeEach】
?[1mSTEP?[0m: 当测试不通过时，我会在这里打印一个消息 【AfterEach】

------------------------------
+ Failure [0.001 seconds]
Gopher
D:/workspace/gopath/src/github.com/go-all/go_other/bdd_test/gopher1/gopher_test.go:35
  校验输入数据
  D:/workspace/gopath/src/github.com/go-all/go_other/bdd_test/gopher1/gopher_test.go:42
    当获取的数据校验失败时
    D:/workspace/gopath/src/github.com/go-all/go_other/bdd_test/gopher1/gopher_test.go:50
      当数据校验返回错误为：名字太短，不能小于3 时 [It]
      D:/workspace/gopath/src/github.com/go-all/go_other/bdd_test/gopher1/gopher_test.go:51

      Expected
          <*errors.errorString | 0xc000204290>: {
              s: "名字太短，不能小于3",
          }
      to match error
          <string>: 名字太短，不能小于33

      D:/workspace/gopath/src/github.com/go-all/go_other/bdd_test/gopher1/gopher_test.go:52
------------------------------
+++

Summarizing 1 Failure:

[Fail] Gopher 校验输入数据 当获取的数据校验失败时 [It] 当数据校验返回错误为：名字太短，不能小于3 时
D:/workspace/gopath/src/github.com/go-all/go_other/bdd_test/gopher1/gopher_test.go:52

Ran 5 of 5 Specs in 0.069 seconds
FAIL! -- 4 Passed | 1 Failed | 0 Pending | 0 Skipped
--- FAIL: TestGopher (0.09s)
FAIL
exit status 1
FAIL    bdd_test/gopher1        0.423s



*/