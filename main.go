package main

import (
	"fmt"
	"zgolang/Algorithm"
	"zgolang/Tgo"
)

/*
type human struct{
	Sex int
}
type person struct{
	human
	Name string
	Age int
}
*/

func main() {

	fmt.Println("奇偶数排序")
	array := []int{12, 3, 45, 12, 44, 56, 89, 97, 87, 66, 57, 85, 63, 69, 27, 54, 54, 36, 83, 346, 34534, 15, 75645, 234235, 67845, 23446, 2346, 13, 3, 7, 2346, 24}
	fmt.Println(array)
	array =Algorithm.OddEven(array)
	fmt.Println(array)


	//*************
	fmt.Println("两个数组合并")
	arr01 :=[]int{1,3,4,6,8}
	arr02 :=[]int{2,5,7,0,9}

	arr03 := Algorithm.GroupArray(arr01,arr02)

	fmt.Println(arr03)

	fmt.Println(Tgo.Add(2,3))

	//快速排序
	fmt.Println("快速排序~")
	Algorithm.QuickSort(array,0,len(array)-1)
	fmt.Println(array)

	str_person := Tgo.Person{
		Name:"zhangsan",
		Age:32,
		Human:Tgo.Human{
			1,
		},

	}
	fmt.Println(str_person)

}
