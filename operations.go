package main

import("math")

type Operation struct{
	Num1 float64	`json: "num1"`
	Num2 float64 	`json: "num2"`
	TypeOp string	`json: "typeop`
}
type Sum Operation
type Sub Operation
type Mul Operation
type Div Operation
type Pow Operation
type Rot Operation

func (sum Sum) Calculate(num1 float64, num2 float64) float64{
	return num1 + num2
}

func (sub Sub)  Calculate(num1 float64, num2 float64) float64{
	return num1 - num2
}

func (mul Mul) Calculate(num1 float64, num2 float64) float64{
	return num1 * num2
}

func (div Div) Calculate(num1 float64, num2 float64) float64{
	return num1 / num2
}

func (pow Pow) Calculate(num1 float64, num2 float64) float64{
	return math.Pow(num1,num2)
}

func (rot Rot) Calculate(num1 float64, num2 float64) float64{
	return math.Pow(num1,1/num2)
}


