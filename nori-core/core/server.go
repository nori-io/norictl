// Copyright © 2018 Secure2Work info@secure2work.com
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of the GNU General Public License
// as published by the Free Software Foundation; either version 2
// of the License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package core

import (
	"errors"
	"net"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/secure2work/nori-cli/proto"
	"github.com/secure2work/nori/core/plugins/manager"
)

var NotSecure = errors.New("Need safe gRPC connect")

const (
	MaxMessageSize int = 100 * 1024 * 1024
)

type Server struct {
	pluginDirs  []string
	gRPCAddress string
	gRPCEnable  bool

	pemFile string
	keyFile string

	pluginManager manager.PluginManager
	passkey       *Passkey
	grpcServer    *grpc.Server
	logger        *logrus.Logger
	gShutdown     chan struct{}
	secure        bool
	done          bool
}

func NewServer(dirs []string, addr string, enable bool) *Server {
	return &Server{
		pluginDirs:  dirs,
		gRPCAddress: addr,
		gRPCEnable:  enable,
		gShutdown:   make(chan struct{}, 1),
	}
}

func (s *Server) SetCertificates(pem, key string) {
	s.pemFile = pem
	s.keyFile = key
}

func (s *Server) Run() error {
	var err error

	s.logger = logrus.New()

	s.pluginManager = manager.GetPluginManager(s.logger)

	for _, dir := range s.pluginDirs {
		s.pluginManager.Load(dir)
	}

	s.passkey, err = NewPasskey()
	if err != nil {
		return err
	}

	logrus.Infof("Passkey: %s", s.passkey)

	wg := new(sync.WaitGroup)

	wg.Add(1)
	go func(s *Server) {
		defer wg.Done()
		for !s.done {
			listener, _ := net.Listen("tcp", s.gRPCAddress)

			var opts []grpc.ServerOption

			opts = append(opts, grpc.MaxMsgSize(MaxMessageSize))

			if opt, err := s.CheckTLS(); err == nil {
				opts = append(opts, opt)
				s.secure = true
			}
			s.grpcServer = grpc.NewServer(opts...)
			commands.RegisterCommandsServer(s.grpcServer, s)
			logrus.WithField("Secure", s.secure).Infof("Starting gRPC server on %s", s.gRPCAddress)
			s.grpcServer.Serve(listener)
		}
	}(s)

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt, syscall.SIGTERM)

	go func(s *Server) {
		for {
			select {
			case sig := <-signalCh:
				s.done = true
				logrus.Infof("Graceful stop gRPC server with signal: %s", sig)
				s.gShutdown <- struct{}{}
			case <-s.gShutdown:
				s.grpcServer.GracefulStop()
			}
		}
	}(s)

	wg.Wait()

	return nil
}

func (s Server) ListCommand(_ context.Context, _ *commands.ListRequest) (*commands.ListReply, error) {
	if !s.secure {
		return nil, NotSecure
	}
	reply := new(commands.ListReply)
	reply.Data = make([]*commands.List, 0)

	for _, plug := range s.pluginManager.Plugins() {
		reply.Data = append(reply.Data, &commands.List{
			Namespace: plug.Plugin().GetMeta().GetNamespace(),
			Name:      plug.Plugin().GetMeta().GetPluginName(),
			Author:    plug.Plugin().GetMeta().GetAuthor(),
		})
	}
	return reply, nil
}

func (s Server) GetCommand(_ context.Context, c *commands.GetRequest) (*commands.ErrorReply, error) {
	if !s.secure {
		return nil, NotSecure
	}
	toolchain, err := SetupToolChain()
	if err != nil {
		return &commands.ErrorReply{
			Status: false,
			Error:  err.Error(),
		}, err
	}

	toolchain.InstallDependencies = c.GetInstallDependencies()
	err = toolchain.Do(c.GetUri())
	if err != nil {
		return &commands.ErrorReply{
			Status: false,
			Error:  err.Error(),
		}, err
	} else {
		return &commands.ErrorReply{
			Status: true,
			Error:  "",
		}, nil
	}
}

func (s Server) RemoveCommand(_ context.Context, c *commands.RemoveRequest) (*commands.ErrorReply, error) {
	return nil, nil
}

func (s Server) EnableCommand(_ context.Context, c *commands.EnableRequest) (*commands.ErrorReply, error) {
	return nil, nil
}

func (s Server) DisableCommand(_ context.Context, c *commands.DisableRequest) (*commands.ErrorReply, error) {
	return nil, nil
}

func (s Server) UploadCommand(_ context.Context, c *commands.UploadRequest) (*commands.ErrorReply, error) {
	path := filepath.Join(s.pluginDirs[0], c.Name)
	if fileExists(path) {
		s.logger.Info("File exist, overwrites")
	}

	err := os.MkdirAll(s.pluginDirs[0], os.ModePerm)
	if err != nil {
		return &commands.ErrorReply{
			Status: false,
			Error:  err.Error(),
		}, err
	}

	f, err := os.Create(path)
	if err != nil {
		return &commands.ErrorReply{
			Status: false,
			Error:  err.Error(),
		}, err
	}

	f.Write(c.So)
	f.Close()

	//s.pluginManager.Load(path)

	return &commands.ErrorReply{
		Status: true,
		Error:  "",
	}, nil
}

func (s Server) UploadCertsCommand(_ context.Context, c *commands.UploadCertsRequest) (*commands.ErrorReply, error) {
	size := int(c.Key[:1][0])
	hmac := c.Key[1 : size+1]
	c.Key = c.Key[size+1:]

	keyBody, err := s.passkey.Decrypt(c.Key, hmac)
	if err != nil {
		return &commands.ErrorReply{
			Status: false,
			Error:  err.Error(),
		}, err
	}

	size = int(c.Pem[:1][0])
	hmac = c.Pem[1 : size+1]
	c.Pem = c.Pem[size+1:]
	pemBody, err := s.passkey.Decrypt(c.Pem, hmac)
	if err != nil {
		return &commands.ErrorReply{
			Status: false,
			Error:  err.Error(),
		}, err
	}

	fKey, err := os.Create(s.keyFile)
	if err != nil {
		return &commands.ErrorReply{
			Status: false,
			Error:  err.Error(),
		}, err
	}
	defer fKey.Close()

	_, err = fKey.Write(keyBody)
	if err != nil {
		return &commands.ErrorReply{
			Status: false,
			Error:  err.Error(),
		}, err
	}

	fPem, err := os.Create(s.pemFile)
	if err != nil {
		return &commands.ErrorReply{
			Status: false,
			Error:  err.Error(),
		}, err
	}
	defer fPem.Close()

	_, err = fPem.Write(pemBody)
	if err != nil {
		return &commands.ErrorReply{
			Status: false,
			Error:  err.Error(),
		}, err
	}

	s.gShutdown <- struct{}{}

	return &commands.ErrorReply{
		Status: true,
		Error:  "",
	}, nil
}

func (s Server) CheckTLS() (grpc.ServerOption, error) {
	if len(s.pemFile) > 0 && len(s.keyFile) > 0 &&
		fileExists(s.pemFile) && fileExists(s.keyFile) {
		creds, err := credentials.NewServerTLSFromFile(s.pemFile, s.keyFile)
		if err != nil {
			return nil, err
		} else {
			return grpc.Creds(creds), nil
		}
	}
	return nil, errors.New("Bad certs")
}

func fileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
