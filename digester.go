package main

const (
	senate = iota
	house  = iota
)

// Congress is
type Congress struct {
	Type    int
	Members []*Member
	Number  int
}

// Member describes the properties of the members of congress
type Member struct {
	ID        string
	FirstName string
	LastName  string
	Funders   []*Funder
	Bills     []*Bill
}

// Bill describes the property of a bill
type Bill struct {
	ID               string
	Type             string
	CongressNumber   int
	Status           string
	LastModifiedDate string
	Description      string
	DisplayNumber    string
	BillID           int
}

// Funder describes the properties of a funder
type Funder struct {
}

// Digester is
type Digester struct{}

// Consume is
func (d *Digester) Consume() {
	// First get all members of congress senate and house
	// look up member's top funders
	// grab all the bills sponsored or voted on by member
	// Get the category for each bill
	// set up our structs to match the consumed data
	// persist data
}
