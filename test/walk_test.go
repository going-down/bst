package test

import (
	"fmt"
	"github.com/going-down/bst/pkg/assoc"
	"github.com/going-down/bst/pkg/binary_tree/visit"
	"github.com/going-down/bst/pkg/binary_tree/walk"
	"github.com/going-down/bst/pkg/bst"
	"github.com/going-down/bst/pkg/comparable"
	"github.com/going-down/bst/pkg/element"
	"reflect"
	"testing"
)

func IntMapFromSliceTester(t *testing.T, walkSequence walk.Sequence, input, expected [][]int) {
	intMap := bst.Tree{}
	intMap.Add(assoc.MakeElement(comparable.MakeInteger(9), 97))
	for _, value := range input {
		intMap.Add(assoc.MakeElement(comparable.MakeInteger(value[0]), value[1]))
	}
	var result [][]int
	intMap.Iterate(&bst.Walker{
		Visitor: &visit.SideEffect{Callable: func(element element.Interface) {
			result = append(result, []int{element.GetKey().(*comparable.Integer).Value, (*element.GetValue()).(int)})
		}},
		WalkSequence: walkSequence,
	})
	//fmt.Printf("%#v\n", result)
	if !reflect.DeepEqual(result, expected) {
		t.Fail()
	}
}

type IntTraceByWalker struct {
	walkSequence walk.Sequence
	trace        [][]int
}

func TestIntMapFromSlice(t *testing.T) {
	input := [][]int{{10, 1}, {16, 3}, {12, 5}, {11, 10}, {15, -1}, {14, -100}}
	config := []IntTraceByWalker{
		{
			walkSequence: walk.Sequence{walk.Data, walk.Left, walk.Right},
			trace:        [][]int{{9, 97}, {10, 1}, {16, 3}, {12, 5}, {11, 10}, {15, -1}, {14, -100}},
		},
		{
			walkSequence: walk.Sequence{walk.Left, walk.Data, walk.Right},
			trace:        [][]int{{9, 97}, {10, 1}, {11, 10}, {12, 5}, {14, -100}, {15, -1}, {16, 3}},
		},
		{
			walkSequence: walk.Sequence{walk.Left, walk.Right, walk.Data},
			trace:        [][]int{{11, 10}, {14, -100}, {15, -1}, {12, 5}, {16, 3}, {10, 1}, {9, 97}},
		},
		{
			walkSequence: walk.Sequence{walk.Data, walk.Right, walk.Left},
			trace:        [][]int{{9, 97}, {10, 1}, {16, 3}, {12, 5}, {15, -1}, {14, -100}, {11, 10}},
		},
		{
			walkSequence: walk.Sequence{walk.Right, walk.Data, walk.Left},
			trace:        [][]int{{16, 3}, {15, -1}, {14, -100}, {12, 5}, {11, 10}, {10, 1}, {9, 97}},
		},
		{
			walkSequence: walk.Sequence{walk.Right, walk.Left, walk.Data},
			trace:        [][]int{{14, -100}, {15, -1}, {11, 10}, {12, 5}, {16, 3}, {10, 1}, {9, 97}},
		},
	}
	for _, i := range config {
		IntMapFromSliceTester(t, i.walkSequence, input, i.trace)
	}
}

func StringMapFromSliceTester(t *testing.T, walkSequence walk.Sequence, input, expected [][]interface{}) {
	intMap := bst.Tree{}
	for _, value := range input {
		intMap.Add(assoc.MakeElement(comparable.MakeString(value[0].(string)), value[1]))
	}
	var result [][]interface{}
	intMap.Iterate(&bst.Walker{
		Visitor: &visit.SideEffect{Callable: func(element element.Interface) {
			result = append(result, []interface{}{element.GetKey().(*comparable.String).Value, fmt.Sprint(*element.GetValue())})
		}},
		WalkSequence: walkSequence,
	})
	//fmt.Printf("%#v\n", result)
	if !reflect.DeepEqual(result, expected) {
		t.Fail()
	}
}

type TraceByWalker struct {
	walkSequence walk.Sequence
	trace        [][]interface{}
}

func TestStringMapFromSlice(t *testing.T) {
	input := [][]interface{}{{"g", "genny"}, {"k", "kate"}, {"a", "apollo"}, {"r", "regex"}, {"b", -1}, {"x", -100}}
	config := []TraceByWalker{
		{
			walkSequence: walk.Sequence{walk.Data, walk.Left, walk.Right},
			trace:        [][]interface{}{{"g", "genny"}, {"a", "apollo"}, {"b", "-1"}, {"k", "kate"}, {"r", "regex"}, {"x", "-100"}},
		},
		{
			walkSequence: walk.Sequence{walk.Left, walk.Data, walk.Right},
			trace:        [][]interface{}{{"a", "apollo"}, {"b", "-1"}, {"g", "genny"}, {"k", "kate"}, {"r", "regex"}, {"x", "-100"}},
		},
		{
			walkSequence: walk.Sequence{walk.Left, walk.Right, walk.Data},
			trace:        [][]interface{}{{"b", "-1"}, {"a", "apollo"}, {"x", "-100"}, {"r", "regex"}, {"k", "kate"}, {"g", "genny"}},
		},
		{
			walkSequence: walk.Sequence{walk.Data, walk.Right, walk.Left},
			trace:        [][]interface{}{{"g", "genny"}, {"k", "kate"}, {"r", "regex"}, {"x", "-100"}, {"a", "apollo"}, {"b", "-1"}},
		},
		{
			walkSequence: walk.Sequence{walk.Right, walk.Data, walk.Left},
			trace:        [][]interface{}{{"x", "-100"}, {"r", "regex"}, {"k", "kate"}, {"g", "genny"}, {"b", "-1"}, {"a", "apollo"}},
		},
		{
			walkSequence: walk.Sequence{walk.Right, walk.Left, walk.Data},
			trace:        [][]interface{}{{"x", "-100"}, {"r", "regex"}, {"k", "kate"}, {"b", "-1"}, {"a", "apollo"}, {"g", "genny"}},
		},
	}
	for _, i := range config {
		StringMapFromSliceTester(t, i.walkSequence, input, i.trace)
	}
}
