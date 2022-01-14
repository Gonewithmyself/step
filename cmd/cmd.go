package cmd

type baseInfo struct {
	head int32
	lv   int32
}
type user struct {
	id   int64
	base *baseInfo
}
type player struct {
	id   int64
	base baseInfo
}

func test() {
	var u = new(user)
	u.base = new(baseInfo)
	var p = new(player)

	_ = p
}
