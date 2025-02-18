package countingsort

func sortColors(nums []int) {
	const (
		red   = 0
		white = 1
		blue  = 2
	)

	r, w, b := 0, 0, 0

	for i := 0; i < len(nums); i++ {
		switch nums[i] {
		case red:
			r++
		case white:
			w++
		case blue:
			b++
		}
	}

	for i := 0; i < r; i++ {
		nums[i] = red
	}
	for i := r; i < w+r; i++ {
		nums[i] = white
	}
	for i := r + w; i < w+r+b; i++ {
		nums[i] = blue
	}
}
