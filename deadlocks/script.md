## Goal

- teach different architectures to prevent deadlocks and foot guns with mutexes
  - The scenario with a lot of mutexes: A, B, C and D (use sorting to make sure lock and unlock order is always consistent)
  - One mutex for an object in memory (use a worker + queue, benchmarks)

## 1. Intro

- Ever written a parallel system? -> Then you probably had to share data between parallel tasks.
- Explain content of the video

## 2. Why?

- Explain mutex: Only one person can lock
- Specifically explain how Go's parallelism goes (one goroutine / thread per request, etc.)
- Sharing data is not so simple (show counter example)
- Show how you can work around it with a mutex

## 3. The simple, but problematic solution

- Problem with mutexes: Show examples of two cases
  - The classic
  - Leaving one open

## 4. Single object: Just use channels, faster anyway

- Explain the channel architecture based on a diagram
- Show an example with the counter (even though doesn't make much sense there)
- Use opportunity to also explain how atomic.Int32 is a hardware primitive and therefore much faster than anything you can do
- Show benchmark to highlight how with a little bit of work the channel is actually faster
  - Worker count actually doesn't really matter a whole lot (results don't change very much)
  - The worker closes in more and more on the mutex implementation
  - **Conclusion:** If you do real work while the mutex is locked, you might as well use the worker (most of your code will likely do more than 2600 ns of work)
- This is not for highly parallel systems, but for people searching an alternative that is still thread-safe

## 5. Multiple objects: Just sort them and be fine

_TODO: Create proper example, Ideas: multiple counters, but we protect them with mutex just to make the point?_
