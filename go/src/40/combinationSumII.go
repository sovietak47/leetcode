package combinationSumII

import "sort"

/*

算法说明：

还是使用"求解集一定是若干个candidates相加得到的"的思路,f(target) = f(candidate) + f(subTarget).由于candidates 中的每个数字在每个组合中只能使用一次,这反而降低了难度.
只需要将candidates从大到小依次试减即可.

以输入: candidates = [10,1,2,7,6,1,5], target = 8,
所求解集为:
[
  [1, 7],
  [1, 2, 5],
  [2, 6],
  [1, 1, 6]
]
为例.

可以使
8=10-2,(X)
8=7+1,1=1
8=6+2,2=2
8=6+2,2=1+1,1=1
8=5+3,3=2+1,1=1
8=2+6,6=1+5,5=1+4,(数不够用了,X)

可以看出,这个和39基本没有太大差别,无非就是递归的时候传给下层的参数里,是否包含正在使用的candidate(传递表示可以重复使用,不传递表示不可重复使用)
唯一的区别是:
	由于candidate可以重复,在循环中多加一个重复剔除步骤

*/

func combinationSum2(candidates []int, target int) [][]int {
	var result [][]int

	sort.Ints(candidates)

	for i:= len(candidates)-1; i>=0; i--{
		var tmp []int

		if i < len(candidates)-1 && candidates[i] == candidates[i+1] {
			continue
		}

		if target < candidates[i] {
			continue
		}

		subTarget := target - candidates[i]
		if subTarget > 0 {
			subResult := combinationSum2(candidates[:i], subTarget)
			if subResult != nil {
				for _, sub := range subResult {
					tmp = append(append(tmp, candidates[i]),sub...)
					result = append(result, tmp)
					tmp = []int{}
				}
			}
		}

		if subTarget == 0 {
			tmp = append(tmp, candidates[i])
			result = append(result, tmp)
		}
	}

	return result
}
