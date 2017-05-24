package main

import "fmt"

type ListGeneric []interface{}

func (list *ListGeneric) removeIndex(index int) interface{} {
    l := *list
    removed := l[index]
    *list = append(l[0:index], l[index+1:]...)
    return removed
}

func (list * ListGeneric) removeBegin() interface{} {
    return list.removeIndex(0)
}

func (list * ListGeneric) removeLast() interface{} {
    return list.removeIndex(len(*list)-1)
}

func main() {
    list := ListGeneric{1, "Coffee", 42, true, 23, "Ball", 3.14, false}

    fmt.Printf("List original:\n%v\n\n", list)


    fmt.Printf("Removing of the begin: %v, after removal:\n%v\n", list.removeBegin(), list)


    fmt.Printf("Removing of the last: %v, after removal:\n%v\n", list.removeLast() ,list)


    fmt.Printf("Removing of the index 3: %v, after removal:\n%v\n", list.removeIndex(3), list)


    fmt.Printf("Removing of the index 9: %v, after removal:\n%v\n", list.removeIndex(0), list)

  
    fmt.Printf("Removing of the last index: %v, after removal:\n%v\n", list.removeIndex(len(list)-1), list)  
}
