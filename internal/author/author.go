package author

import "fmt"

type Author struct {
	Name       string
	Aliases    []string
	Website    string
	AmazonLink string
}

func (a Author) Link() string {
	return fmt.Sprintf("[%s](%s)", a.Name, a.Website)
}
