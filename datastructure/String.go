package TestProject

type MyString string

func BruteForceStringMatch(s MyString, t MyString) int {
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

//KMP
//abaabbabaab
func BuildFArray(s MyString) []int {
	F := make([]int, len(s))
	F[0] = -1
	for i := 1; i < len(s); i++ {
		var j = F[i-1]
		for s[j+1] != s[i] && j >= 0 {
			j = F[j]
		}
		if s[j+1] == s[i] {
			F[i] = j + 1
		} else {
			F[i] = -1
		}
	}
	return F
}

func (s MyString) IndexKMP(b MyString) int {
	var i, j = 0, -1
	f := BuildFArray(b)
	for ; i < len(s); i++ {
		for j > -1 && s[i] != b[j+1] {
			j = f[j]
		}
		if s[i] == b[j+1] {
			j++
		}
		if j == len(b)-1 {
			return i - len(b) + 1
		}
	}
	return -1
}
