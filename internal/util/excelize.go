package util

import (
	"fmt"
	"os"
	"strconv"

	db "github.com/dpurbosakti/booknest-grpc/internal/db/sqlc"
	"github.com/xuri/excelize/v2"
)

func writeUserToExcel(user db.User, f *excelize.File) {
	emptyRow := findEmptyRow(f)
	if emptyRow == 1 {
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(emptyRow), "ID")
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(emptyRow), "Name")
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(emptyRow), "Phone")
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(emptyRow), "Email")
		f.SetCellValue("Sheet1", "E"+strconv.Itoa(emptyRow), "HashedPassword")
		f.SetCellValue("Sheet1", "F"+strconv.Itoa(emptyRow), "Role.String")
		f.SetCellValue("Sheet1", "G"+strconv.Itoa(emptyRow), "PasswordChangedAt")
		f.SetCellValue("Sheet1", "H"+strconv.Itoa(emptyRow), "CreatedAt")
		f.SetCellValue("Sheet1", "I"+strconv.Itoa(emptyRow), "UpdatedAt")
		f.SetCellValue("Sheet1", "J"+strconv.Itoa(emptyRow), "IsEmailVerified")

		emptyRow = 2
		// Set value of cells with user information
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(emptyRow), user.ID.String())
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(emptyRow), user.Name)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(emptyRow), user.Phone)
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(emptyRow), user.Email)
		f.SetCellValue("Sheet1", "E"+strconv.Itoa(emptyRow), "-")
		f.SetCellValue("Sheet1", "F"+strconv.Itoa(emptyRow), user.Role.String)
		f.SetCellValue("Sheet1", "G"+strconv.Itoa(emptyRow), user.PasswordChangedAt.String())
		f.SetCellValue("Sheet1", "H"+strconv.Itoa(emptyRow), user.CreatedAt.String())
		f.SetCellValue("Sheet1", "I"+strconv.Itoa(emptyRow), user.UpdatedAt.String())
		f.SetCellValue("Sheet1", "J"+strconv.Itoa(emptyRow), user.IsEmailVerified)
	} else {
		// Set value of cells with user information
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(emptyRow), user.ID.String())
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(emptyRow), user.Name)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(emptyRow), user.Phone)
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(emptyRow), user.Email)
		f.SetCellValue("Sheet1", "E"+strconv.Itoa(emptyRow), "-")
		f.SetCellValue("Sheet1", "F"+strconv.Itoa(emptyRow), user.Role.String)
		f.SetCellValue("Sheet1", "G"+strconv.Itoa(emptyRow), user.PasswordChangedAt.String())
		f.SetCellValue("Sheet1", "H"+strconv.Itoa(emptyRow), user.CreatedAt.String())
		f.SetCellValue("Sheet1", "I"+strconv.Itoa(emptyRow), user.UpdatedAt.String())
		f.SetCellValue("Sheet1", "J"+strconv.Itoa(emptyRow), user.IsEmailVerified)
	}
}

func WriteToExcel(user db.User) error {
	var f *excelize.File
	var err error

	// Check if the file exists
	if _, err = os.Stat("assets/user.xlsx"); os.IsNotExist(err) {
		// File does not exist, create a new one
		f = excelize.NewFile()
		// Create a new sheet
		f.NewSheet("Sheet1")
	} else {
		// File exists, open it
		f, err = excelize.OpenFile("assets/user.xlsx")
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	// Write user information to Excel
	writeUserToExcel(user, f)
	defer f.Close()

	// Get the index of the "Sheet1"
	sheet1Index, err := f.GetSheetIndex("Sheet1")
	if err != nil {
		fmt.Println("Failed to get sheet index:", err)
		return err
	}

	// Set active sheet of the workbook to "Sheet1"
	f.SetActiveSheet(sheet1Index)

	// Save spreadsheet in assets directory
	if err := f.SaveAs("assets/user.xlsx"); err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Excel file created and saved successfully!")
	return nil
}

func findEmptyRow(f *excelize.File) int {
	// Specify the sheet name
	sheetName := "Sheet1"

	// Start iterating over rows to find the first empty row
	emptyRowIndex := 1 // Start from the first row
	for {
		cellValue, _ := f.GetCellValue(sheetName, fmt.Sprintf("A%d", emptyRowIndex))
		if cellValue == "" {
			// Found an empty row, break the loop
			break
		}
		emptyRowIndex++
	}
	return emptyRowIndex
}
