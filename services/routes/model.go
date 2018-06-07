package routes

// Route ...
type Route struct {
	ID        string
	Modality  int
	Threshold int
	CreatedAt string
	UpdatedAt string
}

// Routes ...
type Routes []Route
