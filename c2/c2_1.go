package c2

import (
	"fmt"
	"time"
	"math/rand"
	)

//实现了 Comparable 接口的数据类型，可以使用sort类的排序算法
type Comparable interface {
	CompareTo(p interface{}) bool
}

func Less(w, v Comparable) bool {
	return w.CompareTo(v)
}

func Exch(a []Comparable, i, j int) {
	a[i], a[j] = a[j], a[i]
}

func IsSorted(a []Comparable) bool {
	for i:= 1; i < len(a); i++ {
		if Less(a[i],a[i-1]) {
			return false
		}
	}
	return true
}

type Int int

func (s Int) CompareTo(p interface{}) bool {
	//i := reflect.ValueOf(p).Interface().(int)
	//fmt.Println(i)

	i := p.(Int)
	return s < i
}


type Sort struct {
}

//选择排序
func (s *Sort) Slection(a []Comparable) {
	n := len(a)
	for i := 0; i < n; i++ {
		min := i
		for j := i + 1; j < n; j++ {
			if Less(a[j], a[min]) {
				min = j
			}
		}
		Exch(a, i, min)
	}
}

// 插入排序
func (s *Sort) Insert(a []Comparable) {
	n := len(a)
	for i:= 1; i < n; i++ {
		for j := i; j > 0 && Less(a[j], a[j-1]); j-- {
			Exch(a, j, j-1)
		}
	}
}

//希尔排序
func (s *Sort) Shellsort(a []Comparable) {
	n := len(a)
	h := 1
	for ;h < n/3; h = h*3 + 1 {
	}
	for h >= 1 {
		for i := 1; i < n; i ++ {
			for j := i; j >= h && Less(a[j], a[j-h]); j -=h {
				Exch(a, j, j-h)
			}
		}
		h = h/3
	}
}


func (s *Sort) Merge(a []Comparable, lo, mid, hi int) {
	i := lo
	j := mid + 1
	aux := make([]Comparable, len(a))
	copy(aux, a)

	//for k:=0; k < len(a); k++ {
	//	aux[k] = a[k]
	//}
	for k := lo; k <= hi; k++ {
		if i > mid {
			a[k] = aux[j]
			j++
		}else if j > hi {
			a[k] = aux[i]
			i++
		}else if Less(aux[i], aux[j]) {
			a[k] = aux[i]
			i++
		}else {
			a[k] = aux[j]
			j++
		}
	}
}

// 归并排序 自顶向下
func  (s *Sort) Mergesort(a []Comparable, lo, hi int) {
	// 优化1：对于小数组，使用插入排序，减少函数调用，性能提升明显
	if hi - lo < 15 {
		s.Insert(a[lo:hi+1])
		return
	}
	if hi <= lo {
		return
	}
	mid := lo + (hi - lo)/2
	s.Mergesort(a, lo, mid)
	s.Mergesort(a, mid+1, hi)
	// 优化2：如果a[mid]< a[mid+1], 说明a[lo...hi]已经有序，不再归并
	if Less(a[mid], a[mid+1]) {
		return
	}
	s.Merge(a, lo, mid, hi)
	return
}

//自底向上
func (s *Sort) MergeSortDownToUp(a []Comparable) {
	n := len(a)
	for sz := 1; sz < n; sz += sz {
		for lo := 0; lo < n - sz; lo += sz +sz {
			hi := lo + sz + sz -1
			if hi > n -1 {
				hi = n -1
			}
			s.Merge(a, lo, lo + sz - 1, hi)
		}
	}
}

func (s *Sort) Time(alg string, a []Comparable) time.Duration {
	t1 := time.Now()
	switch alg {
	case "selection":
		s.Slection(a)
	case "insert":
		s.Insert(a)
	case "shell":
		s.Shellsort(a)
	case "merge":
		s.Mergesort(a, 0 ,len(a)-1)
	case "MergeSortDownToUp":
		s.MergeSortDownToUp(a)
	}

	t := time.Since(t1)
	return t
}

// 随机生成n 个序列，排序， 重复times 次
func Performance(alg string, n, times int) {
	s := new(Sort)
	var t time.Duration = 0
	for i:=0; i<times; i++  {
		rand.Seed(time.Now().UnixNano())
		IntSlice := make([]Comparable, 0)
		for j:=0; j<n; j++ {
			x := rand.Intn(10000)
			IntSlice = append(IntSlice, Int(x))
		}
		//fmt.Println(IntSlice)
		t1 := s.Time(alg, IntSlice)
		t = t + t1
		fmt.Println(IsSorted(IntSlice))
	}
	fmt.Println(alg,": ",t )

}

// merge two sorted slice to one sorted slice
func Merge1(a, b []int,) []int{
	// if a or b is null
	var c = []int{}
	n := len(a) + len(b)
	var i, j, k int
	for k=0; k<n; k++ {
		if i >= len(a) {
			c = append(c, b[j:]...)
			break
		}
		if j >= len(b) {
			c = append(c, a[i:]...)
				break
		}
		if a[i] < b[j] {
			c = append(c, a[i])
			i ++
		}else {
			c = append(c, b[j])
			j ++
		}
	}
	return c
}

