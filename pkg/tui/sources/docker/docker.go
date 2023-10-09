package docker

import (
	"strings"

	"github.com/trufflesecurity/trufflehog/v3/pkg/tui/common"
	"github.com/trufflesecurity/trufflehog/v3/pkg/tui/components/textinputs"
)

type dockerCmdModel struct {
	textinputs.Model
}

func GetFields() dockerCmdModel {
	images := textinputs.InputConfig{
		Label:       "Docker image(s)",
		Key:         "images",
		Required:    true,
		Help:        "Separate by space if multiple.",
		Placeholder: "trufflesecurity/secrets",
	}

	username := textinputs.InputConfig{
		Label: "Docker Hub username",
		Key: "username",
		Required: true,		
		Placeholder: "legenda-cooma",
	}

	password := textinputs.InputConfig{
		Label: "Docker Hub password",
		Key: "password",
		Required: true,		
		Placeholder: "armpitenjoyer228",
	}

	return dockerCmdModel{textinputs.New([]textinputs.InputConfig{images, username, password})}
}

func (m dockerCmdModel) Cmd() string {

	var command []string
	command = append(command, "trufflehog", "docker")

	inputs := m.GetInputs()
	vals := inputs["images"].Value

	if vals != "" {
		images := strings.Fields(vals)
		for _, image := range images {
			command = append(command, "--image="+image)
		}
	}

	command = append(command, "--username="+inputs["username"].Value)
	command = append(command, "--password="+inputs["password"].Value)

	return strings.Join(command, " ")
}

func (m dockerCmdModel) Summary() string {
	inputs := m.GetInputs()
	labels := m.GetLabels()
	keys := []string{"images"}

	return common.SummarizeSource(keys, inputs, labels)
}
