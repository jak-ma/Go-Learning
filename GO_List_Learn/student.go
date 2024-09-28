package main

import "fmt"

type Book struct {
	isbn  string
	name  string
	price float64
}

type LNode struct {
	elem Book
	next *LNode
}

type List *LNode

// 初始化链表
func InitList(L *List) {
	*L = new(LNode)
	(*L).next = nil
}

// 前插法创建链表
func CreatList(L *List, n int) {
	for i := 0; i < n; i++ {
		p := new(LNode)
		fmt.Println("input isbn:")
		fmt.Scan(&p.elem.isbn)
		fmt.Println("input name:")
		fmt.Scan(&p.elem.name)
		fmt.Println("input price:")
		fmt.Scan(&p.elem.price)
		p.next = (*L).next
		(*L).next = p
	}
}

// 遍历链表
func Pt(L List) {
	if L.next == nil {
		fmt.Println("empty list...")
		return
	}
	p := new(LNode)
	p = L.next
	for p != nil {
		fmt.Printf("isbn:%s\tname:%s\tprice:%f\n", p.elem.isbn, p.elem.name, p.elem.price)
		p = p.next
	}
}

// 插入节点
func insert(L *List, b Book, i int) {
	p := new(LNode)
	newp := new(LNode)
	p = *L
	j := 1
	for j < i && p != nil {
		j++
		p = p.next
	}
	newp.elem = b
	newp.next = p.next
	p.next = newp
}

// 删除节点
func delete(L *List, isbn string) {
	p := new(LNode)
	p = (*L)
	for p.next != nil {
		if p.next.elem.isbn == isbn {
			p.next = p.next.next
			return
		} else {
			p = p.next
		}
	}
	fmt.Println("输入的isbn号查找不到...")
}

// 取值下标
func getindex(L List, isbn string) Book {
	p := new(LNode)
	p = L.next
	for p != nil {
		if p.elem.isbn == isbn {
			return (p.elem)
		} else {
			p = p.next
		}
	}
	fmt.Println("isbn没有找到...")
	return Book{}
}
