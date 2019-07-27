// Copyright 2016 ego authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package com

// Config search config options
type Config struct {
	Engine Engine
	Etcd   Etcd
	Rpc    Rpc

	// Has  [][]string
	Url []string
}

// Engine search engine options
type Engine struct {
	Mode  string
	Using int

	StoreShards int    `toml:"store_shards"`
	StoreEngine string `toml:"store_engine"`
	StoreFolder string `toml:"store_folder"`

	NumShards    int `toml:"num_shards"`
	OutputOffset int `toml:"output_offset"`
	MaxOutputs   int `toml:"max_outputs"`

	GseDict       string `toml:"gse_dict"`
	GseMode       string `toml:"gse_mode"`
	StopTokenFile string `toml:"stop_token_file"`

	Relation string
	Time     string
	Ts       int64
}

// Rpc search rpc options
type Rpc struct {
	GrpcPort []string `toml:"grpc_port"`
	// GrpcPort []int    `toml:"grpc_port"`
	DistPort []string `toml:"dist_port"`
	Port     string
}

// Etcd search etcd options
type Etcd struct {
	Addr     string
	SverName string `toml:"sver_name"`
	Port     []int  `toml:"port"`
}
