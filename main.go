package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"regexp"
)

const filePath = "tables.xml"

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
	Content string `xml:"p"`
}

func main() {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Read the file contents
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	fileSize := fileInfo.Size()
	xmlData := make([]byte, fileSize)
	_, err = file.Read(xmlData)
	if err != nil {
		log.Fatal(err)
	}

	cleanedXML := cleanXML(string(xmlData))
	// fmt.Println(cleanedXML)

	var body Body
	if err := xml.Unmarshal([]byte(cleanedXML), &body); err != nil {
		log.Fatal(err)
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))
}

func cleanXML(content string) string {

	body := regexp.MustCompile(`(?s)<w:body>(.*?)<\/w:body>`).FindString(content)
	// fmt.Println(body)
	re := regexp.MustCompile(`<(/?)w:([^/>\s]+)[^/>]*(/?)>`)

	cleanedBody := re.ReplaceAllString(body, "<$1$2$3>")

	fmt.Println(cleanedBody)

	return cleanedBody
}
