package main

import (
	"archive/zip"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const docxFile = "tables.docx"

type Body struct {
	Tables []Table `xml:"tbl"`
}

type Table struct {
	Rows []Row `xml:"tr"`
}

type Row struct {
	Cells []Cell `xml:"tc"`
}

type Cell struct {
	Content []string `xml:"p>r>t"`
}

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

func main() {
	// Open the .docx file
	docx, err := os.Open(docxFile)
	e(err)
	defer docx.Close()

	// Get information about the .docx file
	docxInfo, err := docx.Stat()
	e(err)

	// Read the document.xml file from the .docx file
	r, err := zip.NewReader(docx, docxInfo.Size())
	e(err)

	var xmlFile *zip.File
	for _, f := range r.File {
		if f.Name == "word/document.xml" {
			xmlFile = f
			break
		}
	}

	if xmlFile == nil {
		log.Fatal("word/document.xml not found in the .docx file")
	}

	file, err := xmlFile.Open()
	e(err)
	defer file.Close()

	// Read the content of the document.xml file into a byte slice
	var xmlData []byte
	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if err != nil && err.Error() != "EOF" {
			log.Fatal(err)
		}
		if n == 0 {
			break
		}
		xmlData = append(xmlData, buffer[:n]...)
	}

	// fmt.Println(string(xmlData))

	cleanedXML := cleanXML(string(xmlData))
	// fmt.Println(cleanedXML)
	fmt.Println("Cleaned XML!")

	var body Body
	err = xml.Unmarshal([]byte(cleanedXML), &body)
	e(err)

	// outputJSON(body)

	var questionTables []*Table

	for _, table := range body.Tables {
		if len(table.Rows) >= 6 { // Limit to tables with at least 6 rows
			for _, row := range table.Rows[1:3] { // Limit to the second and third rows
				// Check if the first cell contains "Question #"
				if len(row.Cells) > 0 && len(row.Cells[0].Content) > 0 &&
					strings.Contains(row.Cells[0].Content[0], "Question #") {
					questionTables = append(questionTables, &table)
					break
				}
			}
		}
	}

	// outputJSON(learningObjectiveTables)

	for _, table := range questionTables {
		// Extract questions from the table
		extractQuestions(table)
	}

	outputJSON(questions)

	// Combine all strings in Cell.Content into one string
	for t, table := range body.Tables {
		for r, row := range table.Rows {
			for c, cell := range row.Cells {
				body.Tables[t].Rows[r].Cells[c].Content = []string{strings.Join(cell.Content, "")}
			}
		}
	}

	// https://community.articulate.com/series/articulate-storyline-360/articles/storyline-360-importing-questions-from-excel-spreadsheets-and-text-files#text
	/*
		MC
		5
		Who was the first President of the United States?
		*George Washington | That's correct!
		John Adams | So close! John Adams was the second President and the first Vice President.
		Thomas Jefferson | Actually, Thomas Jefferson was the first Secretary of State. He was also the third President.
		Abraham Lincoln | Sorry, Abraham Lincoln was the sixteenth President.
	*/
}

// func outputJSON(structure Body) {
// func outputJSON(structure []*Table) {
func outputJSON(structure []Question) {
	// JSON Output for comparison only. Format JSON before inspecting.
	// Not needed, use the XML data unmarshalled into the structs directly.

	jsonData, err := json.Marshal(structure)
	e(err)

	jsonFile, err := os.Create(strings.Split(docxFile, ".")[0] + "-min.json")
	e(err)

	_, err = jsonFile.Write(jsonData)
	e(err)

	fmt.Println("Output JSON!")
	// fmt.Println(string(jsonData))
}

func cleanXML(content string) string {
	// fmt.Println(content)
	body := regexp.MustCompile(`(?s)<w:body>(.*?)<\/w:body>`).FindString(content)
	// fmt.Println(body)
	re := regexp.MustCompile(`<(/?)w:([^/>\s]+)[^/>]*(/?)>`)
	cleanedBody := re.ReplaceAllString(body, "<$1$2$3>")
	// fmt.Println(cleanedBody)
	return cleanedBody
}

func extractQuestions(table *Table) {
	var learningObjective string
	var questionNumber int
	var questionText string
	var options []Option

	for _, row := range table.Rows {

		rowValid := len(row.Cells) > 0 && len(row.Cells[0].Content) > 0

		// Check if the first cell contains "Learning Objective"
		if rowValid && strings.Contains(row.Cells[0].Content[0], "Learning Objective") {
			// Extract learning objective from the next cell(s)
			if len(row.Cells) > 1 && len(row.Cells[1].Content) > 0 {
				learningObjective = strings.Join(row.Cells[1].Content, "")
			}
		}

		// Check if the first cell contains "Question #"
		if rowValid && strings.Contains(row.Cells[0].Content[0], "Question #:") {
			// Extract question number from the next cell
			if len(row.Cells) > 1 && len(row.Cells[1].Content) > 0 {
				questionNumber, _ = strconv.Atoi(row.Cells[1].Content[0])
			}
		}

		// Check if the first cell contains "Question Text:"
		if rowValid && strings.Contains(row.Cells[0].Content[0], "Question Text:") {
			// Extract question text from the next cell(s)
			for i := 1; i < len(row.Cells); i++ {
				questionText += strings.Join(row.Cells[i].Content, "")
			}
		}

		// Check if the first cell contains "Correct Answer:"
		if rowValid && strings.Contains(row.Cells[0].Content[0], "Correct Answer:") {
			// Extract correct answer from the next cell(s)
			correctAnswer := ""
			for i := 1; i < len(row.Cells); i++ {
				correctAnswer += strings.Join(row.Cells[i].Content, "")
			}

			// Create an Option object
			option := Option{
				Text:     correctAnswer,
				IsAnswer: true,
				// Set Feedback later if available
			}

			// Append the option to the options slice
			options = append(options, option)
		}

		// Check if the first cell contains "Option:"
		if rowValid && strings.Contains(row.Cells[0].Content[0], "Option:") {
			// Extract options from the next cell(s)
			optionText := ""
			for i := 1; i < len(row.Cells); i++ {
				optionText += strings.Join(row.Cells[i].Content, "")
			}

			// Create an Option object
			option := Option{
				Text:     optionText,
				IsAnswer: false,
				// Set Feedback later if available
			}

			// Append the option to the options slice
			options = append(options, option)
		}
	}

	// Create a Question object
	question := Question{
		QuestionNumber:    questionNumber,
		QuestionText:      questionText,
		LearningObjective: learningObjective,
		Options:           options,
	}

	// Append the question to the questions slice
	questions = append(questions, question)
}

func storeRow(row *Row, text string) (rowContent string) {
	rowValid := len(row.Cells) > 0 && len(row.Cells[0].Content) > 0

	// Check if the first cell contains the specified text
	if rowValid && strings.Contains(row.Cells[0].Content[0], text) {
		for i := 1; i < len(row.Cells); i++ {
			rowContent += strings.Join(row.Cells[i].Content, "")
		}
		// fmt.Println(rowContent)

		return rowContent
	}

	return ""
}
