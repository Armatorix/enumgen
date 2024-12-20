// DO NOT EDIT: auto generated from ddwrap

package example

import (
	"encoding/json"
	"fmt"
)

func (e IntConstType) String() string {
	switch e {
	case Const4:
		return "Const4"
	case Const5:
		return "Const5"
	case Const6:
		return "Const6"
	default:
		return "Unknown"
	}
}

func (e IntConstType) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())
}

func (e *IntConstType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Const4":
		*e = Const4
	case "Const5":
		*e = Const5
	case "Const6":
		*e = Const6
	default:
		return fmt.Errorf("Unknown IntConstType value: %+v", s)
	}
	return nil
}

func IntConstTypeValues() []IntConstType {
	return []IntConstType{
		Const4,
		Const5,
		Const6,
	}
}

func (e IntConstType) IsValid() bool {
	switch e {
	case Const4:
		return true
	case Const5:
		return true
	case Const6:
		return true
	default:
		return false
	}
}
