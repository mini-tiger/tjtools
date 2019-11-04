package LinkedMap

import (
	"sort"
	"sync"
)

// 有序MAP 不建议超过2000个key
type LinkedMap struct {
	sync.RWMutex
	MData map[string]interface{}
	MLink map[int]string
}

func NewLinkedMap() *LinkedMap {
	return &LinkedMap{
		MData: make(map[string]interface{}),
		MLink: make(map[int]string),
	}
}
func (this *LinkedMap) Put(key string, val interface{}) {
	this.Lock()
	this.MData[key] = val
	this.Unlock()
	this.MLink[this.Max()] = key

}

func (this *LinkedMap) Max() int {
	this.RLock()
	defer func() {
		this.RUnlock()

	}()
	var keys []int
	for k, _ := range this.MLink {
		keys = append(keys, k)
	}
	if len(keys) == 0 {
		return 1
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	//fmt.Printf("current max %d\n", keys[0]+1)
	return keys[0] + 1
}

func (this *LinkedMap) Get(key string) (interface{}, bool) {
	this.RLock()
	val, exists := this.MData[key]
	this.RUnlock()
	return val, exists
}

func (this *LinkedMap) Remove(key string) {
	this.Lock()
	delete(this.MData, key)
	for lkey, value := range this.MLink {
		if value == key {
			delete(this.MLink, lkey)
		}
	}
	this.Unlock()
}

func (this *LinkedMap) GetAndRemove(key string) (v interface{}, b bool) {
	this.RLock()
	v, b = this.MData[key]
	this.RUnlock()
	this.Remove(key)
	return
}
func (this *LinkedMap) Clear() {
	this.Lock()
	this.MData = make(map[string]interface{})
	this.MLink = make(map[int]string)
	this.Unlock()
}

func (this *LinkedMap) Keys() []string {
	this.RLock()
	defer this.RUnlock()

	keys := make([]string, 0)
	for key, _ := range this.MData {
		keys = append(keys, key)
	}
	return keys
}

func (this *LinkedMap) ContainsKey(key string) bool {
	this.RLock()
	_, exists := this.MData[key]
	this.RUnlock()
	return exists
}

func (this *LinkedMap) Size() int {
	this.RLock()
	len := len(this.MData)
	this.RUnlock()
	return len
}

func (this *LinkedMap) IsEmpty() bool {
	this.RLock()
	empty := (len(this.MData) == 0)
	this.RUnlock()
	return empty
}
func (this *LinkedMap) SortLinkMap() (keys []string) {
	linkdata := make([]int, 0)
	this.RLock()
	for k, _ := range this.MLink {
		linkdata = append(linkdata, k)
	}
	if len(linkdata) == 0 {
		return
	}
	sort.Sort(sort.IntSlice(linkdata))
	//fmt.Printf("%+v\n",this.MLink)
	//fmt.Printf("%+v\n",linkdata)
	for _,keynum := range linkdata {
		keys = append(keys, this.MLink[keynum])
	}
	//fmt.Printf("%+v\n",keys)
	this.RUnlock()
	return
}

//func main() {
//	a := NewLinkedMap();
//	a.Put("a", 1)
//	a.Put("b", 2)
//	a.Put("c", 3)
//	//fmt.Println(a.Max())
//	//fmt.Printf("%+v\n", a.MData)
//	//fmt.Printf("%+v\n", a.MLink)
//	for _, key := range a.SortLinkMap() {
//		//fmt.Println(key)
//		if v, e := a.Get(key); e {
//			fmt.Printf("key:%s,value:%v\n", key, v)
//		}
//	}
//
//	for _, key := range a.SortLinkMap() {
//		//fmt.Println(key)
//		if v, e := a.Get(key); e {
//			fmt.Printf("key:%s,value:%v\n", key, v)
//		}
//	}
//}
