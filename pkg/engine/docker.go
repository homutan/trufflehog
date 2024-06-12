package engine

import (
	"runtime"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"

	"github.com/trufflesecurity/trufflehog/v3/pkg/context"
	"github.com/trufflesecurity/trufflehog/v3/pkg/pb/credentialspb"
	"github.com/trufflesecurity/trufflehog/v3/pkg/pb/sourcespb"
	"github.com/trufflesecurity/trufflehog/v3/pkg/sources"
	"github.com/trufflesecurity/trufflehog/v3/pkg/sources/docker"
)

// ScanDocker scans a given docker connection.
func (e *Engine) ScanDocker(ctx context.Context, c sources.DockerConfig) error {
	connection := &sourcespb.Docker{Images: c.Images, Cache: c.Cache, CacheDb: c.CacheDb, LocalCache: c.LocalCache, LocalCacheDir: c.LocalCacheDir}

	switch {
	case c.UseDockerKeychain:
		connection.Credential = &sourcespb.Docker_DockerKeychain{DockerKeychain: true}
	case len(c.BearerToken) > 0:
		connection.Credential = &sourcespb.Docker_BearerToken{BearerToken: c.BearerToken}
	case len(c.Username) > 0 && len(c.Password) > 0:
		connection.Credential = &sourcespb.Docker_BasicAuth{BasicAuth: &credentialspb.BasicAuth{Username: c.Username, Password: c.Password}}
	default:
		connection.Credential = &sourcespb.Docker_Unauthenticated{}
	}

	var conn anypb.Any
	err := anypb.MarshalFrom(&conn, connection, proto.MarshalOptions{})
	if err != nil {
		ctx.Logger().Error(err, "failed to marshal gitlab connection")
		return err
	}

	sourceName := "trufflehog - docker"
	sourceID, jobID, _ := e.sourceManager.GetIDs(ctx, sourceName, docker.SourceType)

	dockerSource := &docker.Source{}
	if err := dockerSource.Init(ctx, sourceName, jobID, sourceID, true, &conn, runtime.NumCPU()); err != nil {
		return err
	}
	_, err = e.sourceManager.Run(ctx, sourceName, dockerSource)
	return err
}
