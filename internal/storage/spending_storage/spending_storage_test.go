package spending_storage

import (
	"route256/internal/model/spending"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var timeNow = time.Now()
var userId int64 = 1

func Test_SpendingStore_Add_NewSpending_Create_NewUser_Store(t *testing.T) {

	model := New()

	err := model.SendSpending(userId, 12000, "еда", timeNow)

	assert.NoError(t, err)
	assert.Contains(t, model.HistorySpendingsUsers[userId], spending.Spending{Sum: 12000, SpendingType: "еда", Date: timeNow})
}

func Test_SpendingStore_Add_NewSpending_Minus_Sum(t *testing.T) {
	model := New()

	err := model.SendSpending(userId, -12000, "еда", timeNow)

	assert.Error(t, err)
}

func Test_SpendingStore_Add_NewSpending_Date_After_DateNow(t *testing.T) {
	model := New()

	err := model.SendSpending(userId, 12000, "еда", timeNow.AddDate(0, 0, 1))

	assert.Error(t, err)
}

func Test_GetUserSpendingHistory_Base_Get_Week(t *testing.T) {
	model := New()
	model.SendSpending(userId, 1, "еда", timeNow)
	model.SendSpending(userId, 11, "развлечения", timeNow)
	model.SendSpending(userId, 1, "еда", timeNow)
	model.SendSpending(userId, 1, "еда", timeNow.AddDate(0, 0, -8))
	model.SendSpending(userId, 1, "еда", timeNow.AddDate(0, -1, 0))
	model.SendSpending(userId, 1, "еда", timeNow.AddDate(-1, 0, 0))

	totalSum, err := model.GetUserSpendingHistory(userId, spending.Week)

	assert.NoError(t, err)
	assert.Equal(
		t,
		map[spending.SpendingType]int{
			spending.SpendingTypeFood:          2,
			spending.SpendingTypeEntertainment: 11,
			spending.SpendingTypeEducation:     0,
		},
		totalSum,
	)
}

func Test_GetUserSpendingHistory_Base_Get_Month(t *testing.T) {
	model := New()
	model.SendSpending(userId, 1, "еда", timeNow)
	model.SendSpending(userId, 1, "развлечения", timeNow)
	model.SendSpending(userId, 1, "еда", timeNow)
	model.SendSpending(userId, 1, "еда", timeNow.AddDate(0, 0, -8))
	model.SendSpending(userId, 1, "еда", timeNow.AddDate(0, -1, 0))
	model.SendSpending(userId, 1, "развлечения", timeNow.AddDate(0, -1, 0))
	model.SendSpending(userId, 1, "развлечения", timeNow.AddDate(0, -2, 0))
	model.SendSpending(userId, 1, "еда", timeNow.AddDate(-1, 0, 0))

	totalSum, err := model.GetUserSpendingHistory(userId, spending.Month)

	assert.NoError(t, err)
	assert.Equal(
		t,
		map[spending.SpendingType]int{
			spending.SpendingTypeFood:          4,
			spending.SpendingTypeEntertainment: 2,
			spending.SpendingTypeEducation:     0,
		},
		totalSum,
	)
}

func Test_GetUserSpendingHistory_Base_Get_Year(t *testing.T) {
	model := New()
	model.SendSpending(userId, 1, "еда", timeNow)
	model.SendSpending(userId, 1, "развлечения", timeNow)
	model.SendSpending(userId, 1, "еда", timeNow)
	model.SendSpending(userId, 1, "еда", timeNow.AddDate(0, 0, -8))
	model.SendSpending(userId, 1, "еда", timeNow.AddDate(0, -1, 0))
	model.SendSpending(userId, 1, "развлечения", timeNow.AddDate(0, -1, 0))
	model.SendSpending(userId, 1, "развлечения", timeNow.AddDate(0, -2, 0))
	model.SendSpending(userId, 1, "еда", timeNow.AddDate(-1, 0, 0))
	model.SendSpending(userId, 1, "еда", timeNow.AddDate(-2, 0, 0))
	model.SendSpending(userId, 1, "еда", timeNow.AddDate(-3, 0, 0))

	totalSum, err := model.GetUserSpendingHistory(userId, spending.Year)

	assert.NoError(t, err)
	assert.Equal(
		t,
		map[spending.SpendingType]int{
			spending.SpendingTypeFood:          5,
			spending.SpendingTypeEntertainment: 3,
			spending.SpendingTypeEducation:     0,
		},
		totalSum,
	)
}
