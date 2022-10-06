package timebasedkeyvaluestore

type TimeMap struct {
	TimeIndex map[string][]int
	KVStore   map[string][]string
}

func Constructor() TimeMap {
	return TimeMap{
		TimeIndex: make(map[string][]int),
		KVStore:   make(map[string][]string),
	}
}

func (t *TimeMap) Set(key string, value string, timestamp int) {
	t.TimeIndex[key] = append(t.TimeIndex[key], timestamp)
	t.KVStore[key] = append(t.KVStore[key], value)
}

func (t *TimeMap) Get(key string, timestamp int) string {
	if timeArray, exist := t.TimeIndex[key]; exist {
		timeIndex := t.findFirstIndex(timestamp, timeArray)
		if timeIndex < 0 {
			return ""
		}

		return t.KVStore[key][timeIndex]
	}

	return ""
}

func (t *TimeMap) findFirstIndex(timestamp int, indexArray []int) int {
	l, r := 0, len(indexArray)-1

	for l <= r {
		m := (l + r) >> 1
		if indexArray[m] > timestamp {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return r
}
