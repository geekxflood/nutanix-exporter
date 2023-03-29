package main

type MetadataObjectResponse struct {
	Kind          string `json:"kind"`
	TotalMatches  int    `json:"total_matches"`
	SortAttribute string `json:"sort_attribute"`
	Filter        string `json:"filter"`
	Length        int    `json:"length"`
	SortOrder     string `json:"sort_order"`
	Offset        int    `json:"offset"`
}

type EntitiesObject struct {
	Status     StatusObject   `json:"status"`
	Spec       SpecObject     `json:"spec"`
	APIVersion string         `json:"api_version"`
	Metadata   MetadataObject `json:"metadata"`
}
