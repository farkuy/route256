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

func Test_GetUserSpendingHistory_Base_Get_SUCCESS(t *testing.T) {
	model := New()
	model.SendSpending(userId, 12000, "еда", timeNow)

	userHistory, err := model.GetUserSpendingHistory(userId)

	assert.NoError(t, err)
	assert.Equal(t, []spending.Spending{{Sum: 12000, SpendingType: "еда", Date: timeNow}}, userHistory)
}

func Test_GetUserSpendingHistory_Base_Get_NotFoundUser(t *testing.T) {
	model := New()
	model.SendSpending(userId, 12000, "еда", timeNow)

	userHistory, err := model.GetUserSpendingHistory(2)

	assert.Error(t, err)
	assert.Equal(t, 0, len(userHistory))
}
