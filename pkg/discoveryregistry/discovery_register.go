// Copyright © 2023 OpenIM. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package discoveryregistry

import (
	"context"

	"google.golang.org/grpc"
)

type Conn interface {
	GetConns(ctx context.Context, serviceName string, opts ...grpc.DialOption) ([]grpc.ClientConnInterface, error)
	GetConn(ctx context.Context, serviceName string, opts ...grpc.DialOption) (grpc.ClientConnInterface, error)
	AddOption(opts ...grpc.DialOption)
	CloseConn(conn grpc.ClientConnInterface)
	// do not use this method for call rpc
	GetClientLocalConns() map[string][]grpc.ClientConnInterface
}

type SvcDiscoveryRegistry interface {
	Conn
	Register(serviceName, host string, port int, opts ...grpc.DialOption) error
	UnRegister() error
	CreateRpcRootNodes(serviceNames []string) error
	RegisterConf2Registry(key string, conf []byte) error
	GetConfFromRegistry(key string) ([]byte, error)
}
