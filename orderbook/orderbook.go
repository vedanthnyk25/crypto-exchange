package orderbook

import (
	"fmt"
	"sort"
)

type Orderbook struct {
	Asks      []*Limit
	Bids      []*Limit
	AskLimits map[float64]*Limit
	BidLimits map[float64]*Limit
}

func NewOrderbook() *Orderbook {
	return &Orderbook{
		Asks:      []*Limit{},
		Bids:      []*Limit{},
		AskLimits: map[float64]*Limit{},
		BidLimits: map[float64]*Limit{},
	}
}

func (ob *Orderbook) AskTotalVolume() float64 {
	totalVolume := 0.0

	for i:=0; i < len(ob.Asks); i++ {
		totalVolume += ob.Asks[i].TotalVolume
	}
	return totalVolume
}

func (ob *Orderbook) BidTotalVolume() float64 {
	totalVolume := 0.0

	for i:=0; i < len(ob.Bids); i++ {
		totalVolume += ob.Bids[i].TotalVolume
	}
	return totalVolume
}

func (ob *Orderbook) PlaceMarketOrder(o *Order) []Match {
	matches := []Match{}

	if o.Bid {
		if o.Size > ob.AskTotalVolume() {
			panic(fmt.Errorf("order size %v is greater than total volume %v", o.Size, ob.AskTotalVolume()))
		}
		sort.Sort(ByBestAsk{ob.Asks}) 
		for _, limit := range ob.Asks {
			limitMatches := limit.Fill(o)
			matches = append(matches, limitMatches...)
			if o.IsFilled() {
				break
			}
		}
	} else {
		if o.Size > ob.BidTotalVolume() {
			panic(fmt.Errorf("order size %v is greater than total volume %v", o.Size, ob.BidTotalVolume()))
		}
		sort.Sort(ByBestBid{ob.Bids})
		for _, limit := range ob.Bids {
			limitMatches := limit.Fill(o)
			matches = append(matches, limitMatches...)
			if o.IsFilled() {
				break
			}
		}
	}
	return matches
}


func (ob *Orderbook) PlaceLimitOrder(price float64, o *Order) {
	var limit *Limit

	if o.Bid {
		limit = ob.BidLimits[price]
	} else {
		limit = ob.AskLimits[price]
	}

	if limit == nil {

		limit = NewLimit(price)
		limit.AddOrder(o)

		if o.Bid {
			ob.Bids = append(ob.Bids, limit)
			ob.BidLimits[price] = limit
		} else {
			ob.Asks = append(ob.Asks, limit)
			ob.AskLimits[price] = limit
		}
	}
}

//Temporary
func (ob *Orderbook) GetAsks() []*Limit{
    sort.Sort(ByBestAsk{ob.Asks})

		return ob.Asks
}

func (ob *Orderbook) GetBids() []*Limit{
		sort.Sort(ByBestBid{ob.Bids})

		return ob.Bids
}

