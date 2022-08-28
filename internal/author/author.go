package author

import "fmt"

type Author struct {
	Name       string   `json:"name"`
	Aliases    []string `json:"aliases"`
	Website    string   `json:"website"`
	AmazonLink string   `json:"amazon_link"`
}

func (a Author) Link() string {
	return fmt.Sprintf("[%s](%s)[A](%s)", a.Name, a.Website, a.AmazonLink)
}
