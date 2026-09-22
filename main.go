package main

import (
	"fmt"
	"os"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	_, err := os.Stat("example.txt")
	if os.IsNotExist(err) {
		fmt.Println("File not found!")
	} else if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("File exists.")
	}

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
	zipField := tview.NewInputField().
		SetLabel("Please enter your 5-digit zip code: ").
		SetFieldWidth(20)

	form.AddFormItem(zipField)

	form.AddButton("Enter", func() {
		zip := zipField.GetText()
		if err != nil {
			panic(err)
		}
		app.Stop()
		fmt.Println("You entered" + zip)
		data := []byte(zip)
		_, err = file.Write(data)
		if err != nil {
			fmt.Println(err)
			file.Close()
			return
		}
	})

	flex := tview.NewFlex()
	flex.AddItem(form, 0, 1, true)

	if err := app.SetRoot(flex, true).Run(); err != nil {
		panic(err)
	}
}
