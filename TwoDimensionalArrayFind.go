package leetcode

import "fmt"

func find(array [][]int, int num) [2]int{
	
	if (len(array) == 0){
		return [2]int{-1, -1}
	}
	if (len(array[0] == 0)){
		return [2]int{-1, -1}
	}
	length := len(array[0])
	height := len(array)
	i := length - 1
	j := 0
	for ;i >= 0; i--{
		if array[j][i] >  num{
			continue
		}else if array[j][i] < num{
			for ;j < height; j++{
				if array[j][i] > num{
					break
				}else if array[j][i] < num{
					continue
				}else{
					return [2]int{j, i}
				}
			}
		}else{
			return [2]int{j, i}
		}
	}
}

func FindNum(array [][]int, num int)ret bool{
	result := find(array, num)
	if result[0] != -1 && result [1] != -1
	{
		ret = true
	}else{
		ret = false
	}
}