package part

type LayoutCfg struct {
	FrontTemplate string `yaml:"frontTemplate"` // path to
	BackTemplate  string `yaml:"backTemplate"`  // path to template
}

type Component struct {
	Kind         string
	Name         string
	MachineName  string `yaml:"machineName"`
	Description  string
	Instructions string
	Attributes   map[string]string
	Layout       LayoutCfg
	Quantity     int
}

type Suit struct {
	Component    Component
	InitialValue float64
}

type Card struct {
	Component    Component
	InitialValue float64 `yaml:"initialValue"`
	Suit         Suit
	Associations []Association
}

type Association struct {
	Cards []Card
	Value float64
}

type Deck struct {
	Component Component
	Cards     []Card
}

type Project struct {
	Component Component
	Decks     []Deck
}
