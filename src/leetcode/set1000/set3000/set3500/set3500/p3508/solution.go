package p3508

type Router struct {
	que     [][3]int
	records map[[3]int]int
	head    int
	tail    int
	full    bool
	dests   map[int][]int
}

func Constructor(memoryLimit int) Router {
	que := make([][3]int, memoryLimit)
	records := make(map[[3]int]int)
	return Router{
		que:     que,
		records: records,
		head:    0,
		tail:    0,
		full:    false,
		dests:   make(map[int][]int),
	}
}

func (this *Router) removePacket(key [3]int) {
	delete(this.records, key)
	v := this.dests[key[1]]
	this.dests[key[1]] = v[1:]
}

func (this *Router) AddPacket(source int, destination int, timestamp int) bool {
	key := [3]int{source, destination, timestamp}
	if _, ok := this.records[key]; ok {
		return false
	}
	if this.full {
		// 删除最早的那个
		this.removePacket(this.que[this.tail])
		this.tail = (this.tail + 1) % len(this.que)
		this.full = false
	}
	next := (this.head + 1) % len(this.que)
	if next == this.tail {
		this.full = true
	}
	this.records[key] = this.head
	this.que[this.head] = key
	this.head = next

	if v, ok := this.dests[destination]; !ok {
		this.dests[destination] = []int{timestamp}
	} else {
		this.dests[destination] = append(v, timestamp)
	}

	return true
}

func (this *Router) ForwardPacket() []int {
	if this.head == this.tail && !this.full {
		return nil
	}
	first := this.que[this.tail]
	this.removePacket(first)
	this.tail = (this.tail + 1) % len(this.que)
	this.full = false
	return first[:]
}

func (this *Router) GetCount(destination int, startTime int, endTime int) int {
	if v, ok := this.dests[destination]; !ok {
		return 0
	} else {
		i := search(0, len(v), func(i int) bool {
			return v[i] > endTime
		})
		j := search(0, len(v), func(j int) bool {
			return v[j] >= startTime
		})
		return i - j
	}
}

func search(l int, r int, f func(int) bool) int {
	for l < r {
		mid := (l + r) / 2
		if f(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
