package main

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)


// 今天主要提到了第一种方案，这是在编码时就完全确定键和值的类型，然后利用 Go 语言的编译器帮我们做检查。
// todo 改进：这样做很方便，不是吗？不过，虽然方便，但是却让这样的字典类型缺少了一些灵活性。
//  如果我们还需要一个键类型为uint32并发安全字典的话，那就不得不再如法炮制地写一遍代码了。因此，在需求多样化之后，工作量反而更大，甚至会产生很多雷同的代码。
// IntStrMap 代表键类型为int、值类型为string的并发安全字典。
type IntStrMap struct {
	m sync.Map
}

func (iMap *IntStrMap) Delete(key int) {
	iMap.m.Delete(key)
}

func (iMap *IntStrMap) Load(key int) (value string, ok bool) {
	v, ok := iMap.m.Load(key)
	if v != nil {
		value = v.(string)
	}
	return
}

func (iMap *IntStrMap) LoadOrStore(key int, value string) (actual string, loaded bool) {
	a, loaded := iMap.m.LoadOrStore(key, value)
	actual = a.(string)
	return
}

func (iMap *IntStrMap) Range(f func(key int, value string) bool) {
	f1 := func(key, value interface{}) bool {
		return f(key.(int), value.(string))
	}
	iMap.m.Range(f1)
}

func (iMap *IntStrMap) Store(key int, value string) {
	iMap.m.Store(key, value)
}

/**
TODO 第一种方案存在一个很明显的缺陷，那就是无法灵活地改变字典的键和值的类型。
    一旦需求出现多样化，编码的工作量就会随之而来。
    第二种方案很好地弥补了这一缺陷，但是，那些反射操作或多或少都会降低程序的性能。
   我们往往需要根据实际的应用场景，通过严谨且一致的测试，来获得和比较程序的各项指标，并以此作为方案选择的重要依据之一。
 */

// 在第二种方案中，我们封装的结构体类型的所有方法，都可以与sync.Map类型的方法完全一致（包括方法名称和方法签名）。
// ConcurrentMap todo 代表可自定义键类型和值类型的并发安全字典。
type ConcurrentMap struct {
	m         sync.Map
	keyType   reflect.Type
	valueType reflect.Type
}
// 这两个字段的类型都是reflect.Type，我们可称之为反射类型。


// todo  传入 key，value的类型, 初始化map
func NewConcurrentMap(keyType, valueType reflect.Type) (*ConcurrentMap, error) {
	if keyType == nil {
		return nil, errors.New("nil key type")
	}
	if !keyType.Comparable() {
		return nil, fmt.Errorf("incomparable key type: %s", keyType)
	}
	if valueType == nil {
		return nil, errors.New("nil value type")
	}
	cMap := &ConcurrentMap{
		keyType:   keyType,
		valueType: valueType,
	}
	return cMap, nil
}

func (cMap *ConcurrentMap) Delete(key interface{}) {
	if reflect.TypeOf(key) != cMap.keyType {
		return
	}
	cMap.m.Delete(key)
}

// 在m字段的值中查找键值对
func (cMap *ConcurrentMap) Load(key interface{}) (value interface{}, ok bool) {
	if reflect.TypeOf(key) != cMap.keyType {
		return
	}
	return cMap.m.Load(key)
	// 我们把一个接口类型值传入reflect.TypeOf函数，就可以得到与这个值的实际类型对应的反射类型值。
	// 因此，如果参数值的反射类型与keyType字段代表的反射类型不相等，那么我们就忽略后续操作，并直接返回。

}

func (cMap *ConcurrentMap) LoadOrStore(key, value interface{}) (actual interface{}, loaded bool) {
	if reflect.TypeOf(key) != cMap.keyType {
		panic(fmt.Errorf("wrong key type: %v", reflect.TypeOf(key)))
	}
	if reflect.TypeOf(value) != cMap.valueType {
		panic(fmt.Errorf("wrong value type: %v", reflect.TypeOf(value)))
	}
	actual, loaded = cMap.m.LoadOrStore(key, value)
	return
}

func (cMap *ConcurrentMap) Range(f func(key, value interface{}) bool) {
	cMap.m.Range(f)
}

// 当参数key或value的实际类型不符合要求时，Store方法会立即引发 panic。
//
func (cMap *ConcurrentMap) Store(key, value interface{}) {
	if reflect.TypeOf(key) != cMap.keyType {
		panic(fmt.Errorf("wrong key type: %v", reflect.TypeOf(key)))
	}
	if reflect.TypeOf(value) != cMap.valueType {
		panic(fmt.Errorf("wrong value type: %v", reflect.TypeOf(value)))
	}
	cMap.m.Store(key, value)
}

