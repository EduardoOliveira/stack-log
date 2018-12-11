package core

import (
	"cloud.google.com/go/logging"
	"context"
	"flag"
	"log"
	"os"
	"strings"
)

func Run() {
	ctx := context.Background()

	// Creates a client.
	client, err := logging.NewClient(ctx, os.Args[1])
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	flag.Parse()

	m := make(map[string]string)

	for _, a := range flag.Args() {
		s := strings.Split(a, "=")
		if len(s) == 2 {
			m[s[0]] = s[1]
		}
	}

	err = client.Logger(os.Args[2]).LogSync(ctx, logging.Entry{
		Payload: m,
	})

	if err != nil {
		log.Println(err)
	}

}
