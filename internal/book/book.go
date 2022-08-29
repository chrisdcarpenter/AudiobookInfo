package book

import (
	"fmt"
)

type Book struct {
	Name        string            `json:"name"`
	AuthorName  string            `json:"author_name"`
	SeriesName  string            `json:"series_name"`
	Links       map[string]string `json:"links"`
	ReleaseDate string            `json:"release_date"`
	AudioDate   string            `json:"audio_date"`
}

func (b Book) ToMarkdownLink() string {
	return fmt.Sprintf("[%s](%s)[A](%s)", b.Name, b.Links["website"], b.Links["amazon"])
}
