package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/tealeg/xlsx"
	"log"
	"net/http"
	// "os"
)

func dbregHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "dbreg.html")
}

// export LANG=ru_RU.UTF-8
// go run main.go  | less -S
func dbregJsonHandler(w http.ResponseWriter, r *http.Request) {
	// excelFileName := "data/dbReg.xlsx"

	// strconv.ParseFloat: parsing "11.0.2.0.4.7": invalid syntax rowIdx:  274 cellIdx:  2
	// return error but still return value
	xlFile, err := xlsx.OpenFile("data/dbReg.xlsx")
	// same error
	// xlFile, err := xlsx.OpenFile("data/dbRegText.xlsx")

	if err != nil {
		log.Fatal(err)
	}

	// line := [][]string{}

	type db struct {
		Dbname              string `json:"dbname"`
		Dbtype              string `json:"dbtype"`
		Dbversion           string `json:"dbversion"`
		DbOptimizerFeatures string `json:"dboptimizerfeatures"`
		DbCompatible        string `json:"dbcompatible"`
		DbListenerPort      string `json:"dblistenerport"`
		Dbos                string `json:"dbos"`
		DbHardware          string `json:"dbhardware"`
		DbServerName        string `json:"dbservername"`
		Dbip                string `json:"dbip"`
		DbisVirtual         string `json:"dbisvirtual"`
		DbDataCenter        string `json:"dbdatacenter"`
		DbBackup            string `json:"dbbackup"`
		DbTsmClient         string `json:"dbtsmclient"`
		DbTsmServer         string `json:"dbtsmserver"`
		DbDBA               string `json:"dbdba"`
		DbStorage           string `json:"dbstorage"`
		DbProvisioning      string `json:"dbprovisioning"`
		DbRecreate          string `json:"dbrecreate"`
		Dbhome              string `json:"dbhome"`
		Dbbase              string `json:"dbbase"`
		DbPurpose           string `json:"dbpurpose"`
	}
	line := []db{}

	for _, sheet := range xlFile.Sheets {
		for rowIdx, row := range sheet.Rows {

			// debug
			// read first section and break
			if rowIdx == 13 {
				// break
			}

			// fmt.Printf("rowIdx: %#v ", rowIdx)
			// fmt.Printf("len(row.Cells): %#v \n", len(row.Cells))
			tmpLine := []string{}
			for cellIdx, cell := range row.Cells {

				str, err := cell.String()

				if err != nil {
					// log.Fatal(err)
					log.Println(err, "rowIdx: ", rowIdx, "cellIdx: ", cellIdx)
				}

				// if ferst 3 cell is not empty it is database reocord
				if cellIdx <= 2 && str == "" {
					tmpLine = []string{}
					break
				}
				if cellIdx == 22 {
					break
				}

				// skip headers
				if cellIdx == 0 && str == "Название БД" {
					break
				}

				// print new line if it is the last cell in a row
				if cellIdx != len(row.Cells)-1 {
					// fmt.Printf("%#v ", str)
				} else {
					// fmt.Printf("%#v \n", str)
				}
				tmpLine = append(tmpLine, str)

			}
			// fmt.Printf("\n", "")
			// fmt.Printf("%#v \n", len(tmpLine))
			if len(tmpLine) >= 5 {
				// fmt.Printf("%#v \n", tmpLine)
				line = append(line, db{
					Dbname:              tmpLine[0],
					Dbtype:              tmpLine[1],
					Dbversion:           tmpLine[2],
					DbOptimizerFeatures: tmpLine[3],
					DbCompatible:        tmpLine[4],
					DbListenerPort:      tmpLine[5],
					Dbos:                tmpLine[6],
					DbHardware:          tmpLine[7],
					DbServerName:        tmpLine[8],
					Dbip:                tmpLine[9],
					DbisVirtual:         tmpLine[10],
					DbDataCenter:        tmpLine[11],
					DbBackup:            tmpLine[12],
					DbTsmClient:         tmpLine[13],
					DbTsmServer:         tmpLine[14],
					DbDBA:               tmpLine[15],
					DbStorage:           tmpLine[16],
					DbProvisioning:      tmpLine[17],
					DbRecreate:          tmpLine[18],
					Dbhome:              tmpLine[19],
					Dbbase:              tmpLine[20],
					DbPurpose:           tmpLine[21],
				})
			}
		}
	}
	// fmt.Printf("%#v \n", line)

	j, err := json.Marshal(line)
	if err != nil {
		log.Fatal(err)
	}
	// os.Stdout.Write(j)
	w.Write(j)

}

func webServer() {
	// http://127.0.0.1:8080/dbreg

	routes := mux.NewRouter().StrictSlash(false)

	routes.HandleFunc("/dbreg", dbregHandler).Methods("GET")
	routes.HandleFunc("/api/dbreg", dbregJsonHandler).Methods("GET")

	fmt.Println("Start listening...")
	http.ListenAndServe(":8080", routes)
}

func main() {
	webServer()
}
