package workers

import "sync"

// Pass each of the given todos to the do function and return the resulting dones, in any order,
// because the invocations of the do function are parallelized over numWorkers goroutines.
func Do(todos []int, do func(todo int) (done int), numWorkers int) (dones []int) {
	if len(todos) == 0 { // if no todos then immediately return empty dones
		return []int{}
	}
	if do == nil { // if no do function then default to identity function
		do = func(todo int) int {
			return todo
		}
	}
	if numWorkers <= 0 { // if non-sensible number of workers then default to 1 workers
		numWorkers = 1
	}

	todoChan := make(chan int)
	doneChan := make(chan int)

	workers := forkWorkers(todoChan, do, doneChan, numWorkers)
	go submitTodos(todos, todoChan)
	go waitForWorkers(workers, doneChan)
	return collectDones(doneChan)
}

func forkWorkers(todoChan <-chan int, do func(int) int, doneChan chan<- int, numWorkers int) *sync.WaitGroup {
	workers := &sync.WaitGroup{}
	for n := 0; n < numWorkers; n++ {
		workers.Add(1)
		go worker(todoChan, do, doneChan, workers)
	}
	return workers
}

func worker(todoChan <-chan int, do func(int) int, doneChan chan<- int, workers *sync.WaitGroup) {
	defer workers.Done()
	for todo := range todoChan {
		doneChan <- do(todo)
	}
}

func submitTodos(todos []int, todoChan chan<- int) {
	for _, todo := range todos {
		todoChan <- todo
	}
	close(todoChan)
}

func waitForWorkers(workers *sync.WaitGroup, doneChan chan<- int) {
	workers.Wait()
	close(doneChan)
}

func collectDones(doneChan <-chan int) []int {
	dones := make([]int, 0, len(doneChan))
	for done := range doneChan {
		dones = append(dones, done)
	}
	return dones
}
