package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/tirtahakimpambudhi/Design-Pattern/go/creational_pattern/factory"
)


func main () {
	var (
		ctx context.Context
		err error
		courier factory.Courier
		item	factory.Item
	)
	// Sea Courier With them item
	fmt.Println("==================== START SEA COURIER ==========================")
	item , err = factory.NewSeaItem("FOOD CARGO",factory.Goods,factory.KG(999))
	if err != nil {
		log.Fatal(err.Error())
	}
	courier = factory.NewSeaCourier(factory.BadWeather,false,time.Duration(10 * time.Second),item.(*factory.SeaItem))
	ctx , cancel := context.WithDeadline(context.Background(),time.Now().Add(time.Duration(14 * time.Second)))
	defer cancel()
	item , err = courier.Delivery(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(item.GetItem())
	fmt.Println("==================== END SEA COURIER ==========================")
	// Road Courier With them item
	fmt.Println("==================== START ROAD COURIER ==========================")
	item , err = factory.NewRoadItem("Laptop ACER",factory.Goods,factory.KG(2))
	if err != nil {
		log.Fatal(err.Error())
	}
	courier = factory.NewRoadCourier(false,time.Duration(10 * time.Second),item.(*factory.RoadItem))
	ctx , cancel = context.WithDeadline(context.Background(),time.Now().Add(time.Duration(14 * time.Second)))
	defer cancel()
	item , err = courier.Delivery(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(item.GetItem())	
	fmt.Println("==================== END ROAD COURIER ==========================")
}