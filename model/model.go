package model

import "time"

const (
	ExerciseStatusVeryRare  = 1
	ExerciseStatusRare      = 2
	ExerciseStatusModerate  = 3
	ExerciseStatusOften     = 4
	ExerciseStatusVeryOften = 5

	UserGenderMale   = "Male"
	UserGenderFemale = "Female"

	NutrientsCalories      = "Calories"
	NutrientsFats          = "Fats"
	NutrientsProtein       = "Protein"
	NutrientsSugar         = "Sugar"
	NutrientsCarbohydrates = "Carbohydrates"
)

var mapExerciseStatusToFactor = map[int]float32{
	ExerciseStatusVeryRare:  1.2,
	ExerciseStatusRare:      1.375,
	ExerciseStatusModerate:  1.55,
	ExerciseStatusOften:     1.725,
	ExerciseStatusVeryOften: 1.9,
}

func ConvertExerciseStatusToFactor(status int) float32 {
	return mapExerciseStatusToFactor[status]
}

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
	ID           string        `bson:"_id,omitempty" json:"id"`
	RestaurantID string        `bson:"restaurant_id,omitempty" json:"restaurant_id"`
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

type User struct {
	ID              string           `bson:"_id,omitempty" json:"id"`
	Name            string           `bson:"name,omitempty" json:"name"`
	Email           string           `bson:"email,omitempty" json:"email"`
	Password        string           `bson:"password,omitempty" json:"password"`
	Age             int              `bson:"age,omitempty" json:"age"`
	Gender          string           `bson:"gender,omitempty" json:"gender"`
	Height          float32          `bson:"height,omitempty" json:"height"`
	Weight          float32          `bson:"weight,omitempty" json:"weight"`
	ExerciseStatus  int              `bson:"exercise_status,omitempty" json:"exercise_status"`
	Menus           []UserMenu       `bson:"menus,omitempty" json:"menus"`
	Reports         []Report         `bson:"reports,omitempty" json:"reports"`
	Recommendations []Recommendation `bson:"recommendations,omitempty" json:"recommendations"`
}

type UserMenu struct {
	ID        string    `bson:"id,omitempty" json:"id"`
	Quantity  int       `bson:"quantity,omitempty" json:"quantity"`
	Timestamp time.Time `bson:"timestamp,omitempty" json:"timestamp"`
}

type Report struct {
	Name           string  `json:"name"`
	StandardAmount float32 `json:"standard_amount"`
	CurrentAmount  float32 `json:"current_amount"`
}

type Recommendation struct {
	Name     string  `json:"name"`
	Duration float32 `json:"amount"`
}

type InsertUserMenu struct {
	UserID   string `json:"id"`
	MenuID   string `json:"menu_id"`
	Quantity int    `json:"quantity"`
}

type InsertUserMenuWarning struct {
	Message string `json:"message"`
}
