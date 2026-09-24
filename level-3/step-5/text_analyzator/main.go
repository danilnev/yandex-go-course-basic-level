package main

import (
	"fmt"
	"strings"
)

func getWordStatisticFromText(text string) map[string]int {
	wordCounts := map[string]int{}
	currentWord := ""

	for _, char := range strings.ToLower(text) {
		if char == ' ' || char == ',' || char == '.' || char == '!' || char == '?' {
			if currentWord != "" {
				wordCounts[currentWord]++
				currentWord = ""
			}
		} else {
			currentWord += string(char)
		}
	}

	if currentWord != "" {
		wordCounts[currentWord]++
	}

	return wordCounts
}

func getTopWords(wordMap map[string]int, n int) []string {
	topWords := make([]string, n)
	wordMapCopy := map[string]int{}

	for word, count := range wordMap {
		wordMapCopy[word] = count
	}

	for i := 0; i < n; i++ {
		maxCount := 0
		topWord := ""

		for word, count := range wordMapCopy {
			if count > maxCount {
				maxCount = count
				topWord = word
			}
		}

		topWords[i] = topWord
		delete(wordMapCopy, topWord)
	}

	return topWords
}

func AnalyzeText(text string) {
	textStatistic := getWordStatisticFromText(text)

	wordsCount := 0
	for _, count := range textStatistic {
		wordsCount += count
	}

	uniqueWordsCount := len(textStatistic)

	maxCount := 0
	mostPopularWord := ""
	for word, count := range textStatistic {
		if count > maxCount {
			maxCount = count
			mostPopularWord = word
		}
	}

	resultString := fmt.Sprintf(
		"Количество слов: %d\n"+
			"Количество уникальных слов: %d\n"+
			"Самое часто встречающееся слово: \"%s\" (встречается %d раз)\n"+
			"Топ-5 самых часто встречающихся слов:\n",
		wordsCount, uniqueWordsCount, mostPopularWord, maxCount,
	)

	for _, word := range getTopWords(textStatistic, 5) {
		resultString += fmt.Sprintf("\"%s\": %d раз\n", word, textStatistic[word])
	}

	fmt.Print(resultString)
}

func main() {
	text := "first first! First, second, FIRST, a, 45.    !,.    seCond sadf"
	AnalyzeText(text)
}
