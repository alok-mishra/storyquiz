package storyquiz

import (
	"fmt"
	"os"
	"strings"
)

// https://community.articulate.com/series/articulate-storyline-360/articles/storyline-360-importing-questions-from-excel-spreadsheets-and-text-files#text
// Storyline Text Import Format
/*
	MC
	5
	Who was the first President of the United States?
	*George Washington | That's correct!
	John Adams | So close! John Adams was the second President and the first Vice President.
	Thomas Jefferson | Actually, Thomas Jefferson was the first Secretary of State. He was also the third President.
	Abraham Lincoln | Sorry, Abraham Lincoln was the sixteenth President.
*/

func replaceSmarts(text string) string {
	replacer := strings.NewReplacer("’", "'", "“", "\"", "”", "\"")
	return replacer.Replace(text)
}

func exportStoryline(questions []Question) {
	// Create a new file to write the questions to
	file, err := os.Create("Storyline - " + exportName + ".txt")
	e(err)
	defer file.Close()

	// Write the questions to the file
	for _, question := range questions {
		// Write the question type
		_, err = file.WriteString("MC\n")
		e(err)

		// Write the question points value
		// _, err = file.WriteString(strconv.Itoa(question.QuestionNumber) + "\n") // Question number is not used
		_, err = file.WriteString("1\n")
		e(err)

		// Write the question text
		_, err = file.WriteString(replaceSmarts(question.QuestionText) + "\n")

		e(err)

		// Write the options
		for _, option := range question.Options {
			textPrefix := ""
			feedback := ""

			if option.IsAnswer {
				textPrefix = "*"
			}

			if option.Feedback != "" {
				feedback = " | " + option.Feedback
			}

			_, _ = file.WriteString(textPrefix + replaceSmarts(option.Text) + feedback + "\n")
		}

		// Write a new line
		_, _ = file.WriteString("\n")
	}

	fmt.Print("Exported Storyline text!\n\n")
}
