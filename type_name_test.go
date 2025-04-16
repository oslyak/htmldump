//lint:file-ignore
package htmldump_test

import (
	"reflect"
	"testing"

	"htmldump"

	"github.com/stretchr/testify/require"
)

func TestMapTypeName(t *testing.T) {
	t.Parallel()

	type testCase struct {
		name     string
		input    interface{}
		expected string
	}

	testCases := []testCase{
		{
			name:     `Empty map`,
			input:    map[string]string{},
			expected: `map[string]string`,
		},
		{
			name:     `Map with int keys and string values`,
			input:    map[int]string{1: `one`, 2: `two`},
			expected: `map[int]string`,
		},
		{
			name:     `Map with string keys and int values`,
			input:    map[string]int{`one`: 1, `two`: 2},
			expected: `map[string]int`,
		},
		{
			name:     `Map with string keys and struct values`,
			input:    map[string]struct{ ID int }{`one`: {ID: 1}, `two`: {ID: 2}},
			expected: `map[string]struct { ID int }`,
		},
		{
			name:     `Map with struct keys and string values`,
			input:    map[struct{ ID int }]string{{ID: 1}: `one`, {ID: 2}: `two`},
			expected: `map[struct { ID int }]string`,
		},
		{
			name:     `Map with struct keys and struct values`,
			input:    map[struct{ ID int }]struct{ ID int }{{ID: 1}: {ID: 1}, {ID: 2}: {ID: 2}},
			expected: `map[struct { ID int }]struct { ID int }`,
		},

		{
			name:     `Map with int keys and slice values`,
			input:    map[int][]int{1: {1, 2, 3}, 2: {4, 5, 6}},
			expected: `map[int][]int`,
		},
		{
			name:     `Map with int keys and map values`,
			input:    map[int]map[int]int{1: {1: 1, 2: 2, 3: 3}, 2: {4: 4, 5: 5, 6: 6}},
			expected: `map[int]map[int]int`,
		},
		{
			name:     `Map with int keys and pointer values`,
			input:    map[int]*int{1: new(int), 2: new(int)},
			expected: `map[int]*int`,
		},
		{
			name:     `Map with int keys and interface values`,
			input:    map[int]interface{}{1: 1, 2: `two`},
			expected: `map[int]interface {}`,
		},
		{
			name:     `Map with int keys and channel values`,
			input:    map[int]chan int{1: make(chan int), 2: make(chan int)},
			expected: `map[int]chan int`,
		},
		{
			name:     `Map with int keys and function values`,
			input:    map[int]func(){1: func() {}, 2: func() {}},
			expected: `map[int]func()`,
		},
		{
			name:     `Map with int keys and slice of pointers values`,
			input:    map[int][]*int{1: {new(int), new(int), new(int)}, 2: {new(int), new(int), new(int)}},
			expected: `map[int][]*int`,
		},
		{
			name:     `Map with int keys and slice of interfaces values`,
			input:    map[int][]interface{}{1: {1, `one`}, 2: {2, `two`}},
			expected: `map[int][]interface {}`,
		},
		{
			name:     `Map with int keys and slice of channels values`,
			input:    map[int][]chan int{1: {make(chan int), make(chan int), make(chan int)}, 2: {make(chan int), make(chan int), make(chan int)}}, //nolint:lll
			expected: `map[int][]chan int`,
		},
		{
			name:     `Map with int keys and slice of functions values`,
			input:    map[int][]func(){1: {func() {}, func() {}}, 2: {func() {}, func() {}}},
			expected: `map[int][]func()`,
		},
		{
			name:     `Map with int keys and map of pointers values`,
			input:    map[int]map[int]*int{1: {1: new(int), 2: new(int), 3: new(int)}, 2: {4: new(int), 5: new(int), 6: new(int)}},
			expected: `map[int]map[int]*int`,
		},
		{
			name:     `Map with int keys and map of interfaces values`,
			input:    map[int]map[int]interface{}{1: {1: 1, 2: `two`}, 2: {3: 3, 4: `four`}},
			expected: `map[int]map[int]interface {}`,
		},
		{
			name:     `Map with int keys and map of channels values`,
			input:    map[int]map[int]chan int{1: {1: make(chan int), 2: make(chan int), 3: make(chan int)}, 2: {4: make(chan int), 5: make(chan int), 6: make(chan int)}}, //nolint:lll
			expected: `map[int]map[int]chan int`,
		},
		{
			name:     `Map with int keys and map of functions values`,
			input:    map[int]map[int]func(){1: {1: func() {}, 2: func() {}}, 2: {3: func() {}, 4: func() {}}},
			expected: `map[int]map[int]func()`,
		},
		{
			name:     `Map with int keys and map of pointers to interfaces values`,
			input:    map[int]map[int]*interface{}{1: {1: new(interface{}), 2: new(interface{})}, 2: {3: new(interface{}), 4: new(interface{})}}, //nolint:lll
			expected: `map[int]map[int]*interface {}`,
		},
		{
			name:     `Map with int keys and map of pointers to channels values`,
			input:    map[int]map[int]*chan int{1: {1: new(chan int), 2: new(chan int)}, 2: {3: new(chan int), 4: new(chan int)}},
			expected: `map[int]map[int]*chan int`,
		},
		{
			name:     `Map with int keys and map of pointers to functions values`,
			input:    map[int]map[int]*func(){1: {1: new(func()), 2: new(func())}, 2: {3: new(func()), 4: new(func())}},
			expected: `map[int]map[int]*func()`,
		},
		{
			name:     `Map with int keys and map of slices values`,
			input:    map[int]map[int][]int{1: {1: {1, 2, 3}, 2: {4, 5, 6}}, 2: {3: {7, 8, 9}, 4: {10, 11, 12}}},
			expected: `map[int]map[int][]int`,
		},
		{
			name:     `Map with int keys and map of maps values`,
			input:    map[int]map[int]map[int]int{1: {1: {1: 1, 2: 2, 3: 3}, 2: {4: 4, 5: 5, 6: 6}}, 2: {3: {7: 7, 8: 8, 9: 9}, 4: {10: 10, 11: 11, 12: 12}}}, //nolint:lll
			expected: `map[int]map[int]map[int]int`,
		},
		{
			name:     `Map with int keys and map of pointers to slices values`,
			input:    map[int]map[int]*[]int{1: {1: new([]int), 2: new([]int)}, 2: {3: new([]int), 4: new([]int)}},
			expected: `map[int]map[int]*[]int`,
		},
		{
			name:     `Map with int keys and map of pointers to maps values`,
			input:    map[int]map[int]*map[int]int{1: {1: new(map[int]int), 2: new(map[int]int)}, 2: {3: new(map[int]int), 4: new(map[int]int)}}, //nolint:lll
			expected: `map[int]map[int]*map[int]int`,
		},
		{
			name:     `Map with int keys and map of slices of pointers values`,
			input:    map[int]map[int][]*int{1: {1: {new(int), new(int), new(int)}, 2: {new(int), new(int), new(int)}}, 2: {3: {new(int), new(int), new(int)}, 4: {new(int), new(int), new(int)}}}, //nolint:lll
			expected: `map[int]map[int][]*int`,
		},
		{
			name:     `Map with int keys and map of slices of interfaces values`,
			input:    map[int]map[int][]interface{}{1: {1: {1, `two`}, 2: {3, `four`}}, 2: {3: {5, `six`}, 4: {7, `eight`}}},
			expected: `map[int]map[int][]interface {}`,
		},
		{
			name:     `Map with int keys and map of slices of channels values`,
			input:    map[int]map[int][]chan int{1: {1: {make(chan int), make(chan int), make(chan int)}, 2: {make(chan int), make(chan int), make(chan int)}}, 2: {3: {make(chan int), make(chan int), make(chan int)}, 4: {make(chan int), make(chan int), make(chan int)}}}, //nolint:lll
			expected: `map[int]map[int][]chan int`,
		},
		{
			name:     `Map with int keys and map of slices of functions values`,
			input:    map[int]map[int][]func(){1: {1: {func() {}, func() {}}, 2: {func() {}, func() {}}}, 2: {3: {func() {}, func() {}}, 4: {func() {}, func() {}}}}, //nolint:lll
			expected: `map[int]map[int][]func()`,
		},
		{
			name:     `Map with int keys and map of slices of pointers to interfaces values`,
			input:    map[int]map[int][]*interface{}{1: {1: {new(interface{}), new(interface{}), new(interface{})}, 2: {new(interface{}), new(interface{}), new(interface{})}}, 2: {3: {new(interface{}), new(interface{}), new(interface{})}, 4: {new(interface{}), new(interface{}), new(interface{})}}}, //nolint:lll
			expected: `map[int]map[int][]*interface {}`,
		},
		{
			name:     `Map with int keys and map of slices of pointers to channels values`,
			input:    map[int]map[int][]*chan int{1: {1: {new(chan int), new(chan int), new(chan int)}, 2: {new(chan int), new(chan int), new(chan int)}}, 2: {3: {new(chan int), new(chan int), new(chan int)}, 4: {new(chan int), new(chan int), new(chan int)}}}, //nolint:lll
			expected: `map[int]map[int][]*chan int`,
		},
		{
			name:     `Map with int keys and map of slices of pointers to functions values`,
			input:    map[int]map[int][]*func(){1: {1: {new(func()), new(func())}, 2: {new(func()), new(func())}}, 2: {3: {new(func()), new(func())}, 4: {new(func()), new(func())}}}, //nolint:lll
			expected: `map[int]map[int][]*func()`,
		},
		{
			name:     `Map with int keys and map of maps of pointers values`,
			input:    map[int]map[int]map[int]*int{1: {1: {1: new(int), 2: new(int), 3: new(int)}, 2: {4: new(int), 5: new(int), 6: new(int)}}, 2: {3: {7: new(int), 8: new(int), 9: new(int)}, 4: {10: new(int), 11: new(int), 12: new(int)}}}, //nolint:lll
			expected: `map[int]map[int]map[int]*int`,
		},
	}

	for _, tc := range testCases {
		testCase := tc // capture range variable, for parallel tests
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			reflectedType := reflect.TypeOf(testCase.input)
			actual, err := htmldump.MapTypeName(reflectedType)
			require.NoError(t, err)
			require.Equal(t, testCase.expected, actual)
		})
	}
}

func TestSliceTypeName(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    interface{}
		expected string
	}{
		{
			name:     `[]int`,
			input:    []int{},
			expected: `[]int`,
		},
		{
			name:     `[]*int`,
			input:    []*int{},
			expected: `[]*int`,
		},
		{
			name:     `[]*[]int`,
			input:    []*[]int{},
			expected: `[]*[]int`,
		},
		{
			name:     `[]*[]*int`,
			input:    []*[]*int{},
			expected: `[]*[]*int`,
		},
		{
			name:     `[]*[]*[]int`,
			input:    []*[]*[]int{},
			expected: `[]*[]*[]int`,
		},
		{
			name:     `[]*[]*[]*[]int`,
			input:    []*[]*[]*[]int{},
			expected: `[]*[]*[]*[]int`,
		},
	}

	for _, tc := range testCases {
		testCase := tc // capture range variable, for parallel tests
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			reflectedType := reflect.TypeOf(testCase.input)
			actual, err := htmldump.SliceTypeName(reflectedType)
			require.NoError(t, err)
			require.Equal(t, testCase.expected, actual)
			if actual != testCase.expected {
				t.Errorf(`expected %s, got %s`, testCase.expected, actual)
			}
		})
	}
}
