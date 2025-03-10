package main

import (
	"fmt"
	"os"

	"github.com/MechamJonathan/lotr-companion-app/styles"
)

var exitQuotes = []string{
	"“Yᴏᴜ ʜᴀᴠᴇ ғᴇʟᴛ ɪᴛs ᴘᴏᴡᴇʀ. Iᴛ ʜᴀs ᴛᴏᴜᴄʜᴇᴅ ʏᴏᴜ. Pᴇʀʜᴀᴘs ʏᴏᴜ ʜᴀᴠᴇ sᴇᴇɴ ᴍᴏʀᴇ ᴛʜᴀɴ ʏᴏᴜ sʜᴏᴜʟᴅ.”",
	"(Gandalf rushes over, seizes the Palantír from your hands, and covers it with his cloak.)",
	"“𝘐 𝘴𝘦𝘦 𝘺𝘰𝘶.”",
}

func commandExit(cfg *config, args ...string) error {
	quote := getRandomQuote(exitQuotes)
	fmt.Println(styles.StartUpQuote.Render(quote))
	MoveCursorToBottom()
	os.Exit(0)
	return nil
}
