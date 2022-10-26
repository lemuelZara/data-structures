package arrays

func Merge(nums1, nums2 []int) []int {
	i, j, k := 0, 0, 0

	r := make([]int, len(nums1)+len(nums2))

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			r[k] = nums1[i]
			i++
			k++
		} else {
			r[k] = nums2[j]
			j++
			k++
		}
	}

	for i < len(nums1) {
		r[k] = nums1[i]
		i++
		k++
	}

	for j < len(nums2) {
		r[k] = nums2[j]
		j++
		k++
	}

	return r
}
