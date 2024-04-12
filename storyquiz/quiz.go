package storyquiz

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

var exportFile string

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
		// log.Fatal(err) // I need to see the error message
		fmt.Println(err)
	}
}

func Quiz(decodedBytes []byte, fileName string, fileType string, outputType string) string {

	exportFile = fileName

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
		exportFile = "document" //TODO: get the file name from the stream
	} else {
		exportFile = "data/document"
	}

	outputJSON(questions)

	return "Quiz exported!"
}

func outputJSON(structure []Question) {
	// JSON Output for comparison only. Format JSON before inspecting.
	// Not needed, use the XML data unmarshalled into the structs directly.

	jsonData, err := json.Marshal(structure)
	e(err)

	jsonFile, err := os.Create(strings.Split(exportFile, ".")[0] + ".json")
	e(err)

	_, err = jsonFile.Write(jsonData)
	e(err)

	fmt.Println("Output JSON!")
	// fmt.Println(string(jsonData))
}
