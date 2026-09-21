package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type FAQ struct {
	Question string
	Answer   string
	Keywords []string
}

func main() {
	file, err := os.Open("faq.txt")
	if err != nil {
		fmt.Println("Error reading faq.txt:", err)
		return
	}
	defer file.Close()

	var faqs []FAQ
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "|")
		if len(parts) == 2 {
			faqs = append(faqs, FAQ{
				Question: parts[0],
				Answer:   parts[1],
				Keywords: strings.Fields(strings.ToLower(parts[0])),
			})
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading faq.txt:", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Привет! Задавай вопрос про репетицию (или 'exit' для выхода):")

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		input = strings.TrimSpace(input)
		if strings.ToLower(input) == "exit" || input == "" {
			break
		}

		bestScore := 0.0
		answer := "не знаю"

		userInputFields := strings.Fields(strings.ToLower(input))

		for _, faq := range faqs {
			currentScore := 0.0
			for _, kw := range faq.Keywords {
				cleanKw := strings.Trim(kw, "?.,!")
				if len(cleanKw) < 3 {
					continue
				}

				for _, ui := range userInputFields {
					cleanUi := strings.Trim(ui, "?.,!")
					if len(cleanUi) < 3 {
						continue
					}

					dist := Levenshtein(cleanKw, cleanUi)
					maxLen := math.Max(float64(len([]rune(cleanKw))), float64(len([]rune(cleanUi))))
					similarity := 1.0 - (float64(dist) / maxLen)

					// Only consider high similarity matches
					if similarity > 0.7 {
						if similarity > currentScore {
							currentScore = similarity
						}
					}
				}
			}

			if currentScore > bestScore {
				bestScore = currentScore
				answer = faq.Answer
			}
		}

		// Require a minimum score to consider it a valid match
		if bestScore < 0.7 {
			answer = "не знаю"
		}

		fmt.Println(answer)
	}
}
