package main

import "time"

type Match struct {
	Ask        *Order
	Bid        *Order
	SizeFilled float64
	Price      float64
}
type Order struct {
	Size      float64
	Limit     *Limit
	Bid       bool
	Timestamp int64
}

func NewOrder(bid bool, size float64) *Order {
	return &Order{
		Size:      size,
		Bid:       bid,
		Timestamp: time.Now().UnixNano(),
	}
}

func (l *Limit) AddOrder(o *Order) {
	o.Limit = l
	l.Orders = append(l.Orders, o)
	l.TotalVolume += o.Size
}

func (l *Limit) DeleteOrder(o *Order) {
	for i, order := range l.Orders {
		if order == o {
			l.Orders = append(l.Orders[:i], l.Orders[i+1:]...)
			break
		}
	}
	o.Limit = nil
	l.TotalVolume -= o.Size
}

type Limit struct {
	Price       float64
	Orders      []*Order
	TotalVolume float64
}

func NewLimit(price float64) *Limit {
	return &Limit{
		Price:  price,
		Orders: []*Order{},
	}
}

type Orderbook struct {
	Asks      []Limit
	Bids      []Limit
	AskLimits map[float64]*Limit
	BidLimits map[float64]*Limit
}

func NewOrderbook() *Orderbook {
	return &Orderbook{
		Asks:      []Limit{},
		Bids:      []Limit{},
		AskLimits: map[float64]*Limit{},
		BidLimits: map[float64]*Limit{},
	}
}

func (ob *Orderbook) PlaceOrder(price float64, o *Order) []Match{
  //1. Matching logic

	if o.Size>0{
		ob.add(price, o)
	}

	return []Match{}
}

func (ob *Orderbook) add(price float64, o *Order) {
    var limit *Limit
		
		if o.Bid{
      limit= ob.BidLimits[price]
		}else{
			limit= ob.AskLimits[price]
		}

		if limit==nil {

			limit= NewLimit(price)
			limit.AddOrder(o)

			if o.Bid{
				ob.Bids= append(ob.Asks, *limit)
				ob.BidLimits[price]=limit
			}else{
				ob.Asks= append(ob.Bids, *limit)
				ob.AskLimits[price]=limit
			}
		}
}

