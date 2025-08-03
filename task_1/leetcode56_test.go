package task1

/**
合并区间：以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。

可以先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，
将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中。

示例 1：

输入：intervals = [[1,3],[2,6],[8,10],[15,18]]
输出：[[1,6],[8,10],[15,18]]
解释：区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].

思路
1.​排序​：首先将所有区间按照起始点 starti 进行排序，这样可以确保重叠的区间是连续的。
​2.合并​：遍历排序后的区间，逐个检查当前区间是否可以与前一个合并的区间合并：
	如果当前区间的起始点 starti 小于等于前一个合并区间的结束点 endi，则合并这两个区间（取最大的结束点）。
	否则，将当前区间加入结果集。
**/

// merge 合并所有重叠的区间
// 参数：
//   intervals - 待合并的区间数组，每个区间表示为 [start, end]
// 返回值：
//   [][]int - 合并后的不重叠区间数组

import (
	"reflect"
	"sort"
	"testing"
)

func merge(intervals [][]int) [][]int {
	//判空
	if len(intervals) == 0 {
		return intervals
	}
	//按照区间的起始点给排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	//初始化结果集，先放入第一个区间
	merged := [][]int{intervals[0]}
	//循环后续其他区间
	for i := 1; i < len(intervals); i++ {
		//获取结果集的最后一个区间
		last := merged[len(merged)-1]
		//当前的循环区间
		current := intervals[i]
		//检查循环区间是否与结果集的最后一个区间重叠
		//重复的判断:当前循环区间的start 小于 结果集最后一个区间的 end
		if current[0] < last[1] {
			//合并区间，取end值最大的作为end
			if current[1] > last[1] {
				last[1] = current[1]
			}
		} else {
			//如果不重叠，则将当前循环区间，加入到结果集尾部
			merged = append(merged, current)
		}

	}

	return merged

}

func TestMerge(t *testing.T) {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	result := merge(intervals)
	t.Logf("merge结果: %v", result)
	// 添加一个实际断言
	expected := [][]int{{1, 6}, {8, 10}, {15, 18}}
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("期望得到 %v,但实际得到 %v", expected, result)
	}
}
