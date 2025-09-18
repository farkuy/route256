package spending_storage

import (
	"log/slog"
	"route256/internal/model/spending"
	"time"

	"github.com/pkg/errors"
)

type TimePeriod string

const (
	Week  TimePeriod = "неделя"
	Month TimePeriod = "месяц"
	Year  TimePeriod = "год"
)

type Store struct {
	HistorySpendingsUsers map[int64][]spending.Spending
}

func New() *Store {
	return &Store{HistorySpendingsUsers: make(map[int64][]spending.Spending)}
}

const (
	notFoundUser = "Пользователь с таким id не сущесвтует"
	wrongPeriod  = "Задан не правильный период"
)

func (s *Store) GetUserSpendingHistory(userId int64, timePeriod TimePeriod) (map[spending.SpendingType]int, error) {
	history, isHas := s.HistorySpendingsUsers[userId]
	if !isHas {
		return nil, errors.Errorf(notFoundUser)
	}

	minDate, err := getMinDate(timePeriod)
	if err != nil {
		return nil, err
	}

	categoryTotalSum := map[spending.SpendingType]int{
		spending.SpendingTypeFood:          0,
		spending.SpendingTypeEntertainment: 0,
		spending.SpendingTypeEducation:     0,
	}
	for _, val := range history {
		if val.Date.After(minDate) {
			categoryTotalSum[val.SpendingType] += val.Sum
		}
	}

	return categoryTotalSum, nil
}

func getMinDate(timePeriod TimePeriod) (time.Time, error) {
	now := time.Now()
	minDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	switch timePeriod {
	case Week:
		return minDate.AddDate(0, 0, -7), nil
	case Month:
		return minDate.AddDate(0, -1, 0), nil
	case Year:
		return minDate.AddDate(-1, 0, 0), nil
	default:
		return minDate, errors.Errorf(wrongPeriod)
	}
}

const (
	minusSumErr  = "Трата не может быть отрицательной"
	dateAfterErr = "Заданное время не может быть позже сегодняшнего числа"
)

func (s *Store) SendSpending(userId int64, sum int, spendingType spending.SpendingType, date time.Time) error {
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
