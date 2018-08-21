package ops

import (
	"daker.wang/Azen/Go-execise/dbops"
	"os"
)

func AddRecord(vid string) error {
	return dbops.AddVideoDeletionRecord(vid)
}

func DeleteVideo(vid string) error {
	//  拼接路径拿到视频
	uri := "./videos" + "/" + vid
	//  删除视频
	e := os.Remove(uri); if e != nil {
		return e
	}

	//  删除数据库待删表中记录
	e = dbops.DelVideoDeletionRecord(vid); if e != nil {
		return e
	}

	return nil
}

func GetPaddingVIDs(count int) (vids []string, err error) {
	results, e := dbops.ReadVideoDeletionRecord(count); if e != nil {
		return nil, err
	}
	return results, nil
}

//func _testDeleteVideo(vid string) error {
//	time.Sleep(1 * time.Second)
//	return nil
//}
//
//var _mock = []string{"a","b","c","d","e","f","g","h","i"}
//var _index = 0
//
//func _testGetPaddingVIDs(count int) (vids []string, err error) {
//	_index += 3
//	if _index > 9 {
//		return []string{}, nil
//	}
//	return []string{_mock[_index - 3],_mock[_index - 2],_mock[_index - 1]}, nil
//}
