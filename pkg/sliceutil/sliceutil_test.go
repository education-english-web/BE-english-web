package sliceutil

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRemoveDuplicatedItems(t *testing.T) {
	type args struct {
		numbers []uint64
	}

	tests := []struct {
		name string
		args args
		want []uint64
	}{
		{
			name: "success",
			args: args{
				numbers: []uint64{1, 2, 3, 2, 123, 6, 5, 3, 9},
			},
			want: []uint64{1, 2, 3, 123, 6, 5, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RemoveDuplicatedItems(tt.args.numbers)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestConvert(t *testing.T) {
	type args struct {
		numbers []uint64
	}

	tests := []struct {
		name string
		args args
		want []int64
	}{
		{
			name: "success",
			args: args{
				numbers: []uint64{1, 2, 3},
			},
			want: []int64{1, 2, 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Convert[uint64, int64](tt.args.numbers)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestConvertToString(t *testing.T) {
	type args struct {
		s []int64
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "success",
			args: args{
				s: []int64{-1, 2, 0},
			},
			want: []string{"-1", "2", "0"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConvertToString(tt.args.s)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUnique(t *testing.T) {
	type args[T comparable] struct {
		slice []T
	}

	type testCase[T comparable] struct {
		name string
		args args[T]
		want []T
	}

	testsForInt := []testCase[int]{
		{
			name: "has int duplication",
			args: args[int]{[]int{1, 2, 3, 4, 1, 2, 5, 7, 4}},
			want: []int{1, 2, 3, 4, 5, 7},
		},
		{
			name: "no int duplication",
			args: args[int]{[]int{4, 1, 2, 5, 7}},
			want: []int{4, 1, 2, 5, 7},
		},
	}

	for _, tt := range testsForInt {
		t.Run(tt.name, func(t *testing.T) {
			if got := Unique(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unique() = %v, want %v", got, tt.want)
			}
		})
	}

	testsForString := []testCase[string]{
		{
			name: "has string duplication",
			args: args[string]{[]string{"a", "c", "b", "c", "d"}},
			want: []string{"a", "c", "b", "d"},
		},
		{
			name: "no string duplication",
			args: args[string]{[]string{"a", "c", "b", "d"}},
			want: []string{"a", "c", "b", "d"},
		},
	}

	for _, tt := range testsForString {
		t.Run(tt.name, func(t *testing.T) {
			if got := Unique(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Unique() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDifference(t *testing.T) {
	type testCase[T comparable] struct {
		name    string
		s1      []T
		s2      []T
		diffS12 []T
		diffS21 []T
	}

	tests := []testCase[uint32]{
		{
			name:    "diffs12 is empty, diffs21 is empty",
			s1:      []uint32{1, 2, 3},
			s2:      []uint32{1, 2, 3},
			diffS21: []uint32{},
			diffS12: []uint32{},
		},
		{
			name:    "s1 contains s2",
			s1:      []uint32{1, 2, 3},
			s2:      []uint32{2, 3},
			diffS12: []uint32{1},
			diffS21: []uint32{},
		},
		{
			name:    "s2 contains s1",
			s1:      []uint32{2, 3},
			s2:      []uint32{1, 2, 3},
			diffS12: []uint32{},
			diffS21: []uint32{1},
		},
		{
			name:    "s1 intersects s2",
			s1:      []uint32{2, 3, 4},
			s2:      []uint32{1, 2, 3},
			diffS12: []uint32{4},
			diffS21: []uint32{1},
		},
		{
			name:    "s1 not intersects s2",
			s1:      []uint32{4, 5, 6},
			s2:      []uint32{1, 2, 3},
			diffS12: []uint32{4, 5, 6},
			diffS21: []uint32{1, 2, 3},
		},
		{
			name:    "s1 is empty",
			s1:      []uint32{},
			s2:      []uint32{1, 2, 3},
			diffS12: []uint32{},
			diffS21: []uint32{1, 2, 3},
		},
		{
			name:    "s2 is empty",
			s1:      []uint32{1, 2, 3},
			s2:      []uint32{},
			diffS12: []uint32{1, 2, 3},
			diffS21: []uint32{},
		},
	}

	for i := range tests {
		t.Run(tests[i].name, func(t *testing.T) {
			gotDiffS12, gotDiffS21 := Difference(tests[i].s1, tests[i].s2)

			assert.Equal(t, tests[i].diffS12, gotDiffS12)
			assert.Equal(t, tests[i].diffS21, gotDiffS21)
		})
	}
}

func TestIntersect(t *testing.T) {
	type testCase[T comparable] struct {
		name   string
		s1     []T
		s2     []T
		result []T
	}

	tests := []testCase[uint32]{
		{
			name:   "no intersect",
			s1:     []uint32{1, 2, 3},
			s2:     []uint32{4, 5, 6},
			result: []uint32{},
		},
		{
			name:   "partial intersect",
			s1:     []uint32{1, 2, 3},
			s2:     []uint32{2, 3, 4},
			result: []uint32{2, 3},
		},
		{
			name:   "full intersect",
			s1:     []uint32{1, 2, 3},
			s2:     []uint32{1, 2, 3, 4},
			result: []uint32{1, 2, 3},
		},
	}

	for i := range tests {
		t.Run(tests[i].name, func(t *testing.T) {
			got := Intersect(tests[i].s1, tests[i].s2)

			assert.Equal(t, tests[i].result, got)
		})
	}
}

func TestConvertStringToInt(t *testing.T) {
	type testCase[T comparable] struct {
		name   string
		input  []string
		result []T
		errStr string
	}

	testUint32s := []testCase[uint32]{
		{
			name:   "success - uint32",
			input:  []string{"1", "2", "3"},
			result: []uint32{1, 2, 3},
			errStr: "",
		},
		{
			name:   "fail - input can not be parsed",
			input:  []string{"choi oi"},
			result: []uint32(nil),
			errStr: "strconv.Atoi: parsing \"choi oi\": invalid syntax",
		},
	}

	testUint64s := []testCase[uint64]{
		{
			name:   "success- uint64",
			input:  []string{"4", "5", "6"},
			result: []uint64{4, 5, 6},
			errStr: "",
		},
	}

	for i := range testUint32s {
		t.Run(testUint32s[i].name, func(t *testing.T) {
			got, gotErr := ConvertStringToInt[uint32](testUint32s[i].input)

			assert.Equal(t, got, testUint32s[i].result)

			if testUint32s[i].errStr != "" {
				require.ErrorContains(t, gotErr, testUint32s[i].errStr)
			} else {
				require.NoError(t, gotErr)
			}
		})
	}

	for i := range testUint64s {
		t.Run(testUint64s[i].name, func(t *testing.T) {
			got, gotErr := ConvertStringToInt[uint64](testUint64s[i].input)

			assert.Equal(t, got, testUint64s[i].result)

			if testUint64s[i].errStr != "" {
				require.ErrorContains(t, gotErr, testUint64s[i].errStr)
			} else {
				require.NoError(t, gotErr)
			}
		})
	}
}
