# one-golang

## Master the Interview:

### Introduce Yourself: 
I have a background in B.Tech IT and graduated in 2020. After college, I joined Sensibol as a Golang Backend Developer. My primary role involved developing microservices using Golang, AWS, MongoDB, and Redis. I resigned in April to take care of my father's cancer treatment. Now that he is well, I am seeking new opportunities.

### What projects do you woked on Sensibol, describe: 
I worked on PDL (Phonographic Digital Limited), a music distribution and royalty management platform. 

Additionally, I worked on Singshala, which is similar to TikTok but with the extra feature of analyzing the audio components of videos and providing rankings based on the analysis. 

Both projects used Golang, MongoDB, Redis, AWS S3, SQS, SNS, Lambda, etc. Both had microservices architectures. PDL transitioned from a domain-driven approach to an event-driven one, while Singshala is domain-driven and complete. 

Since Sensibol is very conservative regarding hiring, only three backend developers worked on both projects, so my involvement was extensive.

### Microservices vs Monolith:
Microservices are better for large projects where scaling and almost zero downtime are required. Bug fixing and maintaining the codebase are easier. A disadvantage of microservices can be inter-service network calls.

### Golang Garbage Collection:
Golang uses automatic garbage collection to manage memory. Developers do not need to allocate or deallocate memory manually, which reduces memory-related errors.

### Goroutine vs Thread:
Goroutines are designed for concurrency, meaning multiple tasks are run using context switching. Threads are designed for parallelism, meaning multiple tasks can run simultaneously on multiple CPU cores.

Goroutines have a dynamic stack size and are managed by the Go runtime. Threads have a fixed-size stack and are managed by the OS kernel.

### What is Closure in golang:
A closure is a special type of anonymous function that can use variables declared outside of the function. Closures treat functions as values, allowing us to assign functions to variables, pass functions as arguments, and return functions from other functions. 

```go
v:= func (f func (i int) int) func()int{
    c:=f(3)
    return func()int{
        c++
        return c
    }
}
f:=v()// v() returns a function address
f()// call the function which v() returns
```

### Interfaces in golang:
Interfaces allow us to define contracts, which are abstract methods, have no body/implementations of the methods.
A Struct which wants to implements the Interface need to write the body of every abstract methods the interface holds.
We can compose interfaces together.
An empty interface can hold any type of values.
name.(type) give us the Type the interface will hold at runtime. or we can use reflect.TypeOf(name)
```go
func ty(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("Integer")
	default:
		fmt.Println("No idea")
	}
	fmt.Println(reflect.TypeOf(i))
}
func main() {
	ty(67.89)
}
```
### Panic Defer Recover combo:
panic is use to cause a Runtime Error and Stop the execution.
When a function return or panicking then Defer blocks are called according to Last in First out manner, the lasr defer will execute first.
Recover is use to regain the execution from a panicking situation and handle it properly then stop execution. Recover usefule for close ant connection like db and websockets etc.
```go
func div(num int) int {
	if num == 0 {
		panic("Not divisible by 0")
	} else {
		return 27 / num
	}
}
func rec() {
	r := recover()
	if r != nil {
		fmt.Println("I am recovering from Panic")
		fmt.Println("I am Fine Now")
	}
}
func main() {
	defer rec()
	fmt.Println(div(0))
	fmt.Println("Main Regained") // Will not executed if divisble by 0
}
```
### Array vs Slice: 
Array are not Grow and Shrink dynamically at runtime, Slice can. Slice is just references to an existing array of a fixed length.

### Method Dispatching:
golang use Receiver function for method sipatching and has 2 way to dispatch methods at runtime.

Pointer receiver function: As obj is refrence of any Struct so any modification inside the function will affect the original Strcut. More memory-efficient and can result in faster execution, especially for large structs.
```go
func (obj *class_name)method_name(argument int) (returns_name bool){
    //body
}
```
Value receiver function: As obj is copy of any Struct so any modification inside the function will not affect the original Strcut. 
```go
func (obj class_name)method_name(argument int) (returns_name bool){
    //body
}
```
### Concurency Primitives:
Concurency Primitives are tools that are provided by any programming languages to handle execution behaviors of Concurent tasks.

In golang we have Mutex, Semaphore, Channels as concurency primitives.

Mutex is used to protect shared resources from being accessed by multiple threads simultaneously.

Semaphore is used to protect shared pool of resources from being accessed by multiple threads simultaneously. Semaphore is a Counter which start from Number of Reosurces. When one thread using the reosurces Semaphote decremented by 1. If semaphore value is 0 then thread will wait untils its value greater than 0. When on thread done with the resources then Semaphore incremented by 1.

Channel is used to communicate via sending and receiving data and provide synchronisation between multiple gorountines. If channel have a value then execution blocked until reade reads from the channel.
Channel can be buffered, allowing goroutines to send multiple values without blocking until the buffer is full. 

Waitgroup is used when we want the function should wait until goroutines complete its task.
Waitgroup has Add() function which increments the wait counter for each goroutine.
Wait() function which is used for wait until wait counter became zero.
Done() decrement wait counter function is called when goroutine complete its task.

### Map Synchronisation:
In golang if multiple goroutines try to acess map at same time, then the operations leads to Panic for RACE or DEADLOCK (fatal error: concurrent map read and map write).
So we need proper codes for handeling Map.
We use MUTEX for LOCK and UNLOCK the Map operations like Read and Write. 
```go
func producer(m *map[int]string, wg *sync.WaitGroup, mu *sync.RWMutex) {
	vm := *m
	for i := 0; i < 5; i++ {
		mu.Lock()
		vm[i] = fmt.Sprint("$", i)
		mu.Unlock()
	}
	m = &vm
	wg.Done()
}
func consumer(m *map[int]string, wg *sync.WaitGroup, mu *sync.RWMutex) {
	vm := *m
	for i := 0; i < 5; i++ {
		mu.RLock()
		fmt.Println(vm[i])
		mu.RUnlock()
	}
	wg.Done()
}
func main() {
	m := make(map[int]string)
	m[0] = "1234"
	m[3] = "2345"
	wg := sync.WaitGroup{}
	mu := sync.RWMutex{}
	for i := 0; i < 5; i++ {
		wg.Add(2)
		go producer(&m, &wg, &mu)
		go consumer(&m, &wg, &mu)
	}
	wg.Wait()
}
```
### Describe Channel comunication with task distributions:
Lets imagine we have a number n we have to calculate 0 to n numbers are prime or not.
```go
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
func primer(a []int, ch chan<- map[int]bool, wg *sync.WaitGroup) {
	time.Sleep(time.Second)
	defer wg.Done()
	m := make(map[int]bool)
	for i := range a {
		m[a[i]] = isPrime(a[i])
	}
	ch <- m
}
func main() {
	startTime := time.Now()
	var wg sync.WaitGroup
	n := 12
	arr := []int{}
	for i := 0; i < n; i++ {
		arr = append(arr, i)
	}
	length := len(arr)
	goroutines := 4
	part := length / goroutines
	ch := make(chan map[int]bool, goroutines)
	ma := make(map[int]bool)
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		s := i * part
		e := s + part
		if e > length {
			e = length
		}
		go primer(arr[s:e], ch, &wg)
	}
	wg.Wait()
	close(ch)
	for i := range ch {
		for k, v := range i {
			ma[k] = v
		}
	}
	fmt.Println(ma)
	fmt.Println("Time Taken: ", time.Since(startTime))
}
```
