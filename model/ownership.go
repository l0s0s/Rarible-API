package model

type Ownership struct {
	ID            string    `json:"id"`
	Blockchain    string    `json:"blockchain"`
	ItemID        string    `json:"itemId"`
	Contract      string    `json:"contract"`
	Collection    string    `json:"collection"`
	TokenID       string    `json:"tokenId"`
	Owner         string    `json:"owner"`
	Value         string    `json:"value"`
	Source        string    `json:"source"`
	CreatedAt     string    `json:"createdAt"`
	LastUpdatedAt string    `json:"lastUpdatedAt"`
	Creators      []Creator `json:"creators"`
	LazyValue     string    `json:"lazyValue"`
	Version       int64     `json:"version"`
}

type Creator struct {
	Account string `json:"account"`
	Value   int64  `json:"value"`
}
