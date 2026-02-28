package http

import (
	"net/http"
	"testing"

	lm "github.com/Rom4eg/gostub/log/mocks"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestDefaultServer(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		host string
		port int
		want IServer
	}{
		{
			name: "PASS: net/http server",
			host: "localhost",
			port: 8080,
			want: &http.Server{
				Addr:    "localhost:8080",
				Handler: nil,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			l := lm.NewMockILogger(ctrl)
			l.EXPECT().Info(gomock.Any()).Times(1)
			got := DefaultServer(tt.host, tt.port, nil, l)
			assert.Equal(t, tt.want, got)
		})
	}
}
