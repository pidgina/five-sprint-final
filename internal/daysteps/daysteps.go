package daysteps

import (
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

func (ds *DaySteps) Parse(datastring string) (err error) {
	temp := strings.Split(datastring, ",")

	if len(temp) != 2 {
		return fmt.Errorf("Введенные данные некорректны. Проверьте ввод.")
	}

	step, err := strconv.Atoi(temp[0])
	if err != nil {
		return fmt.Errorf("invalid steps format: %w", err)
	}
	if step <= 0 {
		return fmt.Errorf("Введенные данные некорректны. Проверьте ввод.")
	}

	ds.Steps = step

	d, err := time.ParseDuration(temp[1])
	if err != nil {
		return err
	}
	ds.Duration = d

	if d <= 0 {
		return fmt.Errorf("Введенные данные некорректны. Проверьте ввод.")
	}

	return nil

}

func (ds DaySteps) ActionInfo() (string, error) {

	if ds.Steps <= 0 || ds.Duration <= 0 || ds.Personal.Height <= 0 || ds.Personal.Weight <= 0 {
		return "", fmt.Errorf("Введенные данные некорректны. Проверьте ввод.")
	}

	dist := spentenergy.Distance(ds.Steps, float64(ds.Personal.Height))
	ccal, err := spentenergy.WalkingSpentCalories(ds.Steps,
		float64(ds.Personal.Weight),
		float64(ds.Personal.Height),
		ds.Duration)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		`Количество шагов: %d.
Дистанция составила %.2f км.
Вы сожгли %.2f ккал.
`,
		ds.Steps,
		dist,
		ccal), nil

}
