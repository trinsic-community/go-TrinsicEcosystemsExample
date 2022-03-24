package main

import (
	"context"
	"log"

	sdk "github.com/trinsic-id/sdk/go/proto"
	trinsic "github.com/trinsic-id/sdk/go/services"
)

func main() {
	options, err := trinsic.NewServiceOptions()
	if err != nil {
		log.Fatal(err)
	}

	accountService, err := trinsic.NewAccountService(options)
	if err != nil {
		log.Fatal(err)
	}

	authToken, _, err := accountService.SignIn(context.Background(), &sdk.SignInRequest{})
	if err != nil {
		log.Fatal(err)
	}

	options.ServiceOptions.AuthToken = authToken

	info, err := accountService.GetInfo(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	println(info.String())
}
