package main

import (
	"fmt"
	"os"
)

func main() {
	var A, B, C H
	var N Hs
	A.Init()
	N.Add(&A)
	B.Init()
	N.Add(&B)
	C.Init()
	N.Add(&C)
	fmt.Println("欢迎进入游戏...")
	var n int
	var h *H
	var h1 *H
	for {
		ShowMenu1()
		fmt.Scan(&n)
		switch n {
		case 1:
			ShowHeros(&N)
		case 2:
			// xuanze
			var yourhero string
			fmt.Println("请输入你要选择的英雄:")
			fmt.Scan(&yourhero)
			h = N.Gethero(yourhero)
			fmt.Printf("your hero is %s\n", h.name)
		case 3:
			ShowMenu2()
			fmt.Println(":")
			fmt.Scan(&n)
			switch n {
			case 1:
				// 攻击
				var rival string
				fmt.Println("请输入你要攻击的英雄:")
				fmt.Scan(&rival)
				h1 = N.Gethero(rival)
				fmt.Printf("your rival is %s\n它的hp:%d", h1.name, h1.hp)
				h.Attack(h1, 1)
				fmt.Printf("your rival is %s\n它被attack后的hp:%d", h1.name, h1.hp)
			case 2:
				// 回血

			case 3:
				// 回蓝

			}
		case 4:
			os.Exit(0)
		}
	}
}
