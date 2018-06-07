package routes

// Route ...
type Route struct {
	ID        string `xml:"id" json:"id"`
	Modality  int    `xml:"modalityId" json:"modalityId"`
	Threshold int    `xml:"threshold"`
	CreatedAt string `xml:"created_at"`
	UpdatedAt string `xml:"updated_at"`
}

// Routes ...
type Routes []Route
