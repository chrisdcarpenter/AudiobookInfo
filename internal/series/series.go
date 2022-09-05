package series

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"os"
)

type Series struct {
	Name       string            `json:"name"`
	AuthorName string            `json:"author_name"`
	Links      map[string]string `json:"links"`
}

func (s Series) ToMarkdownLink() string {
	return fmt.Sprintf("[%s](%s)[[A](%s)]", s.Name, s.Links["website"], s.Links["amazon"])
}

func LoadJson(dataPath string, s *Series) {
	file, err := os.ReadFile(dataPath)
	if err != nil {
		log.Error().Err(err).Msgf("Error trying to load: %s", dataPath)
	}
	_ = json.Unmarshal(file, s)
}
