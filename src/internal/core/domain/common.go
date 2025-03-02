package domain

type Plugin struct {
	ID          string `bson:"_id" json:"id"`
	Name        string `bson:"name" json:"name"`
	Label       string `bson:"label" json:"label"`
	Category    string `bson:"catregory" json:"category"`
	Description string `bson:"description" json:"description"`
}
