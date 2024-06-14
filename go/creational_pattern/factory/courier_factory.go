package factory

import "context"


type  (

	Item interface {
		GetItem() (string)
	}

	Courier interface {
		Delivery(context.Context) (Item,error)	
	}

)
