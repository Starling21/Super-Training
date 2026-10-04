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
	parts := strings.Split(data, ",")
	if len(parts) != 2 {
		fmt.Println("Неверное количество параметров")
		os.Exit(1)
	}
	//получаем количество шагов
	n, err1 := strconv.Atoi(parts[0])
	if err1 != nil {
		fmt.Printf("Строка '%s' не является числом: '%v'\n", parts[0], err1)
	}
	//получаем время тренировки из второго элемента
	m, err2 := time.ParseDuration(parts[1])
	if err2 != nil {
		fmt.Printf("Строка '%s' не является временем: '%v'\n", parts[1], err2)
	}
	if m <= 0 {
		m = 0
		fmt.Println("Время тренировки = 0")
	}
	//возвращаем количество шагов, время тренировки, ошибку
	return n, m, err1
}

func DayActionInfo(data string, weight, height float64) string {
	//
	steps, timeTr, err := parsePackage(data)
	//считаем дистанцию в км
	distance := stepLength * float64(steps) / mInKm
	fmt.Println(steps, timeTr, err)
	cal, err2 := spentcalories.WalkingSpentCalories(steps, weight, height, timeTr)
	training := fmt.Sprintf("Количество шагов: '%d'\nДистанция составила: '%.2f'км\nВы сожгли: '%.2f'\n", steps, distance, cal)
	fmt.Println(err2)
	return training
}
