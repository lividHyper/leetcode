package main

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	"example.com/src"

	"github.com/stretchr/testify/assert"
)

func TestOrderBy(t *testing.T) {
	array := BuildArrayByNum(100)
	assert.Equal(t, len(array), 100)

	sort.Ints(array)
	assert.True(t, CheckArrayAsc(array))
}

func TestBubbleSort(t *testing.T) {
	CreateTestWithHundered(t, new(src.BubbleSort))
}

func TestSelectSort(t *testing.T) {
	CreateTestWithHundered(t, new(src.SelectSort))
	assert.True(t)
}

func CreateTestWithHundered(t *testing.T, iSort src.ISort) {
	for i := 0; i < 1000; i++ {
		CompareAndJudgeFunction(t, i, iSort)
	}
}

func CompareAndJudgeFunction(t *testing.T, number int, iSort src.ISort) {
	array := BuildArrayByNum(number)

	var standardArray []int
	var sortArray []int
	copy(array, standardArray)
	copy(array, sortArray)

	sort.Ints(standardArray)
	iSort.Sort(array)

	assert.EqualValues(t, standardArray, sortArray)
}

func BuildArrayByNum(number int) []int {
	var array []int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < number; i++ {
		array = append(array, rand.Intn(65534))
	}
	return array
}

func CheckArrayAsc(array []int) bool {
	if len(array) <= 1 {
		return true
	}
	for i := 0; i < len(array)-1; i++ {
		if array[i] > array[i+1] {
			return false
		}
	}
	return true
}
