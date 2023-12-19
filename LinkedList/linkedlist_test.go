package linkedlist

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

type testCase struct {
	arg, expected, afterDel []float64
	keyToDelSearch float64
	searchOut, delOut bool
}

var testCases = []testCase {
	{arg: []float64{0, 1, 2, 3, 4}, expected: []float64{0, 1, 2, 3, 4}, keyToDelSearch: 2, afterDel: []float64{0, 1, 3, 4}, searchOut: true, delOut: true},
	{arg: []float64{}, expected: []float64{}, keyToDelSearch: 0, afterDel: []float64{}, searchOut: false, delOut: false},	
}
func TestInsertion(t *testing.T) {
	for _, testCase := range testCases {
		flag := true
		var ll *LinkedList
		for _, key := range testCase.arg {
			if flag {
				ll = CreateLinkedList(key)
				flag = false
			} else {
				ll.InsertKey(key)
			}
		}

		got := ll.ShowKeys()
		want := testCase.expected
		
		if !cmp.Equal(got, want) {
			t.Errorf("Got: %v, Expected: %v", got, want)
		}
		
	}
}

func TestListLenght(t *testing.T) {
	t.Logf("Test: %v", t.Name())
	for _, testCase := range testCases {
		flag := true
		var ll *LinkedList
		for _, key := range testCase.arg {
			if flag {
				ll = CreateLinkedList(key)
			} else {
				ll.InsertKey(key)
			}
		}

		got := ll.ListLenght()
		want := len(ll.ShowKeys())

		if !cmp.Equal(got, want) {
			t.Errorf("Got: %v Expected: %v", got, want)
		}
	}
	
}

func TestTransformArray(t *testing.T) {
	for _, testCase := range testCases {
		ll := TransformArray(testCase.arg)

		got := ll.ShowKeys()
		want := testCase.expected

		if !cmp.Equal(got, want) {
			t.Errorf("Got: %v, Expected: %v", got, want)
		}
	}
}

func TestSearchKey(t *testing.T) {
	for _, testCase := range testCases {
		ll := TransformArray(testCase.arg)
		got := ll.SearchKey(testCase.keyToDelSearch)
		want := testCase.searchOut

		if got != want {
			t.Errorf("Got: %v, Expected: %v", got, want)
		}
	}
}

func TestDeleteKey(t *testing.T) {
	for _, testeCase := range testCases {
		ll := TransformArray(testeCase.arg)
		got := ll.DeleteKey(testeCase.keyToDelSearch)
		want := testeCase.delOut

		if got != want {
			t.Errorf("Got: %v, Expected: %v", got, want)
		} else {
			got2 := ll.ShowKeys()
			want2 := testeCase.afterDel

			if !cmp.Equal(got2, want2) {
				t.Errorf("Got: %v, Expected: %v", got2, want2)
			}
		}
	}
}

func TestTransformLinkedList(t *testing.T) {
	for _, testCase := range testCases {
		ll := TransformArray(testCase.arg)
		got := ll.TransformLinkedList()
		want := testCase.arg

		if !cmp.Equal(got, want) {
			t.Errorf("Got: %v, Expected: %v", got, want)
		}
	}
	
}