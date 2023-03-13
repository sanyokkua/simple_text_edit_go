package main

import (
	"context"
	"github.com/labstack/gommon/log"
	menu2 "github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"os"
)

type App struct {
	ctx context.Context
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) basicMenu() *menu2.Menu {
	menu := menu2.NewMenu()

	file := menu.AddSubmenu("File")

	file.AddText("Open", nil, func(data *menu2.CallbackData) {
		dialog, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
			Title: "Open Text File",
		})
		if err != nil {
			log.Info(err)
			return
		}
		log.Info("File: ", dialog)
		file, err := os.ReadFile(dialog)
		if err != nil {
			log.Error(err)
		}
		text := string(file)
		runtime.EventsEmit(a.ctx, "FileIsChosen", text)
	})

	return menu
}
