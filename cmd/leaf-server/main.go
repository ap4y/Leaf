package main

import (
	"flag"
	"log"

	"github.com/ap4y/leaf"
	"github.com/ap4y/leaf/ui"
)

var (
	db      = flag.String("db", "leaf.db", "database location")
	count   = flag.Int("count", 20, "cards to review")
	addr    = flag.String("addr", ":8000", "addr for Web UI")
	devMode = flag.Bool("dev", false, "use local dev assets")
)

func main() {
	flag.Parse()

	db, err := leaf.OpenBoltStore(*db)
	if err != nil {
		log.Fatal("Failed to open stats DB: ", err)
	}

	defer db.Close()

	dm, err := leaf.NewDeckManager("./", db)
	if err != nil {
		log.Fatal("Failed to initialise deck manager: ", err)
	}

	srv := ui.NewServer(dm, *count)

	if err := srv.Serve(*addr, *devMode); err != nil {
		log.Fatal("Failed to render: ", err)
	}
}