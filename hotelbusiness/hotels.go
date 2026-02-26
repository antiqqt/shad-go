//go:build !solution

package hotelbusiness

import (
	"cmp"
	"slices"
)

type Guest struct {
	CheckInDate  int
	CheckOutDate int
}

type Load struct {
	StartDate  int
	GuestCount int
}

func ComputeLoad(guests []Guest) []Load {
	delta := make(map[int]int)

	for _, g := range guests {
		delta[g.CheckInDate]++
		delta[g.CheckOutDate]--
	}

	events := make([][2]int, 0, len(delta))

	for k, v := range delta {
		events = append(events, [2]int{k, v})
	}

	slices.SortFunc(events, func(a, b [2]int) int {
		return cmp.Compare(a[0], b[0])
	})

	var res []Load
	counter := 0
	for _, e := range events {
		date, change := e[0], e[1]

		if change == 0 {
			// skip zero-delta events
			continue
		}

		counter += change
		res = append(res, Load{StartDate: date, GuestCount: counter})
	}

	return res
}
