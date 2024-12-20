// DO NOT EDIT: auto generated from github.com/Armatorix/enumgen

package example

import (
	"encoding/json"
	"fmt"
)

func (e IOTAConstType) String() string {
	switch e {
	case Const7:
		return "Const7"
	case Const8:
		return "Const8"
	case Const9:
		return "Const9"
	default:
		return "Unknown"
	}
}

func (e IOTAConstType) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())
}

func (e *IOTAConstType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Const7":
		*e = Const7
	case "Const8":
		*e = Const8
	case "Const9":
		*e = Const9
	default:
		return fmt.Errorf("Unknown IOTAConstType value: %+v", s)
	}
	return nil
}

func IOTAConstTypeValues() []IOTAConstType {
	return []IOTAConstType{
		Const7,
		Const8,
		Const9,
	}
}

func (e IOTAConstType) IsValid() bool {
	switch e {
	case Const7:
		return true
	case Const8:
		return true
	case Const9:
		return true
	default:
		return false
	}
}
