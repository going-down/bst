package test

import (
	"github.com/going-down/bst/pkg/assoc"
	"github.com/going-down/bst/pkg/bst"
	"github.com/going-down/bst/pkg/comparable"
	"log"
	"math/rand"
	"testing"
	"time"
)

func generateLargeTree(n int) (bst.Tree, map[int]int) {
	var randomKeys []int
	for i := 0; i < n; i++ {
		randomKeys = append(randomKeys, i)
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(n, func(i, j int) { randomKeys[i], randomKeys[j] = randomKeys[j], randomKeys[i] })
	result := bst.Tree{}
	mapToCheck := map[int]int{}
	for i := 0; i < n; i++ {
		value := rand.Int()
		result.Add(assoc.MakeElement(comparable.MakeInteger(randomKeys[i]), value))
		mapToCheck[randomKeys[i]] = value
	}
	return result, mapToCheck
}

func findTester(tb testing.TB) {
	var len = 1000
	var largeTree, mapToCheck = generateLargeTree(len)
	key := rand.Intn(len)
	expected := mapToCheck[key]
	got := (*largeTree.Find(comparable.MakeInteger(key))).(int)
	if got != expected {
		log.Printf("key: %d, expected: %d, got: %d\n", key, expected, got)
		tb.Fail()
	}
}

func BenchmarkFind(b *testing.B) {
	for n := 0; n < b.N; n++ {
		findTester(b)
	}
}

func TestFind(t *testing.T) {
	findTester(t)
}

func TestFindNonExistent(t *testing.T) {
	var len = 1000
	var largeTree, _ = generateLargeTree(len)
	key := len + rand.Intn(len)
	got := largeTree.Find(comparable.MakeInteger(key))
	if got != nil {
		log.Printf("key: %d, expected: %#v, got: %d\n", key, nil, got)
		t.Fail()
	}
}

func TestAddExistent(t *testing.T) {
	var len = 1000
	var largeTree, _ = generateLargeTree(len)
	key := rand.Intn(len)
	defer func() {
		if r := recover(); r == nil {
			t.Fail()
		}
	}()
	largeTree.Add(assoc.MakeElement(comparable.MakeInteger(key), rand.Int()))
}

func setTester(tb testing.TB, len, key int) {
	var largeTree, _ = generateLargeTree(len)
	expected := rand.Int()
	largeTree.Set(comparable.MakeInteger(key), expected)
	got := (*largeTree.Find(comparable.MakeInteger(key))).(int)
	if got != expected {
		log.Printf("key: %d, expected: %d, got: %d\n", key, expected, got)
		tb.Fail()
	}
}

func BenchmarkSet(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var len = 1000
		setTester(b, len, len+rand.Intn(len))
	}
}

func TestSet(t *testing.T) {
	var len = 1000
	setTester(t, len, len+rand.Intn(len))
}

func TestSetTaken(t *testing.T) {
	var len = 1000
	setTester(t, len, len-rand.Intn(len))
}

/*
func TreeToMap(p bst.Tree) map[interface{}]interface{} {
	var result = map[interface{}]interface{}{}
	p.Iterate(&bst.Walker{
		Visitor: &visit.SideEffect{Callable: func(element element.Interface) {
			result[element.GetKey()] = *element.GetValue()
		}},
		WalkSequence: walk.Sequence{walk.Data, walk.Left, walk.Right},
	})
	return result
}
*/
