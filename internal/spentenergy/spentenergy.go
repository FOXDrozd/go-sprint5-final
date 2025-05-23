package spentenergy

import (
	"time"
)

// Основные константы, необходимые для расчетов.
const (
	lenStep   = 0.65  // средняя длина шага.
	mInKm     = 1000  // количество метров в километре.
	minInH    = 60    // количество минут в часе.
	kmhInMsec = 0.278 // коэффициент для преобразования км/ч в м/с.
	cmInM     = 100   // количество сантиметров в метре.
	speed     = 1.39  // средняя скорость в м/с
)

// Константы для расчета калорий, расходуемых при ходьбе.
const (
	walkingCaloriesWeightMultiplier = 0.035 // множитель массы тела.
	walkingSpeedHeightMultiplier    = 0.029 // множитель роста.
)

// WalkingSpentCalories возвращает количество потраченных калорий при ходьбе.
//
// Параметры:
//
// steps int - количество шагов.
// weight float64 — вес пользователя.
// height float64 — рост пользователя.
// duration time.Duration — длительность тренировки.
//
// Создайте функцию ниже.
func WalkingSpentCaloriess(steps int, weight, height float64, duration time.Duration) float64 {
	if checkLessZero(weight) && checkLessZero(height) {
		return 0
	}

	if checkLessZero(float64(duration)) {
		return 0
	}

	meanSpeed := MeanSpeed(steps, duration)

	calories := ((walkingCaloriesWeightMultiplier * weight) + (meanSpeed*meanSpeed/height)*walkingSpeedHeightMultiplier) * duration.Hours() * minInH
	return calories
}

// Константы для расчета калорий, расходуемых при беге.
const (
	runningCaloriesMeanSpeedMultiplier = 18.0 // множитель средней скорости.
	runningCaloriesMeanSpeedShift      = 20.0 // среднее количество сжигаемых калорий при беге.
)

// RunningSpentCalories возвращает количество потраченных колорий при беге.
//
// Параметры:
//
// steps int - количество шагов.
// weight float64 — вес пользователя.
// duration time.Duration — длительность тренировки.
//
// Создайте функцию ниже.
func RunningSpentCalories(steps int, weight float64, duration time.Duration) float64 {
	if checkLessZero(weight) {
		return 0
	} 

	if checkLessZero(float64(duration)) {
		return 0
	}

	meanSpeed := MeanSpeed(steps, duration)

	calories:= ((runningCaloriesMeanSpeedMultiplier*meanSpeed)-runningCaloriesMeanSpeedShift) * weight
	return calories
}


// МeanSpeed возвращает значение средней скорости движения во время тренировки.
//
// Параметры:
//
// steps int — количество совершенных действий(число шагов при ходьбе и беге).
// duration time.Duration — длительность тренировки.
// 
// Создайте функцию ниже.
func MeanSpeed(steps int, duration time.Duration) float64 {
	//Проверить, что продолжительность duration больше 0. Если это не так — вернуть 0.
	if checkLessZero(float64(duration)) {
		return 0
	}

	distanceUser := Distance(steps)

	return distanceUser / duration.Hours()
}


// Distance возвращает дистанцию(в километрах), которую преодолел пользователь за время тренировки.
//
// Для расчета дистанции нужно шаги умножить на длину шага lenStep и разделить на mInKm
// Параметры:
//
// steps int — количество совершенных действий (число шагов при ходьбе и беге).
// 
// Создайте функцию ниже
func Distance(steps int) float64 {
	if(steps == 0){
		return 0
	}

	return (float64(steps) * lenStep) / mInKm
}


//Фугкция проверки на больше 0 
func checkLessZero(num float64) bool {
	if(num <= 0){
		return true
	}

	return false
}