package day1

import (
	"bufio"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"path/filepath"
)

func readTwoLists(filepath string) ([]int, []int) {
	file, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var list1, list2 []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers := strings.Fields(scanner.Text())

		num1, err1 := strconv.Atoi(numbers[0])
		num2, err2 := strconv.Atoi(numbers[1])

		if err1 != nil || err2 != nil {
			panic("Error converting string to int")
		}

		list1 = append(list1, int(num1))
		list2 = append(list2, int(num2))
	}

	return list1, list2
}

func getInputPath() string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("Failed to get current file path")
	}
	return filepath.Join(filepath.Dir(filename), "input.txt")
}

func Day1Part1() int {
	list1, list2 := readTwoLists(getInputPath())

	if len(list1) != len(list2) {
		panic("Lists are not the same length")
	}

	// Sort and calculate the total distance between lists
	sort.Ints(list1)
	sort.Ints(list2)

	totalDistance := 0

	for i := 0; i < len(list1); i++ {
		distance := list1[i] - list2[i]
		if distance < 0 {
			distance = -distance
		}

		totalDistance += distance
	}

	return totalDistance
}

func Day1Part2() int {
	list1, list2 := readTwoLists(getInputPath())

	// Pre-compute frequency map for list2
	freqMap := make(map[int]int)
	for _, num := range list2 {
		freqMap[num]++
	}

	// Split the work into multiple goroutines
	resultChan := make(chan int, len(list1))
	numWorkers := runtime.GOMAXPROCS(0)
	batchSize := (len(list1) + numWorkers - 1) / numWorkers

	for w := 0; w < numWorkers; w++ {
		start := w * batchSize
		end := start + batchSize
		if end > len(list1) {
			end = len(list1)
		}

		go func(start, end int) {
			partialScore := 0
			for i := start; i < end; i++ {
				num := list1[i]
				count := freqMap[num]
				partialScore += num * count
			}
			resultChan <- partialScore
		}(start, end)
	}

	// Collect results from all workers
	similarityScore := 0
	for i := 0; i < numWorkers; i++ {
		similarityScore += <-resultChan
	}

	return similarityScore
}
