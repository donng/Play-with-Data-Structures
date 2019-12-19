package solution

func firstUniqChar(s string) int {
	var freq [26]int

	for _, v := range s {
		freq[v-'a']++
	}

	for k, v := range s {
		if freq[v-'a'] == 1 {
			return k
		}
	}

	return -1
}
