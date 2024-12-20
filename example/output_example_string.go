// DO NOT EDIT: auto generated from github.com/Armatorix/enumgen

package example

import (
	"encoding/json"
	"fmt"
)

func (e StringConstType) String() string {
	switch e {
	case Const1:
		return "Const1"
	case Const2:
		return "Const2"
	case Const3:
		return "Const3"
	case MixedStringConstTypePlacement:
		return "MixedStringConstTypePlacement"
	default:
		return "Unknown"
	}
}

func (e StringConstType) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())
}

func (e *StringConstType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "Const1":
		*e = Const1
	case "Const2":
		*e = Const2
	case "Const3":
		*e = Const3
	case "MixedStringConstTypePlacement":
		*e = MixedStringConstTypePlacement
	default:
		return fmt.Errorf("Unknown StringConstType value: %+v", s)
	}
	return nil
}

func StringConstTypeValues() []StringConstType {
	return []StringConstType{
		Const1,
		Const2,
		Const3,
		MixedStringConstTypePlacement,
	}
}

func (e StringConstType) IsValid() bool {
	switch e {
	case Const1:
		return true
	case Const2:
		return true
	case Const3:
		return true
	case MixedStringConstTypePlacement:
		return true
	default:
		return false
	}
}
