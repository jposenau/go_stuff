package main

import "fmt"

type Foo struct {
    name string
}

// SetName receives a pointer to Foo so it can modify it.
func (f *Foo) SetName(name string) {
    f.name = name
}

// Name receives a copy of Foo since it doesn't need to modify it.
func (f Foo) Name() string {
    return f.name
}

func main() {
    // Notice the Foo{}. The new(Foo) was just a syntactic sugar for &Foo{}
    // and we don't need a pointer to the Foo, so I replaced it.
    // Not relevant to the problem, though.
    p := Foo{}
    p.SetName("Abc")
    name := p.Name()
    fmt.Println(name)
}