package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/gumby/proverbs"
)

var defaultAddress = ":3000"

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	addr := flag.String("addr", defaultAddress, "address to listen to")
	flag.Parse()

	store := proverbs.NewInMemStore()

	if err := http.ListenAndServe(*addr, proverbs.NewServer(store)); err != nil {
		return err
	}
	return nil
}
