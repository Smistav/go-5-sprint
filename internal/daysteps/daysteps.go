package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

// Parse парсит строку с данными о прогулке.
// Принимает формат "количество шагов,продолжительность", например "678,0h50m".
// Возвращает ошибку если она есть
func (ds *DaySteps) Parse(datastring string) (err error) {
	sliceParse := strings.Split(datastring, ",")
	if len(sliceParse) != 2 {
		return errors.New("invalid input format")
	}
	steps, err := strconv.Atoi(sliceParse[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return errors.New("step count must be positive")
	}
	ds.Steps = steps
	duration, err := time.ParseDuration(sliceParse[1])
	if err != nil {
		return err
	}
	if duration <= 0 {
		return errors.New("walk duration must be positive")
	}
	ds.Duration = duration
	return nil
}

// ActionInfo рассчитывает и возвращает строку с информацией о дневной активности.
// Возвращает отформатированную строку с количеством шагов, пройденной дистанцией
// и количеством сожженных калорий.
// Если входные данные некорректны или пусты, функция выводит ошибку и возвращает пустую строку.
func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	walkingSpentCalories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps,
		distance,
		walkingSpentCalories,
	), nil
}
