package spentenergy

import (
	"fmt"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("Введенные данные некорректны. Проверьте ввод.")
	}

	temp := (weight * MeanSpeed(steps, height, duration) * duration.Minutes()) / minInH
	result := temp * walkingCaloriesCoefficient
	return result, nil

}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("Введенные данные некорректны. Проверьте ввод.")
	}

	temp := (weight * MeanSpeed(steps, height, duration) * duration.Minutes()) / minInH
	return temp, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {

	if steps <= 0 || height <= 0 || duration <= 0 {
		return 0
	}

	res := Distance(steps, height) / duration.Hours()
	return res
}

func Distance(steps int, height float64) float64 {
	stepLensOne := height * stepLengthCoefficient // длина шага
	temp := float64(steps) * stepLensOne
	res := temp / mInKm
	return res
}
