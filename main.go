package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Printf("error opening input file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var leftCol []int
	var rightCol []int
	var line int

	for scanner.Scan() {
		line++

		leftNum, err := strconv.Atoi(strings.Split(scanner.Text(), "   ")[0])
		if err != nil {
			log.Fatalf("error parsing left number on line %d\n", line)
		}

		rightNum, err := strconv.Atoi(strings.Split(scanner.Text(), "   ")[1])
		if err != nil {
			log.Fatalf("error parsing right number on line %d\n", line)
		}

		leftCol = append(leftCol, leftNum)
		rightCol = append(rightCol, rightNum)
	}

	totalDistance, err := getDistance(leftCol, rightCol)
	if err != nil {
		log.Print("error getting total distance: %+v", err)
	}
	log.Printf("total distance: %d\n", totalDistance)

	totalSimilarity, err := getSimilarity(leftCol, rightCol)
	if err != nil {
		log.Printf("error getting similarity: %+v\n", err)
	}
	log.Printf("total similarity: %d\n", totalSimilarity)
}

func getDistance(leftCol []int, rightCol []int) (int, error) {
	var distances []int
	totalDistance := 0

	sort.Slice(leftCol, func(i, j int) bool {
		return leftCol[j] > leftCol[i]
	})

	sort.Slice(rightCol, func(i, j int) bool {
		return rightCol[j] > rightCol[i]
	})

	for i := 0; i < len(leftCol); i++ {
		distances = append(distances, int(math.Abs(float64(leftCol[i])-float64(rightCol[i]))))
	}

	for _, distance := range distances {
		totalDistance += distance
	}

	log.Printf("length of left column: %d", len(leftCol))
	log.Printf("length of right column: %d", len(rightCol))
	log.Printf("length of distances: %d", len(distances))

	return totalDistance, nil
}

func getSimilarity(leftCol []int, rightCol []int) (int, error) {
	var similarities []int
	totalSimilarity := 0

	for _, leftColNumber := range leftCol {
		numOccurrences := 0
		for _, rightColNumber := range rightCol {
			if leftColNumber == rightColNumber {
				numOccurrences++
			}
		}

		similarities = append(similarities, leftColNumber*numOccurrences)
	}

	log.Printf("similarities: %+v\n", similarities)

	for _, similarity := range similarities {
		totalSimilarity += similarity
	}

	return totalSimilarity, nil
}
