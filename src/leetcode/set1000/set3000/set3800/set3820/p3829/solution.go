package p3829

const N = 1010

type RideSharingSystem struct {
	ridersState []int
	riders      []int
	drivers     []int
	head        int
}

func Constructor() RideSharingSystem {
	ridersState := make([]int, N)
	for i := range N {
		ridersState[i] = -1
	}
	return RideSharingSystem{
		ridersState: ridersState,
		head:        0,
	}
}

func (this *RideSharingSystem) AddRider(riderId int) {
	this.ridersState[riderId] = len(this.riders)
	this.riders = append(this.riders, riderId)
}

func (this *RideSharingSystem) AddDriver(driverId int) {
	this.drivers = append(this.drivers, driverId)
}

func (this *RideSharingSystem) MatchDriverWithRider() []int {
	for this.head < len(this.riders) && this.ridersState[this.riders[this.head]] < 0 {
		this.head++
	}
	if this.head == len(this.riders) || len(this.drivers) == 0 {
		return []int{-1, -1}
	}
	a := this.drivers[0]
	this.drivers = this.drivers[1:]
	b := this.riders[this.head]
	this.head++
	this.ridersState[b] = -1
	return []int{a, b}
}

func (this *RideSharingSystem) CancelRider(riderId int) {
	if this.ridersState[riderId] == -1 {
		return
	}
	this.ridersState[riderId] = -1
}

/**
 * Your RideSharingSystem object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddRider(riderId);
 * obj.AddDriver(driverId);
 * param_3 := obj.MatchDriverWithRider();
 * obj.CancelRider(riderId);
 */
