package book

import "fmt"

type Book struct {
	Name       string
	Website    string
	AmazonLink string
}

func (b Book) Link() string {
	return fmt.Sprintf("[%s](%s)[A](%s)", b.Name, b.Website, b.AmazonLink)
}
