package utils

import (
	og "github.com/otiai10/opengraph/v2"
)

func FetchOGP(url string) *og.OpenGraph {
	ogp, _ := og.Fetch(url)
	return ogp
}

func FetchImageURL(ogp *og.OpenGraph) string {
	if len(ogp.Image) == 0 {
		return ""
	}
	ogp.ToAbs()
	return ogp.Image[0].URL
}

func FetchTitle(ogp *og.OpenGraph) string {
	return ogp.Title
}
