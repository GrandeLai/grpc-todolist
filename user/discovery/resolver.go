package discovery

import (
	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/resolver"
)

type Resolver struct {
	schema       string
	EtcdAddrs    []string
	DialTimeout  int
	closeCh      chan struct{}
	watchCh      clientv3.WatchChan
	cli          *clientv3.Client
	keyPrefix    string
	srvAddrsList []string
	cc           resolver.ClientConn
	logger       *logrus.Logger
}
