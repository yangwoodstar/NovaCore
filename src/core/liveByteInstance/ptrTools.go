package liveByteInstance

func StringPtr(v string) *string {
	return &v
}

func Int32Ptr(v int32) *int32 {
	return &v
}

func BoolPtr(v bool) *bool {
	return &v
}

func StringPtrs(vals []string) []*string {
	ptrs := make([]*string, len(vals))
	for i := 0; i < len(vals); i++ {
		ptrs[i] = &vals[i]
	}
	return ptrs
}

func Float32Ptrs(vals []float32) []*float32 {
	ptrs := make([]*float32, len(vals))
	for i := 0; i < len(vals); i++ {
		ptrs[i] = &vals[i]
	}
	return ptrs
}

func Float32Ptr(v float32) *float32 {
	return &v
}
