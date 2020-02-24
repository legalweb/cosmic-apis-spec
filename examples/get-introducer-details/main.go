package main

import (
	"context"
	"log"
	"os"

	identityClient "lwebco.de/cosmic-apis-spec/identity/client"
	identitiesOps "lwebco.de/cosmic-apis-spec/identity/client/identities"
	identityModels "lwebco.de/cosmic-apis-spec/identity/models"
	introducersClient "lwebco.de/cosmic-apis-spec/introducers/client"
	introducersOps "lwebco.de/cosmic-apis-spec/introducers/client/introducers"
	"lwebco.de/cosmic-apis-spec/pkg/auth"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if len(os.Args) < 4 {
		log.Fatal("please provide organisation, username and password args")
	}

	identities := identityClient.Default.Identities
	introducers := introducersClient.Default.Introducers

	loginReq := identitiesOps.NewLoginParams().WithContext(ctx).WithBody(&identityModels.V1LoginRequest{
		EmailAddress:   os.Args[2],
		Password:       os.Args[3],
		OrganisationID: []string{os.Args[1]},
	})
	loginRes, err := identities.Login(loginReq)
	if err != nil {
		log.Fatal("unable to login:", err)
	}

	session := auth.NewTokenProvider(loginRes.GetPayload().Token)
	log.Printf("acting as: <%s> %s\n", loginRes.GetPayload().AccountID, loginRes.GetPayload().EmailAddress)

	introducersReq := introducersOps.NewGetIntroducersParamsWithContext(ctx)
	introducersReq.AccountID = []string{loginRes.GetPayload().AccountID}
	introducersRes, err := introducers.GetIntroducers(introducersReq, session)
	if err != nil {
		log.Fatal("unable to get introducer:", err)
	}
	results := introducersRes.GetPayload().Items
	if len(results) < 1 {
		log.Fatal("account is not associated with an introducer")
	}

	log.Println("Got introducer details:")
	log.Printf("ID: %s\n", results[0].IntroducerID)
	log.Printf("Name: %s\n", results[0].Name)
}
