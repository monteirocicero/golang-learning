package main

import "fmt"

type ShoppingList []string

func printSlice(slice []string) {
    fmt.Println("Slice: ", slice)
}

func printList(list ShoppingList) {
    fmt.Println("Shopping list: ", list);
}

func main() {
    list := ShoppingList{"Lettuce", "Tuna", "Olive oil"};
    slice := []string{"Lettuce", "Tuna", "Olive oil"}

    printSlice([]string(list));
    printList(ShoppingList(slice));

}
