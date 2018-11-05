package util

/*返回item的页数*/
func GetPageNum(totalNum int, num int) int {
	if totalNum%num == 0 {
		return totalNum / num
	} else {
		return totalNum/num + 1
	}
}
