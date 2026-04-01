package trainings

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(datastring string) (err error) {
	temp := strings.Split(datastring, ",")

	if len(temp) != 3 {
		return fmt.Errorf("Введенные данные некорректны. Проверьте ввод.")
	}

	step, err := strconv.Atoi(temp[0])
	if err != nil {
		return err
	}
	t.Steps = step

	t.TrainingType = temp[1]

	d, err := time.ParseDuration(temp[2])
	if err != nil {
		return err
	}
	t.Duration = d

	if d <= 0 || step <= 0 {
		return fmt.Errorf("Введенные данные некорректны. Проверьте ввод.")
	}

	return nil
}

func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps, float64(t.Height))

	average := spentenergy.MeanSpeed(t.Steps, float64(t.Height), t.Duration)

	var temp float64

	switch t.TrainingType {
	case "Ходьба":
		temper, err := spentenergy.WalkingSpentCalories(t.Steps,
			float64(t.Weight), float64(t.Height), t.Duration)
		if err != nil {
			return "", err
		}
		temp = temper
	case "Бег":
		temper, err := spentenergy.RunningSpentCalories(t.Steps,
			float64(t.Weight), float64(t.Height), t.Duration)
		if err != nil {
			return "", err
		}
		temp = temper
	default:
		return "", fmt.Errorf("неизвестный тип тренировки")
	}

	return fmt.Sprintf(
		`Тип тренировки: %s
Длительность: %.2f ч.
Дистанция: %.2f км.
Скорость: %.2f км/ч
Сожгли калорий: %.2f
`,
		t.TrainingType,
		t.Duration.Hours(),
		distance,
		average,
		temp), nil
}
