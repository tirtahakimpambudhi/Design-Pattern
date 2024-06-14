package factory

import "fmt"


type (
	status string  	
	KG int
	RoadItem struct {
		statusItem  	status
		weight 			KG
		name 			string
	}
	SeaItem struct {
		statusItem  	status
		weight 			KG
		name 			string
	}
)	

const (
	Goods 		status		= 	"New" 
	Damaged		status		=   "Second"
	maxWeight 	int 		= 	100

)

func NewRoadItem( name string , statusItem status , weight KG) (*RoadItem,error) {
	if int(weight) >= maxWeight {
		return nil,NewCourierError(ErrOverLoad,fmt.Sprintf("maximal weight in road items %d KG",maxWeight))
	}
	return &RoadItem{name: name,statusItem: statusItem,weight: weight},nil
}

func (itemImpl RoadItem) GetItem() string {
	return fmt.Sprintf("Name Item : %s \nStatus Item : %s ",itemImpl.name,itemImpl.statusItem)
}

func NewSeaItem( name string , statusItem status , weight KG) (*SeaItem,error) {
	if int(weight) >= (maxWeight * 10) {
		return nil,NewCourierError(ErrOverLoad,fmt.Sprintf("maximal weight in sea items %d KG",maxWeight*10))
	}
	return &SeaItem{name: name,statusItem: statusItem,weight: weight},nil
}

func (itemImpl SeaItem) GetItem() string {
	return fmt.Sprintf("Name Item : %s \nStatus Item : %s ",itemImpl.name,itemImpl.statusItem)
}