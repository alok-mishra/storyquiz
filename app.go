package main

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"alokmishra.com/storyquiz/storyquiz"
)

var buildTime string // Will be set at build time

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

func (a *App) shutdown(ctx context.Context) {
}

func (a *App) BuildTime() string {
	return buildTime
}

func (a *App) Quiz(file string, fileName string, fileType string, outputType string) (string, error) {

	// oneMinuteAgo := time.Now().Add(-1 * time.Minute)
	// oneYearAgo := time.Now().AddDate(-1, 0, 0)
	// sixMonthsAgo := time.Now().AddDate(0, -6, 0)
	eightMonthsAgo := time.Now().AddDate(0, -8, 0)

	// Parse the build time
	buildTimeValue, err := time.Parse(time.RFC3339Nano, buildTime)
	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("something went wrong, app not built with proper flags")
	}

	if buildTimeValue.Before(eightMonthsAgo) {
		// log.Fatal("This build is stale")
		return "", errors.New("something went wrong, please update the app")
	}

	decoded, err := base64.StdEncoding.DecodeString(file)
	if err != nil {
		return "", err
	}
	return storyquiz.Quiz(decoded, fileName, fileType, outputType)
}
