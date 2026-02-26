package main

import (
	"fmt"
	"os"
	"text/template"
)

type Page struct {
	Title string,
	Body []byte,
	Player string
}

func loadPage() (*Page) {
}
