package main

import "email-bot/emailbot"

func main() {
	botManager := emailbot.NewManager()

	botManager.DataProfile().
		AddField("firstname", 1).
		WithChunk("bank", "name").
		AddField("lastname", 1).
		WithChunk("bank", "name").
		AddField("fullname", 1).
		WithChunk("derived", "firstanme").
		WithChunk("derived", "firstanme").
		AddField("email", 30).
		WithModifiedChunk("derived", "fullname", "slang").
		WithChunk("literal", "@hotmail.com").
		AddField("day", 1).
		WithChunk("bank", "dayvault").
		AddField("month", 1).
		WithChunk("bank", "monthvault").
		AddField("year", 1).
		WithChunk("bank", "yearvault").
		AddField("password", 20).
		WithChunk("bank", "passvault")

	botManager.AddBot(8081)
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

	// setup := database.NewSetup()
	// setup.Setup()
}
