package global

import (
	sf "github.com/bwmarrin/snowflake"
	"time"
)

var (
	Node *sf.Node
)

func SnowFlakeInit() (err error) {
	var st time.Time
	startTime := SnowFlakeSetting.StartTime
	var machineID = SnowFlakeSetting.MachineId

	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	sf.Epoch = st.UnixNano() / 1000000
	Node, err = sf.NewNode(int64(machineID))
	//userId := GenID()
	//fmt.Printf("userid:%v\n", userId)
	return
}
