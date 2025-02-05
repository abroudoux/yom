package types

type Item struct {
	Name string
	Price string
	Quantity int
}

type Person struct {
	Name string
	Amount float64
	HasPaid bool
}

type Format int

const (
    Solo Format = iota
    Duo
    Trio
)

type Choice struct {
    Name    string
    Persons []*Person
    Format  Format
}

func (f Format) String() string {
    return [...]string{"Solo", "Duo", "Trio"}[f]
}
