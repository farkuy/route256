package spending

import (
	"time"
)

type SpendingType string

const (
	SpendingTypeFood          SpendingType = "еда"
	SpendingTypeEntertainment SpendingType = "развлечения"
	SpendingTypeEducation     SpendingType = "учёба"
)

type Spending struct {
	Sum          int
	SpendingType SpendingType
	Date         time.Time
}

type SpendingAction interface {
	SendSpending(userId int64, sum int, spendingType SpendingType, date time.Time) error
	GetUserSpendingHistory(userId int64) (*[]Spending, error)
}

type SpendingsUsersStorage struct {
	Store SpendingAction
}

func New(spending SpendingAction) *SpendingsUsersStorage {
	return &SpendingsUsersStorage{Store: spending}
}

func (s *SpendingsUsersStorage) AddSpending(userId int64, newSpend Spending) error {
	return s.Store.SendSpending(userId, newSpend.Sum, newSpend.SpendingType, newSpend.Date)
}

func (s *SpendingsUsersStorage) GetUserSpendingHistory(userId int64) (*[]Spending, error) {
	return s.Store.GetUserSpendingHistory(userId)
}
