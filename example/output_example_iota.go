// DO NOT EDIT: auto generated from github.com/Armatorix/enumgen

package example

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

var IOTAConstTypeToString = map[IOTAConstType]string{
	Const7: "Const7",
	Const8: "Const8",
	Const9: "Const9",
}

var StringToIOTAConstType = map[string]IOTAConstType{
	"Const7": Const7,
	"Const8": Const8,
	"Const9": Const9,
}

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
	if v, ok := StringToIOTAConstType[s]; ok {
		*e = v
		return nil
	}
	return fmt.Errorf("Unknown IOTAConstType value: %s; %w", s, ErrIOTAConstTypeMissingValue)
}

func IOTAConstTypeValues() []IOTAConstType {
	return []IOTAConstType{
		Const7,
		Const8,
		Const9,
	}
}

func (e IOTAConstType) IsValid() bool {
	_, ok := IOTAConstTypeToString[e]
	return ok
}

// valuer and scanner for database/sql

func (e IOTAConstType) Value() (driver.Value, error) {
	return e.String(), nil
}

var ErrIOTAConstTypeMissingValue = errors.New("missing value")

func (e *IOTAConstType) Scan(value interface{}) error {
	if value == nil {
		return ErrIOTAConstTypeMissingValue
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
