package gofaker

import (
	"math/rand"
	"sort"
)

// Choice is a generic wrapper that can be used to add weights for any object
type Choice struct {
	Item   interface{}
	Weight uint
}

// A Chooser caches many possible Choices in a structure designed to improve
// performance on repeated calls for weighted random selection.
type Chooser struct {
	data   []Choice
	totals []int
	max    int
}

// NewChooser initializes a new Chooser consisting of the possible Choices.
func NewChooser(cs ...Choice) Chooser {
	sort.Slice(cs, func(i, j int) bool {
		return cs[i].Weight < cs[j].Weight
	}) // 核心点，对元素进行递增排序。
	totals := make([]int, len(cs))
	runningTotal := 0
	for i, c := range cs {
		runningTotal += int(c.Weight)
		totals[i] = runningTotal
	}
	return Chooser{data: cs, totals: totals, max: runningTotal}
}

// Pick returns a single weighted random Choice.Item from the Chooser.
func (chs Chooser) Pick() interface{} {
	r := rand.Intn(chs.max) + 1         // 使用最大值获取随机数，避免超过范围，随机生成的数需要排除0，故加1
	i := sort.SearchInts(chs.totals, r) // 核心点该方法使用二分法，找到对应的下标，如果没有则为大于该数的+1 下标，可能为len(a)即数组长度。
	return chs.data[i].Item
}
