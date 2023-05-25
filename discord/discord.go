package discord

import (
	"log"
	"math/big"

	"github.com/gtuk/discordwebhook"
)

func SendDiscordMessage(tokenName string, ammount *big.Int, txHash string, url string) {

	username := "ForgeLiquidation"

	USDCAmmount := ammount.Div(ammount, big.NewInt(1000000))

	content := "Liquidation Found in " + tokenName + "\n" + USDCAmmount.String() + " : USDC liquidated" + "\n" + " https://arbiscan.io/tx" + txHash

	message := discordwebhook.Message{
		Username: &username,
		Content:  &content,
	}

	errSend := discordwebhook.SendMessage(url, message)
	if errSend != nil {
		log.Fatal(errSend)
	}
}

func SendErrorMessage(err error, url string) {

	username := "ForgeLiquidation"

	content := "Error: " + err.Error()

	message := discordwebhook.Message{
		Username: &username,
		Content:  &content,
	}

	errSend := discordwebhook.SendMessage(url, message)
	if errSend != nil {
		log.Fatal(errSend)
	}
}
