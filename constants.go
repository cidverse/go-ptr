package ptr

func True() *bool {
	b := true
	return &b
}

func False() *bool {
	b := false
	return &b
}
