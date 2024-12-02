package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func LoadData(path string) ([]int, []int) {

	file, err := os.Open(path)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sliceOne []int
	var sliceTwo []int

	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Fields(line)
		firstNum, _ := strconv.Atoi(tokens[0])
		secondNum, _ := strconv.Atoi(tokens[1])
		sliceOne = append(sliceOne, firstNum)
		sliceTwo = append(sliceTwo, secondNum)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return sliceOne, sliceTwo
}

func FindDistance(firstSlice []int, secondSlice []int) int {
	sum := 0
	sort.Slice(firstSlice, func(i, j int) bool {
		return firstSlice[i] < firstSlice[j]
	})
	sort.Slice(secondSlice, func(i, j int) bool {
		return secondSlice[i] < secondSlice[j]
	})
	for idx := range firstSlice {
		distance := firstSlice[idx] - secondSlice[idx]
		if distance < 0 {
			distance = -1 * distance
		}
		sum += distance
	}
	return sum
}

func CountOccurences(slice []int) map[int]int {
	dict := make(map[int]int)
	for _, num := range slice {
		dict[num] = dict[num] + 1
	}
	return dict
}

func FindSimilarity(firstSlice []int, secondSlice []int) int {
	similarity := 0
	occurences := CountOccurences(secondSlice)
	for _, num := range firstSlice {
		similarity += num * occurences[num]
	}
	return similarity
}

func main() {
	sliceOne, sliceTwo := LoadData("./day1_input.txt")
	distance := FindDistance(sliceOne, sliceTwo)
	similarity := FindSimilarity(sliceOne, sliceTwo)
	fmt.Printf("Sum of distances: %d\n", distance)
	fmt.Printf("Similarity score: %d\n", similarity)
}
