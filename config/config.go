package config

import "sync"

// common config key names
const (
	HashAlgName string = "HashAlgName"
)

// GlobalConfig project's global config
var GlobalConfig *sync.Map = new(sync.Map)
