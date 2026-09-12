package techpalace
import (
    "strings"
    "fmt"
    )
// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, " + strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	
    stars := strings.Repeat("*",numStarsPerLine) 
    message := stars+"\n"+welcomeMsg+"\n"+stars
    fmt.Println(message)
    return message
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	cleanedMessage := strings.ReplaceAll(strings.ReplaceAll(oldMsg, "*",""), "\n","")
	fmt.Println("str",strings.Trim(cleanedMessage, " "))
    return strings.Trim(cleanedMessage, " ")
}
