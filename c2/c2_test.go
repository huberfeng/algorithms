package c2

import (
	"testing"
	"fmt"
)

func TestSort_Slection(t *testing.T) {
	a := []Comparable{Int(1),Int(8),Int(7),Int(4),Int(5)}
    a[0].CompareTo(a[1])
	sort := new(Sort)
	sort.Slection(a)
	fmt.Println(a)
}

func TestSort_Insert(t *testing.T) {
	a := []Comparable{Int(1),Int(8),Int(7),Int(4),Int(5)}
	sort := new(Sort)
	sort.Insert(a)
	fmt.Println(a)
}

func TestSort_Shellsort(t *testing.T) {
	a := []Comparable{Int(1),Int(8),Int(7),Int(4),Int(5)}
	sort := new(Sort)
	sort.Shellsort(a)
	fmt.Println(a)
}

func TestSort_Merge(t *testing.T) {
	a := []Comparable{Int(1),Int(3),Int(7),Int(4),Int(5)}
	sort := new(Sort)
	sort.Merge(a,0,2,4)
	fmt.Println(a)
}

func TestSort_Mergesort(t *testing.T) {
	a := []Comparable{Int(1),Int(3),Int(7),Int(4),Int(5)}
	sort := new(Sort)
	sort.Mergesort(a,0,4)
	fmt.Println(a)
}

//归并排序的算法有问题, 怕不是个假的归并
func TestPerformance(t *testing.T) {
	Performance("selection", 10000, 1)
	Performance("insert", 10000, 1)
	Performance("shell", 10000, 1)
	Performance("merge", 10000,1)
	Performance("MergeSortDownToUp", 10000,1)
}


func TestMerge1(t *testing.T) {
	a := []int{1,3,5}
	b := []int{2,4}
	c := Merge1(a, b)
	fmt.Println(c)
}



