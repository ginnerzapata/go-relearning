# Learning Go

Notes from working through [Go: The Complete Developer's Guide](https://www.udemy.com/course/go-the-complete-developers-guide/). This repo is a collection of small exercises, not one application. I'm keeping this here so I can remember what I practiced.

## What I learned

- A Go program starts in `main()`. `helloworld/` was the simplest version: import `fmt` and print something.
- Slices hold a list of values. I used `range` to loop through them, `append` to build them, and slice expressions to split a deck into a hand and the remaining cards. Slices can share the same underlying data, so changing elements can affect the original slice too.
- Maps store values by key. In `map/`, I made a map of colors to hex codes and another map of IDs to `person` structs.
- Structs group related fields. In `structs/`, a `person` includes contact information. Methods can belong to a type, and a pointer receiver lets a method change the original value instead of a copy.
- Interfaces describe what a value can do. The greeting bots and shapes both show this: different types can be passed to the same function as long as they have the required method. I didn't have to explicitly declare that they implement an interface.
- Go often returns an `error` alongside a result. The file and HTTP exercises check errors before using the result; this is something I need to keep doing when I write my own code.
- `io.Copy` moves data from a reader to a writer. In `args/`, it copies a file to standard output. In `http/`, I made a type with a `Write` method so it could receive an HTTP response body.
- Goroutines let work happen concurrently, and channels let goroutines send values to each other. In `area/`, I tried checking several websites at once and sending their URLs through a channel before checking again later.
- Tests go in `_test.go` files and use the `testing` package. The cards tests check the contents of a new deck and whether saving and loading a deck preserves it.

## Exercises

| Folder                       | What I tried                                                                                                                          |
| ---------------------------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| [`helloworld/`](helloworld/) | Printing a message.                                                                                                                   |
| [`evenorodd/`](evenorodd/)   | Looping through numbers and using `%` with an `if` statement. Note to self: the current `Even` and `Odd` labels are swapped.          |
| [`cards/`](cards/)           | Creating a custom `deck` type from a slice, adding methods, dealing with slices, shuffling, saving/loading a file, and writing tests. |
| [`structs/`](structs/)       | Struct fields, embedded contact info, methods, and pointer receivers.                                                                 |
| [`map/`](map/)               | Maps with strings and structs as values.                                                                                              |
| [`interfaces/`](interfaces/) | Sharing behavior between greeting bots.                                                                                               |
| [`channels/`](channels/)     | Sharing behavior between shapes through a `getArea` interface. Despite the folder name, this one is about interfaces, not channels.   |
| [`args/`](args/)             | Reading a filename from command-line arguments and printing its contents.                                                             |
| [`http/`](http/)             | Making an HTTP request and writing the response through a custom writer.                                                              |
| [`area/`](area/)             | Goroutines and a channel for repeated website checks. Despite the folder name, this one isn't about calculating area.                 |

## Running things

Each folder is its own little program. For example:

```sh
go run helloworld/main.go
go run cards/main.go cards/deck.go
```

The cards exercise has a Go module and tests:

```sh
cd cards
go test ./...
```

`args/` needs a filename argument. `http/` and `area/` make real network requests; `area/` keeps checking the sites rather than finishing on its own.
