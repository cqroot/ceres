<div align="center">
  <h1>Ceres</h1>

  <p><i>Manage your project templates.</i></p>

  <p>
    <a href="https://github.com/cqroot/ceres/actions">
      <img src="https://github.com/cqroot/ceres/workflows/test/badge.svg" alt="Action Status" />
    </a>
    <a href="https://github.com/cqroot/ceres/blob/main/LICENSE">
      <img src="https://img.shields.io/github/license/cqroot/ceres" />
    </a>
    <a href="https://github.com/cqroot/ceres/issues">
      <img src="https://img.shields.io/github/issues/cqroot/ceres" />
    </a>
  </p>
</div>

## Create a template repository

The ceres template repository must contain a `ceres.toml` file and a `template` folder.

### ceres.toml

`ceres.toml` contains the following sections:

1. common
2. variables
3. include_path_rules
4. exclude_path_rules
5. scripts

#### common

```toml
variables = ["use_config", "use_logger"]
```

| key         | description                                                                                                                         |
| ----------- | ----------------------------------------------------------------------------------------------------------------------------------- |
| `variables` | list of variables. ceres will ask the user for data in the order of the list. Variables must be generated in the variables section. |

#### variables

Ceres will ask the user one by one in the order of the `common.variables` list.
The result is stored in `map[string]string` for data passed into the template.

```
[variables]
[variables.module_name]
message = "Your module name:"
type = "input"
meta = ["github.com/cqroot/ceres"]
```

The available values for type are:

1. `input`: meta only accepts a string, which is the default value of input.
2. `toggle`: meta receives multiple strings, which are options provided to the user (Usually used for yes or no choices).
3. `choose`: meta receives multiple strings, which are options provided to the user.

#### include_path_rules

Generate the file only if the condition is met.

The filename can be a directory or a file, but do not end with '/'.

```toml
"src/config" = { key = "use_config", value = "Yes" }
```

#### exclude_path_rules

Same as `include_path_rules`.

#### scripts

##### all

```toml
after = ["scripts/init.sh"]
```

| key     | description                             |
| ------- | --------------------------------------- |
| `after` | Script to execute after all generation. |
