package main

import (
	"embed"
	"encoding/base64"
	"log"
	"os"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func e(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "StoryQuiz",
		Width:  800,
		Height: 600,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		OnShutdown:       app.shutdown,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error starting Wails app:", err.Error())

		// Read the .docx file into a byte slice
		docxBytes, err := os.ReadFile("data/document.docx")
		e(err)

		// Encode the byte slice to base64
		encodedData := base64.StdEncoding.EncodeToString(docxBytes)

		println(app.Quiz(encodedData, "document.docx", "docx", "storyline"))

	}

}
