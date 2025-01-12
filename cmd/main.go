package main

import (
	"html/template"
	"os"
	"time"
)

func main() {
	t := template.Must(template.ParseFiles("templates/index.html"))
	data := struct{ MyAge int }{
		MyAge: calculateMyAge(),
	}

	fp, err := os.Create("static/index.html")
	check(err)
	defer fp.Close()

	check(t.Execute(fp, data))
	check(fp.Sync())
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func calculateMyAge() int {
	now := time.Now()
	myBirthDate, _ := time.Parse(time.DateOnly, "2001-04-21")
	years := now.Year() - myBirthDate.Year()

	if now.YearDay() < myBirthDate.YearDay() {
		years--
	}

	return years
}
