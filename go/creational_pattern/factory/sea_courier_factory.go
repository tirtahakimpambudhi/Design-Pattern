package factory

import (
	"context"
	"errors"
	"fmt"
	"time"
)
//START_FIELD OMIT
type weather int

const (
	GoodWeather 	weather		= iota
	BadWeather		
)

type SeaCourier struct {
	weather 	weather
	tidalWaves	bool
	estimate	time.Duration
	item 		*SeaItem
}
//END_FIELD OMIT

func NewSeaCourier(weather weather,tidalWaves bool,estimate time.Duration,item *SeaItem) *SeaCourier {
	return &SeaCourier{
		weather: weather,
		tidalWaves: tidalWaves,
		estimate: estimate,
		item: item,
	}
}

func (seaCourier SeaCourier) Delivery (ctx context.Context) (Item,error) {
	done := make(chan bool)
	ticker := time.NewTicker(time.Second)
	
	go func() {
		defer close(done)
		
		if seaCourier.tidalWaves {
			time.Sleep(4 *time.Second)
		}

		if seaCourier.weather == BadWeather {
			time.Sleep(2 *time.Second)
		}

		time.Sleep(seaCourier.estimate)
		done <- true
	}()
	
	for {
		select {
		case <-done:
			ticker.Stop()
			fmt.Println("Your Item Has Arrived")
			return seaCourier.item, nil
		case <-ticker.C:
			fmt.Println("Your Item Has Been Delivered...")
		case <-ctx.Done():
			ticker.Stop()
			// Handle context cancellation and deadline exceeded errors
			if errors.Is(ctx.Err(), context.Canceled) {
				fmt.Println("Delivery was cancelled.")
				return nil, NewCourierError(ErrCancel, "item has been canceled by client or customer")
			}
			if errors.Is(ctx.Err(), context.DeadlineExceeded) {
				fmt.Println("Delivery deadline exceeded.")
				return nil, NewCourierError(ErrDeadline, "item delivery has exceeded the deadline set by client or customer")
			}
		}
	}
}