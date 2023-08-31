package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("My Application")
	myLabel := widget.NewLabel("Hello, World!")
	myWindow.SetContent(myLabel)
	myWindow.ShowAndRun()
}
