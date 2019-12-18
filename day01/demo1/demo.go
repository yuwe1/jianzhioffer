package demo1

type S1 []int
type S2 []S1

func FindNum(data S2, row, col, num int) bool {
	if len(data) == 0 || row <= 0 || col <= 0 {
		return false
	}
	//
	i := 0
	j := col - 1
	for {
		if i < 0 || j < 0 {
			break
		}
		if data[i][j] > num {
			j--
		} else if data[i][j] < num {
			i++
		} else {
			return true
		}
	}
	return false
}
