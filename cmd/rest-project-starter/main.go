package main

import (
	"fmt"

	"github.com/typical-go/rest-project-starter/internal/app"
	_ "github.com/typical-go/rest-project-starter/internal/generated/typical"
	"github.com/typical-go/typical-go/pkg/typapp"
)

func main() {
	fmt.Printf("%s %s\n", typapp.Name, typapp.Version)
	typapp.Start(app.Start)
}
