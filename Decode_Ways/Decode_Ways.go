package decodeways

func NumDecodings(s string) int {

	if s[0] == '0' {
		return 0
	}

	oneword, twoword := 1, 1

	for i := 1; i < len(s); i++ {
		count := 0
		if s[i] != '0' {
			count = oneword
		}

		if s[i-1:i+1] >= "10" && s[i-1:i+1] <= "26" {
			count += twoword
		}

		twoword = oneword
		oneword = count
	}

	return oneword
}
