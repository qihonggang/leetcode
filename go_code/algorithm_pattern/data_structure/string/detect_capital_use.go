package string

func detectCapitalUse(word string) bool {
	if word[0] <= 'Z' {
		for i := 2; i < len(word); i++ {
			if !((word[i] <= 'Z' && word[i-1] <= 'Z') || (word[i] > 'Z' && word[i-1] > 'Z')) {
				return false
			}
		}
		return true
	} else {
		for i := 1; i < len(word); i++ {
			if word[i] <= 'Z' {
				return false
			}
		}
		return true
	}
}
