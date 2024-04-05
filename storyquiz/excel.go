package storyquiz

import (
	"bytes"
	"strings"

	"github.com/xuri/excelize/v2"
)

func ProcessExcel(decodedBytes []byte) []Question {
	// Open the Excel file

	var questions []Question

	println("Processing Excel file")

	f, err := excelize.OpenReader(bytes.NewReader(decodedBytes))
	e(err)

	// get all sheets in the excel file
	// sheets := f.GetSheetList()

	// read the fourth sheet named "Module 1"
	rows, err := f.GetRows("Module 1")
	e(err)

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

	return questions

}
