package main

import (
	"container/heap"
	"fmt"
	"io"
)

type (
	Process struct {
		ProcessID     string
		ArrivalTime   int64
		BurstDuration int64
		Priority      int64
		RemainingTime int64
	}
	TimeSlice struct {
		PID   string
		Start int64
		Stop  int64
	}
	ProcessHeap []*Process
)

func (ph ProcessHeap) Len() int { return len(ph) }

func (ph ProcessHeap) Less(i, j int) bool {
	return ph[i].BurstDuration < ph[j].BurstDuration
}

func (ph ProcessHeap) Swap(i, j int) {
	ph[i], ph[j] = ph[j], ph[i]
}

func (ph *ProcessHeap) Push(x interface{}) {
	*ph = append(*ph, x.(*Process))
}

func (ph *ProcessHeap) Pop() interface{} {
	old := *ph
	n := len(old)
	x := old[n-1]
	*ph = old[0 : n-1]
	return x
}

//region Schedulers

func FCFSSchedule(w io.Writer, title string, processes []Process) {
	// Code for FCFS scheduler (already implemented)
}

func SJFSchedule(w io.Writer, title string, processes []Process) {
	var (
		serviceTime     int64
		totalWait       float64
		totalTurnaround float64
		lastCompletion  float64
		schedule        = make(map[string][]string)
		gantt           = make([]TimeSlice, 0)
		ph              ProcessHeap
	)

	for i := range processes {
		processes[i].RemainingTime = processes[i].BurstDuration
		heap.Push(&ph, &processes[i])
	}

	for ph.Len() > 0 {
		curr := heap.Pop(&ph).(*Process)
		waitingTime := serviceTime - curr.ArrivalTime
		if waitingTime < 0 {
			waitingTime = 0
		}

		totalWait += float64(waitingTime)

		start := waitingTime + curr.ArrivalTime
		turnaround := curr.BurstDuration + waitingTime
		totalTurnaround += float64(turnaround)
		completion := curr.BurstDuration + curr.ArrivalTime + waitingTime
		lastCompletion = float64(completion)

		schedule[curr.ProcessID] = []string{
			fmt.Sprint(curr.ProcessID),
			fmt.Sprint(curr.Priority),
			fmt.Sprint(curr.BurstDuration),
			fmt.Sprint(curr.ArrivalTime),
			fmt.Sprint(waitingTime),
			fmt.Sprint(turnaround),
			fmt.Sprint(completion),
		}

		serviceTime += curr.BurstDuration

		gantt = append(gantt, TimeSlice{
			PID:   curr.ProcessID,
			Start: start,
			Stop:  serviceTime,
		})

		for ph.Len() > 0 {
			next := heap.Pop(&ph).(*Process)
			if next.ArrivalTime <= serviceTime {
				next.RemainingTime -= serviceTime - next.ArrivalTime
				heap.Push(&ph, next)
			} else {
				heap.Push(&ph, next)
				break
			}
		}
	}

	count := float64(len(processes))
	aveWait := totalWait / count
	aveTurnaround := totalTurnaround / count
	aveThroughput := count / lastCompletion

	outputTitle(w, title)
	outputGantt(w, gantt)
	outputSchedule(w, convertScheduleToSlice(schedule), aveWait, aveTurnaround, aveThroughput)
}

