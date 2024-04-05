package utils

import "strings"

const (
	ImagePrefix = "https://image.tmdb.org/t/p/original"
)

func FillURI(uri string) string {
	return ImagePrefix + uri
}

func SplitSeries(str string) []string {
	return strings.Split(str, ",")
}
