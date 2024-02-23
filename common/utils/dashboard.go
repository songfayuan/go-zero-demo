package utils

import "go-zero-demo/common/types"

// CockTailSort 鸡尾酒排序
func CockTailSort(array []*types.Pie) {
	for i := 0; i < len(array)/2; i++ { //外层循环控制轮次
		isSorted := true //是否有序的标记
		//先进行从左到右的排序
		for j := i; j < len(array)-1-i; j++ { //内层循环开始两两元素进行比较，即控制比较的元素边界
			if array[j].Value < array[j+1].Value {
				temp := array[j]
				array[j] = array[j+1]
				array[j+1] = temp

				isSorted = false
			}
		}
		if isSorted { //如果已经有序则，退出排序
			break
		}

		//再进行从右到左的排序
		isSorted = true                           //排序之前将isSorted重置
		for j := len(array) - 1 - i; j > i; j-- { //内层循环开始两两元素进行比较，即控制比较的元素边界
			if array[j].Value > array[j-1].Value {
				temp := array[j]
				array[j] = array[j-1]
				array[j-1] = temp
				isSorted = false
			}
		}
		if isSorted { //如果已经有序则，退出排序
			break
		}
	}
}

// CockTailSortByRequestNum 鸡尾酒排序
func CockTailSortByRequestNum(array []*types.Analysis) {
	for i := 0; i < len(array)/2; i++ { //外层循环控制轮次
		isSorted := true //是否有序的标记
		//先进行从左到右的排序
		for j := i; j < len(array)-1-i; j++ { //内层循环开始两两元素进行比较，即控制比较的元素边界
			if array[j].Request < array[j+1].Request {
				temp := array[j]
				array[j] = array[j+1]
				array[j+1] = temp

				isSorted = false
			}
		}
		if isSorted { //如果已经有序则，退出排序
			break
		}

		//再进行从右到左的排序
		isSorted = true                           //排序之前将isSorted重置
		for j := len(array) - 1 - i; j > i; j-- { //内层循环开始两两元素进行比较，即控制比较的元素边界
			if array[j].Request > array[j-1].Request {
				temp := array[j]
				array[j] = array[j-1]
				array[j-1] = temp
				isSorted = false
			}
		}
		if isSorted { //如果已经有序则，退出排序
			break
		}
	}
}

// ClassLevelTop 前几位
func ClassLevelTop(StructArray, UnStructArray []*types.Pie, Limit int) (Struct, UnStruct []*types.Pie) {
	tempStruct := make([]*types.Pie, 0)
	tempUnStruct := make([]*types.Pie, 0)

	var StructValue int64 = 0
	var UnStructValue int64 = 0
	for i := 0; i < len(StructArray); i++ {
		if i < Limit {
			tempStruct = append(tempStruct, StructArray[i])
		} else {
			StructValue += StructArray[i].Value
		}
	}
	for i := 0; i < len(UnStructArray); i++ {
		if i < Limit {
			tempUnStruct = append(tempUnStruct, UnStructArray[i])
		} else {
			UnStructValue += UnStructArray[i].Value
		}
	}
	// 结构化
	StructOther := &types.Pie{
		Name:  "其他",
		Value: StructValue,
	}

	// 非结构化
	UnStructOther := &types.Pie{
		Name:  "其他",
		Value: UnStructValue,
	}
	tempStruct = append(tempStruct, StructOther)
	tempUnStruct = append(tempUnStruct, UnStructOther)

	return tempStruct, tempUnStruct
}
