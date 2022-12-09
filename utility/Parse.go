package utility

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func assembleFilePathToDay(day string) string {
	path, _ := os.Getwd()
	dayMunge := fmt.Sprintf("day%s", day)

	/*
		necessary to conveniently correct working directory issues present in goland vs terminal.
	*/
	if !strings.Contains(path, dayMunge) {
		path = filepath.Join(path, dayMunge)
	}
	return path
}

func ParseInputFileIntoStringRows(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data := make([]string, 0)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for i := 0; fileScanner.Scan(); i++ {
		aString := fileScanner.Text()
		data = append(data, aString)
	}
	file.Close()
	return data, nil
}

func parseInputFileIntoNumberRows(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data := make([]int, 0)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for i := 0; fileScanner.Scan(); i++ {
		aString := fileScanner.Text()
		aNumber, err := strconv.Atoi(aString)
		if err == nil {
			log.Fatalln(err)
		}
		data = append(data, aNumber)
	}
	file.Close()
	return data, nil
}

func ParseDayForInputIntoStringRows(day string, path string) ([]string, error) {
	return ParseInputFileIntoStringRows(filepath.Join(assembleFilePathToDay(day), path))
}

func ParseDayForInputIntoNumberRows(day string, path string) ([]int, error) {
	return parseInputFileIntoNumberRows(filepath.Join(assembleFilePathToDay(day), path))
}

func ParseRowsToInts(data []string) [][]int {

	array := make([][]int, len(data))

	for y, row := range data {
		anArrayOfNumbers := make([]int, len(row))
		strings := strings.Split(row, "")

		for x, aString := range strings {
			aNumber, _ := strconv.Atoi(aString)
			anArrayOfNumbers[x] = aNumber
		}

		array[y] = anArrayOfNumbers
	}
	return array
}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func IsLower(s string) bool {
	for _, r := range s {
		if !unicode.IsLower(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func IntersectionOfTwoStrings(a string, b string) []rune {

	return IntersectingCharaactersOfStrings([]string{a, b})
}

func IntersectingCharaactersOfStrings(arrayOfStrings []string) []rune {

	var commons []rune
	intersectionMap := make(map[rune]int)

	for _, aString := range arrayOfStrings {
		aChars := strings.Split(aString, "")
		sort.Strings(aChars)
		aChars = RemoveDuplicates(aChars)
		aRunes := []rune(strings.Join(aChars, ""))

		for _, aRune := range aRunes {
			intersectionMap[aRune]++
		}

	}

	for key, element := range intersectionMap {
		if element == len(arrayOfStrings) {
			commons = append(commons, key)
		}
	}
	return commons
}

type SliceType interface {
	~string | ~int | ~float64 // add more *comparable* types as needed
}

func RemoveDuplicates[T SliceType](s []T) []T {
	if len(s) < 1 {
		return s
	}

	// sort
	sort.SliceStable(s, func(i, j int) bool {
		return s[i] < s[j]
	})

	prev := 1
	for curr := 1; curr < len(s); curr++ {
		if s[curr-1] != s[curr] {
			s[prev] = s[curr]
			prev++
		}
	}

	return s[:prev]
}
