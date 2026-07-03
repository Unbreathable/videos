## Goal

- teach different architectures to prevent deadlocks and foot guns with mutexes
  - The scenario with a lot of mutexes: A, B, C and D (use sorting to make sure lock and unlock order is always consistent)
  - One mutex for an object in memory (use a worker + queue, benchmarks)

## 1. Intro

- Ever run into deadlocks?
-

## 2. Why?

- Specifically explain how Go's parallelism goes (one goroutine per request, etc.)
- Sharing data is not so simple (show counter example)
- Show how you can work around it with a mutex

## 3. The simple, but problematic solution

- Explain mutex: Only one person can lock _TODO: Create diagram_
- Problem with mutexes: Show examples of two cases
  - The classic _TODO: Create code example_
  - Leaving one open _TODO: Create code example_

## 4. Single object: Just use channels, faster anyway

- Explain the channel architecture based on a diagram _TODO: Create_
- Show an example with the counter (even though doesn't make much sense there)
- Use opportunity to also explain how atomic.Int32 is a hardware primitive and therefore much faster than anything you can do
- Show benchmark to highlight how with a little bit of work the channel is actually faster
  - Worker count actually doesn't really matter a whole lot (results don't change very much)
  -

## 5. Multiple objects: Just sort them and be fine

_TODO: Create proper example, Ideas: multiple counters, but we protect them with mutex just to make the point?_
