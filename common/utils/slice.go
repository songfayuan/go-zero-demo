package utils

import (
	"bytes"
	"sort"
	"strconv"

	"github.com/deckarep/golang-set/v2"
)

// SliceUnique 切片去重
func SliceUnique[T string | int | int64 | int32](from []T) []T {
	pos := make(map[T]struct{}, 0)
	for _, v := range from {
		pos[v] = struct{}{}
	}

	ret := make([]T, 0)
	for k := range pos {
		ret = append(ret, k)
	}

	return ret
}

// SliceUniqueSortToString 数组去重排序，然后转为字符串
func SliceUniqueSortToString[T int | int64 | int32](ids []T) []string {
	s := mapset.NewSet[int]()
	for _, v := range ids {
		s.Add(int(v))
	}

	arr := s.ToSlice()

	sort.Ints(arr)

	ret := make([]string, len(arr))
	for i, v := range arr {
		ret[i] = strconv.Itoa(v)
	}

	var buffer bytes.Buffer
	for _, val := range s.ToSlice() {
		buffer.WriteString(strconv.Itoa(val))
	}

	return ret
}

// SliceUniqueSort 数组去重排序
func SliceUniqueSort[T int | int64 | int32](ids []T) []int {
	s := mapset.NewSet[int]()
	for _, v := range ids {
		s.Add(int(v))
	}

	arr := s.ToSlice()

	sort.Ints(arr)

	return arr
}

// DiffSlice 返回两个切片的差集，slice1里出现但是slice2里没有出现的情况
func DiffSlice[T int | int32 | string | int64](slice1, slice2 []T) []T {
	// 创建一个 map 用于存储第二个切片中的元素
	exists := make(map[T]bool)
	for _, v := range slice2 {
		exists[v] = true
	}
	// 创建一个空切片用于存储差集结果
	var diff []T
	// 遍历第一个切片，将不在第二个切片中的元素加入到结果切片
	for _, v := range slice1 {
		if _, found := exists[v]; !found {
			diff = append(diff, v)
		}
	}

	return diff
}

// SplitArray 分隔数组
func SplitArray(arr []string, chunkSize int) [][]string {
	var result [][]string
	length := len(arr)

	for i := 0; i < length; i += chunkSize {
		end := i + chunkSize
		if end > length {
			end = length
		}
		result = append(result, arr[i:end])
	}

	return result
}
