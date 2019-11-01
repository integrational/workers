package workers

import (
	"reflect"
	"sort"
	"testing"
)

func TestDo(t *testing.T) {
	primes := []int{2, 3, 5, 7, 11, 13}
	times1000 := func(todo int) int { return todo * 1000 }
	primesTimes1000 := []int{2000, 3000, 5000, 7000, 11000, 13000} // sorted

	cases := []struct {
		todos      []int
		do         func(int) int
		numWorkers int
		want       []int
	}{
		{nil, nil, -1, []int{}},
		{nil, nil, 0, []int{}},
		{nil, nil, 1, []int{}},
		{nil, nil, 10, []int{}},
		{[]int{}, nil, -1, []int{}},
		{[]int{}, nil, 0, []int{}},
		{[]int{}, nil, 1, []int{}},
		{[]int{}, nil, 10, []int{}},
		{[]int{}, times1000, -1, []int{}},
		{[]int{}, times1000, 0, []int{}},
		{[]int{}, times1000, 1, []int{}},
		{[]int{}, times1000, 10, []int{}},
		{primes, times1000, -1, primesTimes1000},
		{primes, times1000, 0, primesTimes1000},
		{primes, times1000, 1, primesTimes1000},
		{primes, times1000, 10, primesTimes1000},
	}
	for _, c := range cases {
		got := Do(c.todos, c.do, c.numWorkers)
		if sort.Ints(got); !reflect.DeepEqual(got, c.want) {
			t.Errorf("DoWorkers(%v, %v) == %v, want %v", c.todos, c.numWorkers, got, c.want)
		}
	}
}
