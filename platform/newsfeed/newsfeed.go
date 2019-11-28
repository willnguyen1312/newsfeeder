package newsfeed

// Getter do something
type Getter interface {
	GetAll() []Item
}

// Adder do something
type Adder interface {
	Add(item Item)
}

// Item struct
type Item struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

// Repo struct
type Repo struct {
	Items []Item
}

// New create a new Repo
func New() *Repo {
	return &Repo{
		Items: []Item{},
	}
}

// Add add new item to the repo
func (r *Repo) Add(item Item) {
	r.Items = append(r.Items, item)
}

// GetAll return all items in the repo
func (r *Repo) GetAll() []Item {
	return r.Items
}
