# Intensity Segment

[中文](README-zh.md)

## Problem Background

**Problem Set**  
We are looking for a program that manages “intensity” by segments. Segments are intervals from -infinity to infinity.  
You are required to implement functions that update intensity by an integer amount for a given range.  
All intensity starts with 0.

Please implement these three functions:
- `Add(from, to, amount int) error`: Add an intensity value to a segment range.
- `Set(from, to, amount int) error`: Set the intensity value for a segment range, ensuring no overlap.
- `ToString() string`: Output the current segments as a string.

## Directory Structure

```
intensitysegment/
    segment.go         # Core implementation of IntensitySegment
    demo/
        main.go       # Demonstration of IntensitySegment usage
    test/
        segment_test.go # Unit tests for IntensitySegment
    go.mod
    go.sum
    README.md
```

## Demo

A demonstration of the IntensitySegment API is provided in [demo/main.go](demo/main.go).

To run the demo, use:

```sh
cd intensitysegment/demo
go run main.go
```

You will see output illustrating how segments are added and updated, and how the segment list changes after each operation.

## Unit Tests

Comprehensive unit tests are provided in [test/segment_test.go](test/segment_test.go).

To run the tests, use:

```sh
cd intensitysegment/test
go test segment_test.go
```

This will execute all test cases to verify the correctness of the segment operations, including edge cases and overlapping scenarios.

## Implementation

The core logic is implemented in [segment.go](segment.go) as the [`IntensitySegment`](segment.go) type, which supports adding, setting, and displaying intensity segments.  
The implementation ensures that segments are efficiently managed and merged as needed.

## How It Works

- **Add**: Adds the given intensity to the specified range, merging and splitting segments as necessary.
- **Set**: Sets the intensity for the specified range, ensuring no overlap with existing segments.
- **ToString**: Returns a string representation of all current segments.

## Example Output

Example output from the demo:

```
init segments:  []
After adding [10,30] with intensity 1:  [[10,1], [30,0]]
After adding [20,40] with intensity 1:  [[10,1], [20,2], [30,1], [40,0]]
After adding [10,40] with intensity -1:  [[20,1], [30,0]]
After adding [10,40] with intensity -1:  [[10,-1], [20,0], [30,-1], [40,0]]
```

## Dependencies

- [github.com/stretchr/testify](https://github.com/stretchr/testify) for unit testing assertions.

## License

This project is for interview problem solutions and is provided as-is