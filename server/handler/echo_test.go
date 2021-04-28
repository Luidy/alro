package handler

import (
	model "alro/model"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEcho(t *testing.T) {
	tests := []struct {
		name    string
		req     *model.EchoRequest
		want    *model.EchoResponse
		wantErr bool
	}{
		{
			name: "normal",
			req: &model.EchoRequest{
				Content: "hello",
			},
			want: &model.EchoResponse{
				Content: "hello",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(*testing.T))
		ctx := context.Background()
		got, err := Echo()(ctx, tt.req)
		assert.Equal(t, tt.want, got)

		if tt.wantErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}
