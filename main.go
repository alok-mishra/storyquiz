package main

import (
	"archive/zip"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"regexp"
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
	QuestionNumber int      // The question number
	QuestionText   string   // The question text
	Options        []Option // The list of options for the question
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

	// Combine all strings in Cell.Content into one string
	for t, table := range body.Tables {
		for r, row := range table.Rows {
			for c, cell := range row.Cells {
				body.Tables[t].Rows[r].Cells[c].Content = []string{strings.Join(cell.Content, "")}
			}
		}
	}

	outputJSON(body)

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

func outputJSON(body Body) {
	// JSON Output for comparison only. Format JSON before inspecting.
	// Not needed, use the XML data unmarshalled into the structs directly.

	jsonData, err := json.Marshal(body)
	e(err)

	jsonFile, err := os.Create(strings.Split(docxFile, ".")[0] + ".json")
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
