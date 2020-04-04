package main

import (
	"log"
	"net/http"

	"github.com/shurcooL/vfsgen"
)

func main() {
	fs := http.Dir("./templates")
	err := vfsgen.Generate(fs, vfsgen.Options{
		Filename:     "generated/assets/assets.go",
		PackageName:  "assets",
		VariableName: "Assets",
	})
	if err != nil {
		log.Fatalln(err)
	}
}
