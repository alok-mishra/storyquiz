# Storyline Quiz - Text Format

> scans for questions in a docx file and writes them to a plain text file in a format that can be used by storyline

## Go / Wails

- Setup

  - we need to add go to path
  - Add this to .zshrc for wsl/linux and restart terminal

    ```sh

    ## windows
    scoop install go
    go install github.com/wailsapp/wails/v2/cmd/wails@latest
    wails init -n storyquiz -t svelte-ts
    wails build

    ## arch
    paru -S go
    export PATH="$PATH:$(go env GOBIN):$(go env GOPATH)/bin"
    go install github.com/wailsapp/wails/v2/cmd/wails@latest

    paru -Sy
    paru -S webkit2gtk

    wails init -n storyquiz -t svelte-ts
    wails build -platform windows
    ```

- Install

  ```sh
  wails init -n storyquiz -t svelte-ts
  ```

- Build
  ```sh
  wails build
  wails build -platform windows # for cross platform build on linux
  ```
  - linux executable: `./build/bin/storyquiz`
  - windows executable: `./build/bin/storyquiz.exe`

## About

This is the official Wails Svelte-TS template.

## Live Development

To run in live development mode, run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build`.
