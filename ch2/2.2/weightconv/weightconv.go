package weightconv

import (
	"fmt"
)

type Pound float64
type Kilo float64

func (p Pound) String() string { return fmt.Sprintf("%g pounds", p) }
func (k Kilo) String() string  { return fmt.Sprintf("%g kilo", k) }

func (p Pound) PoundToKilo() Kilo { return Kilo(p * 0.453592) }
func (k Kilo) KiloToPound() Pound { return Pound(k * 2.20462) }
