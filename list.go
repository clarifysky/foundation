package foundation

func Sort(arr []int) []int {
	var i, j, final int
	var res []int
	for i = 0; i < len(arr); i++ {
		final = arr[i]
		for j = i + 1; j < len(arr); j++ {
			if final > arr[j] {
				tmp := final
				final = arr[j]
				arr[j] = tmp
			}
		}
		res = append(res, final)
	}
	return res
}

