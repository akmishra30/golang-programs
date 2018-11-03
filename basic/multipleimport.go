package main

import (
	p "fmt"
	s "strings"
	t "time"
)

func main() {
	p.Println("Multiple Package Import Example.")

	p.Println(t.Now())

	p.Println("String lowercase: " + s.ToLower("Strings"))
}
