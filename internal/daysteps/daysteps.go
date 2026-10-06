package daysteps

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/spentcalories"
)

const (
	// Длина одного шага в метрах
	stepLength = 0.65
	// Количество метров в одном километре
	mInKm = 1000
)

func parsePackage(data string) (int, time.Duration, error) {
	//получаем из строки элементы для дальнейшей работы
	if data == "" {
		return 0, 0, fmt.Errorf("пустая строка")
	}
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("неверное количество параметров")
		//fmt.Println("Неверное количество параметров")
		os.Exit(1)
	}
	//получаем количество шагов
	steps, err1 := strconv.Atoi(parts[0])
	if err1 != nil {
		return 0, 0, fmt.Errorf("строка не является числом: %w", err1)
		//fmt.Printf("Строка %s не является числом: %v\n", parts[0], err1)
	}
	if steps <= 0 {
		return 0, 0, fmt.Errorf("шаги должны быть положительным числом")
	}
	//получаем время тренировки из второго элемента
	time, err2 := time.ParseDuration(parts[1])
	if err2 != nil {
		return 0, 0, fmt.Errorf("строка не является временем: %w", err2)
		//fmt.Printf("Строка %s не является временем: %v\n", parts[1], err2)
	}
	if time <= 0 {
		return 0, 0, fmt.Errorf("время тренировки равно нулю")
		//fmt.Println("Время тренировки = 0")
	}
	//возвращаем количество шагов, время тренировки, ошибку
	return steps, time, err1
}

func DayActionInfo(data string, weight, height float64) string {
	//
	steps, timeTr, err := parsePackage(data)
	if err != nil {
		return ""
	}
	if steps <= 0 {
		return ""
	}
	//считаем дистанцию в км
	distance := stepLength * float64(steps) / mInKm
	fmt.Println(steps, timeTr, err)
	cal, err2 := spentcalories.WalkingSpentCalories(steps, weight, height, timeTr)
	training := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n", steps, distance, cal)
	fmt.Println(err2)
	return training
}
