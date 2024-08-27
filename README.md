# one-golang

## Master the Interview:

### Introduce Yourself: 
I have a background in B.Tech IT and passout of	 2020. After college, I joined Sensibol as a Golang Backend Developer. My primary role involved developing microservices using Golang, AWS, MongoDB, and Redis. I resigned in April to take care of my father's cancer treatment. Now that he is well, I am seeking new opportunities.

### What projects do you woked on Sensibol, describe: 
I worked on PDL (Phonographic Digital Limited), a music distribution and royalty management platform. 

Additionally, I worked on Singshala, which is similar to TikTok but with extra feature of analyzing the audio components of videos and providing rankings based on the analysis. 

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
When a function return or panicking then Defer blocks are called according to Last in First out manner, the last defer will execute first.
Recover is use to regain the execution from a panicking situation and handle it properly then stop execution. Recover is usefule for close any connection like db and websockets etc.
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
Array can not Grow and Shrink dynamically at runtime, Slice can. Slice is just references to an existing array of a fixed length.

### Method Dispatching:
golang use Receiver function for method dispatching and has 2 way to dispatch methods at runtime.

Pointer receiver function: As obj is refrence of the Struct so any modification inside the function will affect the original Struct. More memory-efficient and can result in faster execution, especially for large structs.
```go
func (obj *class_name)method_name(argument int) (returns_name bool){
    //body
}
```
Value receiver function: As obj is copy of the Struct so any modification inside the function will not affect the original Struct. 
```go
func (obj class_name)method_name(argument int) (returns_name bool){
    //body
}
```
### Concurency Primitives:
Concurency Primitives are tools that are provided by any programming languages to handle execution behaviors of Concurent tasks.

In golang we have Mutex, Semaphore, Channels as concurency primitives.

Mutex is used to protect shared resources from being accessed by multiple threads simultaneously.

Semaphore is used to protect shared pool of resources from being accessed by multiple threads simultaneously. Semaphore is a Counter which start from Number of Reosurces. When one thread using the reosurces Semaphote decremented by 1. If semaphore value is 0 then thread will wait untils its value greater than 0. When one thread done with the resources then Semaphore incremented by 1.

Channel is used to communicate via sending and receiving data and provide synchronisation between multiple gorountines. If channel have a value then execution blocked until reader reads from the channel.
Channel can be buffered, allowing goroutines to send multiple values without blocking until the buffer is full. 

