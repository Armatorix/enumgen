// DO NOT EDIT: auto generated from github.com/Armatorix/enumgen

package example

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

var StringConstTypeToString = map[StringConstType]string{
	Const1:                        "Const1",
	Const2:                        "Const2",
	Const3:                        "Const3",
	MixedStringConstTypePlacement: "MixedStringConstTypePlacement",
}

var StringToStringConstType = map[string]StringConstType{
	"Const1":                        Const1,
	"Const2":                        Const2,
	"Const3":                        Const3,
	"MixedStringConstTypePlacement": MixedStringConstTypePlacement,
}

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
	if v, ok := StringToStringConstType[s]; ok {
		*e = v
		return nil
	}
	return fmt.Errorf("Unknown StringConstType value: %s; %w", s, ErrStringConstTypeMissingValue)
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
	_, ok := StringConstTypeToString[e]
	return ok
}

// valuer and scanner for database/sql

func (e StringConstType) Value() (driver.Value, error) {
	return e.String(), nil
}

var ErrStringConstTypeMissingValue = errors.New("missing value")

func (e *StringConstType) Scan(value interface{}) error {
	if value == nil {
		return ErrStringConstTypeMissingValue
	}
	switch v := value.(type) {
	case []byte:
		return e.UnmarshalJSON(v)
	case string:
		return e.UnmarshalJSON([]byte(v))
	default:
		return fmt.Errorf("Unsupported type: %T", v)
	}
}
