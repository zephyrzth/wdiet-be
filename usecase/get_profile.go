package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/zephyrzth/wdiet-be/model"
)

type userNutrients struct {
	calories      float32
	sugar         float32
	carbohydrates float32
	protein       float32
	fats          float32
}

func (uc *usecase) GetProfile(ctx context.Context, userID string) (model.User, error) {
	user, err := uc.mongoRepo.GetUser(ctx, userID)
	if err != nil {
		return model.User{}, err
	}
	standardNutrients := uc.getUserStandardNutrients(user)
	userNutrients := userNutrients{}

	curTimestamp := time.Now()
	for _, userMenu := range user.Menus {
		if userMenu.Timestamp.Day() != curTimestamp.Day() &&
			userMenu.Timestamp.Month() != curTimestamp.Month() &&
			userMenu.Timestamp.Year() != curTimestamp.Year() {
			continue
		}
		quantity := userMenu.Quantity
		menu, err := uc.mongoRepo.GetMenuByID(ctx, userMenu.ID)
		if err != nil {
			fmt.Println("[usecase][GetProfile] fail get menu with id:", userMenu.ID)
			continue
		}
		for _, composition := range menu.Compositions {
			for _, nutrient := range composition.Nutrients {
				switch nutrient.Name {
				case model.NutrientsCalories:
					userNutrients.calories += (nutrient.Amount * float32(quantity))
				case model.NutrientsCarbohydrates:
					userNutrients.carbohydrates += (nutrient.Amount * float32(quantity))
				case model.NutrientsFats:
					userNutrients.fats += (nutrient.Amount * float32(quantity))
				case model.NutrientsProtein:
					userNutrients.protein += (nutrient.Amount * float32(quantity))
				case model.NutrientsSugar:
					userNutrients.sugar += (nutrient.Amount * float32(quantity))
				}
			}
		}
	}

	user.Reports = []model.Report{
		{
			Name:           model.NutrientsCalories,
			StandardAmount: standardNutrients.calories,
			CurrentAmount:  userNutrients.calories,
		},
		{
			Name:           model.NutrientsCarbohydrates,
			StandardAmount: standardNutrients.carbohydrates,
			CurrentAmount:  userNutrients.carbohydrates,
		},
		{
			Name:           model.NutrientsFats,
			StandardAmount: standardNutrients.fats,
			CurrentAmount:  userNutrients.fats,
		},
		{
			Name:           model.NutrientsProtein,
			StandardAmount: standardNutrients.protein,
			CurrentAmount:  userNutrients.protein,
		},
		{
			Name:           model.NutrientsSugar,
			StandardAmount: standardNutrients.sugar,
			CurrentAmount:  userNutrients.sugar,
		},
	}

	user.Recommendations = []model.Recommendation{}

	if userNutrients.calories > standardNutrients.calories {
		diffCalories := userNutrients.calories - standardNutrients.calories
		user.Recommendations = []model.Recommendation{
			{
				Name:     "Jogging",
				Duration: diffCalories / 9.8,
			},
			{
				Name:     "Sprint",
				Duration: diffCalories / 12.467,
			},
			{
				Name:     "Bicycling",
				Duration: diffCalories / 8.567,
			},
			{
				Name:     "Swimming",
				Duration: diffCalories / 14,
			},
		}
	}

	return user, nil
}

// Count standard daily nutrients for a user, based on age, gender, weight, height
func (uc *usecase) getUserStandardNutrients(userData model.User) userNutrients {
	// formula to count BMP
	var standardCalories float32
	if userData.Gender == model.UserGenderMale {
		standardCalories = (88.4 + (13.4 * userData.Weight)) + (4.8 * userData.Height) - (5.68 * float32(userData.Age))
	} else if userData.Gender == model.UserGenderFemale {
		standardCalories = (447.6 + (9.25 * userData.Weight)) + (3.1 * userData.Height) - (4.33 * float32(userData.Age))
	}

	// add multiplier by exercise status
	standardCalories *= model.ConvertExerciseStatusToFactor(userData.ExerciseStatus)

	return userNutrients{
		calories:      standardCalories,
		sugar:         (0.1 * standardCalories) / 4,
		fats:          (0.15 * standardCalories) / 9,
		carbohydrates: (0.6 * standardCalories) / 4,
		protein:       (0.15 * standardCalories) / 4,
	}
}
