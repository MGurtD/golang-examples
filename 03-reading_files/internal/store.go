package beerscli

// Store representation of store into data struct
type Store struct {
	StoreID int
	Name    string
	Address string
	Beers   []Beer
}
