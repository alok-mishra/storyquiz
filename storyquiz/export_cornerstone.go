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

	const header = "Question ID,Question Text,Question Type,Correct Answer,Category Id,Default Language,Active,Randomize Answer Choices,Answer Explanation,# of Answer Choices,All of the Above,None of the Above,Answer 1,Answer 2,Answer 3,Answer 4,Answer 5,Answer 6,Answer 7,Answer 8,Answer 9,Answer 10,Image Filename,Answer Coordinates,Apply partial scoring,Author"

	// loop through the header and write to the sheet
	for i, v := range strings.Split(header, ",") {
		f.SetCellValue(sheetName, fmt.Sprintf("%c%d", 'A'+i, 1), v)
	}

	// delete the default sheet
	f.DeleteSheet("Sheet1")

	// Set active sheet of the workbook.

	err = f.SaveAs("LMS - " + exportName + ".xlsx")
	e(err)

	fmt.Print(len(questions), "Exported Cornerstone Excel!\n\n")
}
