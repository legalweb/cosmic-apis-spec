package main

import (
	"context"
	"log"
	"os"

	"lwebco.de/cosmic-apis-spec/domainconfig/client"
	"lwebco.de/cosmic-apis-spec/domainconfig/client/domain_config"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if len(os.Args) < 2 {
		log.Fatalln("please provide a domain name to lookup")
	}

	c := client.Default.DomainConfig
	res, err := c.Resolve(&domain_config.ResolveParams{
		DomainName: os.Args[1],
		Context:    ctx,
	})

	if err != nil {
		log.Fatal("unable to lookup: ", err.Error())
	}

	log.Printf("Response: %+v", res.GetPayload().Config)
}
