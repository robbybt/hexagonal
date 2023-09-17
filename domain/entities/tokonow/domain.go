package tokonow

type TokonowCategory struct {
	Id         string     `json:"id"`
	Name       string     `json:"name"`
	Identifier string     `json:"identifier"`
	Url        string     `json:"url"`
	Applinks   string     `json:"applinks"`
	Tree       int32      `json:"tree"`
	Child      []string   `json:"child"`
	Navigation Navigation `json:"navigation"`
	ImageUrl   string     `json:"imageUrl"`
	ParentID   string     `json:"parentID"`
	RootID     string     `json:"rootID"`
	RootName   string     `json:"rootName"`
	IsAdult    int32      `json:"isAdult"`
	IsKyc      bool       `json:"isKyc"`
	MinAge     int32      `json:"minAge"`
}

type Navigation struct {
	PrevID string `json:"prev"`
	NextID string `json:"next"`
}
