package spending

import (
	mock_spending "route256/internal/mock/spending"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	layout     = "2006-01-02 15:04:05"
	dateString = "2025-02-02 12:30:45"
)

func Test_SendSpender_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	modelMock := mock_spending.NewMockSpendingAction(ctrl)
	model := New(modelMock)

	parsedTime, err := time.Parse(layout, dateString)
	assert.NoError(t, err)

	modelMock.EXPECT().SendSpending(int64(1), int64(32), "учёба", parsedTime)
	err = model.AddSpending(1, Spending{Sum: 32, SpendingType: "учёба", Date: parsedTime})

	assert.NoError(t, err)
}
