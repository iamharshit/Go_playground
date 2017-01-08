#go:IRC
An IRC bot written in go that displays the IRC msgs from other users and you can too participate in chat.Also the gobot echos what you have you have send.
It uses husio's IRC library.To install it : `go get github.com/husio/irc`

###Usage:
* Type `go run bot.go <freenode-name>` and get connected to your freenode.
* To private message a user type `PRIVMSG <nick-name> <message-to-be-sent>`
* To quit a freenode type `PART <freenode-name>` 
 
