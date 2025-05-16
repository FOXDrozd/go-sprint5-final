package actioninfo

import (
	"fmt"
)

// создайте интерфейс DataParser
type DataParser interface {
	Parse(dataString string) (err error) 
	ActionInfo() (string, error)
}

// создайте функцию Info()
func Info(dataset []string, dp DataParser) {
	for  _, v := range dataset{
		err := dp.Parse(v)
		if err != nil {
			errorStr := fmt.Errorf("Error function Info Parse:%v\n", err)
			fmt.Println(errorStr)
			continue
		}

		str, err := dp.ActionInfo()
		if err != nil {
			errorStr := fmt.Errorf("Error ActionInfo:%v\n", err)
			fmt.Println(errorStr)
			continue
		}
		fmt.Printf(str)
	}
}
