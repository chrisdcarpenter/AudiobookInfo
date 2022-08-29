package series

import (
	"fmt"
)

type Series struct {
	Name       string            `json:"name"`
	AuthorName string            `json:"author_name"`
	Links      map[string]string `json:"links"`
}

func (s Series) ToMarkdownLink() string {
	return fmt.Sprintf("[%s](%s)[A](%s)", s.Name, s.Links["website"], s.Links["amazon"])
}
