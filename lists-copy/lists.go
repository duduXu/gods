// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package lists provides an abstract List interface.
//
// In computer science, a list or sequence is an abstract data type that represents an ordered sequence of values, where the same value may occur more than once. An instance of a list is a computer representation of the mathematical concept of a finite sequence; the (potentially) infinite analog of a list is a stream.  Lists are a basic example of containers, as they contain other values. If the same value occurs multiple times, each occurrence is considered a distinct item.
//
// Reference: https://en.wikipedia.org/wiki/List_%28abstract_data_type%29
package arraylist_copy

type List struct {
	element []interface{}
	size int
}

// 初始化
func New() *List {
	list := &List{}

	return list
}

// 添加元素
func (list *List) Add(values ...interface{}) {
	for _, value := range values {
		list.element[list.size] = value
		list.size ++
	}
}

// 获取元素
func (list *List) Get(index int) (interface{}, bool) {
	if ! list.WithInRange(index) {
		return nil, false
	}

	return list.element[index], true
}

// 移除元素
func (list *List) Remove(index int) {
	if ! list.WithInRange(index) {
		return
	}

	copy(list.element[index:], list.element[index+1:list.size])
	list.size--
}

// 判断是否包含指定元素
func (list *List) Contains(values ... interface{}) bool {
	return true
}

// 判断某元素的下标
func (list *List) IndexOf(value interface{}) (int) {
	return 0
}

// 交换两个元素
func (list *List) Swap (i, j int) {

}

// 在某个位置插入一个或多个元素
func (list *List) Insert (index int, values ...interface{})  {

}

// 设置某个下标的对应元素
func (list *List) Set(index int, value interface{}) {

}

// 判断
func (list *List) WithInRange(index int) bool {
	return index >=0 && list.size < index
}