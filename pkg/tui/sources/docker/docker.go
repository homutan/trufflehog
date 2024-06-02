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
		Label:       "Dockerhub username",
		Key:         "username",
		Required:    false,
		Placeholder: "legenda-cooma",
	}

	password := textinputs.InputConfig{
		Label:       "Dockerhub password",
		Key:         "password",
		Required:    false,
		Placeholder: "legenda-cooma",
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

	username := inputs["username"].Value

	if username != "" {
		command = append(command, "--username="+username)
	}

	password := inputs["password"].Value

	if password != "" {
		command = append(command, "--password="+password)
	}

	return strings.Join(command, " ")
}

func (m dockerCmdModel) Summary() string {
	inputs := m.GetInputs()
	labels := m.GetLabels()
	keys := []string{"images"}

	return common.SummarizeSource(keys, inputs, labels)
}
