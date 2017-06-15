package main

import "fmt"


type Operation interface {
    Calculate() int
}

type Sum struct {
    operator1, operator2 int
}

func (s Sum) Calculate() int {
    return s.operator1 + s.operator2
}

func (s Sum) String() string {
    return fmt.Sprintf("%d + %d", s.operator1, s.operator2)
}

type Subtraction struct {
    operator1, operator2 int
}

func (s Subtraction) Calculate() int {
    return s.operator1 - s.operator2
}

func (s Subtraction) String() string {
    return fmt.Sprintf("%d - %d", s.operator1, s.operator2)
}

func main() {
    operations := make([]Operation, 4)
    operations[0] = Sum{10, 20}
    operations[1] = Subtraction{30, 15}
    operations[2] = Subtraction{10, 50}
    operations[3] = Sum{5, 2}

    accumulator := 0
    for _, op := range operations {
        value := op.Calculate()
        fmt.Printf("%v = %d\n", op, value)
        accumulator += value
    }

    fmt.Println("Calculated value = ", accumulator)

    
}
