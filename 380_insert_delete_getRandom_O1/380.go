package main

import (
	"math/rand"
	"slices"
)

type RandomizedSet struct {
	main map[int]struct{}

	navigator map[int]int
	arr       []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		main:      make(map[int]struct{}),
		navigator: make(map[int]int),
		arr:       make([]int, 0),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.main[val]; ok {
		return false
	}
	this.main[val] = struct{}{}
	this.arr = append(this.arr, val)
	this.navigator[val] = len(this.arr) - 1

	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.main[val]; !ok {
		return false
	}

	delete(this.main, val)
	i := this.navigator[val]
	l := this.arr[:i]
	r := this.arr[i+1:]
	n := slices.Concat(l, r)
	this.arr = n
	delete(this.navigator, val)
	for k, v := range this.navigator {
		if v > i {
			this.navigator[k]--
		}
	}
	return true
}

func (this *RandomizedSet) GetRandom() int {
	if len(this.arr) == 0 {
		return 0
	}
	r := rand.Intn(len(this.arr))
	return this.arr[r]
}

func main() {
	st := Constructor()
	var b bool
	b = st.Insert(1) //t
	b = st.Insert(2) //t
	b = st.Insert(2) //f
	b = st.Insert(3) //t
	b = st.Insert(3) //f
	b = st.Insert(3) //f
	b = st.Insert(4) //t
	b = st.Remove(5) //f
	b = st.Remove(2) //t

	b = st.Remove(1) //t
	b = st.Remove(3) //t
	b = st.Remove(4) //t

	_ = b
}
