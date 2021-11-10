package dspattn

type tcpServer struct {
	addr string
	o    *Option
}

type Option struct {
	compress bool
	encrypt  bool
}

type OptionFunc func(o *Option)

func withCompress() OptionFunc {
	return func(o *Option) {
		o.compress = true
	}
}
func newTCPServer(addr string, opts ...OptionFunc) *tcpServer {
	o := &Option{}
	for _, opt := range opts {
		opt(o)
	}

	return &tcpServer{
		addr: addr,
		o:    o,
	}
}
