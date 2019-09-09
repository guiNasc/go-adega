package service

import (
	"github.com/guiNasc/go-adega/adega/model"
	"strconv"
)

func InitServiceMock() {
	model.Wines = append(model.Wines, model.Wine{ID: "1", Name: "Tenuta San Guido Bolgheri-Sassicaia Sassicaia", Country: "France", Description: "Rich and concentrated, this red features black currant, blackberry, violet, mineral and spice flavors. Dense yet lively, structured yet impeccably balanced, with vibrant acidity driving the long, fruit-filled aftertaste. The oak is beautifully integrated. Cabernet Sauvignon and Cabernet Franc. Best from 2023 through 2042. 17,200 cases made.", Year: 2015, Amount: 10})
	model.Wines = append(model.Wines, model.Wine{ID: "2", Name: "Domaine Giraud Châteauneuf-du-Pape", Country: "Denmarc", Description: "Enticing, with bright, engaging raspberry, cherry and red currant fruit infused with gently roasted sandalwood and apple wood, garrigue and licorice root notes. A swath of tobacco underscores the finish, giving this an old-school hint. Best from 2020 through 2040. 2,500 cases made." , Year: 2015, Amount: 10})
}

func filter(wines []model.Wine, id string  , f func(w model.Wine, id string) bool) []model.Wine {
    vsf := make([]model.Wine, 0)
    for _, v := range wines {
        if f(v, id) {
            vsf = append(vsf, v)
        }
    }
    return vsf
}

func exist(wine model.Wine, id string) bool {
	return wine.ID == id
}

func notExist(wine model.Wine, id string) bool {
	return wine.ID != id
}

func GetAll_() []model.Wine {
	return model.Wines
}

func Get_(id string) []model.Wine {
	return filter(model.Wines, id, exist)
}

func Create_(wine model.Wine) model.Wine {
	model.CountId = model.CountId + 1
	wine.ID = strconv.Itoa(model.CountId)
	model.Wines = append(model.Wines, wine)
	return wine
}

func Delete_(id string) []model.Wine {
	model.Wines = filter(model.Wines, id, notExist)
	return model.Wines
}

func Update_(wine model.Wine) model.Wine {
	model.Wines = filter(model.Wines, wine.ID, notExist)
	model.Wines = append(model.Wines, wine)
	return wine
}
