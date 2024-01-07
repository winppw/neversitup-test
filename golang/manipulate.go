package main

func Manipulate(text []string) []string {
	// TODO : start your code here
	if len(text) <= 1 {
		if len(text) == 0 {
			return []string{}
		}
		return []string{text[0]}
	}

	permMap := make(map[string]bool)
	var result []string

	for i, char := range text {
		remainingChars := append(append([]string{}, text[:i]...), text[i+1:]...)
		for _, permutation := range Manipulate(remainingChars) {
			combined := char + permutation
			if _, exists := permMap[combined]; !exists {
				result = append(result, combined)
				permMap[combined] = true
			}
		}
	}

	return result
}
