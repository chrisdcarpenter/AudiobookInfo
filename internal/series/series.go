package series

import "fmt"

type Series struct {
	Name       string
	Website    string
	AmazonLink string
}

func (s Series) Link() string {
	return fmt.Sprintf("[%s](%s)[A](%s)", s.Name, s.Website, s.AmazonLink)
}