func SJFPrioritySchedule(w io.Writer, title string, processes []Process) {
	var (
		serviceTime     int64
		totalWait       float64
		totalTurnaround float64
		lastCompletion  float64
		schedule        = make(map[string][]string)
		gantt           = make([]TimeSlice, 0)
		ph              ProcessHeap
	)

	for i := range processes {
		processes[i].RemainingTime = processes[i].BurstDuration
		heap.Push(&ph, &processes[i])
	}

	for ph.Len() > 0 {
		curr := heap.Pop(&ph).(*Process)
		waitingTime := serviceTime - curr.ArrivalTime
		if waitingTime < 0 {
			waitingTime = 0
		}

		totalWait += float64(waitingTime)

		start := waitingTime + curr.ArrivalTime
		turnaround := curr.BurstDuration + waitingTime
		totalTurnaround += float64(turnaround)
		completion := curr.BurstDuration + curr.ArrivalTime + waitingTime
		lastCompletion = float64(completion)

		schedule[curr.ProcessID] = []string{
			fmt.Sprint(curr.ProcessID),
			fmt.Sprint(curr.Priority),
			fmt.Sprint(curr.BurstDuration),
			fmt.Sprint(curr.ArrivalTime),
			fmt.Sprint(waitingTime),
			fmt.Sprint(turnaround),
			fmt.Sprint(completion),
		}

		serviceTime += curr.BurstDuration

		gantt = append(gantt, TimeSlice{
			PID:   curr.ProcessID,
			Start: start,
			Stop:  serviceTime,
		})

		for ph.Len() > 0 {
			next := heap.Pop(&ph).(*Process)
			if next.ArrivalTime <= serviceTime && next.Priority < curr.Priority {
				next.RemainingTime -= serviceTime - next.ArrivalTime
				heap.Push(&ph, next)
			} else {
				heap.Push(&ph, next)
				break
			}
		}
	}

	count := float64(len(processes))
	aveWait := totalWait / count
	aveTurnaround := totalTurnaround / count
	aveThroughput := count / lastCompletion

	outputTitle(w, title)
	outputGantt(w, gantt)
	outputSchedule(w, convertScheduleToSlice(schedule), aveWait, aveTurnaround, aveThroughput)
}

func RRSchedule(w io.Writer, title string, processes []Process) {
	var (
		serviceTime     int64
		totalWait       float64
		totalTurnaround float64
		lastCompletion  float64
		schedule        = make(map[string][]string)
		gantt           = make([]TimeSlice, 0)
		readyQueue      []Process
		quantum         = int64(4)
	)

	for i := range processes {
		processes[i].RemainingTime = processes[i].BurstDuration
	}

	for len(processes) > 0 || len(readyQueue) > 0 {
		for i := 0; i < len(processes); {
			if processes[i].ArrivalTime <= serviceTime {
				readyQueue = append(readyQueue, processes[i])
				processes = append(processes[:i], processes[i+1:]...)
			} else {
				i++
			}
		}

		if len(readyQueue) > 0 {
			curr := readyQueue[0]
			readyQueue = readyQueue[1:]

			waitingTime := serviceTime - curr.ArrivalTime
			if waitingTime < 0 {
				waitingTime = 0
			}

			totalWait += float64(waitingTime)

			start := waitingTime + curr.ArrivalTime
			turnaround := curr.BurstDuration + waitingTime
			totalTurnaround += float64(turnaround)
			completion := curr.BurstDuration + curr.ArrivalTime + waitingTime
			lastCompletion = float64(completion)

			schedule[curr.ProcessID] = []string{
				fmt.Sprint(curr.ProcessID),
				fmt.Sprint(curr.Priority),
				fmt.Sprint(curr.BurstDuration),
				fmt.Sprint(curr.ArrivalTime),
				fmt.Sprint(waitingTime),
				fmt.Sprint(turnaround),
				fmt.Sprint(completion),
			}

			if curr.RemainingTime <= quantum {
				serviceTime += curr.RemainingTime
				gantt = append(gantt, TimeSlice{
					PID:   curr.ProcessID,
					Start: start,
					Stop:  serviceTime,
				})
			} else {
				serviceTime += quantum
				curr.RemainingTime -= quantum
				readyQueue = append(readyQueue, curr)
				gantt = append(gantt, TimeSlice{
					PID:   curr.ProcessID,
					Start: start,
					Stop:  serviceTime,
				})
			}
		} else {
			serviceTime++
		}
	}

	count := float64(len(processes))
	aveWait := totalWait / count
	aveTurnaround := totalTurnaround / count
	aveThroughput := count / lastCompletion

	outputTitle(w, title)
	outputGantt(w, gantt)
	outputSchedule(w, convertScheduleToSlice(schedule), aveWait, aveTurnaround, aveThroughput)
}

//endregion

func convertScheduleToSlice(schedule map[string][]string) [][]string {
	result := make([][]string, len(schedule))
	i := 0
	for _, row := range schedule {
		result[i] = row
		i++
	}
	return result
}

