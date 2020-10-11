package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/typical-go/rest-project-starter/internal/app"
	_ "github.com/typical-go/rest-project-starter/internal/generated/typical"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typgo"
)

func main() {
	fmt.Printf("%s %s\n", typgo.ProjectName, typgo.ProjectVersion)
	if err := typapp.Run(app.Start); err != nil {
		logrus.Fatal(err.Error())
	}
}
