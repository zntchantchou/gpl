package main

import (
	"flag"
	"fmt"
)

// flag.Value implements stringer (String())
/* celsius flag has
- String (from float64)(flag.Value)
- Set (flag.Value)
-
*/
// the goal is to create a flag name "temp"
// the program must output the value as Celsius degrees following a format
//  we provide the set method to satisfy the flag.value

type Celsius float64

type celsiusFlag struct{ Celsius }

var f = createFlag("temp", 0.0, "temperature to be displayed as celsius")

func main() {
	fmt.Println("Celsius")
	flag.Parse()
	fmt.Println(*f)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%.2f°C", c)
}

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	// note that the computed value is stored inside the flag
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil

	case "F", "°F":
		f.Celsius = FtoC(value)
		return nil

	case "K", "°K":
		f.Celsius = KtoC(value)
		return nil

	default:
		return fmt.Errorf("invalid temperature %q", s)
	}
}

func FtoC(v float64) Celsius {
	return Celsius((v - 32) * (5.00 / 9.00))
}

func KtoC(v float64) Celsius {
	return Celsius(v - 273.15)
}

func createFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

/*
type Flag struct {
	Name     string // name as it appears on command line
	Usage    string // help message
	Value    Value  // value as set
	DefValue string // default value (as text); for usage message
}

type Value interface {
	String() string
	Set(string) error
}

*/
