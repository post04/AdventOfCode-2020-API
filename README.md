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

```
go get github.com/postrequest69/AdventOfCode-2020-API
```

```go
package main

import (
	"fmt"
	"strconv"
	"strings"

	adventapi "github.com/postrequest69/AdventOfCode-2020-API"
)

func main() {
	body := adventapi.GetBody()
	numToMul := adventapi.GetNumToMultiply(body)
	fmt.Println("Number:", numToMul)
	currentday := adventapi.GetCurrentDay(body)
	fmt.Println("Current Day:", currentday)
	allStats := adventapi.GetAllStats(body)
	for _, stat := range allStats {
		stats := strings.Split(stat, "-")

		total, err := strconv.Atoi(stats[1])
		if err != nil {
			fmt.Println(err)
		}
		oneStar, err := strconv.Atoi(stats[2])
		if err != nil {
			fmt.Println(err)
		}
		twostar := total - oneStar
		fmt.Println("Day: " + stats[0] + "\n   - Total Submittions: " + strconv.Itoa(total) + "\n   - Two Star Submittions: " + strconv.Itoa(twostar) + "\n   - One Star Submittions: " + strconv.Itoa(oneStar))
	}
}
```

Output:
```
Number: 241
Current Day: 1
Day: 1
   - Total Submittions: 8674
   - Two Star Submittions: 7747
   - One Star Submittions: 927
```
