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
	fmt.Println(cleanedXML)

	var body Body
	err = xml.Unmarshal([]byte(cleanedXML), &body)
	e(err)

	// Combine all strings in Cell.Content into one string
	for i, table := range body.Tables {
		for j, row := range table.Rows {
			for k, cell := range row.Cells {
				body.Tables[i].Rows[j].Cells[k].Content = []string{strings.Join(cell.Content, "")}
			}
		}
	}

	jsonData, err := json.Marshal(body)
	e(err)

	fmt.Println(string(jsonData))
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
