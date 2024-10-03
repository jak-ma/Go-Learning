package main

import "fmt"

// 定义英雄属性
type H struct {
	hp         int
	attacktion int
	cd         int
	tools      int
	name       string
}

// 初始化英雄属性
func (h *H) Init() {
	h.attacktion = 5
	h.cd = 4
	h.hp = 10
	h.tools = 5
	fmt.Println("Input the hero's name :")
	fmt.Scan(&h.name)
	fmt.Println("初始化完成...")
}

// 定义英雄池
type Hs struct {
	hs     []H
	hsnums int
}

// 增加英雄
func (s *Hs) Add(h *H) {
	s.hs = append(s.hs, *h)
	s.hsnums++
}

// 显示英雄列表
func ShowHeros(n *Hs) {
	for _, v := range n.hs {
		fmt.Println(v.name)
	}
}

// 根据name选择英雄
func (s *Hs) Gethero(n string) *H {
	for _, v := range s.hs {
		if v.name == n {
			return &v
		}
	}
	return &H{} // 解决空接口匹配的问题
}

// 定义英雄接口
type hero interface {
	Attack(rival H)
	UseSkills()
	UseTools()
}

// 攻击实现
func (h H) Attack(rival *H, n int) {
	if n == 1 {
		h.Useskills()
		rival.hp--
	} else {
		h.UseTools()
		rival.hp -= 2
	}
}

// 使用技能
func (h H) Useskills() {
	h.cd--
}

// 使用道具
func (h H) UseTools() {
	h.attacktion++
	h.tools--
}

func ShowMenu1() {
	fmt.Println("请输入你的命令:\n",
		"1.获取英雄列表\n",
		"2.选择你的英雄\n",
		"3.选择你要进行的操作\n",
		"4.退出游戏\n")
}
func ShowMenu2() {
	fmt.Println("1.攻击英雄\n",
		"2.回复血量\n",
		"3.回复蓝量\n")
}
