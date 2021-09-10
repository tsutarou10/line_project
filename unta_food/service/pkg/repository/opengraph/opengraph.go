package opengraph

import (
	og "github.com/otiai10/opengraph/v2"
	"github.com/tsutarou10/line_project/service/pkg/entity"
)

type OpenGraph struct {
	ogp *og.OpenGraph
}

func NewOpenGraph() *OpenGraph {
	return &OpenGraph{}
}

func (o *OpenGraph) FetchOGPTag(url string) entity.OGPTag {
	ogp, _ := og.Fetch(url)
	o.ogp = ogp

	return entity.OGPTag{
		Title:    o.FetchTitle(),
		ImageURL: o.FetchImageURL(),
	}
}

func (o *OpenGraph) FetchImageURL() string {
	if len(o.ogp.Image) == 0 {
		return ""
	}
	o.ogp.ToAbs()
	return o.ogp.Image[0].URL
}

func (o *OpenGraph) FetchTitle() string {
	return o.ogp.Title
}
