package main

import (
    "flag"
    "github.com/ennm/go-utils/generate"
)

func main() {

    var name string

    flag.StringVar(&name, "name", "", "generate go file")

    flag.Parse()

    if len(name) > 0 {
        //file
        generate.Do(name)
    }
}