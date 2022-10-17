package checkifthesentenceispangram

func CheckIfPangram(sentence string) bool {
	appearArray := make([]int, 26)
	AppearCount := 26

	for _, b := range sentence {
		if appearArray[int(b-'a')] == 0 {
			appearArray[int(b-'a')] = 1
			AppearCount--
		}

		if AppearCount == 0 {
			return true
		}
	}

	return AppearCount == 0
}
