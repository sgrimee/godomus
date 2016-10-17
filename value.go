package godomus

type Value struct {
	Index       string `json:"index"`
	Value       string `json:"value"`
	Description string `json:"description"`
	Unit        string `json:"unit"`
	Label       string `json:"label"`
}
