package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"log"
)

func main() {
	// excelFileName := "data/dbReg.xlsx"
	xlFile, err := xlsx.OpenFile("data/dbReg.xlsx")
	if err != nil {
		log.Fatal(err)
	}

	for _, sheet := range xlFile.Sheets {
		for idx, row := range sheet.Rows {

			// debug
			// read first section and break
			if idx == 13 {
				break
			}

			for _, cell := range row.Cells {
				str, err := cell.String()
				if err != nil {
					log.Fatal(err)
				}

				// fmt.Printf("%s ", str)
				fmt.Printf("%#v ", str)
			}
			fmt.Printf("\n", "")
		}
	}
}
