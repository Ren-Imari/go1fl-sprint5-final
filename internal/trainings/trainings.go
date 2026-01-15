package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps int
	TrainingType string
	Duration time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	list := strings.Split(datastring, ",")
	if len(list) != 3 {
		return errors.New("incorrect number of parameters")
	}
	steps, err := strconv.Atoi(list[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return errors.New("incorrect steps count")
	}
	t.Steps = steps
	t.TrainingType = list[1]
	duration, err := time.ParseDuration(list[2])
	if err != nil {
		return err
	}
	if duration <= 0 {
		return errors.New("incorrect duration")
	}
	t.Duration = duration
	return nil
}

func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)
	var calories float64
	var err error
	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(t.Steps, t.Weight, t.Height, t.Duration)	
	default: 
		return "", errors.New("unknown training type")	
	}
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", t.TrainingType, t.Duration.Hours(), distance, speed, calories), nil
}
