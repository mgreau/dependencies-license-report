package types

// Summary of the project to display in the report
type Summary struct {
	ProjectName  string       `json:"projectName"`
	Dependencies []Dependency `json:"dependencies"`
}

// License
type License struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	URL  string `json:"url"`
}

// Dependency used by the project
type Dependency struct {
	Name    string  `json:"name"`
	File    string  `json:"file"`
	License License `json:"license"`
}

// Params to generate a report
type Params struct {
	Name    string
	Format  string
	Project string
	Path    string
	// Output path to generate the report
	Output string
}
