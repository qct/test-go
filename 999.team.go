package main

import "fmt"

type Player struct {
    name  string
    level int
}

type Team *[]Player

func (t *Team) Join(p *Player) {
    *t = append(*t, p)
}

func (t *Team) Quit(p *Player)  {
    for i, v := range *t {
        if(v.name == p.name) {
            copy((*t)[i:], (*t)[i+1:])
            *t = (*t)[:len(*t)-1]
            return
        }
    }
}

func main() {
    a := &Player{name:"a", level:1}
    fmt.Println(a)
}
