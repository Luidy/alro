package handler

import (
	model "alro/model"
	"context"
)

type EchoHandlerFunc func(ctx context.Context, req *model.EchoRequest) (*model.EchoResponse, error)
func Echo() EchoHandlerFunc {
	return func(ctx context.Context, req *model.EchoRequest) (*model.EchoResponse, error) {
		return &model.EchoResponse{
			Content: req.Content,
		}, nil
	}
}
