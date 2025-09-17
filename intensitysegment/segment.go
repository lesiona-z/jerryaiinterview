package intensitysegment

import (
	"fmt"
	"strings"
)

type IntensitySegment struct {
	Segments *interval
}

type interval struct {
	Start int
	End   int
	Value int
	Next  *interval
}

func (is *IntensitySegment) Add(from, to, amount int) error {
	if is == nil {
		return fmt.Errorf("IntensitySegment is nil")
	}
	if from >= to {
		return fmt.Errorf("Invalid segment range")
	}
	// if there are no segments yet, create the first one
	if is.Segments == nil {
		is.Segments = &interval{Start: from, End: to, Value: amount}
		return nil
	}
	// insert the new segment and merge overlapping segments
	newSegs, err := insertSegment(*is.Segments, from, to, amount)
	if err != nil {
		return err
	}
	mergedSegs := mergeSegments(newSegs)
	is.Segments = &mergedSegs
	return nil
}

func (is *IntensitySegment) Set(from, to, amount int) error {
	if is == nil {
		return fmt.Errorf("IntensitySegment is nil")
	}
	if from >= to {
		return fmt.Errorf("Invalid segment range")
	}
	if is.Segments == nil {
		is.Segments = &interval{Start: from, End: to, Value: amount}
		return nil
	}

	// find the position to set the new segment
	var dummy interval
	dummy.Next = is.Segments
	prev := &dummy
	for curr := prev.Next; curr != nil; curr = curr.Next {
		if !(to <= curr.Start || from >= curr.End) {
			return fmt.Errorf("New segment overlaps with old ones")
		}

		// insert before curr
		if to <= curr.Start {
			newSeg := &interval{Start: from, End: to, Value: amount, Next: curr}
			prev.Next = newSeg
			is.Segments = dummy.Next
			break
		}

		// insert after curr
		if from >= curr.End && curr.Next == nil {
			newSeg := &interval{Start: from, End: to, Value: amount, Next: nil}
			curr.Next = newSeg
			is.Segments = dummy.Next
			break
		}
	}

	segs := mergeSegments(*is.Segments)
	is.Segments = &segs
	return nil
}

func (is *IntensitySegment) ToString() string {
	if is == nil || is.Segments == nil {
		return "[]"
	}

	var segStrs []string
	for seg := is.Segments; seg != nil; seg = seg.Next {
		segStrs = append(segStrs, fmt.Sprintf("[%d,%d]", seg.Start, seg.Value))
		// if seg is the last segment or there is a gap to the next segment, add an end marker with value 0
		if seg.Next == nil || seg.End != seg.Next.Start {
			segStrs = append(segStrs, fmt.Sprintf("[%d,0]", seg.End))
		}
	}
	return "[" + strings.Join(segStrs, ", ") + "]"
}

func insertSegment(segs interval, from, to, amount int) (retSegs interval, err error) {
	var dummy interval
	dummy.Next = &segs
	prev := &dummy
	curr := prev.Next
	// find the position to insert the new segment
	for curr != nil {
		// insert before curr
		// 	prev---[from, to]---curr
		if to <= curr.Start {
			newSeg := &interval{Start: from, End: to, Value: amount, Next: curr}
			prev.Next = newSeg
			return *dummy.Next, nil
		}
		// insert after curr
		// 	prev---curr---[from, to]
		if from >= curr.End && (curr.Next == nil || to <= curr.Next.Start) {
			newSeg := &interval{Start: from, End: to, Value: amount, Next: curr.Next}
			curr.Next = newSeg
			return *dummy.Next, nil
		}
		// curr segment completely covers the new segment
		// split curr into 3 sub-segments if needed:
		// 	prev---[curr.Start, from]---[from, to]---[to, curr.End]---curr.Next
		if curr.Start <= from && to <= curr.End {
			if curr.Start < from {
				prev.Next = &interval{Start: curr.Start, End: from, Value: curr.Value}
				prev = prev.Next
			}
			prev.Next = &interval{Start: from, End: to, Value: curr.Value + amount}
			prev = prev.Next
			if to < curr.End {
				prev.Next = &interval{Start: to, End: curr.End, Value: curr.Value}
				prev = prev.Next
			}
			prev.Next = curr.Next
			return *dummy.Next, nil
		}
		// new segment completely covers curr
		// split the new segment into 3 sub-segments if needed:
		// 	prev---[from, curr.Start]---[curr.Start, curr.End]---...
		if from < curr.Start && curr.End < to {
			prev.Next = &interval{Start: from, End: curr.Start, Value: amount}
			prev = prev.Next
			prev.Next = &interval{Start: curr.Start, End: curr.End, Value: curr.Value + amount}
			prev = prev.Next

			// the '[from, to]' interval may covers curr.Next too, so
			// shrink the new segment to the remaining part
			// continue to check the next segment
			from = curr.End
			curr = curr.Next
			if curr == nil {
				// reached the end, insert the remaining part
				prev.Next = &interval{Start: from, End: to, Value: amount}
				return *dummy.Next, nil
			}
			continue
		}

		// overlap on the left side of curr
		// split into 2 sub-segments if needed:
		// 	prev---[from, curr.Start]---[curr.Start, to]---[to, curr.End]---curr.Next
		if from < curr.Start && curr.Start < to && to <= curr.End {
			prev.Next = &interval{Start: from, End: curr.Start, Value: amount}
			prev = prev.Next
			prev.Next = &interval{Start: curr.Start, End: to, Value: curr.Value + amount}
			prev = prev.Next
			if to < curr.End {
				prev.Next = &interval{Start: to, End: curr.End, Value: curr.Value}
				prev = prev.Next
			}
			prev.Next = curr.Next
			return *dummy.Next, nil
		}

		// overlap on the right side of curr
		// split into 2 sub-segments if needed:
		// 	prev---[curr.Start, from]---[from, curr.End]---...
		if curr.Start <= from && from < curr.End && curr.End < to {
			if curr.Start < from {
				prev.Next = &interval{Start: curr.Start, End: from, Value: curr.Value}
				prev = prev.Next
			}
			prev.Next = &interval{Start: from, End: curr.End, Value: curr.Value + amount, Next: curr.Next}
			prev = prev.Next

			// the '[from, to]' interval may covers curr.Next too, so
			// shrink the new segment to the remaining part
			// continue to check the next segment
			from = curr.End
			curr = curr.Next
			if curr == nil {
				// reached the end, insert the remaining part
				prev.Next = &interval{Start: from, End: to, Value: amount}
				return *dummy.Next, nil
			}
			continue
		}

		// pace to next segment
		prev = curr
		curr = curr.Next
	}
	return segs, nil
}

func mergeSegments(segs interval) interval {
	var dummy interval
	dummy.Next = &segs
	prev := &dummy
	curr := prev.Next
	for curr != nil {
		// remove zero-intensity segments
		if curr.Value == 0 {
			prev.Next = curr.Next
			curr = prev.Next
			continue
		}

		if *prev == dummy {
			// pace to next segment
			prev = curr
			curr = curr.Next
			continue
		}

		if prev.Value == curr.Value && prev.End >= curr.Start {
			// merge segments, extend the end of prev seg and drop curr seg
			if curr.End > prev.End {
				prev.End = curr.End
			}
			prev.Next = curr.Next
			curr = prev.Next
			continue
		}
		// pace to next segment
		prev = curr
		curr = curr.Next
	}
	return *dummy.Next
}
