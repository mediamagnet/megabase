package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/sirupsen/logrus"
	"image/color"
	"os"
)

var log = &logrus.Logger{
	Out:       os.Stderr,
	Formatter: new(logrus.TextFormatter),
	Hooks:     make(logrus.LevelHooks),
	Level:     logrus.InfoLevel,
}
var Path string

func main() {
	a := app.New()
	w := a.NewWindow("Megabase")


	tabs := container.NewAppTabs(
		container.NewTabItem("Installed", widget.NewLabel("Installed Here")),
		container.NewTabItem("Download", widget.NewLabel("Downloads Here")),
		container.NewTabItem("Mods", widget.NewLabel("Mod Downloads")),
	)

	rect := canvas.NewRectangle(color.White)

	buttons := container.New(
		layout.NewVBoxLayout(),
		widget.NewButton("Login", func() {log.Infoln("Clicked")}),
		widget.NewButton("New Instance", func() {log.Infoln("Clicked")}),
		widget.NewButton("Delete Instance", func() {log.Infoln("Clicked")}),
		widget.NewButton("Path", func() {filebox(w); log.Infof("Path: %v", Path)}),
		widget.NewButton("Working Dir", func(){log.Infoln(Path)}),
		)

	tabs.SetTabLocation(container.TabLocationTop)

	rect.Resize(fyne.NewSize(150,150))
	w.Resize(fyne.NewSize(300,500))
	w.SetContent(container.New(layout.NewHBoxLayout(), container.NewVBox(tabs, rect), buttons))
	w.ShowAndRun()
}

func filebox(window fyne.Window) {
	dialog.ShowFolderOpen(func(dir fyne.ListableURI, err error){
		if err != nil {
			dialog.ShowError(err, window)
			return
		}
		getPath(dir)
	}, window)
}

func getPath(uri fyne.ListableURI) string {
	Path = uri.String()
	return uri.String()
}