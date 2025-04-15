package canplaceflowers

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	
	l := len(flowerbed)
	if l == 1 {
		return flowerbed[0] == 0 && n == 1
	}

	for i := 0; i < l; i++ {
		if flowerbed[i] != 0 {
			continue
		}

		if i == 0 {
			if flowerbed[i+1] != 0 {
				continue
			}
		} else if i == l-1 {
			if flowerbed[i-1] != 0 {
				continue
			}
		} else if flowerbed[i-1] != 0 || flowerbed[i+1] != 0 {
			continue
		}

		n--
		flowerbed[i] = 1

		if n == 0 {
			return true
		}
	}

	return false
}
