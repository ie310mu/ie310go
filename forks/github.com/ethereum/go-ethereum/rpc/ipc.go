// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package rpc

import (
	"context"
	"net"

	"github.com/ie310mu/ie310go/forks/github.com/ethereum/go-ethereum/log"
	"github.com/ie310mu/ie310go/forks/github.com/ethereum/go-ethereum/p2p/netutil"
)

// ServeListener accepts connections on l, serving JSON-RPC on them.
//开始监听
func (s *Server) ServeListener(l net.Listener) error {
	for {
		//这个在没有连接时会一直阻塞，即使server.Stop()也一样，
		//但stop后这里还是可以accept，但无法ServeCodec了，所以server.stop实际上是有点问题的
		conn, err := l.Accept()
		if netutil.IsTemporaryError(err) {
			log.Warn("RPC accept error", "err", err)
			continue
		} else if err != nil {
			return err
		}
		log.Trace("Accepted RPC connection", "conn", conn.RemoteAddr())
		go s.ServeCodec(NewJSONCodec(conn), OptionMethodInvocation|OptionSubscriptions) //使用json编码解码器，并启动编码解码
	}
}

//add by ie310 at 2019.02.22
func IpcListen(endpoint string) (net.Listener, error) {
	return ipcListen(endpoint)
}

// DialIPC create a new IPC client that connects to the given endpoint. On Unix it assumes
// the endpoint is the full path to a unix socket, and Windows the endpoint is an
// identifier for a named pipe.
//
// The context is used for the initial connection establishment. It does not
// affect subsequent interactions with the client.
//创建一个ipc client并连接到server  unix系统中是一个全路径的socket，windows是命名管道的名称
func DialIPC(ctx context.Context, endpoint string) (*Client, error) {
	return newClient(ctx, func(ctx context.Context) (ServerCodec, error) {
		conn, err := newIPCConnection(ctx, endpoint)
		if err != nil {
			return nil, err
		}
		return NewJSONCodec(conn), err
	})
}
