package main

import "log"

func main() {
	cfg := config{
		addr: ":8000",
	}
	app := &application{
		config: cfg,
	}

	log.Fatal(app.run())
}
