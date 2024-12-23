package enumerable

import (
	"fmt"
	"testing"
)

func TestFilterInts(t *testing.T) {
	t.Parallel()
	tests := []struct {
		filter   func(x int32) bool
		input    []int32
		expected []int32
	}{
		{
			func(x int32) bool { return x != 12345 },
			[]int32{12345, 123, 125},
			[]int32{123, 125},
		},
		{
			func(x int32) bool { return x != 12345 },
			[]int32{12345},
			[]int32{},
		},
		{
			func(x int32) bool { return x != 0 },
			[]int32{12345, 123, 125},
			[]int32{12345, 123, 125},
		},
		{
			func(x int32) bool { return x != 0 },
			[]int32{},
			[]int32{},
		},
	}

	for _, test := range tests {
		test := test
		res := New(test.input).Filter(test.filter).Value()

		if len(res) != len(test.expected) {
			t.Errorf("Wrong number of elements in array. Expected: %d found: %d", len(test.expected), len(res))
		}

		for i, expected := range test.expected {
			result := res[i]
			if expected != result {
				t.Errorf("Wrong element found. Expected: '%d' found '%d'", result, expected)
			}
		}
	}
}

func TestFilterObjects(t *testing.T) {
	t.Parallel()
	type Book struct {
		title string
	}

	tests := []struct {
		filter   func(x Book) bool
		input    []Book
		expected []Book
	}{
		{
			func(x Book) bool { return x.title != "book1" },
			[]Book{{title: "book1"}, {title: "book2"}, {title: "book3"}},
			[]Book{{title: "book2"}, {title: "book3"}},
		},
		{
			func(x Book) bool { return x != Book{title: "book1"} },
			[]Book{{title: "book1"}},
			[]Book{},
		},
		{
			func(x Book) bool { return x.title != "" },
			[]Book{{title: "book1"}, {title: "book2"}, {title: "book3"}},
			[]Book{{title: "book1"}, {title: "book2"}, {title: "book3"}},
		},
		{
			func(x Book) bool { return x.title != "" },
			[]Book{},
			[]Book{},
		},
	}

	for _, test := range tests {
		test := test
		res := New(test.input).Filter(test.filter).Value()

		if len(res) != len(test.expected) {
			t.Errorf("Wrong number of elements in array. Expected: %d found: %d", len(test.expected), len(res))
		}

		for i, expected := range test.expected {
			result := res[i]
			if expected != result {
				t.Errorf("Wrong element found. Expected: '%s' found '%s'", result, expected)
			}
		}
	}
}

func TestMap(t *testing.T) {
	t.Parallel()
	tests := []struct {
		mutation func(x int32) string
		input    []int32
		expected []string
	}{
		{
			func(x int32) string { return fmt.Sprint(x) },
			[]int32{12345, 123, 125},
			[]string{"12345", "123", "125"},
		},
	}

	for _, test := range tests {
		test := test
		res := Map(test.input, test.mutation)

		if len(res) != len(test.expected) {
			t.Errorf("Wrong number of elements in array. Expected: %d found: %d", len(test.expected), len(res))
		}

		for i, expected := range test.expected {
			result := res[i]
			if expected != result {
				t.Errorf("Wrong element found. Expected: '%s' found '%s'", result, expected)
			}
		}
	}
}

func TestMapStringToString(t *testing.T) {
	t.Parallel()
	tests := []struct {
		mutation func(x string) string
		input    []string
		expected []string
	}{
		{
			func(x string) string { return x },
			[]string{"12345", "123", "125"},
			[]string{"12345", "123", "125"},
		},
	}

	for _, test := range tests {
		test := test
		res := Map(test.input, test.mutation)

		if len(res) != len(test.expected) {
			t.Errorf("Wrong number of elements in array. Expected: %d found: %d", len(test.expected), len(res))
		}

		for i, expected := range test.expected {
			result := res[i]
			if expected != result {
				t.Errorf("Wrong element found. Expected: '%s' found '%s'", result, expected)
			}
		}
	}
}

