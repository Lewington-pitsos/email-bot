package main

import "email-bot/emailbot"

func main() {
	botManager := emailbot.NewManager()

	botManager.DataProfile.
		AddField("firstname", 1).
		__WithChunk("bank", "name").
		AddField("lastname", 1).
		__WithChunk("bank", "name").
		AddField("fullname", 1).
		__WithChunk("derived", "firstanme").
		__WithChunk("derived", "firstanme").
		AddField("email", 30).
		__WithModifiedChunk("derived", "fullname", "slang").
		__WithChunk("literal", "@hotmail.com").
		AddField("day", 1).
		__WithChunk("bank", "dayvault").
		AddField("month", 1).
		__WithChunk("bank", "monthvault").
		AddField("year", 1).
		__WithChunk("bank", "yearvault").
		AddField("password", 20).
		__WithChunk("bank", "passvault")

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
