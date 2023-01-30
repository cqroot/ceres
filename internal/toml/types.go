package toml

type CommonItem struct {
	Output    string   `toml:"output"`
	Variables []string `toml:"variables"`
}

type VariableItem struct {
	Message string   `toml:"message"`
	Type    string   `toml:"type"`
	Meta    []string `toml:"meta"`
}

type Rule struct {
	Key   string `toml:"key"`
	Value string `toml:"value"`
}

type ScriptItem struct {
	AfterScripts []string `toml:"after"`
}

type ConfigObject struct {
	CommonItem       CommonItem              `toml:"common"`
	Variable         map[string]VariableItem `toml:"variables"`
	IncludePathRules map[string]Rule         `toml:"include_path_rules"`
	ExcludePathRules map[string]Rule         `toml:"exclude_path_rules"`
	Scripts          ScriptItem              `toml:"scripts"`
}
