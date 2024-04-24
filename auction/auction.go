package auction

import (
	"container/heap"
	"fmt"
	"github.com/brendontj/sturdy-umbrella/bidder"
)

var (
	ErrNoBidders = fmt.Errorf("there's no bidder in the auction")
	EmptyBidder  = bidder.Bidder{}
)

type Auction struct {
	bidders bidder.Heap
	winner  bidder.Bidder
}

func New(bidders []bidder.Bidder) *Auction {
	bidderHeap := &bidder.Heap{}
	heap.Init(bidderHeap)
	for _, b := range bidders {
		heap.Push(bidderHeap, b)
	}

	return &Auction{
		bidders: *bidderHeap,
		winner:  EmptyBidder,
	}
}

func (a *Auction) GetWinner() bidder.Bidder {
	return a.winner
}

func (a *Auction) Run() error {
	if a.bidders.Len() == 0 {
		return ErrNoBidders
	}

	for !a.hasWinner() {
		if a.bidders.Len() == 1 {
			a.winner = heap.Pop(&a.bidders).(bidder.Bidder)
			continue
		}

		bidderWithSmallBid := heap.Pop(&a.bidders).(bidder.Bidder)
		if !bidderWithSmallBid.CanIncrement() {
			continue
		}
		bidderWithSmallBid.IncrementCurrentBid()
		heap.Push(&a.bidders, bidderWithSmallBid)
	}
	return nil
}

func (a *Auction) hasWinner() bool {
	return a.winner != EmptyBidder
}
