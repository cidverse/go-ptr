package ptr

func Int64ToInt(i *int64) *int {
	if i != nil {
		v := int(*i)
		return &v
	}

	return nil
}

func IntToInt64(i *int) *int64 {
	if i != nil {
		v := int64(*i)
		return &v
	}

	return nil
}
