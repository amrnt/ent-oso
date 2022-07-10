package main

import (
	"log"

	"github.com/amrnt/ent-oso/ent"
	"github.com/osohq/go-oso"
)

func main() {
	o, err := oso.NewOso()
	if err != nil {
		log.Print("failed to set up Oso\n")
	}

	log.Print("start: RegisterClass Tag")
	if err := o.RegisterClass(ent.Tag{}, nil); err != nil {
		log.Printf("failed to register class: Tag: %v\n", err)
	}
	log.Print("finish: RegisterClass Tag")

	log.Print(o)
}
