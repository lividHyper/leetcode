package test

import (
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPrint(t *testing.T) {
	array := BuildArrayByNum(100)

	assert.Equal(t, len(array), 100)

	sort.Ints(array)
	assert.True(t, CheckArrayAsc(array))

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
