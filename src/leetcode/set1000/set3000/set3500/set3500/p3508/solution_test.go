package p3508

import (
	"reflect"
	"testing"
)

func TestSample1(t *testing.T) {
	router := Constructor(3)

	// Add packets and check return values
	if !router.AddPacket(1, 4, 90) {
		t.Fatalf("Expected AddPacket(1, 4, 90) to return true")
	}

	if !router.AddPacket(2, 5, 90) {
		t.Fatalf("Expected AddPacket(2, 5, 90) to return true")
	}

	if router.AddPacket(1, 4, 90) {
		t.Fatalf("Expected AddPacket(1, 4, 90) to return false (duplicate)")
	}

	if !router.AddPacket(3, 5, 95) {
		t.Fatalf("Expected AddPacket(3, 5, 95) to return true")
	}

	if !router.AddPacket(4, 5, 105) {
		t.Fatalf("Expected AddPacket(4, 5, 105) to return true")
	}

	// Forward packet and check result
	packet := router.ForwardPacket()
	expected := []int{2, 5, 90}
	if !reflect.DeepEqual(packet, expected) {
		t.Fatalf("Expected ForwardPacket() to return %v, got %v", expected, packet)
	}

	// Add another packet
	if !router.AddPacket(5, 2, 110) {
		t.Fatalf("Expected AddPacket(5, 2, 110) to return true")
	}

	// Get count for specific destination and time range
	count := router.GetCount(5, 100, 110)
	if count != 1 {
		t.Fatalf("Expected GetCount(5, 100, 110) to return 1, got %d", count)
	}
}

func TestSample2(t *testing.T) {
	router := Constructor(2)

	// Add packet and check return value
	if !router.AddPacket(7, 4, 90) {
		t.Fatalf("Expected AddPacket(7, 4, 90) to return true")
	}

	// Forward first packet and check result
	packet := router.ForwardPacket()
	expected := []int{7, 4, 90}
	if !reflect.DeepEqual(packet, expected) {
		t.Fatalf("Expected ForwardPacket() to return %v, got %v", expected, packet)
	}

	// Forward second packet (should be empty)
	packet = router.ForwardPacket()
	if packet != nil {
		t.Fatalf("Expected ForwardPacket() to return nil, got %v", packet)
	}
}
