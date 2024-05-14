package storyquiz

import (
	"fmt"
	"strings"

	"github.com/xuri/excelize/v2"
)

func exportCornerstone(questions []Question) {

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	var sheetName = "Questions"

	index, err := f.NewSheet(sheetName)
	e(err)

	f.SetActiveSheet(index)

	const header = "Question Reference Number,Question Text,Question Type,Correct Answer,Category Id,Default Language,Active,Randomize Answer Choices,Answer Explanation,# of Answer Choices,All of the Above,None of the Above,Answer 1,Answer 2,Answer 3,Answer 4,Answer 5,Answer 6,Answer 7,Answer 8,Answer 9,Answer 10,Image Filename,Answer Coordinates,Apply partial scoring,Author"

	// loop through the header and write to the sheet
	for i, v := range strings.Split(header, ",") {
		f.SetCellValue(sheetName, fmt.Sprintf("%c%d", 'A'+i, 1), v)
	}

	fmt.Println("Code:", courseCode)

	// for each question, write to the sheet
	for i, question := range questions {
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", i+2), strings.TrimSpace(fmt.Sprintf("%s %d", courseCode, question.QuestionNumber)))
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", i+2), question.QuestionText)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", i+2), "Multiple Choice / Single Answer")
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", i+2), 1)
		f.SetCellValue(sheetName, fmt.Sprintf("E%d", i+2), "")
		f.SetCellValue(sheetName, fmt.Sprintf("F%d", i+2), "en-US")
		f.SetCellValue(sheetName, fmt.Sprintf("G%d", i+2), 1)
		f.SetCellValue(sheetName, fmt.Sprintf("H%d", i+2), 1)
		f.SetCellValue(sheetName, fmt.Sprintf("I%d", i+2), "")
		f.SetCellValue(sheetName, fmt.Sprintf("J%d", i+2), len(question.Options))
		f.SetCellValue(sheetName, fmt.Sprintf("K%d", i+2), "")
		f.SetCellValue(sheetName, fmt.Sprintf("L%d", i+2), "")
		for j, option := range question.Options {
			f.SetCellValue(sheetName, fmt.Sprintf("%c%d", 'M'+j, i+2), option.Text)
		}
	}

	// delete the default sheet
	f.DeleteSheet("Sheet1")

	err = f.SaveAs(exportName + " (LMS).xlsx")
	e(err)

	fmt.Print("Exported Cornerstone Excel!\n\n")
}
