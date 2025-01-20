# tzutils

`tzutils` is a lightweight Go library for time zone conversions and utilities. It helps you work with time zones easily and efficiently.

## Installation

To install `tzutils`, run:

```bash
go get github.com/mawais54013/tzutils
```

# Features
- Get the current time in a specific time zone.
- Convert a given time between time zones.
- Retrieve the current system time zone and offset.
- Get the offset for a specific time zone.

# Usage 
Current Time in a Time Zone 
```
package main

import (
	"fmt"
	"github.com/mawais54013/tzutils"
)

func main() {
	timeInNY, err := tzutils.CurrentTimeIn("America/New_York")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Time in New York:", timeInNY)
	}
}
```

Convert Time Between Time Zones
```
timeInLondon, err := tzutils.ConvertTime(timeInNY, "America/New_York", "Europe/London")
if err != nil {
	fmt.Println("Error:", err)
} else {
	fmt.Println("Time in London:", timeInLondon)
}

```

Get Current System Time Zone
```name, offset := tzutils.CurrentTimeZone()
fmt.Printf("Current Time Zone: %s, Offset: %d seconds\n", name, offset)

```

Get Offset for a Time Zone 
```
offset, err := tzutils.Offset("Asia/Kolkata")
if err != nil {
	fmt.Println("Error:", err)
} else {
	fmt.Printf("Offset for Asia/Kolkata: %d seconds\n", offset)
}

```

# Contributing 
Contributions are welcome! Please fork the repository and submit a pull request.

# License
This project is licensed under the MIT License. See the LICENSE file for details.