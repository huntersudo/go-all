package main

import "fmt"

type Button struct {
	Label // Embedding (delegation)
}

type ListBox struct {
	Widget          // Embedding (delegation)
	Texts  []string // Aggregation
	Index  int      // Aggregation
}

// todo 然后，我们需要两个接口：用 Painter 把组件画出来；Clicker 用于表明点击事件。
type Painter interface {
	Paint()
}

type Clicker interface {
	Click()
}
//当然，对于 Lable 来说，只有 Painter ，没有Clicker；
//    对于 Button 和 ListBox来说，Painter 和Clicker都有。
//

func (label Label) Paint() {
	fmt.Printf("%p:Label.Paint(%q)\n", &label, label.Text)
}

//因为这个接口可以通过 Label 的嵌入带到新的结构体，
//所以，可以在 Button 中重载这个接口方法
func (button Button) Paint() { // Override
	fmt.Printf("Button.Paint(%s)\n", button.Text)
}
func (button Button) Click() {
	fmt.Printf("Button.Click(%s)\n", button.Text)
}
func (listBox ListBox) Paint() {
	fmt.Printf("ListBox.Paint(%q)\n", listBox.Texts)
}
func (listBox ListBox) Click() {
	fmt.Printf("ListBox.Click(%q)\n", listBox.Texts)
}
// todo 说到这儿，我要重点提醒你一下，Button.Paint() 接口可以通过 Label 的嵌入带到新的结构体，
//   如果 Button.Paint() 不实现的话，会调用 Label.Paint() ，所以，在 Button 中声明 Paint() 方法，相当于 Override。

func use2(){

	label:= Label{Widget{10, 70},"label"}
	button1 := Button{Label{Widget{10, 70}, "button1"}}
	button2 := Button{Label{Widget{50, 100}, "button2"}}
	listBox := ListBox{Widget{10, 40},
		[]string{"AL", "AK", "AZ", "AR"}, 0}

	for _, painter := range []Painter{label, listBox, button1, button2} {
		painter.Paint()
	}

	for _, widget := range []interface{}{label, listBox, button1, button2} {

		widget.(Painter).Paint()
		// todo 如果是click类型，则调用
		//  我们可以使用接口来多态，也可以使用泛型的 interface{} 来多态，但是需要有一个类型转换。
		if clicker, ok := widget.(Clicker); ok {
			clicker.Click()
		}
		fmt.Println() // print a empty line
	}
}

