package spentcalories

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	//"log"
)

// Основные константы, необходимые для расчетов.
const (
	//lenStep                    = 0.65 // средняя длина шага.
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе
)

func parseTraining(data string) (int, string, time.Duration, error) {
	//обрабатываем ошибки ввода
	if data == "" {
		return 0, "", 0, fmt.Errorf("пустая строка")
	}
	parts := strings.Split(data, ",")
	if len(parts) != 3 {
		//fmt.Println("Неверное количество параметров")
		//os.Exit(1)
		return 0, "", 0, fmt.Errorf("неверное количество параметров")
	}
	//делим строку на переменные, чтобы использовать их далее
	steps, err1 := strconv.Atoi(parts[0])
	if err1 != nil {
		return 0, "", 0, fmt.Errorf("шаги должны быть положительным числом: %w", err1)
		//fmt.Printf("Строка %s не является числом: %v\n", parts[0], err1)
	}
	if steps <= 0 {
		return 0, "", 0, fmt.Errorf("шаги должны быть больше нуля")
	}
	activity := parts[1]
	if activity == "" {
		return 0, "", 0, fmt.Errorf("тип тренировки не указан")
	}
	duration, err3 := time.ParseDuration(parts[2])
	if err3 != nil {
		return 0, "", 0, fmt.Errorf("неверный формат времени: %w", err3)
	}
	if duration <= 0 {
		return 0, "", 0, fmt.Errorf("время должно быть больше нуля")
	}
	return steps, activity, duration, nil
}

func distance(steps int, height float64) float64 {
	//считаем длина шага, опираясь на высоту человека
	lengthStep := height * stepLengthCoefficient
	//считаем длину дистанции
	distTr := (float64(steps) * (lengthStep)) / float64(mInKm)
	return distTr
}

func meanSpeed(steps int, height float64, duration time.Duration) float64 {
	//проверяем отрицательную продолжительность тренировки
	if duration <= 0 {
		return 0.0
	}
	//считаем среднюю скорость в км/ч
	distanceTr := distance(steps, height)
	meanSp := distanceTr / duration.Hours()
	return meanSp
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	//проверка корректности введённых данных
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		err := fmt.Errorf("Несоответствие введённых данных заданному типу")
		return 0, err
	}
	//считаем среднюю скорость
	meanSp := meanSpeed(steps, height, duration)
	//считаем потраченные калории
	spentCal := (weight * meanSp * duration.Minutes()) / minInH
	return spentCal, nil
}

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	//проверка корректности введённых данных
	if steps <= 0 || weight <= 0 || height <= 0 || duration <= 0 {
		return 0, fmt.Errorf("Несоответствие введённых данных заданному типу")
	}
	meanSp := meanSpeed(steps, height, duration)
	//if meanSp > 10 {
	//	err := fmt.Errorf("\nЭто уже не ходьба! %.2f км/ч", meanSp)
	//	result := (weight * meanSp * duration.Minutes()) / minInH * walkingCaloriesCoefficient
	//	return result, err
	//}
	result := ((weight * meanSp * duration.Minutes()) / float64(minInH)) * walkingCaloriesCoefficient
	return result, nil
}

func TrainingInfo(data string, weight, height float64) (string, error) {
	//задаём переменные для упрощения кода
	var cal float64
	var errCal error
	//парсим вводные данные
	stepsQn, typeTr, timeTr, err := parseTraining(data)
	//обрабатываем ошибку
	if err != nil {
		return "", fmt.Errorf("Ошибка парсинга: %w", err)
	}
	//перебираем варианты тренировки и ошибку
	switch typeTr {

	case "Ходьба":
		cal, errCal = WalkingSpentCalories(stepsQn, weight, height, timeTr)

	case "Бег":
		cal, errCal = RunningSpentCalories(stepsQn, weight, height, timeTr)

	default:
		return "", fmt.Errorf("неизвестный тип тренировки: %s", typeTr)
	}
	if errCal != nil {
		return "", fmt.Errorf("ошибка расчёта калорий: %w", errCal)
	}
	//вывод результата
	dist := distance(stepsQn, height)
	speed := meanSpeed(stepsQn, height, timeTr)
	//timeTr = duration.Hours(timeTr)
	report := fmt.Sprintf("Тип тренировки: %s\nДлительность: %.2f ч.\nДистанция: %.2f км.\nСкорость: %.2f км/ч\nСожгли калорий: %.2f\n", typeTr, timeTr.Hours(), dist, speed, cal)

	return report, nil
}
