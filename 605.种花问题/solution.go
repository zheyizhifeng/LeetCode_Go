// @algorithm @lc id=605 lang=golang
// @title can-place-flowers

package main

// @test([1,0,0,0,1],1)=true
// @test([1,0,0,0,1],2)=false
func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	length := len(flowerbed)
	for i, planted := range flowerbed {
		if planted == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == length-1 || flowerbed[i+1] == 0) {
			n--
			flowerbed[i] = 1 // * 种花
			if n <= 0 {
				return true
			}
		}
	}
	return false
}
/* 
func canPlaceFlowers(flowerbed []int, n int) bool {
	flo := []int{}
	flo = append(flo, 0)
	flo = append(flo, flowerbed...)
	flo = append(flo, 0)

	for i := 1; i < len(flo) - 1; i++ {
			if n == 0 {
					return true
			}

			if flo[i] == 0 && flo[i-1] == 0 && flo[i+1] == 0 {
					flo[i] = 1
					n--
			}
	}

	return n == 0
}
*/