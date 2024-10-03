# for statements

Other programming languages have a lot of different types of loops (while, do, until, foreach, ...) but Go only has one that can be used in a variety of different ways.

The for statement allows us to repeat a list of statements (a block) multiple times.

for initialization; condition; post {
}

For example:
for i := 1; i <= 10; i++ {
fmt.Println(i)
}

# if statement

```
if condition {
    statement
} else if condition{

} else {

}
```

# Switch statement

A switch statement starts with the keyword switch followed by an expression (in this case, i) and then a series of cases. The value of the expression is compared to the expression following each case keyword. If they are equivalent, then the statements following the : are executed.
Like an if statement, each case is checked top down and the first one to succeed is chosen. A switch also supports a default case that will happen if none of the cases matches the value (similar to how the else works in an if statement).
for, if, and switch are the main control flow statements.

```
switch i {
case 0: fmt.Println("Zero")
case 1: fmt.Println("One")
case 2: fmt.Println("Two")
case 3: fmt.Println("Three")
case 4: fmt.Println("Four")
case 5: fmt.Println("Five")
default: fmt.Println("Unknown Number") }
```
