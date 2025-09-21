package spending

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	layout     = "2006-01-02 15:04:05"
	dateString = "2025-02-02 12:30:45"
)

// func Test_SendSpender_Success(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	modelMock := mock_spending.NewMockSpendingAction(ctrl)
// 	model := New(modelMock)

// 	parsedTime, err := time.Parse(layout, dateString)
// 	assert.NoError(t, err)

// 	modelMock.EXPECT().SendSpending(int64(1), int64(32), "учёба", parsedTime)
// 	err = model.AddSpending(1, Spending{Sum: 32, SpendingType: "учёба", Date: parsedTime})

// 	assert.NoError(t, err)
// }

func Test_ParseSendSpendingComand(t *testing.T) {
	data, err := ParseSendSpendingComand("/addSum 15000 месяц 22.09.2025")

	timeCorrect, _ := time.ParseInLocation("02.01.2006", "22.09.2025", time.Local)
	assert.NoError(t, err)
	assert.Equal(t, data, &Spending{
		Sum:          15000,
		SpendingType: "месяц",
		Date:         timeCorrect,
	})
}

func Test_ParseSendSpendingComand_Empty_Text(t *testing.T) {
	_, err := ParseSendSpendingComand("/addSum ")

	assert.EqualError(t, err, ErrWrongStructure)
}

func Test_ParseSendSpendingComand_Uncorrect_Sum_Format(t *testing.T) {
	_, err := ParseSendSpendingComand("/addSum 232swqe2 месяц 22.09.2025")

	assert.EqualError(t, err, ErrConvertNumber)
}

func Test_ParseSendSpendingComand_Uncorrect_Sum(t *testing.T) {
	_, err := ParseSendSpendingComand("/addSum -15000 месяц 22.09.2025")

	assert.EqualError(t, err, ErrSumNonPositive)
}

func Test_ParseSendSpendingComand_Uncorrect_Period(t *testing.T) {
	_, err := ParseSendSpendingComand("/addSum 15000 день 22.09.2025")

	assert.EqualError(t, err, ErrWrongPeriod)
}

func Test_ParseSendSpendingComand_Uncorrect_Date(t *testing.T) {
	_, err := ParseSendSpendingComand("/addSum 15000 месяц 22-09-2025")

	assert.EqualError(t, err, ErrWrongTimeFormat)
}

func Test_ParseGetUserSpendingHistory(t *testing.T) {
	data, err := ParseGetUserSpendingHistory("/getSpending месяц")

	assert.NoError(t, err)
	assert.Equal(t, data, "месяц")
}

func Test_ParseGetUserSpendingHistory_Uncorrect_Period(t *testing.T) {
	data, err := ParseGetUserSpendingHistory("/getSpending годы")

	assert.EqualError(t, err, ErrWrongPeriod)
	assert.Equal(t, data, "")
}
