package main

import "math/rand"

func main() {

}

type RandomizedSet struct {
	rset  map[int]int // maps value to index
	rlist map[int]int // maps index to value
	size  int         // int of length of entries
}

// Constructor initializes a new instance of RandomizedSet
func Constructor() RandomizedSet {
	return RandomizedSet{
		rset:  make(map[int]int),
		rlist: make(map[int]int),
		size:  0,
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	_, exists := this.rset[val]
	if !exists {
		this.rset[val] = this.size
		this.rlist[this.size] = val
		this.size++
		return true
	} else {
		return false
	}
}

func (this *RandomizedSet) Remove(val int) bool {
	idx, exists := this.rset[val]
	if !exists {
		delete(this.rset, val)
		delete(this.rlist, idx)
		this.size--
		return true
	} else {
		return false
	}
}

func (this *RandomizedSet) GetRandom() int {
	val, _ := this.rlist[rand.Intn(this.size)]
	return val
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
