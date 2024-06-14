package factory

import (
	"context"
	"errors"
	"fmt"
	"time"
)

//START_FIELD OMIT
type RoadCourier struct {
	isMissing	bool
	estimate	time.Duration
	item		*RoadItem
}
//END_FIELD OMIT

func NewRoadCourier(isMissing bool,estimate time.Duration,item *RoadItem) *RoadCourier {
	return &RoadCourier{
		isMissing: isMissing,
		estimate: estimate,
		item: item,
	}
}
func (roadCourier RoadCourier) Delivery (ctx context.Context) (	Item,error) {
	done := make(chan bool)
	ticker := time.NewTicker(time.Second)
	
	go func() {
		defer close(done)
		time.Sleep(roadCourier.estimate)
		done <- true
	}()
	for {
		select {
//START_DELIV OMIT	
		case <-done:
			ticker.Stop()
			if roadCourier.isMissing {
				return nil,NewCourierError(ErrSteal,"Your Item Has Been Steal By Courier")
			}
			fmt.Println("Your Item Has Arrived")
			return roadCourier.item, nil
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
//END_DELIV OMIT
	}
}
