package main

import "regexp"

func CountSmilyFace(text []string) int {
	// TODO : start your code here
	count := 0
	validSmileyRegex := regexp.MustCompile(`^[:;][-~]?[)D]$`)

	for _, face := range text {
		if validSmileyRegex.MatchString(face) {
			count++
		}
	}

	return count
}
