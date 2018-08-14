package main

import (
	"fmt"
	"sync"
)

type App struct {
    version int
    sync.Mutex
}
func (app *App) increment() {
    app.Lock()
    defer app.Unlock()
    app.version++
}

func main() {
	app := &App{version: 10}
	app.increment()
	fmt.Printf("app-version: %d\n", app.version)
}
