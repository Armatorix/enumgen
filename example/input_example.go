//go:generate go run github.com/Armatorix/enumgen@v0.0.1 input_example.go output_example_string.go StringConstType
//go:generate go run github.com/Armatorix/enumgen@v0.0.1 input_example.go output_example_int.go IntConstType
//go:generate go run github.com/Armatorix/enumgen@v0.0.1 input_example.go output_example_iota.go IOTAConstType

package example

type StringConstType string

const (
	Const1 StringConstType = "const1"
	Const2 StringConstType = "const2"
	Const3 StringConstType = "const3"
)

type IntConstType int

const (
	Const4 IntConstType = 1
	Const5 IntConstType = 2
	Const6 IntConstType = 3
)

type IOTAConstType int

const (
	Const7 IOTAConstType = iota
	Const8
	Const9
)

const MixedStringConstTypePlacement StringConstType = "const10"
