package sorting

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

type addTests struct {
	arg, expected []float64
}

var addTest = []addTests {
	{arg: []float64{5, 6, 1, 2, 7, 9, 0, -99}, expected: []float64{-99, 0, 1, 2, 5, 6, 7, 9}},
	{arg: []float64{0, 0, 0, 0, 0, 0, 0}, expected: []float64{0, 0, 0, 0, 0, 0, 0}},
	{arg: []float64{0, 1, 2, 3, 4, 5, 6}, expected: []float64{0, 1, 2, 3, 4, 5, 6}},
}
func TestMergeSort(t *testing.T) {
	for _, testCase := range addTest {
		got := MergeSort(testCase.arg)
		want := testCase.expected
		if !cmp.Equal(got, want) {
			t.Errorf("Got %v, want %v", got, want)
		}
	}
}

func TestBubbleSort(t *testing.T) {
	for _, testCase := range addTest {
		got := BubbleSort(testCase.arg)
		want := testCase.expected
		if !cmp.Equal(got, want) {
			t.Errorf("Got %v, want %v", got, want)
		}
	}
}

func TestSelectionSort(t *testing.T) {
	for _, testCase := range addTest {
		got := SelectionSort(testCase.arg)
		want := testCase.expected
		if !cmp.Equal(got, want) {
			t.Errorf("Got %v, want %v", got, want)
		}
	}
}

func TestInsertionSort(t *testing.T) {
	for _, testCase := range addTest {
		got := InsertionSort(testCase.arg)
		want := testCase.expected
		if !cmp.Equal(got, want) {
			t.Errorf("Got %v, want %v", got, want)
		}
	}
}

func TestHeapSort(t *testing.T) {
	for _, testCase := range addTest {
		got := testCase.arg
		want := testCase.expected
		HeapSort(got)
		if !cmp.Equal(got, want) {
			t.Errorf("Got %v, want %v", got, want)
		}
	}
}