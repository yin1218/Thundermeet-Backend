package helper

import (
	"strconv"
)

func ConvertToTimeblockId(timeString string, eventId int64) string {
	timeblock_id := timeString + "A" + strconv.Itoa(int(eventId))
	return timeblock_id
}
