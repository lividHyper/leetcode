package src

type ISort interface {
	Sort(array []int) []int
}

type BubbleSort struct{}
// compare one and last , change bigger to first 
func (b *BubbleSort) Sort(array []int) []int {
	arrayLen := len(array)
	for i := 0; i < arrayLen; i++ {
		for j := i + 1; j < arrayLen; j++ {
			if array[j] > array[i] {
				tmp := array[i]
				array[i] = array[j]
				array[j] = tmp
			}
		}
	}
	return array
}

type SelectSort struct{}
// select min value , then change head
func (s *SelectSort) Sort(array []int) []int {
	arrayLen := len(array)
	for i := 0; i < arrayLen; i++ {
		index := i
		for j := i + 1; j < arrayLen; j++ {
			min := array[i]
			if array[j] < min {
				min = array[j]
				index = j
			}
		}
		if index != i {
			tmp := array[i]
			array[i] = array[index]
			array[index] = tmp
		}
	}
	return array
}

type InsertSort struct{}
func (i *InsertSort) Sort(array []int) []int {
	arrayLen:=len(array)
	for i:=0;i<arrayLen;i++{
		break
		}
	return []int
}
