package models

type HeroAPIResponse struct {
	Response string           `json:"response"`
	Results  []HeroAPIResults `json:"results"`
}

type HeroAPIResults struct {
	ID          string      `json:"id"`
	Name        string      `json:"name"`
	Powerstats  Powerstats  `json:"powerstats"`
	Biography   Biography   `json:"biography"`
	Work        Work        `json:"work"`
	Connections Connections `json:"connections"`
	Image       Image       `json:"image"`
}

type Powerstats struct {
	Intelligence string `json:"intelligence"`
	Power        string `json:"power"`
}

type Biography struct {
	FullName  string `json:"full-name"`
	Alignment string `json:"alignment"`
}

type Work struct {
	Occupation string `json:"occupation"`
}

type Connections struct {
	GroupAffiliations string `json:"group-affiliation"`
	Relatives         string `json:"relatives"`
}

type Image struct {
	URL string `json:"url"`
}

type Super struct {
	HeroAPIResults
	UUID string `json:"uuid"`
}