func TestFirstOrDefault(t *testing.T) {
	t.Parallel()
	tests := []struct {
		comparator   func(x int32) bool
		inputArr     []int32
		defaultValue int32
		expected     int32
	}{
		{
			func(x int32) bool { return x == 12345 },
			[]int32{12345, 123, 125},
			0,
			12345,
		},
		{
			func(x int32) bool { return x == 0 },
			[]int32{12345, 123, 125},
			0,
			0,
		},
		{
			func(x int32) bool { return x == 12345 },
			[]int32{},
			0,
			0,
		},
		{
			func(x int32) bool { return x == 12345 },
			[]int32{},
			5,
			5,
		},
	}

	for _, test := range tests {
		test := test
		res := New(test.inputArr).FirstOrDefault(test.comparator, test.defaultValue)

		if res != test.expected {
			t.Errorf("Wrong value returned. Expected: %d found: %d", test.expected, res)
		}
	}
}

func TestSome(t *testing.T) {
	t.Parallel()
	tests := []struct {
		comparator func(x int32) bool
		inputArr   []int32
		expected   bool
	}{
		{
			func(x int32) bool { return x == 12345 },
			[]int32{12345, 123, 125},
			true,
		},
		{
			func(x int32) bool { return x == 123 },
			[]int32{12345, 123, 125},
			true,
		},
		{
			func(x int32) bool { return x == 123 },
			[]int32{},
			false,
		},
		{
			func(x int32) bool { return x == 10 },
			[]int32{12345, 123, 125},
			false,
		},
		{
			func(x int32) bool { return x%2 == 0 },
			[]int32{12345, 123, 124},
			true,
		},
	}

	for _, test := range tests {
		test := test
		res := Some(test.inputArr, test.comparator)

		if res != test.expected {
			t.Errorf("Wrong value returned. Expected: %t found: %t", test.expected, res)
		}
	}
}

func TestEvery(t *testing.T) {
	t.Parallel()
	tests := []struct {
		comparator func(x int32) bool
		inputArr   []int32
		expected   bool
	}{
		{
			func(x int32) bool { return x == 12345 },
			[]int32{12345, 123, 125},
			false,
		},
		{
			func(x int32) bool { return x > 1 },
			[]int32{12345, 123, 125},
			true,
		},
		{
			func(x int32) bool { return x > 9999 },
			[]int32{12345, 123, 125},
			false,
		},
		{
			func(x int32) bool { return x == 10 },
			[]int32{},
			true,
		},
	}

	for i, test := range tests {
		test := test
		res := New(test.inputArr).Every(test.comparator)

		if res != test.expected {
			t.Errorf("Wrong value returned for test %d. Expected: %t found: %t", i+1, test.expected, res)
		}
	}
}

func TestFind(t *testing.T) {
	t.Parallel()
	tests := []struct {
		comparator    func(x int32) bool
		inputArr      []int32
		expectedIndex int
		expected      int32
		errors        bool
	}{
		{
			func(x int32) bool { return x == 12345 },
			[]int32{12345, 123, 125},
			0,
			12345,
			false,
		},
		{
			func(x int32) bool { return x == 125 },
			[]int32{12345, 123, 125},
			2,
			125,
			false,
		},
		{
			func(x int32) bool { return x > 99999 },
			[]int32{12345, 123, 125},
			-1,
			0,
			true,
		},
		{
			func(x int32) bool { return x == 10 },
			[]int32{},
			-1,
			0,
			true,
		},
	}

	for i, test := range tests {
		test := test
		res, index := New(test.inputArr).Find(test.comparator)

		if res != test.expected {
			t.Fatalf("Wrong value returned for test %d. Expected: %d found: %d", i+1, test.expected, res)
		}

		if index != test.expectedIndex {
			t.Fatalf("Wrong index returned for test %d. Expected: %d found: %d", i+1, test.expected, res)
		}
	}
}

func TestSort(t *testing.T) {
	t.Parallel()
	tests := []struct {
		comparator func(a, b int32) bool
		inputArr   []int32
		expected   []int32
	}{
		{
			func(a, b int32) bool { return a >= b },
			[]int32{123, 125, 12345},
			[]int32{12345, 125, 123},
		},
		{
			func(a, b int32) bool { return a <= b },
			[]int32{7, 9, 5, 10, 3, 6, 1, 8, 2, 4},
			[]int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		},
	}

	for i, test := range tests {
		test := test
		res := New(test.inputArr).Sort(test.comparator).Value()

		if len(res) != len(test.expected) {
			t.Fatalf("Wrong value returned for test %d. Expected: %d found: %d", i+1, test.expected, res)
		}

		for i, exp := range test.expected {
			r := res[i]
			if r != exp {
				t.Fatalf("Wrong value in array inex %d. Expected: %d found: %d", i+1, test.expected, res)
			}
		}
	}
}
