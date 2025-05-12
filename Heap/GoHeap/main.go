package main

import "fmt"

type MaxHeap struct {
	arr []int
}

func NewMaxHeap(len int) *MaxHeap {
	return &MaxHeap{
		arr: make([]int, 1, len+1),
	}
}

func (mh *MaxHeap) push(item int) {
	if len(mh.arr) == 1 {
		mh.arr = append(mh.arr, item)
		return
	}

	mh.arr = append(mh.arr, item)

	currIdx := len(mh.arr) - 1

	for currIdx > 1 {
		currIdxParent := currIdx / 2

		if mh.arr[currIdx] > mh.arr[currIdxParent] {
			tmp := mh.arr[currIdx]
			mh.arr[currIdx] = mh.arr[currIdxParent]
			mh.arr[currIdxParent] = tmp
		}
		currIdx = currIdxParent
	}
}

func (mh *MaxHeap) pop() int {
	if len(mh.arr) == 1 {
		return 0
	}
	if len(mh.arr) == 2 {
		popItem := mh.arr[1]
		mh.arr = mh.arr[:1]
		return popItem
	}

	popItem := mh.arr[1]
	mh.arr[1] = mh.arr[len(mh.arr)-1]
	mh.arr = mh.arr[:len(mh.arr)-1]

	currIdx := 1

	for currIdx < len(mh.arr) {
		leftChild := currIdx * 2
		rightChild := (currIdx * 2) + 1

		if leftChild < len(mh.arr) && mh.arr[leftChild] > mh.arr[currIdx] {
			tmp := mh.arr[leftChild]
			mh.arr[leftChild] = mh.arr[currIdx]
			mh.arr[currIdx] = tmp

			currIdx = leftChild
		} else if rightChild < len(mh.arr) && mh.arr[rightChild] > mh.arr[currIdx] {
			tmp := mh.arr[rightChild]
			mh.arr[rightChild] = mh.arr[currIdx]
			mh.arr[currIdx] = tmp

			currIdx = rightChild
		} else {
			break
		}
	}

	return popItem

}

func main() {
	heap := NewMaxHeap(10)

	heap.push(3)
	fmt.Println(heap.arr)

	heap.push(5)
	fmt.Println(heap.arr)

	heap.push(2)
	fmt.Println(heap.arr)

	heap.push(55)
	fmt.Println(heap.arr)

	heap.push(6)
	fmt.Println(heap.arr)

	heap.push(1)
	fmt.Println(heap.arr)

	item := heap.pop()
	fmt.Println(heap.arr, item)

	item = heap.pop()
	fmt.Println(heap.arr, item)

	item = heap.pop()
	fmt.Println(heap.arr, item)

	item = heap.pop()
	fmt.Println(heap.arr, item)

	item = heap.pop()
	fmt.Println(heap.arr, item)

	item = heap.pop()
	fmt.Println(heap.arr, item)

	item = heap.pop()
	fmt.Println(heap.arr, item)
}