Waitgroup is used when we want the function should wait until goroutines complete its task.
Waitgroup has Add() function which increments the wait counter for each goroutine.
Wait() is used for wait until wait counter became zero.
Done() decrement wait counter and it called when goroutine complete its task.

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
Lets imagine we have a number n, we have to find 0 to n numbers are prime or not.
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
func primeHelper(a []int, ch chan<- map[int]bool, wg *sync.WaitGroup) {
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
		go primeHelper(arr[s:e], ch, &wg)
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

### Select Statement:
Assume a development scenerio where we have 3 s3 Buckets. We spawn 3 GO-Routines each one uploading a File on each S3 bucket at same time. We have to Return SignedUrl of the file so user can stream the File as soon as possible. Now we do not have to wait for 3 S3 Upload operation, when one s3 upload done we can send the SignedUrl of the File to the User so he can Stream. And Other two S3 Upload will continue at same time. This is the Scenerio when Select Statement will work as a Charm.

Select statement is used for Concurency coomunication between multiple goroutines. Select have multiple Case statement related to channel operations. Select block the execution unitl one of its case return. If multiple case returns at same time, then one random case is selected for returns. If no case is ready and there's a default case, it executes immediately. If there's no default case, select blocks until at least one case is ready.
```go
func work(ctx context.Context, ch chan<- string) {
	rand.NewSource(time.Now().Unix())
	r := rand.Intn(6)
	t1 := time.Duration(r) * time.Second
	ctx, cancel := context.WithTimeout(ctx, t1)
	defer cancel()
	select {
	case <-time.After(t1):
		ch <- "Connection established"
	case <-ctx.Done():
		ch <- "Context Expired"
	}
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go work(context.Background(), ch1)
	go work(context.Background(), ch2)
	select {
	case res := <-ch1:
		fmt.Println("ch1 ", res)
	case res := <-ch2:
		fmt.Println("ch2 ", res)
	}
}
```

### SOLID Principles:
SOLID priciples are guidelines for designing Code base that are easy to understand maintain adn extend over time.

Single Responsibility:- A Struct/Class should only a single reason to change. Fields of Author shoud not placed inside Book Struct.
```go
type Book struct{
  ISIN string
  Name String
  AuthorID string
}
type Author struct{
  ID string
  Name String
}
```
Assume One Author decided later, he does not want to Disclose its Real Name to Spread. So we can Serve Frontend by Alias instead of Real Name. Without Changing Book Class/Struct, we can add Alias in Author Struct. By that, Existing Authors present in DB will not be affected as Frontend will Change Name only when it Founds that Alias field is not empty.
```go
type Book struct{
  ISIN string
  Name String
  AuthorID string
}
type Author struct{
  ID string
  Name String
  Alias String
}

```
Open Close:- Struct and Functions should be open for Extension but closed for modifications. New functionality to be added without changing existing Code.
```go
type Shape interface{
	Area() float64
}
type Rectangle struct{
	W float64
	H float64
}
type Circle struct{
	R float64
}
```
Now we want to Calculate Area of Rectangle and Circle, so Rectangle and Circle both can Implements Shape Interface by Write Body of the Area() Function.
```go
func (r Rectangle) Area()float64{
	return r.W * r.H
}
func (c Circle)Area()float64{
	return 3.14 * c.R * c.R
}
```
Now we can create a Function PrintArea() which take Shape as Arguments and Calculate Area of that Shape. So here Shape can be Rectangle, Circle. In Future we can add Triangle Struct which implements Shape interface by writing Body of Area. Now Traingle can be passed to PrintArea() with out modifing the PrintArea() Function.
```go
func PrintArea(shape Shape) {
	fmt.Printf("Area of the shape: %f\n", shape.Area())
}

// In Future
type Triangle struct{
	B float64
	H float54
}
func (t Triangle)Area()float64{
	return 1/2 * t.B * t.H
}

func main(){
	rect:= Rectangle{W:5,H:3}
	cir:=Circle{R:3}
	PrintArea(rect)
	PrintArea(cir)
	// In Future
	tri:=Triangle{B:4,H:8}
	PrintArea(tri)
}
```
Liskov Substitution:- Super class Object can be replaced by Child Class object without affecting the correctness of the program.
```go
type Bird interface{
	Fly() string
}
type Sparrow struct{
	Name string
}
type Penguin struct{
	Name string
}
```
Sparrow and Pengin both are Bird, But Sparrow can Fly, Penguin Not. ShowFly() function take argument of Bird type and call Fly() function. Now as Penguin and Sparrow both are types of Bird, they should be passed as Bird within ShowFly() function.
```go
func (s Sparrow) Fly() string{
	return "Sparrow is Flying"
}
func (p Penguin) Fly() string{
	return "Penguin Can Not Fly"
}

func ShowFly(b Bird){
	fmt.Println(b.Fly())
}
func main() {
	sparrow := Sparrow{Name: "Sparrow"}
	penguin := Penguin{Name: "Penguin"}
  // SuperClass is Bird,  Sparrow, Penguin are the SubClass
	ShowFly(sparrow)
	ShowFly(penguin)
}
```
Interface Segregation:- A class should not be forced to implements interfaces which are not required for the class. Do not couple multiple interfaces together if not necessary then. 
```go
// The Printer interface defines a contract for printers with a Print method.
type Printer interface {
	Print()
}
// The Scanner interface defines a contract for scanners with a Scan method.
type Scanner interface {
	Scan()
}
// The NewTypeOfDevice interface combines Printer and Scanner interfaces for
// New type of devices which can Print and Scan with it new invented Hardware.
type NewTypeOfDevice interface {
	Printer
	Scanner
}
```

Dependecy Inversion:- Class should depends on the Interfaces not the implementations of methods.

```go
// The MessageSender interface defines a contract for 
//sending messages with a SendMessage method.
type MessageSender interface {
	SendMessage(msg string) error
}
// EmailSender and SMSClient structs implement 
//the MessageSender interface with their respective SendMessage methods.
type EmailSender struct{}

func (es EmailSender) SendMessage(msg string) error {
	fmt.Println("Sending email:", msg)
	return nil
}
type SMSClient struct{}

func (sc SMSClient) SendMessage(msg string) error {
	fmt.Println("Sending SMS:", msg)
	return nil
}
type NotificationService struct {
	Sender MessageSender
}
```
The NotificationService struct depends on MessageSender interface, not on concrete implementations (EmailSender or SMSClient). This adheres to Dependency Inversion, because high-level modules (NotificationService) depend on abstractions (MessageSender) rather than details.
```go
func (ns NotificationService) SendNotification(msg string) error {
	return ns.Sender.SendMessage(msg)
}
func main() {
	emailSender := EmailSender{}

	emailNotification := NotificationService{Sender: emailSender}

	emailNotification.SendNotification("Hello, this is an email notification!")
}
```
### Some Coding:
#### Reverse a String:
```go
func reverse(s string) string {
	arr := []rune(s)
	l, r := 0, len(arr)-1
	for l < r {
		temp := arr[l]
		arr[l] = arr[r]
		arr[r] = temp
		l++
		r--
	}
	return string(arr)
}
```
Modified for Channel example:
```go
func reverse(s string, ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	arr := []rune(s)
	l, r := 0, len(arr)-1
	for l < r {
		temp := arr[l]
		arr[l] = arr[r]
		arr[r] = temp
		l++
		r--
	}
	ch <- string(arr)
}
func main() {
	arr := []string{"Boomer", "Golang", "LiL", "NebulA"}
	ch := make(chan string, len(arr))
	wg := sync.WaitGroup{}
	for i := range arr {
		wg.Add(1)
		go reverse(arr[i], ch, &wg)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for i := range ch {
		fmt.Println(i)
	}
}
```
#### Factorial:
```go
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
```
Modified for Channel example:
```go
func factorials(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorials(n-1)
}
func factorialHelper(n int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- factorials(n)
}

func main() {
	arr := []int{5, 10, 25, 31}
	ch := make(chan int, len(arr))
	wg := sync.WaitGroup{}
	for i := range arr {
		wg.Add(1)
		go factorialHelper(arr[i], ch, &wg)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for i := range ch {
		fmt.Println(i)
	}
}
```
#### Fibonacci:
```go
func fibonnaci(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return fibonnaci(n-1) + fibonnaci(n-2)
}
func fibonnaciHelper(n int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- fibonnaci(n)
}

func main() {
	arr := []int{5, 10, 25, 31}
	ch := make(chan int, len(arr))
	wg := sync.WaitGroup{}
	for i := range arr {
		wg.Add(1)
		go fibonnaciHelper(arr[i], ch, &wg)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for i := range ch {
		fmt.Println(i)
	}
}
```
#### Palindrome:
```go
func palindrome(s string, ch chan<- bool, wg *sync.WaitGroup) {
	defer wg.Done()
	arr := []rune(s)
	l := 0
	r := len(arr) - 1
	for l < r {
		if arr[l] != arr[r] {
			ch <- false
			return
		}
		l++
		r--
	}
	ch <- true
}
func main() {
	arr := []string{"lola", "boob", "fidrat", "Ninja"}
	ch := make(chan bool, len(arr))
	wg := sync.WaitGroup{}
	for i := range arr {
		wg.Add(1)
		go palindrome(arr[i], ch, &wg)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for i := range ch {
		fmt.Println(i)
	}
}
```
Using Map
```go
func palindrome(s string, ch chan<- map[string]bool, wg *sync.WaitGroup) {
	defer wg.Done()
	arr := []rune(s)
	l := 0
	r := len(arr) - 1
	m := make(map[string]bool)
	for l < r {
		if arr[l] != arr[r] {
			m[s] = false
			ch <- m
			return
		}
		l++
		r--
	}
	m[s] = true
	ch <- m
}
func main() {
	arr := []string{"lola", "dood", "fidrat", "ninja"}
	ch := make(chan map[string]bool, len(arr))
	wg := sync.WaitGroup{}
	for i := range arr {
		wg.Add(1)
		go palindrome(arr[i], ch, &wg)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	m := make(map[string]bool)
	for i := range ch {
		for k, v := range i {
			m[k] = v
		}
	}
	fmt.Println(m)
}
```