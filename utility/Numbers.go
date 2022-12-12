package utility

import (
	"log"
	"sort"
	"strconv"
	"unicode"
)

func Average(numbers []int) int {
	return Mean(numbers)
}

func Mean(numbers []int) int {
	sum := 0
	for _, aNumber := range numbers {
		sum += aNumber
	}
	result := sum / len(numbers)
	return result
}

func Median(numbers []int) float32 {
	median := float32(0)
	middle := len(numbers) / 2
	if (len(numbers) % 2) == 1 {
		median = float32(numbers[middle])
	} else {
		temp := numbers[middle] + numbers[middle-1]
		median = float32(float32(temp) / float32(2))
	}

	return median
}

func LeastAndMax(numbers []int) (int, int) {
	if numbers == nil {
		numbers = make([]int, 0)
	}

	least := numbers[0]
	max := numbers[0]

	for _, aNumber := range numbers {
		if aNumber < least {
			least = aNumber
		}
		if aNumber > max {
			max = aNumber
		}

	}
	return least, max
}

func OrderNumbersStartingWithAndEndingWith(numbers []int, start int, end int) []int {

	if numbers[0] == start && numbers[len(numbers)-1] == end {
		return numbers
	} else {

		sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
		return numbers
	}
	return numbers
}

func OrderNumbersSort(numbers []int) []int {

	sort.Sort(sort.IntSlice(numbers))
	return numbers

}

func OrderNumbersSortReversed(numbers []int) []int {

	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
	return numbers

}

func NumbersBetween(start int, end int) []int {
	least, max := LeastAndMax([]int{start, end})

	numbers := make([]int, 0)
	for x := least; x <= max; x++ {
		numbers = append(numbers, x)
	}
	return numbers
}

func StringToInt(aString string) int {
	aNumber, err := strconv.Atoi(aString)
	if err != nil {
		panic("Expected a number and got: " + aString)
	}
	return aNumber
}

func BytesToInt(byteArray []byte) int {
	/*
		lets convert to a string
	*/
	bytesAsString := string(byteArray)
	//for len(bytesAsString) < 8{
	//	bytesAsString = "0" + bytesAsString
	//}

	anInt, err := strconv.ParseInt(bytesAsString, 2, 64)

	if err != nil {
		log.Fatalln(err)
	}
	return int(anInt)
}

func SliceToInt(aSlice []int) int {
	result := 0
	power := 1
	for i := len(aSlice) - 1; i >= 0; i-- {
		result += aSlice[i] * power
		power *= 10
	}
	return result
}

func AssignNumbersToLetters(lowerCaseStart int, upperCaseStart int) map[rune]int {

	priorities := make(map[rune]int)
	index := lowerCaseStart
	for aLowerCaseLetter := 'a'; aLowerCaseLetter <= 'z'; aLowerCaseLetter++ {
		aCapitalLetter := unicode.ToUpper(aLowerCaseLetter)
		priorities[aLowerCaseLetter] = index
		priorities[aCapitalLetter] = upperCaseStart + index
		index++
	}

	return priorities

}

func SubSetPresent(a []int, b []int) bool {
	aLeast, aMax := LeastAndMax(a)
	bLeast, bMax := LeastAndMax(b)

	if bLeast >= aLeast && bMax <= aMax {
		return true
	} else if aLeast >= bLeast && aMax <= bMax {
		return true
	}

	return false

}

func IntersectionOfNumbers(numberArrays [][]int) []int {

	var commons []int
	intersectionMap := make(map[int]int)

	outerLength := len(numberArrays)

	for i := 0; i < outerLength; i++ {
		for j := 0; j < len(numberArrays[i]); j++ {
			aNumber := numberArrays[i][j]
			intersectionMap[aNumber]++
		}
	}

	for key, element := range intersectionMap {
		if element == outerLength {
			commons = append(commons, key)
		}
	}
	return commons
}
