package main

func main() {

}

type RandomizedSet struct {
	rset map[int]struct{}
}

func Constructor() RandomizedSet {
}

func (this *RandomizedSet) Insert(val int) bool {
	_, exists := this.rset[val]
	if !exists {
		this.rset[val] = struct{}{}
		return true
	} else {
		return false
	}
}

func (this *RandomizedSet) Remove(val int) bool {
	_, exists := this.rset[val]
	if !exists {
		delete(this.rset, val)
		return true
	} else {
		return false
	}
}

func (this *RandomizedSet) GetRandom() int {

}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
