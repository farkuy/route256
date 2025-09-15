package spending

import (
	"errors"
	"time"
)

type SpendingSpender interface {
	SendSpending(userId int64, sum int64, spendingType string, date time.Time) error
}

type Spending struct {
	Sum          int64
	SpendingType string //сделать енум
	Date         time.Time
}

type SpendingsUsersStorage struct {
	Historys        map[int64][]Spending
	spendingSpender SpendingSpender
}

func New(spendingSender SpendingSpender) *SpendingsUsersStorage {
	return &SpendingsUsersStorage{Historys: make(map[int64][]Spending), spendingSpender: spendingSender}
}

func (s *SpendingsUsersStorage) AddSpending(userId int64, newSpend Spending) error {
	return s.spendingSpender.SendSpending(userId, newSpend.Sum, newSpend.SpendingType, newSpend.Date)
}

func (s *SpendingsUsersStorage) GetUserSpendingHistory(userId int64) (*[]Spending, error) {
	userHistory, isHave := s.Historys[userId]
	if !isHave {
		return nil, errors.New("история пользователя не найдена")
	}

	return &userHistory, nil
}
