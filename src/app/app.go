package app

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"github.com/darley-raffael/app-age-golang/src/calculator"
)

type Application struct {
	fyneApp fyne.App
	window  fyne.Window
}

func NewApp() *Application {
	fyneApp := app.New()
	window := fyneApp.NewWindow("Calculadora Idade")

	calc := calculator.NewAgeCalculator()
	ui := calculator.NewAgeCalculatorUI(calc)

	content := container.NewVBox(ui.CanvasObject())

	window.SetContent(content)

	return &Application{
		fyneApp: fyneApp,
		window:  window,
	}

}

func (a *Application) Run() {
	a.window.ShowAndRun()
}
