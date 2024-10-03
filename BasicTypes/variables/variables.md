# Variables

## var

Variables in Go are created by first using the `var` keyword, then specifying the variable name (x) and the type (string)

## const

Go also has support for constants. Constants are essentially variables whose values cannot be changed later. They are created in the same way you create variables, but instead of using the var keyword we use the `const` keyword.

There are several ways to declare a string variable; these are all equivalent:

- s := ""
- var s string
- var s = ""
- var s string = ""

The first form, a short variable declaration, is the most compact, but it may be used only within a function, not for package-level variables.

The second form relies on default initialization to the zero value for strings, which is "".

The third form is rarely used except when declaring multiple variables.

The fourth form is explicit about the variable’s type, which is redundant when it is the same as that of the initial value but necessary in other cases where they are not of the same type. In practice, you should generally use one of the first two forms, with explicit initialization to say that the initial value is important and implicit initialization to say that the initial value doesn’t matter.

Naming a variable properly is an important part of software development. Names must start with a letter and may contain letters, numbers, or the underscore symbol `_`.
