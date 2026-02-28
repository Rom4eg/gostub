package http

import (
	"context"
	"testing"

	"github.com/Rom4eg/gostub/internal/service/http/mocks"
	lm "github.com/Rom4eg/gostub/log/mocks"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestService_Stop(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		prepare   func(t *testing.T, l *lm.MockILogger, srv *mocks.MockIServer)
		expectErr func(*testing.T, error)
	}{
		{
			name: "FAIL: error returned",
			prepare: func(t *testing.T, l *lm.MockILogger, srv *mocks.MockIServer) {
				l.EXPECT().Debug(gomock.Any()).Times(2)
				srv.EXPECT().Shutdown(gomock.Any()).Return(context.Canceled)
			},
			expectErr: func(t *testing.T, err error) {
				assert.Error(t, err)
			},
		},
		{
			name: "PASS: no errors",
			prepare: func(t *testing.T, l *lm.MockILogger, srv *mocks.MockIServer) {
				l.EXPECT().Debug(gomock.Any()).Times(2)
				srv.EXPECT().Shutdown(gomock.Any()).Return(nil)
			},
			expectErr: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			l := lm.NewMockILogger(ctrl)
			srv := mocks.NewMockIServer(ctrl)

			opts := ServiceOpts{
				Server: srv,
			}

			s := New("test", l, opts)
			tt.prepare(t, l, srv)
			err := s.Stop()

			tt.expectErr(t, err)
		})
	}
}
