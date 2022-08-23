package sol

func checkInclusion(s1 string, s2 string) bool {
	s1Len, s2Len := len(s1), len(s2)
	if s1Len > s2Len {
		return false
	}
	left, freq, count := 0, make([]int, 26), s1Len
	for idx := 0; idx < s1Len; idx++ {
		freq[s1[idx]-'a']++
	}
	for right := 0; right < s2Len; right++ {
		if freq[s2[right]-'a'] > 0 {
			count--
		}
		freq[s2[right]-'a']--
		if count == 0 {
			return true
		}
		if right-left+1 >= s1Len { // slide-window out of range and not found
			if freq[s2[left]-'a'] >= 0 {
				count++
			}
			freq[s2[left]-'a']++
			left++
		}
	}
	return false
}
