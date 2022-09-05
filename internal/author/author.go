package author

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"os"
)

type Author struct {
	Name    string            `json:"name"`
	Aliases []string          `json:"aliases"`
	Links   map[string]string `json:"links"`
}

func (a Author) ToMarkdownLink() string {
	return fmt.Sprintf("[%s](%s)[[A](%s)]", a.Name, a.Links["website"], a.Links["amazon"])
}

func LoadJson(dataPath string, a *Author) {
	file, err := os.ReadFile(dataPath)
	if err != nil {
		log.Error().Err(err).Msgf("Error trying to load: %s", dataPath)
	}
	_ = json.Unmarshal(file, a)
}
