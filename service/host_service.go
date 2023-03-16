package service

import (
	"code/service/dto"
	"context"
	"fmt"
	"github.com/apenella/go-ansible/pkg/adhoc"
	"github.com/apenella/go-ansible/pkg/options"
	"github.com/spf13/viper"
)

var hostService *HostService

type HostService struct {
	BaseService
}

func NewHostService() *HostService {
	if hostService == nil {
		hostService = &HostService{}
	}
	return hostService
}

func (m *HostService) Shutdown(iShutdownHostDTO dto.ShutdownHostDTO) error {
	var errResult error
	stHostIP := iShutdownHostDTO.HostIP
	fmt.Println(stHostIP)
	ansibleConnectionOptions := &options.AnsibleConnectionOptions{
		Connection: "ssh",
		User:       viper.GetString("ansible.user.name"),
	}

	ansibleAdhocOptions := &adhoc.AnsibleAdhocOptions{
		Inventory:  fmt.Sprintf("%s,", stHostIP),
		ModuleName: "command",
		Args:       viper.GetString("ansible.shutdown.args"),
		ExtraVars: map[string]any{
			"ansible-password": viper.GetString("ansible.user.password"),
		},
	}

	adhocCmd := &adhoc.AnsibleAdhocCmd{
		Pattern:           "all",
		Options:           ansibleAdhocOptions,
		ConnectionOptions: ansibleConnectionOptions,
		StdoutCallback:    "oneline",
	}

	errResult = adhocCmd.Run(context.TODO())

	return errResult
}
