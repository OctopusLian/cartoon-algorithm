/*
 * @Author: your name
 * @Date: 2021-06-29 11:38:56
 * @LastEditTime: 2021-06-29 11:56:44
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \cartoon-algorithm\src\chapter2\part1\MyArray.go
 */
package main

import "fmt"

func insert(src []int, data int) {
	src = append(src, data)
	fmt.Println("插入后:", src)
}

func delete(src []int, index int) {
	src = append(src[:index], src[index+1:]...)
	fmt.Println("删除后：", src)
}

func output(src []int) {
	fmt.Println(src)
}

func main() {
	init := []int{1, 2, 3, 4, 5, 6, 7}
	insert(init, 9)
	delete(init, 3)
	output(init)
}

/*
插入后: [1 2 3 4 5 6 7 9]
删除后： [1 2 3 5 6 7]
[1 2 3 5 6 7 7]
*/
