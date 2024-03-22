package main

import (
	"changeme/storyquiz"
	"context"
	"encoding/base64"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Quiz(file string) string {
	decoded, err := base64.StdEncoding.DecodeString(file)
	if err != nil {
		return err.Error()
	}
	return storyquiz.Quiz(decoded)
}
