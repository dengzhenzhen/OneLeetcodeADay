package main

func countWords(words1 []string, words2 []string) int {
	wordCount := make(map[string]int8)
	for _, word := range words1 {
		_, exist := wordCount[word]
		if !exist {
			wordCount[word] = 1
		}else {
			wordCount[word]++
		}
	}
	for _, word := range words2 {
		count, exist := wordCount[word]
		if !exist || count > 1 {
			continue
		}
		wordCount[word]--
	}
	ret := 0
	for _, v := range wordCount {
		if v == 0 {
			ret++
		}
	}
	return ret
}
