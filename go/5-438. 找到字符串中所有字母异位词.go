package main

func findAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}

	count := make([]int, 26, 26)
	for i, ch := range p {
		count[s[i]-'a']++
		count[ch-'a']--
	}

	differ := 0
	for _, v := range count {
		if v != 0 {
			differ++
		}
	}

	res := make([]int, 0)
	if differ == 0 {
		res = append(res, 0)
	}

	left := 0
	right := len(p) - 1

	for right < len(s)-1 {
		if count[s[left]-'a'] == 0 {
			differ++
		} else if count[s[left]-'a'] == 1 {
			differ--
		}

		count[s[left]-'a']--

		if count[s[right+1]-'a'] == 0 {
			differ++
		} else if count[s[right+1]-'a'] == -1 {
			differ--
		}

		count[s[right+1]-'a']++

		right++
		left++

		if differ == 0 {
			res = append(res, left)
		}
	}

	return res
}
