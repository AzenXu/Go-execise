package ops

import (
	"time"
)

func AddRecord(vid string) error {
	return nil
}

func DeleteVideo(vid string) error {
	time.Sleep(1 * time.Second)
	return nil
}

func GetAllPaddingVIDs() (vids []string) {
	return nil
}

var mock = []string{"a","b","c","d","e","f","g","h","i"}
var index = 0

func GetPaddingVIDs(count int) (vids []string, err error) {
	index += 3
	if index > 9 {
		return []string{}, nil
	}
	return []string{mock[index - 3],mock[index - 2],mock[index - 1]}, nil
}
