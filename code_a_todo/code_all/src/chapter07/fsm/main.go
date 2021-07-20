package main

import (
	"fmt"
)

/**
todo
有限状态机又简称 FSM（Finite-State Machine 的首字母缩写），也可以称为有限状态自动机。
它是为研究有限内存的计算过程和某些语言类而抽象出的一种计算模型。
有限状态机拥有有限数量的状态，每个状态可以迁移到零个或多个状态，输入字串决定执行哪个状态的迁移。
 */
// 闲置状态
type IdleState struct {
	StateInfo // 使用StateInfo实现基础接口
}

// 重新实现状态开始
func (i *IdleState) OnBegin() {
	fmt.Println("IdleState begin")
}

// 重新实现状态结束
func (i *IdleState) OnEnd() {
	fmt.Println("IdleState end")
}

// 移动状态
type MoveState struct {
	StateInfo
}

func (m *MoveState) OnBegin() {
	fmt.Println("MoveState begin")
}

// 允许移动状态互相转换
func (m *MoveState) EnableSameTransit() bool {
	return true
}

// 跳跃状态
type JumpState struct {
	StateInfo
}

func (j *JumpState) OnBegin() {
	fmt.Println("JumpState begin")
}

// 跳跃状态不能转移到移动状态
func (j *JumpState) CanTransitTo(name string) bool {
	return name != "MoveState"
}

// 封装转移状态和输出日志
func transitAndReport(sm *StateManager, target string) {
	if err := sm.Transit(target); err != nil {
		fmt.Printf("FAILED! %s --> %s, %s\n\n", sm.CurrState().Name(), target, err.Error())
	}
}

func main() {
	// 实例化一个状态管理器
	sm := NewStateManager()

	// 响应状态转移的通知
	sm.OnChange = func(from, to State) {

		// 打印状态转移的流向
		fmt.Printf("%s ---> %s\n\n", StateName(from), StateName(to))
	}

	// 添加3个状态
	sm.Add(new(IdleState))
	sm.Add(new(MoveState))
	sm.Add(new(JumpState))

	// 在不同状态间转移
	transitAndReport(sm, "IdleState")

	transitAndReport(sm, "MoveState")

	transitAndReport(sm, "MoveState")

	transitAndReport(sm, "JumpState")

	transitAndReport(sm, "JumpState")

	transitAndReport(sm, "IdleState")
}
