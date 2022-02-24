package algorythms

func InsertionSort(numberList []int) {
	for keyIndex := 2; keyIndex < len(numberList); keyIndex++ {
		keyElement := numberList[keyIndex]
		previousIndex := keyIndex - 1
		for previousIndex >= 0 && numberList[previousIndex] > keyElement {
			numberList[previousIndex+1] = numberList[previousIndex]
			previousIndex -= 1
		}
		numberList[previousIndex+1] = keyElement
	}
}

func InsertionSortReverse(numberList []int) {
	for keyIndex := 2; keyIndex < len(numberList); keyIndex++ {
		keyElement := numberList[keyIndex]
		previousIndex := keyIndex - 1
		for previousIndex >= 0 && numberList[previousIndex] < keyElement {
			numberList[previousIndex+1] = numberList[previousIndex]
			previousIndex -= 1
		}
		numberList[previousIndex+1] = keyElement
	}
}
