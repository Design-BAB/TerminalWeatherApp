package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

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

func askForZip(form *tview.Form, app *tview.Application, file *os.File) {
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

}

func main() {
	fmt.Println(isThereAZip("zip.txt"))
	file, err := os.Create("zip.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	app := tview.NewApplication()
	app.EnableMouse(true)

	tview.Styles.PrimitiveBackgroundColor = tcell.ColorDefault

	form := tview.NewForm()
	askForZip(form, app, file)

	flex := tview.NewFlex()
	flex.AddItem(form, 0, 1, true)

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}
