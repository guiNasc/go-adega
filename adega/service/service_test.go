package service_test

import (
	"testing"
	"github.com/guiNasc/go-adega/adega/model"
	"github.com/guiNasc/go-adega/adega/service"
) 

func TestShouldCreateWine(t *testing.T) {
	var wine = model.Wine{ Name: "Wine 1", Country: "France", Description: "A Very fine wine.", Year: 2015, Amount: 10}
	var wineSaved = service.Create_(wine)
	if (wineSaved.Name != "Wine 1" || wineSaved.Country != "France" || wineSaved.Description != "A Very fine wine." || wineSaved.Year != 2015 || wineSaved.Amount != 10) {
		 t.Error(wineSaved)
	 }
}

func TestShouldGetWine(t *testing.T) {
	var wine = model.Wine{ Name: "Wine 2", Country: "Belgium", Description: "Long story short.", Year: 1990, Amount: 2}
	var wineSaved = service.Create_(wine)
	var wineReturned = service.Get_("4")[0]
	if (wineReturned.Name != "Wine 2" || wineReturned.Country != "Belgium" || wineReturned.Description != "Long story short." || wineReturned.Year != 1990 || wineReturned.Amount != 2) {
		t.Error(wineSaved)
	}
}

func TestShouldUpdateWine(t *testing.T) {
	var wine = model.Wine{ Name: "Wine 3", Country: "England", Description: "Good grapes, good thoughts, good deeds...", Year: 1986, Amount: 1}
	var wineSaved = service.Create_(wine)
	wine.Name = "THE MOST AWESOME ONE"
	var wineReturned = service.Update_(wine)
	if (wineReturned.Name != "THE MOST AWESOME ONE" || wineReturned.Country != "England" || wineReturned.Description != "Good grapes, good thoughts, good deeds..." || wineReturned.Year != 1986 || wineReturned.Amount != 1) {
		t.Error(wineSaved)
	}
}

func TestShouldDeleteWine(t *testing.T) {
	var wine = model.Wine{ Name: "Wine 4", Country: "Ireland", Description: "Best wine you ever dream of.", Year: 2015, Amount: 10}
	service.Create_(wine)
	var countAll = service.GetAll_()
	var countDeleted = service.Delete_("6")
	var counted = (len(countAll)- len(countDeleted))
	if (counted != 1 ) {
		t.Error(counted)
	}
}