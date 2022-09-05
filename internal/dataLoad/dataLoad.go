package dataLoad

import (
	"fmt"
	"github.com/chrisdcarpenter/bookTracking/internal/author"
	"github.com/chrisdcarpenter/bookTracking/internal/book"
	"github.com/chrisdcarpenter/bookTracking/internal/errors"
	"github.com/chrisdcarpenter/bookTracking/internal/series"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
	"strings"
)

const dataPathBase = "./data"
const BookType = "book"
const AuthorType = "author"
const SeriesType = "series"

func LoadBooks() (map[string]book.Book, error) {
	files, err := getFiles(BookType)
	if err != nil {
		fmt.Printf("Error getting files: %s", BookType)
		return nil, err
	}
	books := make(map[string]book.Book)
	for _, file := range files {
		var b book.Book
		book.LoadJson(file, &b)
		books[b.Name] = b
	}
	if len(books) == 0 {
		log.Error().Msg("Failed to load any books")
		return nil, errors.LoadError{}
	}
	return books, nil
}

func LoadAuthors() (map[string]author.Author, error) {
	files, err := getFiles(AuthorType)
	if err != nil {
		fmt.Printf("Error getting files: %s", AuthorType)
		return nil, err
	}
	authors := make(map[string]author.Author)
	for _, file := range files {
		var a author.Author
		author.LoadJson(file, &a)
		authors[a.Name] = a
	}
	if len(authors) == 0 {
		log.Error().Msg("Failed to load any books")
		return nil, errors.LoadError{}
	}
	return authors, nil
}

func LoadSeries() (map[string]series.Series, error) {
	files, err := getFiles(SeriesType)
	if err != nil {
		fmt.Printf("Error getting files: %s", SeriesType)
		return nil, err
	}
	_series := make(map[string]series.Series)
	for _, file := range files {
		var s series.Series
		series.LoadJson(file, &s)
		_series[s.Name] = s
	}
	if len(_series) == 0 {
		log.Error().Msg("Failed to load any books")
		return nil, errors.LoadError{}
	}
	return _series, nil
}

func getFiles(dataType string) ([]string, error) {
	var files []string
	root := fmt.Sprintf("%s/%s", dataPathBase, dataType)
	absRoot, _ := filepath.Abs(root)
	log.Info().Str("root", root).Str("abs_root", absRoot).Msg("Getting files in this root")
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !strings.HasPrefix(path, dataPathBase) && !strings.HasPrefix(filepath.Base(path), "_") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("Error loading %ss", dataType)
		return nil, err
	}
	return files, nil
}
