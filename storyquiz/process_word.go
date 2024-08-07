package storyquiz

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

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

func ProcessWord(decodedBytes []byte) {
	println("Processing Word file")

	// Clear the questions slice
	questions = questions[:0]

	r, err := zip.NewReader(bytes.NewReader(decodedBytes), int64(len(decodedBytes)))
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

	var questionTables []Table

	for _, table := range body.Tables {
		if len(table.Rows) >= 6 { // Limit to tables with at least 6 rows
			// fmt.Println("Table with at least 6 rows")
			for _, row := range table.Rows[1:4] { // Limit to the second to fourth rows
				// Check if the first cell contains "Question #"
				if len(row.Cells) > 0 && len(row.Cells[0].Content) > 0 &&
					strings.Contains(row.Cells[0].Content[0], "Question #") {
					// fmt.Println(row.Cells[1].Content[0])
					questionTables = append(questionTables, table)
					break
				}
			}
		}
	}

	// // Pause for user input
	// fmt.Println("Press 'Enter' to continue...")
	// bufio.NewReader(os.Stdin).ReadBytes('\n')

	if len(body.Tables[0].Rows[0].Cells) > 3 {
		courseCode = strings.Join(body.Tables[0].Rows[0].Cells[3].Content, "")
	} else {
		courseCode = "CODEMISSING"
	}

	//wait

	for _, table := range questionTables {
		// Extract questions from the table
		extractQuestions(table)
	}

	fmt.Println(len(questions), "Questions extracted from Word!")
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

func extractQuestions(table Table) {
	var learningObjective string
	var questionNumber int
	var questionText string
	var options []Option

	for _, row := range table.Rows {

		rowValid := len(row.Cells) > 0 && len(row.Cells[0].Content) > 0

		if learningObjective == "" {
			findAndStore(&row, &learningObjective, "Learning Objective")
		}

		if questionNumber == 0 {
			var qNum string
			findAndStore(&row, &qNum, "Question #")

			if qNum != "" {
				questionNumber, _ = strconv.Atoi(qNum)
			}
		}

		if questionText == "" {
			findAndStore(&row, &questionText, "Question Text")
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
				Feedback: "", // Set Feedback later if available
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
				Feedback: "", // Set Feedback later if available
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

func findAndStore(row *Row, storeVal *string, text string) {

	// fmt.Println("Finding and storing", text, "for", storeVal)
	rowValid := len(row.Cells) > 0 && len(row.Cells[0].Content) > 0

	// Check if the first cell contains the specified text
	if rowValid && strings.Contains(row.Cells[0].Content[0], text) {
		for i := 1; i < len(row.Cells); i++ {
			*storeVal += strings.Join(row.Cells[i].Content, "")
		}
	}
}
