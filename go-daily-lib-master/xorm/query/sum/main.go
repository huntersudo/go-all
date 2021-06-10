package main

import (
	"fmt"
	"math/rand"

	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

type Sum struct {
	Id    int64
	Money int32
	Rate  float32
}

const ConStr ="sml_user:1qaz@WSX@tcp(192.168.88.132:3306)/sml_db?charset=utf8"

func main() {
	engine, _ := xorm.NewEngine("mysql", ConStr)
	engine.Sync2(&Sum{})

	//Sum/SumInt ：求某个字段的和， Sum 返回 float64 ， SumInt 返回 int64 ；
	//Sums/SumsInt ：分别求某些字段的和， Sums 返回 []float64 ， SumsInt 返回 []int64 。
	var slice []*Sum
	for i := 0; i < 10; i++ {
		slice = append(slice, &Sum{
			Money: rand.Int31n(10000),
			Rate:  rand.Float32(),
		})
	}
	engine.Insert(&slice)

	totalMoney, _ := engine.SumInt(&Sum{}, "money")
	fmt.Println("total money:", totalMoney)

	totalRate, _ := engine.Sum(&Sum{}, "rate")
	fmt.Println("total rate:", totalRate)

	// sums
	totals, _ := engine.Sums(&Sum{}, "money", "rate")
	fmt.Printf("total money:%f & total rate:%f", totals[0], totals[1])

	//total money: 34922
	//total rate: 4.697319328784943
	//total money:34922.000000 & total rate:4.697319

	//mysql> select * from sum;
	//	+----+-------+----------+
	//	| id | money | rate     |
	//		+----+-------+----------+
	//	|  1 |  8081 | 0.940509 |
	//	|  2 |  1847 | 0.437714 |
	//	|  3 |  2081 | 0.686823 |
	//	|  4 |  4425 | 0.156519 |
	//	|  5 |   456 | 0.300912 |
	//	|  6 |   694 |  0.81364 |
	//	|  7 |  8162 | 0.380657 |
	//	|  8 |  4728 |  0.46889 |
	//	|  9 |  1211 | 0.293102 |
	//	| 10 |  3237 | 0.218553 |
	//		+----+-------+----------+
	//		10 rows in set (0.00 sec)



	}
