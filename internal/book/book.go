package book

import (
	"fmt"
	"sort"
	"time"
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
	return fmt.Sprintf("[%s](%s)[A](%s)", b.Name, b.Links["website"], b.Links["amazon"])
}

func (b Book) NextReleaseDate() string {
	if b.ReleaseDate == "" && b.AudioDate == "" {
		return "Unknown"
	}
	var dates []time.Time
	today := time.Now()
	releaseDate, err := time.Parse(layout, b.ReleaseDate)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return "err"
	}
	if today.Before(releaseDate) {
		dates = append(dates, releaseDate)
	}
	audioDate, err := time.Parse(layout, b.AudioDate)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
		return "err"
	}
	if today.Before(audioDate) {
		dates = append(dates, audioDate)
	}
	sort.Slice(dates, func(i, j int) bool {
		return dates[i].Before(dates[j])
	})
	if b.ReleaseDate == "" && len(dates) == 0 {
		return "Release Unknown"
	}
	if b.AudioDate == "" && len(dates) == 0 {
		return "Audiobook Unknown"
	}
	return "Unknown"
}
