package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"go-sprint5-final/internal/personaldata"
	"go-sprint5-final/internal/spentenergy"
)

var (
	ErrorConvData =  errors.New("this data isn't correction conv")
	ErrorCorrectionData =  errors.New("this data isn't correction")
	ErrorUnknownTrainingType =errors.New("unknown training type")
)

// создайте структуру Training
type Training struct {
	personaldata.Personal
	Steps int
	TrainingType string
	Duration time.Duration
}


// создайте метод Parse()
func (t *Training) Parse(dataString string) (err error) {
	data := strings.Split(dataString, ",")

	if len(data) != 3 {
		return  ErrorCorrectionData
	}

	steps, err := strconv.Atoi(data[0])
	if err != nil {
		return ErrorConvData
	}
	t.Steps = steps

	trainingType := data[1]
	if trainingType != "Бег" && trainingType != "Ходьба" {
		return ErrorCorrectionData
	}
	t.TrainingType = trainingType

	duration, err := time.ParseDuration(data[2])
	if err != nil {
		return ErrorConvData
	}
	t.Duration = duration

	return nil
}

// создайте метод ActionInfo()
func (t Training) ActionInfo() (string, error) {
	distance := spentenergy.Distance(t.Steps)

	if t.Duration <= 0 {
		return "", ErrorCorrectionData
	}

	meanSpeed := spentenergy.MeanSpeed(t.Steps,t.Duration)

	var dataTraining float64
	switch t.TrainingType {
		case "Бег": {
			dataTraining = spentenergy.RunningSpentCalories(t.Steps, t.Personal.Weight, t.Duration)
		}
		case "Ходьба": {
			dataTraining = spentenergy.WalkingSpentCaloriess(t.Steps, t.Personal.Weight, t.Personal.Height, t.Duration)
		}
		default:{
			return "неизвестный тип тренировки", ErrorUnknownTrainingType
		}
	}

	str := fmt.Sprintf("Тип тренировки: %s\nДлительность: %0.2f ч.\nДистанция: %0.2f км.\nСкорость: %0.2f км/ч\nСожгли калорий: %0.2f\n\n", t.TrainingType, t.Duration.Hours(), distance, meanSpeed, dataTraining)
	return str, nil
}

