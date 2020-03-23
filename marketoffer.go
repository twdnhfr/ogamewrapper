package ogame

// Marketoffer represent a offer from the Marketplace
type Marketoffer struct {
	Seller    string
	Resources Resources
	Offer     ID
	Amount    int64
}
