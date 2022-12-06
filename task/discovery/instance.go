package discovery

import (
	"encoding/json"
	"errors"
	"fmt"
	"google.golang.org/grpc/resolver"
	"strings"
)

type Server struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Version string `json:"version"`
	Weight  int64  `json:"weight"` //权重
}

func BuildPrefix(server Server) string {
	if server.Version == "" {
		return fmt.Sprintf("/%s/", server.Name)
	}
	return fmt.Sprintf("/%s/%s/", server.Name, server.Version)
}

//拼接注册路径
func BuildRegisterPath(server Server) string {
	return fmt.Sprintf("%s%s", BuildPrefix(server), server.Address)
}

func ParseValue(value []byte) (Server, error) {
	server := Server{}
	if err := json.Unmarshal(value, &server); err != nil {
		return server, err
	}
	return server, nil
}

func SplitPath(path string) (Server, error) {
	server := Server{}
	strs := strings.Split(path, "/")
	if len(strs) == 0 {
		return server, errors.New("invalid path")
	}
	server.Address = strs[len(strs)-1]
	return server, nil
}

func Exist(l []resolver.Address, addr resolver.Address) bool {
	for i := range l {
		if l[i].Addr == addr.Addr {
			return true
		}
	}
	return false
}

// Remove helper function
func Remove(s []resolver.Address, addr resolver.Address) ([]resolver.Address, bool) {
	for i := range s {
		if s[i].Addr == addr.Addr {
			s[i] = s[len(s)-1]
			return s[:len(s)-1], true
		}
	}
	return nil, false
}
