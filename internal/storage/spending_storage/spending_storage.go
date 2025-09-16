package spending_storage

import (
	"log/slog"
	"route256/internal/model/spending"
	"time"

	"github.com/pkg/errors"
)

type Store struct {
	HistorySpendingsUsers map[int64][]spending.Spending
}

func New() *Store {
	return &Store{HistorySpendingsUsers: make(map[int64][]spending.Spending)}
}

const (
	notFountUser = "Пользователь с таким id не сущесвтует"
)

func (s *Store) GetUserSpendingHistory(userId int64, timePeriod string) ([]spending.Spending, error) {
	history, isHas := s.HistorySpendingsUsers[userId]
	if !isHas {
		return nil, errors.Errorf(notFountUser)
	}

	return history, nil
}

const (
	minusSumErr  = "Трата не может быть отрицательной"
	dateAfterErr = "Заданное время не может быть позже сегодняшнего числа"
)

func (s *Store) SendSpending(userId int64, sum int64, spendingType spending.SpendingType, date time.Time) error {
	if sum < 0 {
		return errors.Errorf(minusSumErr)
	}
	if date.After(time.Now()) {
		return errors.Errorf(dateAfterErr)
	}

	newSpending := spending.Spending{Sum: sum, SpendingType: spendingType, Date: date}
	history, isHas := s.HistorySpendingsUsers[userId]
	if !isHas {
		s.HistorySpendingsUsers[userId] = []spending.Spending{newSpending}
		slog.Info("Для пользователя создан стор c", slog.Int64("id", userId))
		return nil
	}
	s.HistorySpendingsUsers[userId] = append(history, newSpending)

	return nil
}
