package book

import (
	"fmt"
)

type Book struct {
	Name       string `json:"name"`
	Website    string `json:"website"`
	AmazonLink string `json:"amazon_link"`
	AuthorName string `json:"author"`
	SeriesName string `json:"series"`
}

func (b Book) Link() string {
	return fmt.Sprintf("[%s](%s)[A](%s)", b.Name, b.Website, b.AmazonLink)
}
