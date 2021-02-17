package tools

func init() {
	Concurrency = make(chan struct{}, parNum)

}
