package model

type Restaurants struct {
	ID        string `bson:"_id,omitempty"`
	Name      string `bson:"name,omitempty"`
	IsWifi    bool   `bson:"is_wifi,omitempty"`
	IsMushola bool   `bson:"is_mushola,omitempty"`
	IsIndoor  bool   `bson:"is_indoor,omitempty"`
	OpenTime  int    `bson:"open_time,omitempty"`
	CloseTime int    `bson:"close_time,omitempty"`
	Phone     string `bson:"phone,omitempty"`
	Address   string `bson:"address,omitempty"`
	Menus     []Menu `bson:"menus,omitempty"`
}

type Menu struct {
	Title        string        `bson:"title,omitempty"`
	Price        int32         `bson:"price,omitempty"`
	Description  string        `bson:"description,omitempty"`
	Compositions []Composition `bson:"compositions,omitempty"`
}

type Composition struct {
	Name      string     `bson:"name,omitempty"`
	Nutrients []Nutrient `bson:"nutrients,omitempty"`
}

type Nutrient struct {
	Name   string  `bson:"name,omitempty"`
	Amount float32 `bson:"amount,omitempty"`
}
