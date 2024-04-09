package storyquiz

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var outputFile string

// Define structs to represent questions and options
type Question struct {
	LearningObjective string   // The learning objective
	QuestionNumber    int      // The question number
	QuestionText      string   // The question text
	Options           []Option // The list of options for the question
}

type Option struct {
	Text     string // The option text
	IsAnswer bool   // Indicates if the option is a correct answer
	Feedback string // The feedback for the option
}

// Initialize a slice to hold the extracted questions
var questions []Question

func e(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Quiz(decodedBytes []byte, fileName string, fileType string) string {

	outputFile = fileName

	isWord := fileType == "doc" || fileType == "docx"
	isExcel := fileType == "xls" || fileType == "xlsx" || fileType == "xlsm"

	if isExcel {
		ProcessExcel(decodedBytes)
		outputStorylineText(questions)
		return "<span class='text-cyan-400'>" + fileName + "</span> exported!"
	}

	if isWord {
		ProcessWord(decodedBytes)
		outputStorylineText(questions)
		return "<span class='text-cyan-400'>" + fileName + "</span> exported!"
	}

	if _, err := os.Stat("data/"); os.IsNotExist(err) {
		outputFile = "document" //TODO: get the file name from the stream
	} else {
		outputFile = "data/document"
	}

	outputJSON(questions)

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

	outputStorylineText(questions)

	return "Quiz exported!"
}

func outputStorylineText(questions []Question) {
	// Create a new file to write the questions to
	file, err := os.Create(strings.Split(outputFile, ".")[0] + ".txt")
	e(err)
	defer file.Close()

	// Write the questions to the file
	for _, question := range questions {
		// Write the question type
		_, err = file.WriteString("MC\n")
		e(err)

		// Write the question number
		_, err = file.WriteString(strconv.Itoa(question.QuestionNumber) + "\n")
		e(err)

		// Write the question text
		_, err = file.WriteString(question.QuestionText + "\n")
		e(err)

		// Write the options
		for _, option := range question.Options {
			// Write the option text

			textPrefix := ""
			feedback := option.Feedback

			if feedback == "" {
				feedback = "That is incorrect."
			}

			if option.IsAnswer {
				textPrefix = "*"
				feedback = "That's correct!"
			}

			optionText := option.Text + " | " + feedback

			_, _ = file.WriteString(textPrefix + optionText)

			// Write the feedback if available
			if option.Feedback != "" {
				_, _ = file.WriteString(option.Feedback)
			}

			// Write a new line
			_, _ = file.WriteString("\n")
		}

		// Write a new line
		_, _ = file.WriteString("\n")
	}

	fmt.Print("Output Storyline text!\n\n")
}

func outputJSON(structure []Question) {
	// JSON Output for comparison only. Format JSON before inspecting.
	// Not needed, use the XML data unmarshalled into the structs directly.

	jsonData, err := json.Marshal(structure)
	e(err)

	jsonFile, err := os.Create(strings.Split(outputFile, ".")[0] + ".json")
	e(err)

	_, err = jsonFile.Write(jsonData)
	e(err)

	fmt.Println("Output JSON!")
	// fmt.Println(string(jsonData))
}
