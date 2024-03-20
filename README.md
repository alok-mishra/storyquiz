# Storyline Quiz - Text Format

> scans for questions in a docx file and writes them to a plain text file in a format that can be used by storyline

## wails

- go install github.com/wailsapp/wails/v2/cmd/wails@latest
- wails init -n storyquiz -t svelte-ts

## About

This is the official Wails Svelte-TS template.

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.
