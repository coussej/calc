#calc

This package provides some math functions on integer, or on floats returning integers. 

As the standard math package only implements functions as Min & Max on floats, it can be cumbersome to use them on ints, resulting in either typecasting or rolling your own functions. 

Secondly, things line Floor and Ceil are supposed to work on float input, but they also return floats! In many cases, this result in post-function type casting, leading too uglier code. Therefore, this packages also includes some wrappers around the standard functions that take care of the casting for you.

Lastly, some functions (like Round) are simply missing in the standard package, and therefore included here.