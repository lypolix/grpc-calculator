package main

import (
	"context"
	"flag"
	"log"
	"strconv"

	"github.com/lypolix/grpc-calculator/pkg/api"
	"google.golang.org/grpc"
)

func main () {
	flag.Parse()
	if flag.NArg() < 2 {
		log.Fatal("not enough arguments")
	}

	x, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	y, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial(":9090", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	
	c := api.NewAdderClient(conn)
	res, err := c.Add(context.Background(), &api.AddRequest{X: int32(x), Y: int32(y)})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(res.GetResult())

}	