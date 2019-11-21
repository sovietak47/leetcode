package asteroidCollision

import (
	"math"
)


/*
算法说明：

由于小行星有向左向右移动的属性。因此只有向右移动的小行星是具有威胁的小行星。
因此，如果上来收到一组负数，那么这些都可以放心放过。
但一旦开始收到正数以后，这些正数就会威胁到后续的来者。因此，在第一次遇到正数时，就标记它的索引，然后后面如果出现冲突（收到正数后又收到负数），那么冲突范围一定是在标记点及以后的。

冲突有三种情况：
// -1 -2 -3 5    -5   在某个位置两败俱伤
// -1 -2 -3 5    -1   在某个位置被撞毁
// -1 -2 -3 5    -7    在某个位置撞毁对方，可以继续往下撞，这个里面有两种情况，一种是前面还有正数：这个可以直接忽略跳过，我们只在达成稳态的时候，才做处理。另一种就是前面一个正数都没有了，这个时候才需要真正做些处理。

分别针对这三种情况予以分别处理即可。
*/

func asteroidCollision(asteroids []int) []int {
	checker := newChecker()

	for _, asteroid := range asteroids {
		// -1 -2 -3 ………………
		if checker.isAsteroidNoNeedToCheck(asteroid) {
			checker.appendToResult(asteroid)
		}

		// -1 -2 -3 5 !!got a trouble
		// -1 -2 -3 5 7 ………………
		if isRightMoveAsteroid(asteroid) {
			// -1 -2 -3 5 !!got a trouble
			if checker.isTroubleAsteroid(asteroid) {
				checker.markTroubleAsteroidIndex()
			}
			checker.appendToResult(asteroid)
		}

		// -1 -2 -3 5    -5
		// -1 -2 -3 5    -1
		// -1 -2 -3 5    -7
		if checker.isAsteroidNeedToCheck(asteroid) {
			checker.computeCollisionResult(asteroid)
		}
	}
	return checker.result
}

func isRightMoveAsteroid(asteroid int) bool {
	return asteroid > 0
}

type checker struct {
	troubleAsteroidIndex int // first right move asteroid Index
	result []int
}

func newChecker() *checker {
	return &checker{
		troubleAsteroidIndex:-1,
		result:[]int{},
	}
}

func (c *checker)isTroubleAsteroidHasBeenMarked() bool {
	return c.troubleAsteroidIndex == -1
}

func (c *checker)markTroubleAsteroidIndex(){
	c.troubleAsteroidIndex = len(c.result)
}

func (c *checker)cleanTroubleAsteroidIndex(){
	c.troubleAsteroidIndex = -1
}

func (c *checker)isCurrentAsteroidIsTroubleAsteroid(currentAsteroidIndex int) bool {
	return c.troubleAsteroidIndex == currentAsteroidIndex
}

func (c *checker)isAsteroidNoNeedToCheck(asteroid int) bool {
	return asteroid < 0 && c.isTroubleAsteroidHasBeenMarked()
}

func (c *checker)isAsteroidNeedToCheck(asteroid int) bool {
	return asteroid < 0 && !c.isTroubleAsteroidHasBeenMarked()
}

func (c *checker)isTroubleAsteroid(asteroid int) bool {
	return isRightMoveAsteroid(asteroid) && c.isTroubleAsteroidHasBeenMarked()
}

func (c *checker)appendToResult(asteroid int){
	c.result = append(c.result, asteroid)
}
func (c *checker)deleteResultFromIndex(index int){
	c.result = c.result[:index]
}

func (c *checker)computeCollisionResult(asteroid int){
	for j := len(c.result)-1; j >= c.troubleAsteroidIndex; j-- {
		// -1 -2 -3 5    -5
		if c.areBothAsteroidsBeenDestroyed(j, asteroid) {
			c.deleteResultFromIndex(j)
			if c.isCurrentAsteroidIsTroubleAsteroid(j) {
				c.cleanTroubleAsteroidIndex()
			}
			break
			// -1 -2 -3 5    -1
		}else if c.isRightMoveAsteroidSurvive(j, asteroid) {
			c.deleteResultFromIndex(j+1)
			break
			// -1 -2 -3 5    -7
		}else if c.isLastRightMoveAsteroidBeenDestroyed(j, asteroid) {
			c.deleteResultFromIndex(c.troubleAsteroidIndex)
			c.appendToResult(asteroid)
			c.cleanTroubleAsteroidIndex()
			break
		}
		// pass the rightMoveAsteroidBeenDestroyed situation, we only care about the last stable situation
	}
}

func (c *checker)areBothAsteroidsBeenDestroyed(rightMoveAsteroidIndex, leftMoveAsteroid int) bool {
	return math.Abs(float64(c.result[rightMoveAsteroidIndex])) == math.Abs(float64(leftMoveAsteroid))
}

func (c *checker)isRightMoveAsteroidSurvive(rightMoveAsteroidIndex, leftMoveAsteroid int) bool {
	return math.Abs(float64(c.result[rightMoveAsteroidIndex])) > math.Abs(float64(leftMoveAsteroid))
}

func (c *checker)isLastRightMoveAsteroidBeenDestroyed(rightMoveAsteroidIndex int, leftMoveAsteroid int) bool {
	return c.isCurrentAsteroidIsTroubleAsteroid(rightMoveAsteroidIndex) && math.Abs(float64(c.result[rightMoveAsteroidIndex])) < math.Abs(float64(leftMoveAsteroid))
}


