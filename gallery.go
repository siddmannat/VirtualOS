package main

import (
	//"fmt"
    "io/ioutil"
    "log"
	"strings"
	"fyne.io/fyne/v2"
	//"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	
)

func showGallery() {
	
	w := myapp.NewWindow("Gallery")
	w.Resize(fyne.NewSize(800,600))
	tabs:=container.NewAppTabs()
	root_src:="E:\\go\\go"
	files, err := ioutil.ReadDir(root_src)
	if err != nil {
        log.Fatal(err)
    }
	
	for _, file := range files {
       // fmt.Println(file.Name(), file.IsDir())
		if !file.IsDir(){
			extensions:= strings.Split(file.Name(),".")[1]
			if extensions=="png" || extensions=="jpeg" || extensions=="jfif" ||extensions=="jpg" {
				
				image:=canvas.NewImageFromFile(root_src+"\\"+ file.Name())
			    image.FillMode = canvas.ImageFillStretch
				
			    tabs.Append(container.NewTabItem(file.Name(),image ))
				
				 
	
			}
		}
    }
	
	tabs.SetTabLocation(container.TabLocationLeading)
	w.SetContent(container.NewBorder(desktopbttn,nil,nil,nil,tabs),)
	w.Show()
}

