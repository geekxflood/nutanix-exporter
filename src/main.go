package main

import (
	"context"
	"log"
	"os"

	nutanix "github.com/tecbiz-ch/nutanix-go-sdk"
)

func main() {

	configCreds := nutanix.Credentials{
		Username: os.Getenv("NUTANIX_USERNAME"),
		Password: os.Getenv("NUTANIX_PASSWORD"),
	}

	log.Printf("Username: %s", configCreds.Username)
	log.Printf("Password: %s", configCreds.Password)
	opts := []nutanix.ClientOption{
		nutanix.WithCredentials(&configCreds),
		nutanix.WithEndpoint("https://PC"),
		nutanix.WithInsecure(),
	}

	client := nutanix.NewClient(opts...)

	ctx := context.Background()
	mycluster, err = client.Cluster.Get(ctx, "mycluster")

	list, err = client.VM.All(ctx)

}
