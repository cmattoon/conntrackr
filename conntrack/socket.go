package conntrack

func NewSocket(addr string, port string) *Socket {
	return &Socket{
		Addr: addr,
		Port: port,
	}
}

func NewSocketPair(src *Socket, dst *Socket) *SocketPair {
	return &SocketPair{
		Src: src,
		Dst: dst,
	}
}
