package book

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"os"
)

const layout = "2006/01/02"

type Book struct {
	Name        string            `json:"name"`
	AuthorName  string            `json:"author_name"`
	SeriesName  string            `json:"series_name"`
	Links       map[string]string `json:"links"`
	ReleaseDate string            `json:"release_date"`
	AudioDate   string            `json:"audio_date"`
}

func (b Book) ToMarkdownLink() string {
	return fmt.Sprintf("[%s](%s)[[A](%s)]", b.Name, b.Links["website"], b.Links["amazon"])
}

func LoadJson(dataPath string, b *Book) {
	file, err := os.ReadFile(dataPath)
	if err != nil {
		log.Error().Err(err).Msgf("Error trying to load: %s", dataPath)
	}
	_ = json.Unmarshal(file, b)
}
