package leetcode

// version 1. time limt exceed
func countGoodTriplets(arr []int, a int, b int, c int) int {
	var (
		count  int
		length = len(arr)
	)

	for i := 0; i < length-3; i++ {
		for j := i + 1; i < length-2; j++ {
			for k := j + 1; j < length-1; j++ {
				if AbsInt(arr[i]-arr[j]) > a || AbsInt(arr[j]-arr[k]) > b || AbsInt(arr[i]-arr[k]) > c {
					continue
				}
				count += 1
			}
		}
	}
	return count
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func countGoodTripletsV2(arr []int, a int, b int, c int) int {
	var (
		count        int
		length       = len(arr)
		temp1, temp2 [][]int
	)

	for i := 0; i < length-3; i++ {
		for j := i + 1; j < length-2; j++ {
			if AbsInt(arr[i]-arr[j]) <= a {
				temp1 = append(temp1, []int{i, j})
			}
		}
	}

	for k := length - 1; k > 1; k-- {
		for j := k - 1; j >= 0; j-- {
			if AbsInt(arr[k]-arr[j]) <= b {
				temp2 = append(temp2, []int{j, k})
			}
		}
	}

	for _, dm1 := range temp1 {
		for _, dm2 := range temp2 {
			i := dm1[0]
			j1 := dm1[1]
			j2 := dm2[0]
			k := dm2[1]
			if j1 == j2 && AbsInt(arr[i]-arr[k]) <= c {
				count += 1
			}
		}
	}
	return count
}
