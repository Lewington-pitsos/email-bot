package main

import (
	"email-bot/database"
)

func main() {
	//master := master.NewMaster()
	//master.Scrape()
	//setup := database.NewSetup()
	//setup.Setup()
	profile := map[string]string{
		"day":       "19",
		"email":     "Ame123hystAudriana@hotmail.com",
		"firstname": "Amethyst",
		"fullname":  "AmethystAudriana",
		"lastname":  "Audriana",
		"month":     "November",
		"password":  "bS0x2(M(OCSL4tR%ZNIF",
		"year":      "1961",
		"birthdate": "02/06/1993",
	}

	archivist := database.NewArchivist()
	archivist.RecordProfile(profile)
}

func mainn() {

}
