package algorythms

func LinearSearch(numberList []int, value int) int {
	for i := 0; i <= len(numberList); i++ {
		if numberList[i] == value {
			return i
		}
	}
	return -1
}
