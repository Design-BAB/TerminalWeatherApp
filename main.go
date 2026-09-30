package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type weatherReport struct {
	zip           int
	place         string
	weatherStatus string
	temp          [5]int
}

func newWeatherReport(Zip int) *weatherReport {
	Temp := [5]int{90, 99, 100, 90, 81}
	theReport := weatherReport{zip: Zip, place: "Cape Coral", weatherStatus: "Sunny", temp: Temp}
	return &theReport
}

func isThereAZip(theFile string) bool {
	_, err := os.Stat(theFile)
	if os.IsNotExist(err) {
		fmt.Println("File not found!")
		return false
	} else if err != nil {
		fmt.Println("Error:", err)
		return false
	} else {
		fmt.Println("File exists.")
		return true
	}
}

func askForZip(flex *tview.Flex, app *tview.Application) {
	file, err := os.Create("zip.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	//defer file.Close()
	form := tview.NewForm()
	zipField := tview.NewInputField().
		SetLabel("Please enter your 5-digit zip code: ").
		SetFieldWidth(20)
	form.AddFormItem(zipField)
	form.AddButton("Enter", func() {
		zip := zipField.GetText()
		app.Stop()
		fmt.Println("You entered " + zip)
		data := []byte(zip)
		_, err := file.Write(data)
		if err != nil {
			fmt.Println(err)
			file.Close()
			return
		}
	})
	flex.AddItem(form, 0, 1, true)
}

func displayWeather(flex *tview.Flex, report *weatherReport) {
	table := tview.NewTable().SetBorders(true)

	const dayRow = 0
	const tempRow = 1

	for column, temp := range report.temp {
		dayNumber := column + 1

		dayText := fmt.Sprintf("%d", dayNumber)
		tempText := fmt.Sprintf("%d°", temp)

		dayCell := tview.NewTableCell(dayText)
		dayCell.SetAlign(tview.AlignCenter)
		dayCell.SetSelectable(false)

		tempCell := tview.NewTableCell(tempText)
		tempCell.SetAlign(tview.AlignCenter)
		tempCell.SetSelectable(false)

		table.SetCell(dayRow, column, dayCell)
		table.SetCell(tempRow, column, tempCell)
	}

	flex.AddItem(table, 0, 1, false)
}

func main() {
	fmt.Println(isThereAZip("zip.txt"))
	app := tview.NewApplication()
	app.EnableMouse(true)

	tview.Styles.PrimitiveBackgroundColor = tcell.ColorDefault

	flex := tview.NewFlex()

	flex.SetDirection(2)
	flex.SetBorder(true)

	if isThereAZip("zip.txt") {
		ReportForToday := newWeatherReport(33904)
		displayWeather(flex, ReportForToday)
	} else {
		askForZip(flex, app)
	}

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}
