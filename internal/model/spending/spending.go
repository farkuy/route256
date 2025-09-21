package spending

import (
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type SpendingType string

const (
	SpendingTypeFood          SpendingType = "еда"
	SpendingTypeEntertainment SpendingType = "развлечения"
	SpendingTypeEducation     SpendingType = "учёба"
)

type TimePeriod string

const (
	Week  TimePeriod = "неделя"
	Month TimePeriod = "месяц"
	Year  TimePeriod = "год"
)

type Spending struct {
	Sum          int
	SpendingType SpendingType
	Date         time.Time
}

type SpendingAction interface {
	SendSpending(userId int64, sum int, spendingType SpendingType, date time.Time) error
	GetUserSpendingHistory(userId int64, timePeriod TimePeriod) (map[SpendingType]int, error)
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

func (s *SpendingsUsersStorage) GetUserSpendingHistory(userId int64, timePeriod TimePeriod) (map[SpendingType]int, error) {
	return s.Store.GetUserSpendingHistory(userId, timePeriod)
}

const (
	ErrWrongStructure  = "Не правильная структура команды"
	ErrConvertNumber   = "Произошла ошибка конвертации числа"
	ErrSumNonPositive  = "Сумма не может быть 0 или отрицательной"
	ErrWrongPeriod     = "Введен некорректный период"
	ErrWrongType       = "Введен некорректный тип"
	ErrWrongTimeFormat = "Введен некорректный формат времени"
)

func ParseGetUserSpendingHistory(reqText string) (string, error) {
	data := strings.Split(reqText, " ")
	if len(data) != 2 {
		return "", errors.Errorf(ErrWrongStructure)
	}

	timePeriod := data[1]
	isCorrectType := isCorrectPeriod(timePeriod)
	if !isCorrectType {
		return "", errors.Errorf(ErrWrongPeriod)
	}

	return timePeriod, nil
}

func ParseSendSpendingComand(reqText string) (*Spending, error) {
	data := strings.Split(reqText, " ")
	if len(data) != 4 {
		return nil, errors.Errorf(ErrWrongStructure)
	}

	sum, err := strconv.Atoi(data[1])
	if err != nil {
		return nil, errors.Errorf(ErrConvertNumber)
	}
	if sum <= 0 {
		return nil, errors.Errorf(ErrSumNonPositive)
	}

	spendingType := data[2]
	isCorrectType := isCorrectSpendingType(spendingType)
	if !isCorrectType {
		return nil, errors.Errorf(ErrWrongType)
	}

	textTime := data[3]
	timeCorrect, err := time.ParseInLocation("02.01.2006", textTime, time.Local)
	if err != nil {
		return nil, errors.Errorf(ErrWrongTimeFormat)
	}

	return &Spending{
		Sum:          sum,
		SpendingType: SpendingType(spendingType), //TODO подумать как переписать ублюдство
		Date:         timeCorrect,
	}, nil
}

func isCorrectPeriod(period string) bool {
	switch period {
	case string(Week):
		return true
	case string(Month):
		return true
	case string(Year):
		return true
	default:
		return false
	}
}

func isCorrectSpendingType(spendType string) bool {
	switch spendType {
	case string(SpendingTypeFood):
		return true
	case string(SpendingTypeEducation):
		return true
	case string(SpendingTypeEntertainment):
		return true
	default:
		return false
	}
}
