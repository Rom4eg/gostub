package manager

import (
	"context"
	"testing"

	"github.com/Rom4eg/gostub/internal/manager/mocks"
	lm "github.com/Rom4eg/gostub/log/mocks"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestManager_StartService(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		prepare  func(*testing.T, *mocks.MockIService, IManager, *lm.MockILogger)
		expected func(*testing.T, error)
	}{
		{
			name: "Fail: error - service already started",
			prepare: func(t *testing.T, s *mocks.MockIService, m IManager, l *lm.MockILogger) {
				m.(*Manager).svc["test"] = s
				l.EXPECT().Debug(gomock.Any()).Times(2)
			},
			expected: func(t *testing.T, err error) {
				assert.ErrorIs(t, err, ErrServiceAlreadyStarted)
			},
		},
		{
			name: "FAIL: error returned",
			prepare: func(t *testing.T, s *mocks.MockIService, m IManager, l *lm.MockILogger) {
				l.EXPECT().Debug(gomock.Any()).Times(2)
				s.EXPECT().Start().Return(context.Canceled)
			},
			expected: func(t *testing.T, err error) {
				assert.Error(t, err)
			},
		},
		{
			name: "Pass: no errors",
			prepare: func(t *testing.T, s *mocks.MockIService, m IManager, l *lm.MockILogger) {
				l.EXPECT().Debug(gomock.Any()).Times(2)
				s.EXPECT().Start().Return(nil)
			},
			expected: func(t *testing.T, err error) {
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
			l.EXPECT().Debug(gomock.Any()).Times(2)

			m := New(l)
			s := mocks.NewMockIService(ctrl)

			tt.prepare(t, s, m, l)
			err := m.StartService("test", s)
			tt.expected(t, err)
		})
	}
}
