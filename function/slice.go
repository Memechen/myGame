package function

func CheckInNumberSlice[T uint64 | int32](a T, s []T) bool {
	for _, val := range s {
		if val == a {
			return true
		}
	}
	return false
}

// DelEleInSlice 只支持 元素不重复的情景
func DelEleInSlice[T uint64 | int32](a T, old []T) (new []T) {
	for i := 0; i < len(old); i++ {
		if old[i] == a {
			new = append(old[0:i], old[i+1:]...)
			return new
		}
	}
	return old
}
