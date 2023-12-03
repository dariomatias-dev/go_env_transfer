package main

import "github.com/dariomatias-dev/go_env_transfer/versions"

func main() {
	targetFilePath := ".env"

	versions.Version3(targetFilePath)
}
