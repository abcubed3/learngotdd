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