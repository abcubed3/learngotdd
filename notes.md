quii.gitbook.io/learn-go-with-tests
godoc -http :8000 takes yo to localhost:8000/pkg

const help improve performance because they are not created every time the func or variable is called

Discipline
- Write a test 
- Make the compiler pass
- Run the test, see that it fails and check the error message is meaningful
- Write enough code to make the test pass
- refactor

If a named return value is added to a func definition, then the zero value is assigned by default. No need to return variable to the func, just call `return`

The strings package has a lot of useful methods to manage strings
_ (underscore) is used to ignore a variable value or index value

Variadic functions are functions you can call with any number of trailing arguments. E.g fmt.Println() is a common variadic funciton. The form is func(slice ...[]int)

make(slice, N) - allows to create a slice with a starting capacity

A Struct is a named collection of fields where you can store data
Methods are very similar to functions but they are called by invoking them on an instance of a particular type

Interfaces are  very powerful concept in statically typed languages like GO, They allow you to make functions that can be used with different types and create highly-decoupled doe whilst still maintaiing type-safety

In GO, interface resolution is implicity. If the type you pass in matches what the interface is asking for, it will compile

You can create anonymous structs in a slice
areaTests := []struct {
		shape Shape
		want  float64
	}
Anonymous slice structs can be good for Table driven tests. You can build a list of test cases based on the interfaces used

In Go when you call a function or a method, the arguements are copied
There are two reasons to use a pointer receiver.
The first is so that the method can modify the value that its receiver points to.
The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct,
Pointers to structs are called struct pointers and they are automatically dereferenced

Apart from structs, Go allows to create new types from exisiting ones. e.g type Bitcoin float32

The fmt package allows you to create how your user defined type is printed when used with a formatter like %s
Use t.Fatal to stop and reduce panic
The var keyword allows to define values global to the package

Follwow github.com/quii for more Go best practices

In maps, the lookup returns 2 values, where the 2nd is a boolean which indicates if the key was found successfully

An interesting property of maps is that you can modify them without passing them as a pointer. This is because map is a reference type.

Reference types hold reference to the underlying data structure much like a pointer. E.g hash table 

A map can be a nil value, and behaves like an empty map when readumg, but attempts to write to a nil map will cause runtime panic. So you can't write var m map[string][string].

To prevent runtime panic, use the syntax below
var dictionary = map[string][string] OR 
var dictionary = make(map[string]string)

Its good practice to return string and error, when fetching data. While its good to return error when adding/updating data
It's good practice to create a Error type for every struct/type you create. That helps you manage errors specific to that struct or object
Also good to have Struct/Type Error implement the Error() interface, to create a customized error method

Go has a builtin function `delete` to delete items of a map. E.g delete(map,key)
An operation that does not block in Go will run in a separate process called a goroutine. Think of a process as reading down the page of Go code from top to bottom, going 'inside' each function when it gets called to read what it does. When a separate process starts it's like another reader begins reading inside the function, leaving the original reader to carry on going down the page.
To tell Go to start a new goroutine we turn a function call into a go statement by putting the keyword go in front of it: go doSomething()

Because the only way to start a goroutine is to put go in front of a function call, we often use anonymous functions when we want to start a goroutine.
Anonymous functions have a number of features which make them useful;  Firstly, they can be executed at the same time that the're declared - this is what the () at the end of the anonymous function is doing;
go func() {
            results[url] = wc(url)
        }()

Secondly they maintain access to the lexical scope they are defined in - all the variables that are available at the point when you declare the anonymous function are also available in the body of the function.

Race condition, a bug that occurs when the output of our software is dependent on the timing and sequence of events that we have no control over.
Go can help us to spot race conditions with its built in race detector. To enable this feature, run the tests with the race flag: go test -race
Channels are a Go data structure that can both receive and send values. These operations, along with their details, allow communication between different processes.
The send operator <- takes a channel on the left and a value on the right:
The receiver expresssion is same as the send operator, but the channel and values reversed

Concurrency in a nutshell
- goroutines, the basic unit of concurrency in Go, which let us check more
than one website at the same time.
- anonymous functions, which we used to start each of the concurrent processes
that check websites.
- channels, to help organize and control the communication between the
different processes, allowing us to avoid a race condition bug.
- the race detector which helped us debug problems with concurrent code

In the standard library, there is a package called net/http/httptest where you can easily create a mock HTTP server.
By prefixing a function call with defer it will now call that function at the end of the containing function

select {
    case <-ping(a):
        return a
    case <-ping(b):
        return b
    }

In concurrency, you can wait for values to be sent to a channel with `myVar := <-ch`. This is a blocking call, as you're waiting for a value. What `select` lets you do is wait on multiple channels. 
`select`
Helps you wait on multiple channels.
Sometimes you'll want to include time.After in one of your cases to prevent your system blocking forever.

httptest
A convenient way of creating test servers so you can have reliable and controllable tests.
Using the same interfaces as the "real" net/http servers which is consistent and less for you to learn.
Useful time functions
 - time.After
 - time.Since
 - time.Seconds
 - time.Sleep(duration)
 - time.Duration

Reflection in computing is the ability of a program to examine its own structure, particularly through types;
Interface{} lets you write functions without knowing the parameter type at compile time. It's of anytype. 
In short only use reflection if you really need to.

The reflect package has a function ValueOf which returns us a Value of a given variable. 

When you're doing a comparison on the same value more than once generally refactoring into a switch will improve readability and make your code easier to extend.

A WaitGroup waits for a collection of goroutines to finish. The main goroutine calls Add to set the number of goroutines to wait for. Then each of the goroutines runs and calls Done when finished. At the same time, Wait can be used to block until all goroutines have finished.
We are using sync.WaitGroup which is a convenient way of synchronising concurrent processes
A Mutex is a mutual exclusion lock. The zero value for a Mutex is an unlocked mutex. A Mutex must not be copied after first use.

Use channels when passing ownership of data 
Use mutexes for managing state