package spentenergy

import (
	"errors"
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

// WalkingSpentCalories рассчитывает количество потраченных каллорий при ходьбе.
// Принимает данные количество, вес пользователя в кг, рост в м, длительностью активности.
// Возвращает количество сожженных калорий и ошибку, если она возникает.
func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if weight <= 0 {
		return 0, errors.New("invalid weight value")
	}
	if height <= 0 {
		return 0, errors.New("invalid height value")
	}
	if steps <= 0 {
		return 0, errors.New("invalid step count")
	}
	if duration <= 0 {
		return 0, errors.New("invalid duration")
	}
	return (weight * MeanSpeed(steps, height, duration) * duration.Minutes()) / minInH * walkingCaloriesCoefficient, nil
}

// RunningSpentCalories рассчитывает количество потраченных каллорий при беге.
// Принимает данные количество, вес пользователя в кг, рост в м, длительностью активности.
// Возвращает количество сожженных калорий и ошибку, если она возникает.
func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if weight <= 0 {
		return 0, errors.New("invalid weight value")
	}
	if height <= 0 {
		return 0, errors.New("invalid height value")
	}
	if steps <= 0 {
		return 0, errors.New("invalid step count")
	}
	if duration <= 0 {
		return 0, errors.New("invalid duration")
	}
	return (weight * MeanSpeed(steps, height, duration) * duration.Minutes()) / minInH, nil
}

// MeanSpeed вычисляет среднюю скорость.
// Принимает количество шагов, рост пользователя в метрах, продолжительность активности.
// Возвращает среднюю скорость в километрах/час.
func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if duration <= 0 {
		return 0
	}
	return Distance(steps, height) / duration.Hours()
}

// Distance вычисляет пройденную дистанцию.
// Принимает количество шагов, рост пользователя в метрах.
// Возвращает дистанцию в километрах.
func Distance(steps int, height float64) float64 {
	stepLength := float64(height) * stepLengthCoefficient
	return float64(steps) * stepLength / mInKm
}
