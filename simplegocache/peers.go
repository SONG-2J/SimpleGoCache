package simplegocache

import pb "SimpleGoCache/cachepb"

// PeerPicker是定位时必须实现的接口
// 拥有特定密钥的peer
type PeerPicker interface {
	PickPeer(key string) (peer PeerGetter, ok bool)
}

// PeerGetter是必须由peer实现的接口
type PeerGetter interface {
	Get(in *pb.Request, out *pb.Response) error
}
