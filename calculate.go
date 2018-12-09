package calculate

import (
	"fmt"
)

const (
	ADD        = "+"
	SUB        = "-"
	MUL        = "*"
	MULx       = "x"
	MULX       = "X"
	MULP       = "×"
	DIV        = "/"
	StartOfExp = "("
	EndOfExp   = ")"
)

var (
	DIGMAP = map[string]struct{}{
		"0": struct{}{},
		"1": struct{}{},
		"2": struct{}{},
		"3": struct{}{},
		"4": struct{}{},
		"5": struct{}{},
		"6": struct{}{},
		"7": struct{}{},
		"8": struct{}{},
		"9": struct{}{},
	}
)

func cal(a float64, op string, b float64) (float64, error) {
	switch op {
	case ADD:
		return a + b, nil
	case SUB:
		return a - b, nil
	case MUL:
		return a * b, nil
	case MULX:
		return a * b, nil
	case MULx:
		return a * b, nil
	case MULP:
		return a * b, nil
	case DIV:
		if b == 0 {
			return 0, fmt.Errorf("zero divide error")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("Unknown operator")
	}
}

type Stack struct {
	data  []interface{}
	index int
}

func (s *Stack) Pop() (interface{}, error) {
	if s.index == 0 || len(s.data) == 0 {
		return 0, fmt.Errorf("Empty Stack")
	}
	v := s.data[s.index-1]
	s.index--
	return v, nil

}

func (s *Stack) Empty() bool {
	return s.index == 0
}

func (s *Stack) Push(f interface{}) error {
	if s.index == len(s.data) {
		s.data = append(s.data, f)
		s.index++
	} else {
		s.data[s.index] = f
		s.index++
	}
	return nil
}

type parser struct {
	cache_v  float64
	stack_v  Stack
	stack_op Stack
}

func (p *parser) PushValue(f float64) error {
	//fmt.Println("In value :", f)
	return p.stack_v.Push(f)
}

func (p *parser) PopValue() (float64, error) {
	f, err := p.stack_v.Pop()
	if err == nil {
		//fmt.Println("Out value:", f)
		return f.(float64), nil
	} else {
		return 0, err
	}
}

func (p *parser) PushOp(o string) error {
	return p.stack_op.Push(o)
}

func (p *parser) PopOp() (string, error) {
	op, err := p.stack_op.Pop()
	if err == nil {
		return op.(string), nil
	} else {
		return "", err
	}
}

func isDigit(s string) bool {
	_, ok := DIGMAP[s]
	return ok
}

func (p *parser) Eval(exp string) (int, float64, error) {
	//l := len(exp)
	exp_rune := []rune(exp)
	l := len(exp_rune)
	var is_p, vop bool
	var dec float64
	i := 0
	for i < l {
		b := string(exp_rune[i])
		//fmt.Println(b)
		switch {
		//case b >= "0" && b <= "9":
		case isDigit(b):
			if is_p {
				p.cache_v = p.cache_v + (float64(b[0])-48)*dec
				dec = dec * 0.1
			} else {
				p.cache_v = p.cache_v*10 + float64(b[0]) - 48
			}
			if i == l-1 {
				p.pushCache()
			}
			i++
		case b == ".":
			if is_p != false {
				return 0, 0, fmt.Errorf("Too many points")
			}
			is_p = true
			dec = 0.1
			i++
		case b == MUL || b == MULx || b == MULX || b == DIV || b == MULP:
			is_p = false
			p.pushCache()
			if vop {
				p.cal()
			}
			p.PushOp(b)
			vop = true
			i++
		case b == ADD || b == SUB:
			is_p = false
			p.pushCache()
			if p.canCal() {
				p.cal()
			}
			vop = false
			p.PushOp(b)
			i++
		case b == StartOfExp:
			p2 := parser{}
			j, v, err := p2.Eval(exp[i+1:])
			if err != nil {
				return i, 0, err
			}
			p.cache_v = v
			i = i + j + 2
			if i == l {
				p.pushCache()
			}
		case b == EndOfExp:
			p.pushCache()
			v, err := p.calAll()
			return i, v, err
		default:
			i++

		}
		//		fmt.Println("value:", p.stack_v)
		//		fmt.Println("op:", p.stack_op)
	}
	v, err := p.calAll()

	return i, v, err

}

func (p *parser) pushCache() {
	p.PushValue(p.cache_v)
	p.cache_v = 0
}

func (p *parser) canCal() bool {
	return p.stack_op.Empty() == false
}

func (p *parser) value() (float64, error) {
	if p.stack_op.Empty() == true && p.stack_v.index == 1 {
		return p.stack_v.data[0].(float64), nil
	}
	fmt.Println("debug value", p)
	return 0, fmt.Errorf("Not yet")
}

func (p *parser) calAll() (float64, error) {
	var rv float64
	for {
		if p.canCal() {
			v, err := p.cal()
			if err != nil {
				return 0, err
			} else {
				rv = v
			}
		} else {
			v, err := p.value()
			if err == nil {
				rv = v
			} else {
				return 0, err
			}
			return rv, nil

		}

	}
}

func (p *parser) cal() (float64, error) {
	op, err_op := p.PopOp()
	if err_op != nil {
		return 0, fmt.Errorf("Pop op error")
	}
	v2, err_v := p.PopValue()
	if err_v != nil {
		return 0, fmt.Errorf("Pop value2 error")
	}
	v1, err_v := p.PopValue()
	if err_v != nil {
		return 0, fmt.Errorf("Pop value1 error")
	}
	v, err := cal(v1, op, v2)
	if err != nil {
		return v, err
	} else {
		p.PushValue(v)
	}
	return v, err
}

func Calculate(exp string) (float64, error) {
	p := parser{}
	_, v, err := p.Eval(exp)

	//fmt.Println("index:", i, "result:", v, "error:", err)
	return v, err
}
