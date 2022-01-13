package main

import "fmt"

// Bitcoin represents a number of Bitcoins.
// Go 允许从现有的类型创建新的类型。
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet stores the number of Bitcoin someone owns.
type Wallet struct {
	balance Bitcoin
}
// 接收者类型是 *Wallet 而不是 Wallet ，

// Deposit will add some Bitcoin to a wallet.
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance returns the number of Bitcoin a wallet has.
// 余额
// 记住，我们可以使用「receiver」变量访问结构体内部的 balance 字段。
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
