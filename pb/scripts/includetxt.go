package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// Reads all .json files in the swagger folder
// and encodes them as strings literals in textfiles.go
func main() {
	fs, _ := ioutil.ReadDir("swagger")
	out, _ := os.Create("swagger.pb.go")
	out.Write([]byte("package pb \n\nconst (\n"))
	for _, f := range fs {
		if strings.HasSuffix(f.Name(), ".json") {
			name := strings.TrimPrefix(f.Name(), "swagger.")
			fmt.Printf("Generate: %s %s\n", f.Name(), name)
			out.Write([]byte(strings.TrimSuffix(name, ".swagger.json") + " = `"))
			f, _ := os.Open("swagger/" + f.Name())
			io.Copy(out, f)
			out.Write([]byte("`\n"))
		}
	}
	out.Write([]byte(")\n"))
}
