package gcloud

import (
	"context"
	"fmt"

	"github.com/samkhawase/testcontainers-go"
	"github.com/samkhawase/testcontainers-go/wait"
)

// RunPubsubContainer creates an instance of the GCloud container type for Pubsub
func RunPubsubContainer(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*GCloudContainer, error) {
	req := testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image:        "gcr.io/google.com/cloudsdktool/cloud-sdk:367.0.0-emulators",
			ExposedPorts: []string{"8085/tcp"},
			WaitingFor:   wait.ForLog("started"),
		},
		Started: true,
	}

	settings, err := applyOptions(&req, opts)
	if err != nil {
		return nil, err
	}

	req.Cmd = []string{
		"/bin/sh",
		"-c",
		"gcloud beta emulators pubsub start --host-port 0.0.0.0:8085 " + fmt.Sprintf("--project=%s", settings.ProjectID),
	}

	container, err := testcontainers.GenericContainer(ctx, req)
	if err != nil {
		return nil, err
	}

	return newGCloudContainer(ctx, 8085, container, settings)
}
