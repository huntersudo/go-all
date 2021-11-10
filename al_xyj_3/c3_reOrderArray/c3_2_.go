package main

// 根据指定闭包对数组切片排序
func reOrderArrayV2(arr []int, orderFunc func(int) bool) []int {
	// 小于等于1个元素无需处理
	if arr == nil || len(arr) <= 1 {
		return arr
	}
	// 设置两个指针，从头尾往中间移动
	i := 0
	j := len(arr) - 1
	// 头指针不能越过尾指针，否则退出
	// 以奇偶数排序为例，i 从左到右寻找偶数，j 从右到左寻找奇数
	// 该循环执行完毕后，i 左侧的都是奇数，j 右侧的都是偶数，也就完成了顺序调整
	for i < j {
		// 如果不符合条件，则头指针后移，否则中断
		// 以 orderFunc 为偶数判断函数为例，返回 false 表示是奇数
		// 题目要求奇数排在前面，因此，当 i 对应值是奇数时，往后移一位，然后继续下一个循环，直到 i==j 或者遇到第一个偶数中断
		for i < j && !orderFunc(arr[i]) {
			i++
		}
		// 如果符合条件，则尾指针前移，否则中断
		// 还是以 orderFunc 为偶数判断函数为例，返回 true 表示是偶数
		// 题目要求偶数排在后面，因此，当 j 对应值是偶数时，往前移一位，然后继续下一个循环，直到 j==i 或者遇到第一个奇数中断
		for i < j && orderFunc(arr[j]) {
			j--
		}
		// 如果 i < j，则交换对应值的位置
		// 以奇偶数为例，此时 arr[i] 是偶数，arr[j] 是奇数，则交换两个值，将奇数放到前面，偶数放到后面
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
		// 继续下一个循环，直到 i==j，此时 i 左侧都是奇数，j 右侧都是偶数，所有奇数都排到了偶数前面
	}
	return arr
}

// 排序条件：是否是偶数
func isEven(num int) bool {
	return num & 1 == 0
}
