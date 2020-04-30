package Error

import "errors"

//Op give me info for whick metot create the current error
type Op string

//Kind give me info to the level of error
type Kind int

//Myerror is extend for add some property to info whre why or when
type Myerror struct {
	Op   Op
	Kind Kind
	Err  error
}

//Error metot for type cast Myerror to error
func (e Myerror) Error() string {
	return string(Ops(*e))
}

func singleError() *Myerror {
	err := &Myerror{Op: "singleError", Kind: 404, Err: errors.New("could not open")}
	//OR--
	// // errNew := E(Op("single"), Kind("infomsu"), errors.New("ds"))// kabuklanmış kendi errorüm
	return err
}

func E(args ...interface{}) *Myerror {
	e := &Myerror{}
	for _, arg := range args {
		switch arg := arg.(type) {
		case Op:
			e.Op = arg
		case error:
			e.Err = arg
		case Kind:
			e.Kind = arg
		default:
			panic("better call type")
		}
	}
	return e
}
func Ops(e *Myerror) []Op {
	res := []Op{e.Op}
	subError, ok := e.Err.(*Myerror)
	if !ok {
		return res
	}
	res = append(res, Ops(subError)...)
	return res
}
