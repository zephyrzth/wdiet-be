package model

type Restaurants struct {
	ID        string `bson:"_id,omitempty" json:"id"`
	Name      string `bson:"name,omitempty" json:"restaurant_name"`
	IsWifi    bool   `bson:"is_wifi,omitempty" json:"is_wifi"`
	IsMushola bool   `bson:"is_mushola,omitempty" json:"is_mushola"`
	IsIndoor  bool   `bson:"is_indoor,omitempty" json:"is_indoor"`
	OpenTime  int    `bson:"open_time,omitempty" json:"open_time"`
	CloseTime int    `bson:"close_time,omitempty" json:"close_time"`
	Phone     string `bson:"phone,omitempty" json:"phone"`
	Address   string `bson:"address,omitempty" json:"address"`
	Menus     []Menu `bson:"menus,omitempty" json:"menus"`
}

type Menu struct {
	Title        string        `bson:"title,omitempty" json:"title"`
	Price        int32         `bson:"price,omitempty" json:"price"`
	Description  string        `bson:"description,omitempty" json:"description"`
	Compositions []Composition `bson:"compositions,omitempty" json:"compositions"`
}

type Composition struct {
	Name      string     `bson:"name,omitempty" json:"name"`
	Nutrients []Nutrient `bson:"nutrients,omitempty" json:"nutrients"`
}

type Nutrient struct {
	Name   string  `bson:"name,omitempty" json:"name"`
	Amount float32 `bson:"amount,omitempty" json:"amount"`
}
