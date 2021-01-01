package pixela

// GraphID is an ID for identifying the pixelation graph.
// Validation rule: ^[a-z][a-z0-9-]{1,16}
type GraphID string

// GraphType  is the type of quantity to be handled in the graph.
// Only int or float are supported.
type GraphType string

const (
	Int   GraphType = "int"
	Float GraphType = "float"
)

// GraphColor defines the display color of the pixel in the pixelation graph.
// shibafu (green)  momiji (red),
// sora (blue), ichou (yellow), ajisai (purple)
// and kuro (black) are supported as color kind.
type GraphColor string

const (
	Shibafu GraphColor = "shibafu"
	Momiji  GraphColor = "momiji"
	// TODO define another color type
)

type SelfSufficientType string

const (
	SelfSufficientIncrement SelfSufficientType = "increment"
	SelfSufficientDecrement SelfSufficientType = "decrement"
	SelfSufficientNone      SelfSufficientType = "none"
)

// GraphDefinition is graph definition.
type GraphDefinition struct {
	ID                  GraphID            `json:"id"`
	Name                string             `json:"name"`
	Unit                string             `json:"unit"`
	Type                GraphType          `json:"type"`
	Color               GraphColor         `json:"color"`
	TimeZone            string             `json:"timezone"`
	PurgeCacheURLs      []string           `json:"purgeCacheURLs"`
	SelfSufficient      SelfSufficientType `json:"selfSufficient"`
	IsSecret            bool               `json:"isSecret"`
	PublishOptionalData bool               `json:"publishOptionalData"`
}

type createGraphParams struct {
	ID                  GraphID            `json:"id"`
	Name                string             `json:"name"`
	Unit                string             `json:"unit"`
	Type                GraphType          `json:"type"`
	Color               GraphColor         `json:"color"`
	Timezone            string             `json:"timezone,omitempty"`
	SelfSufficient      SelfSufficientType `json:"selfSufficient,omitempty"`
	IsSecret            bool               `json:"isSecret,omitempty"`
	PublishOptionalData bool               `json:"publishOptionalData,omitempty"`
}
