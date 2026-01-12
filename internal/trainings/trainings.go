package trainings

import (
	"errors"
	"fmt"
	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
	"strconv"
	"strings"
	"time"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

// Parse парсит строку с данными об тренировки.
// Принимает формат "количество шагов,вид тренировки,продолжительность", например "3456,Ходьба,3h00m".
// Возвращает если корректный формат nil или возвращает описание ошибки.
func (t *Training) Parse(datastring string) (err error) {
	sliceParse := strings.Split(datastring, ",")
	if len(sliceParse) != 3 {
		return errors.New("invalid input format")
	}
	steps, err := strconv.Atoi(sliceParse[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return errors.New("step count must be positive")
	}
	t.Steps = steps
	trainingType := sliceParse[1]
	if trainingType == "" {
		return errors.New("training type not specified")
	}
	t.TrainingType = trainingType
	duration, err := time.ParseDuration(sliceParse[2])
	if err != nil {
		return err
	}
	if duration <= 0 {
		return errors.New("training duration must be positive")
	}
	t.Duration = duration
	return nil
}

// ActionInfo рассчитывает и возвращает строку с информацией о тренировке.
// Возвращает отформатированную строку с типом тренировки, длительностью, дистанцией,
// скоростью, количеством сожженных калорий и ошибкой, если она возникает.
func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)
	meanspeed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	switch t.TrainingType {
	case "Ходьба":
		walkingSpentCalories, err := spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			fmt.Println(err)
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			t.TrainingType, t.Duration.Hours(), distance, meanspeed, walkingSpentCalories), nil
	case "Бег":
		runningSpentCalories, err := spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
		if err != nil {
			fmt.Println(err)
			return "", err
		}
		return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n",
			t.TrainingType, t.Duration.Hours(), distance, meanspeed, runningSpentCalories), nil
	}
	return "", errors.New("неизвестный тип тренировки")
}
