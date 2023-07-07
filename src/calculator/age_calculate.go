package calculator

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type AgeCalculatorUI struct {
	dateEntry    *widget.Entry
	resultLabel  *widget.Label
	calc         AgeCalculator
	calculateBtn *widget.Button
	content      fyne.CanvasObject
}

func NewAgeCalculatorUI(calc AgeCalculator) *AgeCalculatorUI {
	dateEntry := widget.NewEntry()
	resultLabel := widget.NewLabel("")

	calculateButton := widget.NewButton("Calcular Idade", func() {
		dateStr := dateEntry.Text

		age := calc.CalculateAge(dateStr)

		resultLabel.SetText("Idade: " + strconv.Itoa(age))
	})

	content := container.NewVBox(
		widget.NewLabel("Digite sua data de nascimento (dd/mm/yyyy):"),
		dateEntry,
		calculateButton,
		resultLabel,
	)

	return &AgeCalculatorUI{
		dateEntry:    dateEntry,
		resultLabel:  resultLabel,
		calc:         calc,
		calculateBtn: calculateButton,
		content:      content,
	}
}

func (ac *AgeCalculatorUI) CanvasObject() fyne.CanvasObject {
	return ac.content
}
