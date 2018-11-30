package main

import (
	"fmt"
)

func selectionSort(list []int) []int {
	len := len(list)

	for i := 0; i < len-1; i++ {
		max := i
		for j := i + 1; j < len; j++ {
			if list[max] < list[j] {
				max = j
			}
		}

		//交换
		tmp := list[i]
		list[i] = list[max]
		list[max] = tmp
		fmt.Println(list[max], list[i])
	}

	return list
}

func main() {
	//选择排序Selection sort
	//首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置，然后，再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。以此类推，直到所有元素均排序完毕。
	//选择排序的主要优点与数据移动有关。如果某个元素位于正确的最终位置上，则它不会被移动。
	//最坏时间复杂度	О(n²)
	//最优时间复杂度	О(n²)
	//平均时间复杂度	О(n²)

	list := []int{1, 3, 8, 3, 5}

	list = selectionSort(list)
	fmt.Println(list)
}
