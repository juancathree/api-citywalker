package domain

type MatrixCost struct {
	City        string                        `validate:"required,lowecase" json:"city" bson:"city"`
	AverageCost float64                       `validate:"required" json:"averageCost" bson:"averageCost"`
	Costs       map[string]map[string]float64 `validate:"required" json:"costs" bson:"costs"`
	MaxLong     float64                       `validate:"required,longitude" json:"maxLong" bson:"maxLong"`
	MaxLat      float64                       `validate:"required,latitude" json:"maxLat" bson:"maxLat"`
	MinLong     float64                       `validate:"required,longitude" json:"minLong" bson:"minLong"`
	MinLat      float64                       `validate:"required,latitude" json:"minLat" bson:"minLat"`
}
