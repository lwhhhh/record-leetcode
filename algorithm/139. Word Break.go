package algorithm

func wordBreak(s string, wordDict []string) bool {
	match := make([]bool, len(s)+1)
	match[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if match[j] && anyoneMatch(s[j:i], wordDict) {
				match[i] = true
				break
			}
		}
	}

	return match[len(s)]
}

func anyoneMatch(s string, ss []string) bool {
	for _, e := range ss {
		if e == s {
			//fmt.Println("func:", s)
			return true
		}
	}
	return false
}
