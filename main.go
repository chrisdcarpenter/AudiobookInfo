package main

import (
	"fmt"
	"github.com/chrisdcarpenter/bookTracking/internal/book"
	"github.com/chrisdcarpenter/bookTracking/internal/library"
	"github.com/rs/zerolog/log"
	"os"
	"text/template"
)

const booksPath = "./data/book"

func main() {
	b := book.Book{}
	book.LoadJson("./data/book/balefire.json", &b)
	fmt.Printf("Book: %+v\n", b)

	var l library.Library
	library.LoadAllData(&l)

	err := renderBooksMarkdown("./templates/books.md.tmpl", l)
	if err != nil {
		log.Error().Err(err).Msg("Failed to render books markdown")
	}
}

func renderBooksMarkdown(path string, library library.Library) error {
	t, err := template.ParseFiles(path)
	if err != nil {
		log.Print(err)
		return err
	}

	f, err := os.Create("./books.md")
	if err != nil {
		log.Error().Err(err).Msgf("Error creating file")
		return err
	}

	err = t.Execute(f, library)
	if err != nil {
		log.Print("execute: ", err)
		return err
	}
	return nil
}
