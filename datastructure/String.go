package TestProject

func BruteForceStringMatch(s string, t string) int {
	s_len := len(s)
	t_len := len(t)
	for i := 0; i <= s_len-t_len; i++ {
		var j int
		for j < t_len && s[i+j] == t[j] {
			j++
		}
		if j == t_len {
			return i
		}
	}
	return -1
}
