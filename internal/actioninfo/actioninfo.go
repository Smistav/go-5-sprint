package actioninfo

import (
	"fmt"
	"log"
)

type DataParser interface {
	Parse(datastring string) (err error)
	ActionInfo() (string, error)
}

func Info(dataset []string, dp DataParser) {
	if len(dataset) == 0 {
		log.Println("dataset is empty")
		return
	}
	for _, v := range dataset {
		err := dp.Parse(v)
		if err != nil {
			log.Println(err)
		}
	}
	str, err := dp.ActionInfo()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(str)
}
