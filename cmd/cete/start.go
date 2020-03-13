// Copyright (c) 2020 Minoru Osuka
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"github.com/mosuka/cete/http"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/mosuka/cete/kvs"
	cetelog "github.com/mosuka/cete/log"
	"github.com/urfave/cli"
)

func execStart(c *cli.Context) error {
	nodeId := c.String("node-id")
	bindAddr := c.String("bind-addr")
	grpcAddr := c.String("grpc-addr")
	httpAddr := c.String("http-addr")
	dataDir := c.String("data-dir")
	joinAddr := c.String("join-addr")

	logger := cetelog.NewLogger(
		c.String("log-level"),
		"",
		log.LstdFlags|log.Lmicroseconds|log.LUTC,
		c.String("log-file"),
		c.Int("log-max-size"),
		c.Int("log-max-backups"),
		c.Int("log-max-age"),
		c.Bool("log-compress"),
	)

	httpAccessLogger := http.NewLogger(
		c.String("http-access-log-file"),
		c.Int("http-access-log-max-size"),
		c.Int("http-access-log-max-backups"),
		c.Int("http-access-log-max-age"),
		c.Bool("http-access-log-compress"),
	)

	svr, err := kvs.NewServer(nodeId, bindAddr, grpcAddr, httpAddr, dataDir, joinAddr, logger, httpAccessLogger)
	if err != nil {
		return err
	}

	quitCh := make(chan os.Signal, 1)
	signal.Notify(quitCh, os.Kill, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go svr.Start()

	<-quitCh

	svr.Stop()

	return nil
}