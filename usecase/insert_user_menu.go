package usecase

import (
	"context"
	"strings"

	"github.com/zephyrzth/wdiet-be/model"
)

func (uc *usecase) InsertUserMenu(ctx context.Context, insertData model.InsertUserMenu) (model.InsertUserMenuWarning, error) {
	// get current user data
	user, err := uc.mongoRepo.GetUser(ctx, insertData.UserID)
	if err != nil {
		return model.InsertUserMenuWarning{}, err
	}
	standardNutrients := uc.getUserStandardNutrients(user)
	userNutrients := uc.getUserCurrentNutrients(ctx, user)

	// get menu data
	menu, err := uc.mongoRepo.GetMenuByID(ctx, insertData.MenuID)
	if err != nil {
		return model.InsertUserMenuWarning{}, err
	}
	menuNutrients := uc.getMenuNutrients(menu, insertData.Quantity)

	message := uc.checkInsertUserMenuWarning(standardNutrients, userNutrients, menuNutrients)
	warning := model.InsertUserMenuWarning{
		Message: message,
	}

	err = uc.mongoRepo.InsertUserMenu(ctx, insertData)

	return warning, err
}

func (uc *usecase) checkInsertUserMenuWarning(standardNutrients, userNutrients, menuNutrients nutrients) string {
	overNutrients := []string{}
	if userNutrients.calories+menuNutrients.calories > standardNutrients.calories {
		overNutrients = append(overNutrients, model.NutrientsCalories)
	}
	if userNutrients.carbohydrates+menuNutrients.carbohydrates > standardNutrients.carbohydrates {
		overNutrients = append(overNutrients, model.NutrientsCarbohydrates)
	}
	if userNutrients.fats+menuNutrients.fats > standardNutrients.fats {
		overNutrients = append(overNutrients, model.NutrientsFats)
	}
	if userNutrients.protein+menuNutrients.protein > standardNutrients.protein {
		overNutrients = append(overNutrients, model.NutrientsProtein)
	}
	if userNutrients.sugar+menuNutrients.sugar > standardNutrients.sugar {
		overNutrients = append(overNutrients, model.NutrientsSugar)
	}

	var messages string
	if len(overNutrients) > 0 {
		messages = strings.Join(overNutrients, ", ") + " will exceed the health standard if you buy the menu at this quantity"
	}

	return messages
}

func (uc *usecase) getMenuNutrients(menu model.Menu, quantity int) nutrients {
	menuNutrients := nutrients{}
	for _, composition := range menu.Compositions {
		for _, nutrient := range composition.Nutrients {
			switch nutrient.Name {
			case model.NutrientsCalories:
				menuNutrients.calories += (nutrient.Amount * float32(quantity))
			case model.NutrientsCarbohydrates:
				menuNutrients.carbohydrates += (nutrient.Amount * float32(quantity))
			case model.NutrientsFats:
				menuNutrients.fats += (nutrient.Amount * float32(quantity))
			case model.NutrientsProtein:
				menuNutrients.protein += (nutrient.Amount * float32(quantity))
			case model.NutrientsSugar:
				menuNutrients.sugar += (nutrient.Amount * float32(quantity))
			}
		}
	}

	return menuNutrients
}
