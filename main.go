package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"strings"
)

type Package struct {
	XMLName xml.Name `xml:"http://schemas.microsoft.com/office/2006/xmlPackage package"`
	Parts   []Part   `xml:"part"`
}

type Part struct {
	Name        string `xml:"name,attr"`
	ContentType string `xml:"contentType,attr"`
	XMLData     struct {
		Relationships Relationships `xml:"Relationships"`
		Document      Document      `xml:"document"`
	} `xml:"xmlData"`
}

type Relationships struct {
	Relationships []Relationship `xml:"Relationship"`
}

type Relationship struct {
	ID     string `xml:"Id,attr"`
	Type   string `xml:"Type,attr"`
	Target string `xml:"Target,attr"`
}

type Document struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main document"`
	Body    Body     `xml:"body"`
}

type Body struct {
	XMLName    xml.Name    `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main body"`
	Paragraphs []Paragraph `xml:"p"`
}

type Paragraph struct {
	XMLName xml.Name `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main p"`
	Text    string   `xml:",innerxml"`
}

func main() {
	xmlFile, err := os.Open("document.xml")
	if err != nil {
		log.Fatalf("error opening XML file: %s", err)
	}
	defer xmlFile.Close()

	var pkg Package
	decoder := xml.NewDecoder(xmlFile)
	if err := decoder.Decode(&pkg); err != nil {
		log.Fatalf("error decoding XML: %s", err)
	}

	for _, part := range pkg.Parts {
		if part.ContentType == "application/vnd.openxmlformats-officedocument.wordprocessingml.document.main+xml" {
			for _, p := range part.XMLData.Document.Body.Paragraphs {
				// Remove XML tags and trim whitespace to extract text content
				text := strings.TrimSpace(strings.ReplaceAll(p.Text, "<w:t>", ""))
				fmt.Println(text)
			}
		}
	}
}
