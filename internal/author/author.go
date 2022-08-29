package author

import "fmt"

type Author struct {
	Name    string            `json:"name"`
	Aliases []string          `json:"aliases"`
	Links   map[string]string `json:"links"`
}

func (a Author) ToMarkdownLink() string {
	return fmt.Sprintf("[%s](%s)[A](%s)", a.Name, a.Links["website"], a.Links["amazon"])
}
