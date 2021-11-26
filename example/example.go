package main

import (
	"fmt"
	"time"

	"github.com/godruoyi/go-snowflake"
)

func main() {
	// set starttime and machineID for the first time if you wan't to use the default value
	snowflake.SetStartTime(time.Date(2021, 9, 1, 0, 0, 0, 0, time.UTC))
	snowflake.SetMachineID(snowflake.PrivateIPToMachineID()) // testing, not to be used in production

	id := snowflake.ID()
	fmt.Println(id) // 1537200202186752

	sid := snowflake.ParseID(id)
	// SID {
	//     Sequence: 0
	//     MachineID: 0
	//     Timestamp: x
	//     ID: x
	// }
	fmt.Println(sid)
}
