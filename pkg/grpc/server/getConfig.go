package server

import (
	"bytes"
	"context"
	"fmt"
	"strings"

	"github.com/compscore/compscore/pkg/config"
	"github.com/compscore/compscore/pkg/grpc/proto"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

func (*compscoreServer_s) GetConfig(ctx context.Context, in *proto.GetConfigRequest) (*proto.GetConfigResponse, error) {
	logrus.Info("Received get config request")

	heading := strings.Repeat("=", 20) + "\n"

	var conf *bytes.Buffer = new(bytes.Buffer)

	fmt.Fprintf(conf, "Name: %v\n\n", config.Name)

	web, err := yaml.Marshal(config.Web)
	if err != nil {
		return nil, err
	}
	fmt.Fprintf(conf, "Web:\n%v%v\n\n", heading, string(web))

	teams, err := yaml.Marshal(config.Teams)
	if err != nil {
		return nil, err
	}
	fmt.Fprintf(conf, "Teams:\n%v%v\n\n", heading, string(teams))

	scoring, err := yaml.Marshal(config.Scoring)
	if err != nil {
		return nil, err
	}
	fmt.Fprintf(conf, "Scoring:\n%v%v\n\n", heading, string(scoring))

	engine, err := yaml.Marshal(config.Engine)
	if err != nil {
		return nil, err
	}
	fmt.Fprintf(conf, "Engine:\n%v%v\n\n", heading, string(engine))

	checks, err := yaml.Marshal(config.Checks)
	if err != nil {
		return nil, err
	}
	fmt.Fprintf(conf, "Checks:\n%v%v\n\n", heading, string(checks))

	return &proto.GetConfigResponse{Config: conf.String()}, nil
}
