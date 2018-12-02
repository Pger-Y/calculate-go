package calculate

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	return
	s := Stack{}
	fmt.Println(s.Pop())
	fmt.Println(s.Push(1))
	fmt.Println(s.Pop())
	fmt.Println(s.Push(2))
	fmt.Println(s.Pop())
	fmt.Println(s.Push(3))
	fmt.Println(s.Pop())
	fmt.Println(s.data)
	fmt.Println("len:", len(s.data))
	fmt.Println("cap:", cap(s.data))
	fmt.Println(s.data, s.index)

}

func TestCal(t *testing.T) {
	exps := []string{}
	exps = append(exps, "1+2+3")                 //6
	exps = append(exps, "1+2*3")                 //7
	exps = append(exps, "2*2+3")                 //7
	exps = append(exps, "(2+2)+3")               //7
	exps = append(exps, "(2*2+3)")               //7
	exps = append(exps, "2*(2+3)")               //10
	exps = append(exps, "1*2+5*6")               // 32
	exps = append(exps, "1*(3+6)+1")             //10
	exps = append(exps, "1/(2+3)-6")             // -5.8
	exps = append(exps, "2/0")                   //error
	exps = append(exps, "2-(2+3)")               // -3
	exps = append(exps, "(2+(2+3*6)*2)*4-(2+3)") // -3
	exps = append(exps, "(((2-22)+6)*3)*(2+3)")  // -3
	for _, exp := range exps {
		Calculate(exp)
	}
}