// pairs 代表测试用的键值对列表。
var pairs = []struct {
	k int
	v string
}{
	{k: 1, v: "a"},
	{k: 2, v: "b"},
	{k: 3, v: "c"},
	{k: 4, v: "d"},
}

func main() {
	// 示例1。
	var sMap sync.Map
	//sMap.Store([]int{1, 2, 3}, 4) // 这行代码会引发panic。
	_ = sMap

	// 示例2。
	{
		var iMap IntStrMap
		iMap.Store(pairs[0].k, pairs[0].v)
		iMap.Store(pairs[1].k, pairs[1].v)
		iMap.Store(pairs[2].k, pairs[2].v)
		fmt.Println("[Three pairs have been stored in the IntStrMap instance]")

		iMap.Range(func(key int, value string) bool {
			fmt.Printf("The result of an iteration in Range: %d, %s\n",
				key, value)
			return true
		})

		k0 := pairs[0].k
		v0, ok := iMap.Load(k0)
		fmt.Printf("The result of Load: %v, %v (key: %v)\n",
			v0, ok, k0)

		k3 := pairs[3].k
		v3, ok := iMap.Load(k3)
		fmt.Printf("The result of Load: %v, %v (key: %v)\n",
			v3, ok, k3)

		k2, v2 := pairs[2].k, pairs[2].v
		actual2, loaded2 := iMap.LoadOrStore(k2, v2)
		fmt.Printf("The result of LoadOrStore: %v, %v (key: %v, value: %v)\n",
			actual2, loaded2, k2, v2)
		v3 = pairs[3].v
		actual3, loaded3 := iMap.LoadOrStore(k3, v3)
		fmt.Printf("The result of LoadOrStore: %v, %v (key: %v, value: %v)\n",
			actual3, loaded3, k3, v3)

		k1 := pairs[1].k
		iMap.Delete(k1)
		fmt.Printf("[The pair with the key of %v has been removed from the IntStrMap instance]\n",
			k1)
		v1, ok := iMap.Load(k1)
		fmt.Printf("The result of Load: %v, %v (key: %v)\n",
			v1, ok, k1)
		v1 = pairs[1].v
		actual1, loaded1 := iMap.LoadOrStore(k1, v1)
		fmt.Printf("The result of LoadOrStore: %v, %v (key: %v, value: %v)\n",
			actual1, loaded1, k1, v1)

		iMap.Range(func(key int, value string) bool {
			fmt.Printf("The result of an iteration in Range: %d, %s\n",
				key, value)
			return true
		})
	}
	fmt.Println()

	// 示例2。
	{
		cMap, err := NewConcurrentMap(
			reflect.TypeOf(pairs[0].k), reflect.TypeOf(pairs[0].v))
		if err != nil {
			fmt.Printf("fatal error: %s", err)
			return
		}
		cMap.Store(pairs[0].k, pairs[0].v)
		cMap.Store(pairs[1].k, pairs[1].v)
		cMap.Store(pairs[2].k, pairs[2].v)
		fmt.Println("[Three pairs have been stored in the ConcurrentMap instance]")

		cMap.Range(func(key, value interface{}) bool {
			fmt.Printf("The result of an iteration in Range: %d, %s\n",
				key, value)
			return true
		})

		k0 := pairs[0].k
		v0, ok := cMap.Load(k0)
		fmt.Printf("The result of Load: %v, %v (key: %v)\n",
			v0, ok, k0)

		k3 := pairs[3].k
		v3, ok := cMap.Load(k3)
		fmt.Printf("The result of Load: %v, %v (key: %v)\n",
			v3, ok, k3)

		k2, v2 := pairs[2].k, pairs[2].v
		actual2, loaded2 := cMap.LoadOrStore(k2, v2)
		fmt.Printf("The result of LoadOrStore: %v, %v (key: %v, value: %v)\n",
			actual2, loaded2, k2, v2)
		v3 = pairs[3].v
		actual3, loaded3 := cMap.LoadOrStore(k3, v3)
		fmt.Printf("The result of LoadOrStore: %v, %v (key: %v, value: %v)\n",
			actual3, loaded3, k3, v3)

		k1 := pairs[1].k
		cMap.Delete(k1)
		fmt.Printf("[The pair with the key of %v has been removed from the ConcurrentMap instance]\n",
			k1)
		v1, ok := cMap.Load(k1)
		fmt.Printf("The result of Load: %v, %v (key: %v)\n",
			v1, ok, k1)
		v1 = pairs[1].v
		actual1, loaded1 := cMap.LoadOrStore(k1, v1)
		fmt.Printf("The result of LoadOrStore: %v, %v (key: %v, value: %v)\n",
			actual1, loaded1, k1, v1)

		cMap.Range(func(key, value interface{}) bool {
			fmt.Printf("The result of an iteration in Range: %d, %s\n",
				key, value)
			return true
		})
	}
}
