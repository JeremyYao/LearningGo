package main

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	// ?
	return mToSend.sender.name != "" && mToSend.recipient.name != "" && mToSend.recipient.number != 0 && mToSend.sender.number != 0
}
