package main

import (
	"os"

	godist "github.com/paketo-buildpacks/go-dist"
	"github.com/paketo-buildpacks/packit"
	"github.com/paketo-buildpacks/packit/cargo"
	"github.com/paketo-buildpacks/packit/chronos"
	"github.com/paketo-buildpacks/packit/draft"
	"github.com/paketo-buildpacks/packit/postal"
	"github.com/paketo-buildpacks/packit/scribe"
)

func main() {
	logEmitter := godist.NewGoLogger(scribe.NewEmitter(os.Stdout))
	buildpackYAMLParser := godist.NewBuildpackYAMLParser()
	entryResolver := draft.NewPlanner()
	dependencyManager := postal.NewService(cargo.NewTransport())

	packit.Run(
		godist.Detect(buildpackYAMLParser),
		godist.Build(entryResolver, dependencyManager, chronos.DefaultClock, logEmitter),
	)
}
