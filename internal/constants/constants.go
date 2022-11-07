package constants

import (
	"math"
	"strconv"
)

var (
	Values = map[string]string{}

	constants = map[string]string{
		"e":   strconv.FormatFloat(math.E, 'f', -1, 64),
		"pi":  strconv.FormatFloat(math.Pi, 'f', -1, 64),
		"phi": strconv.FormatFloat(math.Phi, 'f', -1, 64),
	}
)

type Const struct {
	Name  string
	Value string
}

func Init() {
	for val, key := range constants {
		newConst := Const{
			Name:  val,
			Value: key,
		}

		Values[newConst.Name] = newConst.Value
	}
}
