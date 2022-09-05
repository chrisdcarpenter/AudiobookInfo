# AudiobookInfo

| Latest Title | Series | Author | Release Date | Audio Date |
|--------------|--------|--------|--------------|------------|
{{- range .Books }}
| {{ .ToMarkdownLink }} | {{$.SeriesToMarkdownLink .SeriesName}} | {{$.AuthorToMarkdownLink .AuthorName}} | {{.ReleaseDate}} | {{.AudioDate}} |
{{- end}}
