package weightconv

import "fmt"

type Pounds float64
type Kilograms float64

func (lb Pounds) String() string    { return fmt.Sprintf("%glb", lb) }
func (kg Kilograms) String() string { return fmt.Sprintf("%gkg", kg) }
