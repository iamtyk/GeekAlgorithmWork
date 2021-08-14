package week8

func findAnagrams(s string, p string) []int {
	pm := map[byte]int{}
	for i := 0; i < len(p); i++ {
		pm[p[i]]++
	}
	sm := map[byte]int{}
	var match int
	var ans []int
	for l, r := 0, 0; r < len(s); r++ {
		sm[s[r]]++
		if _, ok := pm[s[r]]; ok {
			if pm[s[r]] == sm[s[r]] {
				match++
			}
			if match == len(pm) {
				ans = append(ans, l)
			}
		}
		for r-l+1 == len(p) {
			if _, ok := pm[s[l]]; ok && sm[s[l]] == pm[s[l]] {
				match--
			}
			sm[s[l]]--
			l++
		}
	}
	return ans
}
