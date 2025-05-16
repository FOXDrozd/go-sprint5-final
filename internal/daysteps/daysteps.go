package daysteps

import (
	"errors"
	"fmt"
	"go-sprint5-final/internal/personaldata"
	"strconv"
	"strings"
	"time"

	"go-sprint5-final/internal/spentenergy"
)

var (
	ErrorConvData =  errors.New("this data isn't correction conv")
	ErrorCorrectionData =  errors.New("this data isn't correction")
)

const (
	StepLength = 0.65
)

// создайте структуру DaySteps
type DaySteps struct {
	personaldata.Personal
	Steps int
	Duration time.Duration
}


// создайте метод Parse()
func (ds *DaySteps) Parse(datastring string) (err error) {
	data := strings.Split(datastring, ",")

	if len(data) != 2 {
		return ErrorCorrectionData
	}


	steps, err := strconv.Atoi(data[0])
	if err != nil {
		return ErrorConvData
	}
	if steps == 0 {
		return ErrorCorrectionData
	}
	ds.Steps = steps

	duration, err := time.ParseDuration(data[1])

	if err != nil {
		return ErrorConvData
	}
	ds.Duration = duration

	return nil
}

// создайте метод ActionInfo()
func (ds DaySteps) ActionInfo() (string, error) {
	if ds.Duration <=0 {
		return "", ErrorCorrectionData
	}

	distance := spentenergy.Distance(ds.Steps)
	calories := spentenergy.WalkingSpentCaloriess(ds.Steps, ds.Personal.Weight, ds.Personal.Height, ds.Duration)

	str := fmt.Sprintf("Количество шагов: %d.\nДистанция составила %0.2f км.\nВы сожгли %0.2f ккал.\n\n", ds.Steps, distance, calories)

	return str, nil
}
