# Integers

The other difference between the number system we use and the one computers use is that all of the integer types have a definite size.

Go’s integer types are uint8, uint16, uint32, uint64, int8, int16, int32, and int64.
uint means “unsigned integer” while int means “signed integer.”
Unsigned integers only contain positive numbers (or zero).
In addition, there two alias types: byte (which is the same as uint8) and rune (which is the same as int32).

_NB: Generally, if you are working with integers, you should just use the int type._

# Floating-Point Numbers

Floating-point numbers are numbers that contain a decimal component (i.e., real numbers). For example, 1.234, 123.4, 0.00001234, and 12340000 are all floating-point numbers.

Go has two floating-point types: float32 and float64 (also often referred to as single precision and double precision, respectively). Generally, we should stick with float64 when working with floating- point numbers.
