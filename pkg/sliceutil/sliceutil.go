package sliceutil

import "strconv"

func RemoveDuplicatedItems[T comparable](input []T) []T {
	encountered := map[T]struct{}{}
	result := make([]T, 0, len(input))

	for _, v := range input {
		if _, ok := encountered[v]; !ok {
			encountered[v] = struct{}{}

			result = append(result, v)
		}
	}

	return result
}

func Convert[T1, T2 uint32 | uint64 | int32 | int64](input []T1) []T2 {
	output := make([]T2, len(input))
	for i, v := range input {
		output[i] = T2(v)
	}

	return output
}

func ConvertToString[S ~[]T, T uint32 | uint64 | int32 | int64](input S) []string {
	result := make([]string, len(input))
	for i, v := range input {
		result[i] = strconv.Itoa(int(v))
	}

	return result
}

func ConvertStringToInt[T uint32 | uint64 | int32 | int64](input []string) ([]T, error) {
	result := make([]T, len(input))

	for i := range input {
		value, err := strconv.Atoi(input[i])
		if err != nil {
			return nil, err
		}

		result[i] = T(value)
	}

	return result, nil
}

/*
Unique

To return a duplicate-free version of a slice
*/
func Unique[T comparable](slice []T) []T {
	var (
		existingT = make(map[T]bool, len(slice))
		result    = make([]T, 0, len(slice))
	)

	for i := range slice {
		if _, ok := existingT[slice[i]]; !ok {
			existingT[slice[i]] = true

			result = append(result, slice[i])
		}
	}

	return result
}

/*
Difference

To return two slices containing difference between s1 & s2
diffS12 = s1 - s2
diffS21 = s2 - s1
*/
func Difference[T comparable](s1, s2 []T) ([]T, []T) {
	var (
		mS1 = make(map[T]bool, len(s1))
		mS2 = make(map[T]bool, len(s2))

		diffS12 = make([]T, 0)
		diffS21 = make([]T, 0)
	)

	for i := range s1 {
		mS1[s1[i]] = true
	}

	for i := range s2 {
		mS2[s2[i]] = true

		if !mS1[s2[i]] {
			diffS21 = append(diffS21, s2[i])
		}
	}

	for i := range s1 {
		if !mS2[s1[i]] {
			diffS12 = append(diffS12, s1[i])
		}
	}

	return diffS12, diffS21
}

/*
Intersect

To return the intersection between s1 & s2
*/
func Intersect[T comparable](s1, s2 []T) []T {
	var (
		result = make([]T, 0, len(s1))
		mS1    = make(map[T]bool, len(s1))
	)

	for i := range s1 {
		mS1[s1[i]] = true
	}

	for i := range s2 {
		if mS1[s2[i]] {
			result = append(result, s2[i])
		}
	}

	return result
}
