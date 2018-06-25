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

	emailInput := "//input[@id='MemberName']"
	passInput := "//input[@id='PasswordInput']"
	firstInput := "//input[@id='FirstName']"
	lastInput := "//input[@id='LastName']"
	dayInput := "//select[@id='BirthDay']"
	monthInput := "//select[@id='BirthMonth']"
	yearInput := "//select[@id='BirthYear']"
	capchaBox := "//div[@id='hipTemplateContainer']"
	homeBanner := "//div[@id='loaded-home-banner-profile-section']"
	homeHeader := "//div[@id='headerUniversalHeader']"

	submitInput := "//input[@id='iSignupAction']"
	dateOption := "//option[@value='%s']"

	botManager.AddAction().
		AddVisit("https://signup.live.com/signup").
		AddWait(300)

	// ========================================================

	botManager.AddAction().
		AddFillOperation(emailInput, candidateValues["email"]).
		AddWait(200).
		AddSubmitOperation(submitInput).
		AddWait(300).

	// ========================================================

	botManager.AddAction().
		AddFillOperation(passInput, candidateValues["password"]).
		AddWait(200).
		AddSubmitOperation(submitInput).
		AddWait(300).

	// ========================================================

	botManager.AddAction().
		AddFillOperation(firstInput, candidateValues["firstname"]).
		AddFillOperation(lastInput, candidateValues["lastname"]).
		AddWait(200).
		AddSubmitOperation(submitInput).
		AddWait(300).

	// ========================================================

	botManager.AddAction().
		AddSelectOperation(dayInput, dateOption, candidateValues["day"]).
		AddSelectOperation(monthInput, dateOption, candidateValues["month"]).
		AddSelectOperation(yearInput, dateOption, candidateValues["year"]).
		AddWait(300).
		AddToSpec(action.CheckExists(submitInput)).
		AddToInteraction(action.Click(submitInput)).
		AddWait(4000).

	// ========================================================

	botManager.AddAction().
		AddToSpec(action.CheckExists(capchaBox)).
		AddWait(90000).

	// ========================================================

	botManager.AddAction().
		AddToSpec(action.CheckExists(homeBanner)).
		AddToSpec(action.CheckExists(homeHeader)).
		
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
