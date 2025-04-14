// DO NOT EDIT: auto generated from github.com/Armatorix/enumgen

package example

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

var IntConstTypeToString = map[IntConstType]string{
	Const4: "Const4",
	Const5: "Const5",
	Const6: "Const6",
}

var StringToIntConstType = map[string]IntConstType{
	"Const4": Const4,
	"Const5": Const5,
	"Const6": Const6,
}

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
	if v, ok := StringToIntConstType[s]; ok {
		*e = v
		return nil
	}
	return fmt.Errorf("Unknown IntConstType value: %s; %w", s, ErrIntConstTypeMissingValue)
}

func IntConstTypeValues() []IntConstType {
	return []IntConstType{
		Const4,
		Const5,
		Const6,
	}
}

func (e IntConstType) IsValid() bool {
	_, ok := IntConstTypeToString[e]
	return ok
}

// valuer and scanner for database/sql

func (e IntConstType) Value() (driver.Value, error) {
	return e.String(), nil
}

var ErrIntConstTypeMissingValue = errors.New("missing value")

func (e *IntConstType) Scan(value interface{}) error {
	if value == nil {
		return ErrIntConstTypeMissingValue
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
