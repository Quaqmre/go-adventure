package main

import (
	"flag"
	"fmt"
	"math"
	"sync"
	"time"
)

//cmd:go run threadperf.go -summary -iterations 100000 -threads 96 -all-stats

type globals_s struct {
	lock           sync.Mutex
	running        bool
	status         int
	stats          []Stats
	threads_waiter sync.WaitGroup
}

var globals globals_s

type Stats struct {
	min_lock_time int64
	max_lock_time int64
	total_time    int64
	n_locks       int64
	n_flips       int64
}

func update_stats(stats *Stats, elapsed_time int64) {
	stats.n_locks += 1
	if elapsed_time < stats.min_lock_time {
		stats.min_lock_time = elapsed_time
	}
	if elapsed_time > stats.max_lock_time {
		stats.max_lock_time = elapsed_time
	}
	stats.total_time += elapsed_time
}

func fun(check int, set int, stats *Stats) {
	loop := true
	for loop {
		start := time.Now()
		globals.lock.Lock()
		if globals.status == check {
			loop = false
			stats.n_flips += 1
			globals.status = set
		}
		globals.lock.Unlock()
		stop := time.Now()
		update_stats(stats, stop.Sub(start).Nanoseconds())
	}
}

func thread_entry_point(stats *Stats) {
	// run until canceled
	for globals.running {
		fun(1, 0, stats)
	}
}

func print_stats() {
	n_threads := len(globals.stats)
	for i := 0; i < n_threads; i++ {
		var average float64
		stats := globals.stats[i]
		average = float64(stats.total_time) / float64(stats.n_locks)
		if i == n_threads-1 {
			fmt.Printf("server: ")
		} else {
			fmt.Printf("thread %d: ", i)
		}
		fmt.Printf("min=%d, max=%d, average=%f, mutexes_locked=%d, flips=%d\n",
			stats.min_lock_time, stats.max_lock_time, average, stats.n_locks, stats.n_flips)
	}
}

func avg(n []int64) float64 {
	sum := float64(0)

	for i := 0; i < len(n); i++ {
		sum += float64(n[i])
	}
	return sum / float64(len(n))
}

func stdev(n []int64) float64 {
	average := avg(n)
	sd := float64(0)

	for i := 0; i < len(n); i++ {
		sd += math.Pow(float64(n[i])-average, 2)
	}
	return math.Sqrt(sd / float64(len(n)))
}

func max(n []int64) int64 {
	m := n[0]
	for i := 1; i < len(n); i++ {
		if m < n[i] {
			m = n[i]
		}
	}
	return m
}

func min(n []int64) int64 {
	m := n[0]
	for i := 1; i < len(n); i++ {
		if m > n[i] {
			m = n[i]
		}
	}
	return m
}

func calculate_stats(n []int64) string {
	max_val := max(n)
	min_val := min(n)
	range_val := max_val - min_val
	stdev_val := stdev(n)
	return fmt.Sprintf("range: %7d, stdev: %11.3f, max: %7d, min: %7d",
		range_val, stdev_val, max_val, min_val)
}

func print_stats_summary() {
	n_threads := len(globals.stats)
	flips := make([]int64, 0, n_threads)
	mutexes_locked := make([]int64, 0, n_threads)

	for i := 0; i < n_threads; i++ {
		stats := globals.stats[i]
		flips = append(flips, stats.n_flips)
		mutexes_locked = append(mutexes_locked, stats.n_locks)
	}

	fmt.Printf("         flips: %s\n", calculate_stats(flips))
	fmt.Printf("mutexes_locked: %s\n", calculate_stats(mutexes_locked))
}

func main() {

	n_threads_arg := flag.Int("threads", 16, "number of worker threads")
	n_iterations_arg := flag.Int("iterations", 10000, "number of iterations of the server thread")
	summary := flag.Bool("summary", false, "print a summary of stats instead of raw data")
	all_stats := flag.Bool("all-stats", false, "print all raw stats")

	flag.Parse()
	// add one to n_threads because we always have one extra for the "server" thread
	n_threads := *n_threads_arg + 1
	n_iterations := *n_iterations_arg

	// initialize the stats slice
	globals.stats = make([]Stats, n_threads, n_threads)
	for i := 0; i < n_threads; i++ {
		globals.stats[i].min_lock_time = 1000000
	}

	globals.running = true

	// start threads
	for i := 0; i < n_threads-1; i++ {
		go thread_entry_point(&globals.stats[i])
	}
	globals.threads_waiter.Add(0)

	// start the main thread (but re-use the "main" thread)
	for i := 0; i < n_iterations; i++ {
		fun(0, 1, &globals.stats[n_threads-1])
	}

	// stop
	globals.running = false
	globals.threads_waiter.Wait()

	if *summary {
		print_stats_summary()
		if *all_stats {
			print_stats()
		}
	} else {
		print_stats()
	}
}
