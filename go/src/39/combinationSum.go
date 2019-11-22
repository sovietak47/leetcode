package combinationSum

import "sort"

/*

算法说明：

以输入: candidates = [2,3,6,7], target = 7,
所求解集为:
[
  [7],
  [2,2,3]
]
为例.

由于求解集一定是若干个candidates相加得到的.因此,反过来想,我们可以将target用各个candidates分解,能形成通路最终成为解集的分支上的分解值一定是candidates内的成员.
因此f(target) = f(candidate) + f(subTarget).

以7为例,可以使
7=7,
7=6+1,(X)
7=3+4,4=3+1,(X)
7=3+4,4=2+2,2=2
7=2+5,5=3+2,2=2
7=2+5,5=2+3,3=3
7=2+5,5=2+3,3=2+1,(X)
即,如果不能被减到0,且余数小于所有candidates,即宣布此路分支死亡.
如果能减到0,即宣布此路分支走通.

从而得出:
[
  [7],
  [3,2,2],
  [2,3,2],
  [2,2,3]
]

由于后三种情况被视为是同一种情况,需要合并.因此,可以考虑用计数法（7多少次,6多少次,3多少次,2多少次）,而不是记录路径的方式,使后三种情况可被识别并合并.但这终归还是有成本.最好还是能尽早识别重复并停止计算.

由于诸如[2,2,3]这种多种排列组合的情况,其从大到小的顺序排序序列,一定是唯一的.因此,可以要求每一分支的分解值必须从大到小顺序分解.从而避免排列组合带来的多分支发散.
还是以7为例,以各candidates分别肢解target时,
第一分支：7=7,（后续可选分解值：7,6,3,2）
第二分支：7=6+1,（后续可选分解值：6,3,2）
第三分支：7=3+4,（后续可选分解值：3,2）
4=3+1,(X)
4=2+2,（后续可选分解值：2）,2=2
第四分支：7=2+5,（后续可选分解值：2）,5=2+3,（后续可选分解值：2）,3=2+1,（后续可选分解值：2）,(X)

因此,可以通过:
1.target用各个candidates分多个分支分别分解,直到其彻底分解到0.
2.每一分支的分解值必须从大到小,顺序分解.
两个条件联合协同,做到最优分解效果.

处理逻辑可以概括为:
	对于每一个candidate:
		如果 target < candidate
			尝试下一个candidate

		如果 target > candidate
			由公式 f(target) = f(candidate) + f(subTarget). 递归求解subTarget, 如果返回值不为空(应该是一组数组序列), 则在返回值每个序列前加此次的candidate, 添入最后结果集.
			否则, 意味着这个分支走不通, 不做任何处理.

		如果 target = candidate
			意味着这是一条通路,此次的candidate是通路上的最后一个值,添入最后结果集即可.

	结果集在遍历结束后返回即可,由于可以重复取值,多个更小值拼一个通路是客观存在的,因此上述分支如果在某轮循环中求得了通路,均有可能在下一轮循环中通过更小值再次求得通路.因此不要在上述任何分支中return中断程序,除非遍历结束.
*/

func combinationSum(candidates []int, target int) [][]int {
	var result [][]int

	sort.Ints(candidates)

	for i := len(candidates) - 1; i >= 0; i-- {
		var tmp []int

		if target < candidates[i] {
			continue
		}

		subTarget := target - candidates[i]
		if subTarget > 0 {
			subResult := combinationSum(candidates[:i+1], subTarget)
			if subResult != nil {
				for _, sub := range subResult {
					tmp = append(append(tmp, candidates[i]), sub...)
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
