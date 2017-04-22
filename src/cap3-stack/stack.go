package main

import(
    "errors"
    "fmt"
)

func main() {
    stack := Stack{}
    fmt.Println("Stack created with length ", stack.Length())
    fmt.Println("Empty? ", stack.Empty())

    stack.Push("Go")
    stack.Push(2009)
    stack.Push(3.14)
    stack.Push("End")
    fmt.Println("Length after push 4 values: ", stack.Length())
    fmt.Println("Empty? ", stack.Empty())


    for !stack.Empty() {
        value, _ := stack.Pop()
        fmt.Println("Removing ", value)
        fmt.Println("Length: ", stack.Length())
        fmt.Println("Empty? ", stack.Empty())
    }
   
    _, err := stack.Pop()
    if err != nil {
            fmt.Println("Error: ", err)
    }
    
}

type Stack struct {
    values []interface{}
}

func (stack Stack) Length() int {
    return len(stack.values)
}

func (stack Stack) Empty() bool {
    return stack.Length() == 0
}

func (stack *Stack) Push(value interface{}) {
    stack.values = append(stack.values, value)
}

func (stack *Stack) Pop() (interface{}, error) {
    if stack.Empty() {
        return nil, errors.New("Empty Stack!")
    }

    value := stack.values[stack.Length() - 1]
    stack.values = stack.values[:stack.Length() - 1]
    return value, nil
}
