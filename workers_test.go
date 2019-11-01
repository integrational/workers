package workers

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func TestDo(t *testing.T) {
	times1000 := func(todo int) int { return todo * 1000 }
	few, few1000 := []int{2, 3, 5, 7, 11}, []int{2000, 3000, 5000, 7000, 11000} // sorted
	many, many1000 := testDataForMany(times1000)                                // sorted

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
		{few, nil, -1, few},
		{few, nil, 0, few},
		{few, nil, 1, few},
		{few, nil, 10, few},
		{few, times1000, -1, few1000},
		{few, times1000, 0, few1000},
		{few, times1000, 1, few1000},
		{few, times1000, 10, few1000},
		{many, times1000, 10, many1000},
	}
	for _, c := range cases {
		got := Do(c.todos, c.do, c.numWorkers)
		if sort.Ints(got); !reflect.DeepEqual(got, c.want) {
			t.Errorf("DoWorkers(%v, %v) == %v, want %v", c.todos, c.numWorkers, got, c.want)
		}
	}
}

func testDataForMany(times1000 func(int) int) (many []int, manyTimes1000 []int) {
	const howMany = 6789

	many = make([]int, howMany)
	manyTimes1000 = make([]int, howMany)
	for i := 0; i < howMany; i++ {
		many[i] = rand.Intn(howMany * 3)
		manyTimes1000[i] = times1000(many[i])
	}
	sort.Ints(manyTimes1000)
	return
}
