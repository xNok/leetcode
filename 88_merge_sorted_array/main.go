package mergesortedarray

func merge(nums1 []int, m int, nums2 []int, n int) {

	// we process the list backward until all elements of nums2 are place
	// in nums1, since there are padding 0 at the end of nums1 we
	// never have to bother about memorizing the current value
	for n != 0 {
		if m != 0 && nums1[m-1] > nums2[n-1] {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else {
			nums1[m+n-1] = nums2[n-1]
			n--
		}
	}

}
