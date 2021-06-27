# GOnads
Gonads is a monadic package for Golang. Both `result` and `option` are heavily (if not exclusivley) inspired by the Rust standard library ([here](https://doc.rust-lang.org/std/result/enum.Result.html) for `Result` and [here](https://doc.rust-lang.org/stable/std/option/enum.Option.html) for `Option`). Rust has two `Option` structures, one from `std::option` and the other from `std::io`, identical in every way.

## Use Case
Timid developers approaching Go for the first time are often caught by `if err != nil` statements throughout the code. For seasoned Go engineers, this doesn't shock us, and in many ways we have come to appreciate why we [handle errors](https://blog.golang.org/error-handling-and-go) this way. Yet still, for those of us who appreciate both the __science__ and the __art__ of writing code, finding elegant, concise, and expressive solutions to perpetual tangled webs is always on our minds. Enter Happy Path Programming.

### Happy Path Programming
Operating on Monads is what I call __Happy Pathing__, where you compose your functions in a way that __hopes__ for the best, yet is still resiliant and capable of responding to the worst. Monads allow for both of these and offer an elegant alternative to `if err != nil`. Of course, __somewhere__ in your code, you'll have to handle errors this way. But, ideally, this type of check is burried in the lowest layers of your application design and can be written in a minimal amount of places.

## How do they work?
The first thing you need to do is wrap an operation in a Result. Let's say you need to open a file, and that file may or may not exist.
```go
r := result.Handle(os.Open("./filename.jpg"))
```
Now, 
