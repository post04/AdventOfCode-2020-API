# What it does
This package has 4 functions.

1.) GetBody
  - This function gets the main html body for the AdventOfCode stats page.
 
 2.) GetNumToMultiply
  - This function gets the number to multiply by, for example "Each * or * star represents up to ??? users." The number would be that ??? in this example.
 
 3.) GetCurrentDay
  - This function gets the current Advent day, like 1, 2, 3, etc.
 
 4.) GetAllStats
  - This function gets the Total submittions and One star submittions for the advent calander.
  - This function returns an array of strings, each string will look like "1-7700-820".
  - "1" is the day that this data refers to.
  - "7700" is the total submittions.
  - "820" is the one star submittions.
  - To get two star submittions just convert 7700 and 820 to ints and minus 7700 by 820. 
  
 
# How to use

```go
package main
```
