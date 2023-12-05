# Control Structures in Go

Go has a few common structures that control the flow of a program.

## if/else

Go has two keywords that can be used to create conditional code blocks in your program.

- `if`
- `else`

The `if` statement can be used by itself, or with subsequent `else` blocks. It is followed by a _conditional_ that determines if the attached code block should be executed or not.

```go
var isOnline boolean = users.GetStatus("Bryantellius")

if isOnline {
    streams.SendMessage()
} else {
    streams.SendPushNotification()
}
```

In this example, the variable `isOnline` is a boolean value used as the condition for this if/else block. If `isOnline` is true, then `streams.SendMessage()` will run. Otherwise, `streams.SendPushNotification()` will execute.

### Optional `if` Statement

The `if` condition can be preceded by a statement. If any variables are declared in the statement, they can be accessed in the if block and any preceding blocks in the if/else chain.

```go
wallet := 14
price := 12.99
tax := .10

if total := price*tax; wallet > total {
    fmt.Printf("Can afford. Wallet balance is now %f", wallet - total)
} else {
    fmt.Printf("Cannot afford. Balance would need %f", total - wallet)
}
```

## Switch

A `switch` case allows you to evaluate multiple conditionals in one structure. The `switch` keyword is followed by a value, and the block contains `case` statements that are evaluated as equality. Unlike other languages, the `break` functionality is implicitly applied to each case.

> The condition can be omitted to evaluate for `switch true`.

> If you want to enable functionality similar to the `break` keyword in other languages, you can use the `fallthrough` keyword.

```go
birthMonth := 1
var monthStr string

switch {
    case birthMonth == 12 || birthMonth < 3:
        monthStr = "Winter"
    case birthMonth < 6:
        monthStr = "Spring"
    case birthMonth < 9:
        monthStr = "Summer"
    case birthMonth < 12:
        monthStr = "Fall"
}

fmt.Printf("You were born in %s", monthStr)
```

### Type Switch

A special switch block can evaluate the `type` of interface values. Use `T.(type)` in the condition to then evaluate the type of the value in the case statements.

```go
value := GetRandomValue()

switch value.(type) {
    case string:
        fmt.Println("value is a string")
    case int:
        fmt.Println("value is a integer")
    case bool:
        fmt.Println("value is a boolean")
    default:
        fmt.Println("value type is unknown")
}
```
