package Algorithm

import "fmt"

func QuickSort(values []int, left int, right int) {
	if left < right {
		// 设置基准值
		temp := values[left]
		// 设置哨兵
		i, j := left, right
		for {
			// 从右向左找，找到第一个比基准值小的数
			for values[j] >= temp && i < j {
				j--
			}
			// 从左向右找，找到第一个比基准值大的数
			for values[i] <= temp && i < j {
				i++
			}
			// 如果哨兵相遇，则退出循环
			if i >= j {
				break
			}
			// 交换左右两侧的值
			values[i], values[j] = values[j], values[i]
		}
		// 将基准值移到哨兵相遇点
		values[left] = values[i]
		values[i] = temp
		// 递归，左右两侧分别排序
		QuickSort(values, left, i-1)
		QuickSort(values, i+1, right)
	}
}

func GroupArray (arr01 []int,arr02 []int) ([]int){
	var arr03 =make([]int,len(arr01)+len(arr02))
	for i,v := range arr01 {
		arr03[i]=v
	}
	for i1,v1 := range  arr02 {
		arr03[len(arr01)+i1]=v1
	}
	fmt.Println(arr03)
	QuickSort(arr03,0,len(arr03)-1)
	return  arr03
}