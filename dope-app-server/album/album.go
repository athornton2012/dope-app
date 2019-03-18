package album

type Album struct {
	Id      string
	Artists []Artist
	Genres  []string
	Label   string
	Name    string
	Tracks  []Track
	Review  Review
}

type Review struct {
	Stars    int    `json:"id,omitempty"`
	Reviewer string `json:"id,omitempty"`
	Review   string `json:"id,omitempty"`
}

type Artist struct {
}

type Track struct {
}

func (a *Album) fetch(id string) (*Album, error) {
	return nil, nil
}
