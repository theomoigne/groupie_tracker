package helpers

import (
	"encoding/json"
	"errors"
	"net/http"
)

// variabels
var (
	Main      = []ForMainPage{}
	Artists   = []Artist{}
	Relations = Relation{}
)

// GetJSONArtists get data from file
func GetJSONArtists(w http.ResponseWriter) error {
	url := "https://groupietrackers.herokuapp.com/api/artists"
	resp, err := http.Get(url)

	if err != nil {
		return errors.New("500: API unreachable")
	}
	err = json.NewDecoder(resp.Body).Decode(&Artists)
	if err != nil {
		return errors.New("500: data not decoded")
	}

	url = "https://groupietrackers.herokuapp.com/api/relation"
	resp, err = http.Get(url)
	if err != nil {
		return errors.New("500: API unreachable")
	}
	err = json.NewDecoder(resp.Body).Decode(&Relations)
	if err != nil {
		return errors.New("500: data not decoded")
	}

	for i := range Artists {
		Dates := ""
		Loc := ""
		Rel := ""
		for i2, v := range Relations.Index[i].DatesLocations {
			Loc += i2 + "| "
			cap := ""
			for _, v2 := range v {
				cap += v2 + "| "
			}
			Dates += cap
			Rel += i2 + ": " + cap + "\n"
		}
		Artists[i].ConcertDates = Dates
		Artists[i].Locations = Loc
		Artists[i].Relations = Rel
	}
	return nil
}

// PrepMainStruct to index.html
func PrepMainStruct() {
	for _, v := range Artists {
		Main = append(Main, ForMainPage{Image: v.Image, Name: v.Name})
	}
}
