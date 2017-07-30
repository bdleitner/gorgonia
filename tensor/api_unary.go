package tensor

import "github.com/pkg/errors"

/*
GENERATED FILE. DO NOT EDIT
*/

func Neg(a Tensor, opts ...FuncOpt) (retVal Tensor, err error) {
	var e Engine = a.Engine()
	if e == nil {
		e = StdEng{}
	}
	if neger, ok := e.(Neger); ok {
		return neger.Neg(a, opts...)
	}
	err = errors.Errorf("Engine does not perform Neg")
	return
}

func Inv(a Tensor, opts ...FuncOpt) (retVal Tensor, err error) {
	var e Engine = a.Engine()
	if e == nil {
		e = StdEng{}
	}
	if inver, ok := e.(Inver); ok {
		return inver.Inv(a, opts...)
	}
	err = errors.Errorf("Engine does not perform Inv")
	return
}

func Square(a Tensor, opts ...FuncOpt) (retVal Tensor, err error) {
	var e Engine = a.Engine()
	if e == nil {
		e = StdEng{}
	}
	if squareer, ok := e.(Squarer); ok {
		return squareer.Square(a, opts...)
	}
	err = errors.Errorf("Engine does not perform Square")
	return
}

func Cube(a Tensor, opts ...FuncOpt) (retVal Tensor, err error) {
	var e Engine = a.Engine()
	if e == nil {
		e = StdEng{}
	}
	if cubeer, ok := e.(Cuber); ok {
		return cubeer.Cube(a, opts...)
	}
	err = errors.Errorf("Engine does not perform Cube")
	return
}

func Exp(a Tensor, opts ...FuncOpt) (retVal Tensor, err error) {
	var e Engine = a.Engine()
	if e == nil {
		e = StdEng{}
	}
	if exper, ok := e.(Exper); ok {
		return exper.Exp(a, opts...)
	}
	err = errors.Errorf("Engine does not perform Exp")
	return
}

func Tanh(a Tensor, opts ...FuncOpt) (retVal Tensor, err error) {
	var e Engine = a.Engine()
	if e == nil {
		e = StdEng{}
	}
	if tanher, ok := e.(Tanher); ok {
		return tanher.Tanh(a, opts...)
	}
	err = errors.Errorf("Engine does not perform Tanh")
	return
}

func Log(a Tensor, opts ...FuncOpt) (retVal Tensor, err error) {
	var e Engine = a.Engine()
	if e == nil {
		e = StdEng{}
	}
	if loger, ok := e.(Loger); ok {
		return loger.Log(a, opts...)
	}
	err = errors.Errorf("Engine does not perform Log")
	return
}

func Log2(a Tensor, opts ...FuncOpt) (retVal Tensor, err error) {
	var e Engine = a.Engine()
	if e == nil {
		e = StdEng{}
	}
	if log2er, ok := e.(Log2er); ok {
		return log2er.Log2(a, opts...)
	}
	err = errors.Errorf("Engine does not perform Log2")
	return
}

func Log10(a Tensor, opts ...FuncOpt) (retVal Tensor, err error) {
	var e Engine = a.Engine()
	if e == nil {
		e = StdEng{}
	}
	if log10er, ok := e.(Log10er); ok {
		return log10er.Log10(a, opts...)
	}
	err = errors.Errorf("Engine does not perform Log10")
	return
}

func Sqrt(a Tensor, opts ...FuncOpt) (retVal Tensor, err error) {
	var e Engine = a.Engine()
	if e == nil {
		e = StdEng{}
	}
	if sqrter, ok := e.(Sqrter); ok {
		return sqrter.Sqrt(a, opts...)
	}
	err = errors.Errorf("Engine does not perform Sqrt")
	return
}

func Cbrt(a Tensor, opts ...FuncOpt) (retVal Tensor, err error) {
	var e Engine = a.Engine()
	if e == nil {
		e = StdEng{}
	}
	if cbrter, ok := e.(Cbrter); ok {
		return cbrter.Cbrt(a, opts...)
	}
	err = errors.Errorf("Engine does not perform Cbrt")
	return
}

func InvSqrt(a Tensor, opts ...FuncOpt) (retVal Tensor, err error) {
	var e Engine = a.Engine()
	if e == nil {
		e = StdEng{}
	}
	if invsqrter, ok := e.(InvSqrter); ok {
		return invsqrter.InvSqrt(a, opts...)
	}
	err = errors.Errorf("Engine does not perform InvSqrt")
	return
}

func Abs(a Tensor, opts ...FuncOpt) (retVal Tensor, err error) {
	var e Engine = a.Engine()
	if e == nil {
		e = StdEng{}
	}
	if abser, ok := e.(Abser); ok {
		return abser.Abs(a, opts...)
	}
	err = errors.Errorf("Engine does not perform Abs")
	return
}

func Sign(a Tensor, opts ...FuncOpt) (retVal Tensor, err error) {
	var e Engine = a.Engine()
	if e == nil {
		e = StdEng{}
	}
	if signer, ok := e.(Signer); ok {
		return signer.Sign(a, opts...)
	}
	err = errors.Errorf("Engine does not perform Sign")
	return
}
