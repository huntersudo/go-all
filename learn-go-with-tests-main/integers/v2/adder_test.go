package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}
// 请注意，如果删除注释 「//Output: 6」，示例函数将不会执行。虽然函数会被编译，但是它不会执 行。
// 通过添加这段代码，示例将出现在 godoc 的文档中，这将使你的代码更容易理解
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
