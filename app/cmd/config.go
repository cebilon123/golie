package cmd

import cfg2 "github.com/cebilon123/golie/cfg"

//SetPath sets path to directory which should be watched and synced
func SetPath(path string) {
	if len(path) == 0 {
		return
	}

	cfg := cfg2.Configuration{Path: path}
	cfg.Serialize()
}

