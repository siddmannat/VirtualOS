package main

import (
	//"fyne.io/fyne"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)
var myapp fyne.App=app.New()

var mywindow fyne.Window=myapp.NewWindow("Virtual OS")

var b1 fyne.Widget
var b2 fyne.Widget
var b3 fyne.Widget
var b4 fyne.Widget

var img fyne.CanvasObject
var desktopbttn fyne.Widget
var panelContent *fyne.Container

func main(){
	myapp.Settings().SetTheme(theme.LightTheme())
	img=canvas.NewImageFromFile("E:\\go\\295661.jpg")

	b1= widget.NewButtonWithIcon("Weather App",theme.InfoIcon(),func() {
		showWeatherApp(mywindow)})
	b2= widget.NewButtonWithIcon("Calculator",theme.ContentAddIcon(),func() {
		showCalculator(mywindow)})
	b3= widget.NewButtonWithIcon("Gallery App",theme.StorageIcon(),func() {
		showGallery()})
	b4= widget.NewButtonWithIcon("Text Editor",theme.DocumentIcon(),func() {
		showTextEditor(mywindow)})
	
	desktopbttn=widget.NewButtonWithIcon("Home",theme.HomeIcon(),func() {
		mywindow.SetContent(container.NewBorder(panelContent,nil,nil,nil,img))
	})
	panelContent =container.NewVBox((container.NewGridWithColumns(5,desktopbttn,b1,b2,b3,b4)))
	mywindow.Resize(fyne.NewSize(1280,720))
	mywindow.CenterOnScreen()
	mywindow.SetContent(
		container.NewBorder(panelContent,nil,nil,nil,img),
	)
	mywindow.ShowAndRun()
}