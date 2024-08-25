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
```