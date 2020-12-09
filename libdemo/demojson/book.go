package demojson

type Book struct {
	BookId int64   `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float64 `json:"price"`
	Hot    bool    `json:"hot"`
	Weight int64   `json:"-"`
}
