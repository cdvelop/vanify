package server

func (d Server) Write(p []byte) (n int, err error) {
	d.programStartedMessages <- string(p)
	// fmt.Println(string(p))
	return len(p), nil
}
