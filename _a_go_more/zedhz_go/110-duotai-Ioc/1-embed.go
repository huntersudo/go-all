package main

// 结构体嵌入

type Widget struct {
	X, Y int
}
type Label struct {
	Widget        // Embedding (delegation)
	Text   string // Aggregation
}

func use1()  {
	label := Label{Widget{10, 10}, "State:"}

	label.X = 11
	label.Y = 12
}
//如果成员 X 重名，我们就要用 label.X表明是自己的X ，用 label.Wedget.X 表明是嵌入过来的




