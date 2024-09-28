package main

import (
	"fmt"
)

type student struct {
	id    int
	name  string
	class string
}

// 学员管理类型的结构体
type studentMer struct {
	alStudent []*student
}

// 初始化切片的操作
func newStudentMer() *studentMer {
	return &studentMer{
		alStudent: make([]*student, 0, 100),
	}
}

// 0.展示系统菜单
func showMenu() {
	fmt.Println("欢迎进入...")
	fmt.Println("1.添加学员")
	fmt.Println("2.编辑学员信息")
	fmt.Println("3.展示所有学员信息")
	fmt.Println("4.退出...")
}

// 1.添加学员
func (s *studentMer) addStudent(newStu *student) {
	s.alStudent = append(s.alStudent, newStu)
	// fmt.Printf("已添加学员：ID=%d, 姓名=%s, 班级=%s\n", s.alStudent[0].id, s.alStudent[0].name, s.alStudent[0].class)
	// fmt.Printf("当前学生列表长度: %d\n", len(s.alStudent))
}

// 2.编辑学生
func (s *studentMer) editStudent(newStu *student) {
	for i, v := range s.alStudent {
		if newStu.id == v.id {
			s.alStudent[i] = newStu
			return
		}
	}
	fmt.Println("输入学生信息有误...")
}

// 3.展示学生
func (s *studentMer) showStudent() {
	fmt.Println("------------学生信息表------------")
	// fmt.Println("当前学生信息：", s.alStudent) // 打印切片内容
	if len(s.alStudent) == 0 {
		fmt.Println("没有学员信息可展示...")
		return
	}
	for _, v := range s.alStudent {
		fmt.Printf("%d\t%s\t%s\n", v.id, v.name, v.class)
	}
}
func newStudent(id int, name, class string) *student {
	return &student{
		id:    id,
		name:  name,
		class: class,
	}
}
