package main

import (
	"fmt"
	"os"
)

func input() *student {
	var (
		id    int
		name  string
		class string
	)
	fmt.Println("输入学生信息：")
	fmt.Println("输入学号：")
	fmt.Scan(&id)
	fmt.Println("输入姓名：")
	fmt.Scan(&name)
	fmt.Println("输入班级号：")
	fmt.Scan(&class)
	// fmt.Println("输入的学生信息：", id, name, class)
	stu1 := newStudent(id, name, class)
	return stu1
}

func getScanf(s *studentMer) {
	fmt.Println("请输入你要执行的动作的序号：")
	var n int
	fmt.Scan(&n)
	fmt.Printf("用户输入的是%d\n", n)
	switch n {
	case 1: //添加学员
		stu := input()
		s.addStudent(stu)
	case 2: //编辑学员
		stu := input()
		s.editStudent(stu)
	case 3: //展示学员信息
		s.showStudent()
	case 4: //退出...
		os.Exit(0)
	default:
		fmt.Println("输入序号非法！")
	}
	// fmt.Println("当前学生信息：", s.alStudent) // 打印切片内容
}

func main() {
	s1 := newStudentMer()
	for {
		// 1.打印系统菜单
		showMenu()
		// 2.等待用户输入
		getScanf(s1)
		// 3.执行用户选择的动作
	}
}
