package p3815

import "container/heap"

type AuctionSystem struct {
	bids  map[int]*PQ
	items map[int]map[int]*Item
}

func Constructor() AuctionSystem {
	return AuctionSystem{bids: make(map[int]*PQ), items: make(map[int]map[int]*Item)}
}

func (this *AuctionSystem) AddBid(userId int, itemId int, bidAmount int) {
	this.bid(userId, itemId, bidAmount)
}

func (this *AuctionSystem) UpdateBid(userId int, itemId int, newAmount int) {
	this.bid(userId, itemId, newAmount)
}

func (this *AuctionSystem) bid(userId int, itemId int, bidAmount int) {
	if v, ok := this.items[itemId][userId]; ok {
		v.priority = bidAmount
		heap.Fix(this.bids[itemId], v.index)
	} else {
		it := &Item{id: userId, priority: bidAmount}
		if _, ok := this.bids[itemId]; !ok {
			this.bids[itemId] = &PQ{}
			this.items[itemId] = make(map[int]*Item)
		}
		this.items[itemId][userId] = it
		heap.Push(this.bids[itemId], it)
	}
}

func (this *AuctionSystem) RemoveBid(userId int, itemId int) {
	it := this.items[itemId][userId]
	if it == nil {
		return
	}
	heap.Remove(this.bids[itemId], it.index)
	delete(this.items[itemId], userId)
}

func (this *AuctionSystem) GetHighestBidder(itemId int) int {
	if v, ok := this.bids[itemId]; ok && v.Len() > 0 {
		return (*v)[0].id
	}
	return -1
}

type Item struct {
	id       int
	priority int
	index    int
}

type PQ []*Item

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority || pq[i].priority == pq[j].priority && pq[i].id > pq[j].id
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PQ) Push(x any) {
	item := x.(*Item)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PQ) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

/**
 * Your AuctionSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddBid(userId,itemId,bidAmount);
 * obj.UpdateBid(userId,itemId,newAmount);
 * obj.RemoveBid(userId,itemId);
 * param_4 := obj.GetHighestBidder(itemId);
 */
