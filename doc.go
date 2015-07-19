/*
Package irc implements a basic IRC client for Go programs, following RFC 1459.

This does not keep track of channel state.

Usage example:

	package main

	import (
		"github.com/Elemental-IRCd/irc"
	)

	func main() {
		bot := irc.New("MyBot", "foosmith")
		bot.UseTLS = true

		bot.AddCallback("001", func(*irc.Event) {
			bot.Join("#irc")
		})

		bot.Connect("irc.ponychat.net:6697")

		bot.Loop()
	}
*/
package irc
