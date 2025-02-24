package config

import (
	"embed"
	"io/fs"
)

//go:embed frontend/dist/*
var embeddedFiles embed.FS

func ConfigureFrontend() (fs.FS, error) {
	// Create a filesystem from the embedded files, stripping the "frontend/dist" prefix.
	distFS, err := fs.Sub(embeddedFiles, "frontend/dist")

	return distFS, err
}
