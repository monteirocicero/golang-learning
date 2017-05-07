package main

import "fmt"


type State struct {
    name string
    population int
    capital string
}


func main() {
    states := make(map[string]State, 6)
    states["GO"] = State{"Goias", 6434052, "Goiania"}
    states["PB"] = State{"Paraiba", 3914418, "Joao Pessoa"}
    states["PR"] = State{"Parana", 10997462, "Curitiba"}
    states["RN"] = State{"Rio Grande do Norte", 3373960, "Natal"}
    states["AM"] = State{"Amazonas", 3807923, "Manaus"}
    states["SE"] = State{"Sergipe", 2228489, "Aracaju"}

    fmt.Println(states)

    saoPaulo, found := states["SP"]
    if found {
        fmt.Println(saoPaulo)
    }

    for abbreviation, state := range states {
        fmt.Printf("%s (%s) has %d inhabitant.\n", state.name, abbreviation, state.population)
    }


}
