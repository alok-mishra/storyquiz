package storyquiz

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/xuri/excelize/v2"
)

func ProcessExcel(decodedBytes []byte) {
	println("Processing Excel file")

	// Clear the questions slice
	questions = questions[:0]

	f, err := excelize.OpenReader(bytes.NewReader(decodedBytes))
	e(err)

	// get all sheets in the excel file
	// sheets := f.GetSheetList()

	// read the fourth sheet named "Module 1"
	rows, err := f.GetRows("Module 1")
	if err != nil {
		// sheet must be named with roles
		rows, err = f.GetRows(f.GetSheetName(3))
		e(err)
		fmt.Println("Sheet:", f.GetSheetName(3))
		extractQuestionRows(rows)
	}

	if len(rows) == 0 {
		fmt.Println("No rows found in the Excel file")
		return
	}

	fmt.Println(len(questions), "Questions extracted from Excel!")
}

func extractQuestionRows(rows [][]string) {
	// first row contains the headers
	// row[1] contains the question text, and row[2] contains the answer
	// while row[1] is empty, row[2] contains the distractors

	// while loop were row[2] is not empty
	// append the question to the questions slice
	// move to the next row

	questionNumber := 1

	for index, row := range rows {
		if index == 0 {
			continue // skip header row
		}

		if row[1] != "" {
			// This is a question row, with the correct answer
			question := Question{
				LearningObjective: "",
				QuestionNumber:    questionNumber,
				QuestionText:      strings.Split(row[1], "\n")[0],
			}

			answer := Option{
				Text:     row[2],
				IsAnswer: true,
				Feedback: "",
			}

			question.Options = append(question.Options, answer)
			questions = append(questions, question)
			questionNumber++

		} else if row[2] != "" {
			// This is a distractor row
			lastQuestionIndex := len(questions) - 1
			if lastQuestionIndex >= 0 {
				distractor := Option{
					Text:     row[2],
					IsAnswer: false,
					Feedback: "",
				}
				questions[lastQuestionIndex].Options = append(questions[lastQuestionIndex].Options, distractor)
			}
		} else {
			// end of questions
			break
		}
	}

}
