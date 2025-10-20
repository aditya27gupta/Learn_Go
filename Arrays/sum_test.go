package arrays

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5}
		want := 15
		got := Sum(nums)
		if got != want {
			t.Errorf("Got %d Want %d Given %v", got, want, nums)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{4, 5})
	want := []int{3, 9}
	if !slices.Equal(got, want) {
		t.Errorf("Got %v,but Want %v", got, want)
	}
}

func BenchmarkSum(b *testing.B) {
	nums := []int{1, 2, 3, 4, 5}
	for b.Loop() {
		Sum(nums)
	}
}

func BenchmarkSumAll(b *testing.B) {
	for b.Loop() {
		SumAll([]int{3, 4, 5}, []int{7, 2, 9})
	}
}
