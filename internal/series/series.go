package series

import (
	"fmt"
)

type Series struct {
	Name       string `json:"name"`
	Website    string `json:"website"`
	AmazonLink string `json:"amazon_link"`
	AuthorName string `json:"author"`
}

func (s Series) Link() string {
	return fmt.Sprintf("[%s](%s)[A](%s)", s.Name, s.Website, s.AmazonLink)
}
