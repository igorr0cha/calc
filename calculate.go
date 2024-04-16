package main
import ("errors")

type Calculation interface{
	Calculate()func(float64,float64)float64 
}

func (o Operation) Calculate()func(n1 float64, n2 float64)(float64,error){
	
	if result, ok := ValOps[o.TypeOp];ok{
			return func(n1 float64, n2 float64)(float64,error){
				return result(n1,n2), nil
	}
	}

return func(n1 float64, n2 float64)(float64,error){
	return 0, errors.New("Invalid operand request!")
}
}