package bidder

type Bidder struct {
	Name            string
	InitialBid      int
	MaxBid          int
	IncrementAmount int
	currentBid      int
	timesIncreased  int
}

func (b Bidder) CurrentBid() int {
	return b.currentBid
}

func (b Bidder) CanIncrement() bool {
	return b.currentBid+b.IncrementAmount <= b.MaxBid
}

func (b *Bidder) IncrementCurrentBid() {
	if b.currentBid+b.IncrementAmount <= b.MaxBid {
		b.timesIncreased++
		b.currentBid += b.IncrementAmount
	}
}

func New(name string, initialBid, maxBid, incrementAmount int) Bidder {
	return Bidder{
		Name:            name,
		InitialBid:      initialBid,
		MaxBid:          maxBid,
		IncrementAmount: incrementAmount,
		currentBid:      initialBid,
		timesIncreased:  0,
	}
}

type Heap []Bidder

func (h Heap) Less(i, j int) bool {
	if h[i].CurrentBid() == h[j].CurrentBid() {
		return h[i].timesIncreased > h[j].timesIncreased
	}

	return h[i].CurrentBid() < h[j].CurrentBid()
}

func (h Heap) Len() int { return len(h) }

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(newElement any) {
	*h = append(*h, newElement.(Bidder))
}

func (h *Heap) Pop() any {
	tmp := *h
	heapSize := len(tmp)
	elementToRemove := tmp[heapSize-1]
	*h = tmp[0 : heapSize-1]
	return elementToRemove
}
