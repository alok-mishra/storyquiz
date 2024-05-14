package storyquiz

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

var exportName, courseCode, courseTitle string

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

func Quiz(decodedBytes []byte, fileName string, fileType string, outputType string) (string, error) {

	exportName = strings.Split(fileName, ".")[0]

	isWord := fileType == "doc" || fileType == "docx"
	isExcel := fileType == "xls" || fileType == "xlsx" || fileType == "xlsm"

	if isExcel {
		ProcessExcel(decodedBytes)
	} else if isWord {
		ProcessWord(decodedBytes)
	}

	if len(questions) == 0 {
		return "", errors.New("no questions found, do you have the correct file?")
	}

	if outputType == "cornerstone" {
		exportCornerstone(questions)
	} else if outputType == "storyline" {
		exportStoryline(questions)
	} else {
		if _, err := os.Stat("data/"); os.IsNotExist(err) {
			exportName = "document" //TODO: get the file name from the stream
		} else {
			exportName = "data/document"
		}
		outputJSON(questions)
	}

	return fmt.Sprintf("<span class='text-cyan-400'>%d</span> questions exported from <span class='text-cyan-400'>%s</span>!", len(questions), fileName), nil
}

func outputJSON(structure []Question) {
	// JSON Output for comparison only. Format JSON before inspecting.
	// Not needed, use the XML data unmarshalled into the structs directly.

	jsonData, err := json.Marshal(structure)
	e(err)

	jsonFile, err := os.Create(exportName + ".json")
	e(err)

	_, err = jsonFile.Write(jsonData)
	e(err)

	fmt.Println("Output JSON!")
	// fmt.Println(string(jsonData))
}
