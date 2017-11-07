package Algorithm


func OddEven(array []int)  ([]int) {
	i := 0
	j := len(array)-1

	for i<j{
		if array[i]%2 ==1 {
			temp := array[i]
			array[i] = array[j]
			array[j] = temp
			j--
		}else {
			i++
		}
	}

	return  array
}


