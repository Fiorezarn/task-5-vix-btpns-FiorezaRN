package service

import (
	"BTPN-API-go/model/web"
	"context"
)

type ImageService interface {
	Create(ctx context.Context, request web.ImageCreateRequest) []web.ImageResponse
	Delete(ctx context.Context, imageId string)
	FindById(ctx context.Context, imageId string) web.ImageResponse
}
