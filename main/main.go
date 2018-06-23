package main

import "email-bot/emailbot"

func main() {
	botManager := emailbot.NewManager()
	botManager.AddBot(9999)
	botManager.AddBot(8082)
	botManager.ScrapeAll()

	// profile := map[string]string{
	// 	"day":       "19",
	// 	"email":     "Ame123hystAudriana@hotmail.com",
	// 	"firstname": "Amethyst",
	// 	"fullname":  "AmethystAudriana",
	// 	"lastname":  "Audriana",
	// 	"month":     "November",
	// 	"password":  "bS0x2(M(OCSL4tR%ZNIF",
	// 	"year":      "1961",
	// 	"birthdate": "2/6/1993",
	// }

	// archivist := database.NewArchivist()
	// archivist.RecordProfile(profile)

	// profile := profile.NewActiveProfile("floob")
	// profile.Populate()
}

func mainn() {

}
