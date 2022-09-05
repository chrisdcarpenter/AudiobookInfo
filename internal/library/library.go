package library

import (
	"github.com/chrisdcarpenter/bookTracking/internal/author"
	"github.com/chrisdcarpenter/bookTracking/internal/book"
	"github.com/chrisdcarpenter/bookTracking/internal/dataLoad"
	"github.com/chrisdcarpenter/bookTracking/internal/series"
	"github.com/rs/zerolog/log"
)

type Library struct {
	Books   map[string]book.Book
	Series  map[string]series.Series
	Authors map[string]author.Author
}

func LoadAllData(l *Library) {
	var err error
	l.Books, err = dataLoad.LoadBooks()
	if err != nil {
		log.Error().Err(err).Msg("Error happened loading books")
	}
	l.Authors, err = dataLoad.LoadAuthors()
	if err != nil {
		log.Error().Err(err).Msg("Error happened loading authors")
	}
	l.Series, err = dataLoad.LoadSeries()
	if err != nil {
		log.Error().Err(err).Msg("Error happened loading series")
	}
}

func (l Library) AuthorToMarkdownLink(a string) string {
	if val, ok := l.Authors[a]; ok {
		return val.ToMarkdownLink()
	}
	return ""
}

func (l Library) SeriesToMarkdownLink(s string) string {
	if val, ok := l.Series[s]; ok {
		return val.ToMarkdownLink()
	}
	return ""
}

func (l Library) BookToMarkdownLink(b string) string {
	if val, ok := l.Books[b]; ok {
		return val.ToMarkdownLink()
	}
	return ""
}
