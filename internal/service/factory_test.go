package service

import (
	"gostub/internal/service/http"
	lm "gostub/log/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestFactory_MakeService(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		t       ServiceType
		opts    map[string]any
		expect  func(*testing.T, Service, error)
		prepare func(t *testing.T, l *lm.MockILogger)
	}{
		{
			name: "FAIL: unknown type",
			t:    ServiceType("unknown"),
			expect: func(t *testing.T, s Service, err error) {
				assert.Nil(t, s)
				assert.ErrorIs(t, err, ErrUnknownServiceType)
			},
			prepare: func(t *testing.T, l *lm.MockILogger) {

			},
		},
		{
			name: "FAIL: incorrect options",
			t:    ServiceHttp,
			opts: map[string]any{"host": 123},
			expect: func(t *testing.T, s Service, err error) {
				assert.Nil(t, s)
				assert.Error(t, err)
			},
			prepare: func(t *testing.T, l *lm.MockILogger) {

			},
		},
		{
			name: "Pass: http",
			t:    ServiceHttp,
			opts: map[string]any{
				"host": "localhost",
				"port": 8080,
				"name": "test",
				"root": "test",
			},
			expect: func(t *testing.T, s Service, err error) {
				assert.NoError(t, err)
				assert.IsType(t, &http.Service{}, s)
			},
			prepare: func(t *testing.T, l *lm.MockILogger) {
				l.EXPECT().Info(gomock.Any()).Times(1)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			l := lm.NewMockILogger(ctrl)
			tt.prepare(t, l)
			f := NewFactory()
			got, err := f.MakeService("test", FactoryOpt{Type: tt.t, ServiceOpt: tt.opts, Logger: l})
			tt.expect(t, got, err)
		})
	}
}
