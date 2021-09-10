package gateway

import "github.com/tsutarou10/line_project/service/pkg/entity"

type OpenGraphGateway interface {
	FetchOGPTag(string) entity.OGPTag
}
