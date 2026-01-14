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
	Steps int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(datastring string) (err error) {
	list := strings.Split(datastring, ",")
	if len(list) != 2 {
		return errors.New("incorrect number of parameters")
	}
	steps, err := strconv.Atoi(list[0])
	if err != nil {
		return err
	}
	if steps <= 0 {
		return errors.New("incorrect steps count")
	}
	duration, err := time.ParseDuration(list[1])
	if err != nil {
		return err
	}
	if duration <= 0 {
		return errors.New("incorrect duration")
	}
	ds.Duration = duration
	ds.Steps = steps
	return nil
}

func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)
	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", ds.Steps, distance, calories), nil
}
