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
	"net"

	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/secure2work/nori-cli/proto"
	"github.com/secure2work/nori/core/registry"
)

type Server struct {
	pluginDirs    []string
	gRPCAddress   string
	gRPCEnable    bool
	pluginManager registry.PluginManager
}

func NewServer(dirs []string, addr string, enable bool) *Server {
	return &Server{
		pluginDirs:  dirs,
		gRPCAddress: addr,
		gRPCEnable:  enable,
	}
}

func (s *Server) Run() error {
	lis, err := net.Listen("tcp", s.gRPCAddress)
	if err != nil {
		return err
	}

	log := logrus.New()

	s.pluginManager = registry.GetPluginManager(log)

	for _, dir := range s.pluginDirs {
		s.pluginManager.Load(dir)
	}

	srv := grpc.NewServer()
	commands.RegisterCommandsServer(srv, s)
	logrus.Infof("Starting gRPC server on %s", s.gRPCAddress)
	return srv.Serve(lis)
}

func (s Server) ListCommand(_ context.Context, _ *commands.ListRequest) (*commands.ListReply, error) {
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
