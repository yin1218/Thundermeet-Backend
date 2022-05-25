package service

import (
	"testing"
	"time"
)

// a successful case
func TestCreatetimeblocksuccess(t *testing.T) {

	// now we execute our method
	secondsEastOfUTC := int((8 * time.Hour).Seconds())
	beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)
	start_t := time.Date(2021, 1, 1, 9, 0, 0, 0, beijing)
	if err := CreateOneTimeblock("2021-01-01T09:00:00+08:00A304", 304, start_t); err != nil {
		t.Errorf("error was not expected while creating timeblock: %s", err)
	}
}
