package main

var sum Sum
var sub Sub
var mul Mul
var div Div
var pow Pow
var rot Rot

var ValOps = map[string]func(a float64, b float64) float64{
	"sum": sum.Calculate,
	"+"  : sum.Calculate,
	"sub": sub.Calculate,
	"-"  : sub.Calculate,
	"mul": mul.Calculate,
	"*"  : mul.Calculate,
	"div": div.Calculate,
	"/"  : div.Calculate,
	"pow": pow.Calculate,
	"^"  : pow.Calculate,
	"rot": rot.Calculate,
	"&"  : rot.Calculate,
}