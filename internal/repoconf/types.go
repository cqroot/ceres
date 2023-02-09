package repoconf

type Common struct {
	Variables []string `toml:"variables"`
}

type VariableItem struct {
	Message string   `toml:"message"`
	Type    string   `toml:"type"`
	Meta    []string `toml:"meta"`
}

type ScriptItem struct {
	AfterScripts []string `toml:"after"`
}

type RepoConf struct {
	Common           Common                  `toml:"common"`
	Variable         map[string]VariableItem `toml:"variables"`
	IncludePathRules map[string][]string     `toml:"include_path_rules"`
	ExcludePathRules map[string][]string     `toml:"exclude_path_rules"`
	Scripts          ScriptItem              `toml:"scripts"`
}
