package baitap01

/* Tìm số lớn thứ 2 trong dãy số
 */
func Max2Numbers(numbers []float64) (max2 float64) {
	var max1 = numbers[0]
	max2 = numbers[0]
	for _, value := range numbers {
		if max1 < value {
			max2 = max1
			max1 = value
		}
	}
	return max2
}
