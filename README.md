# env2x

`env2x` outputs env variables in `json`, `yaml` or `env` formats.

### Usage

```
env2json [-p] VAR1 [VAR2] [VAR3{=value}] ...
env2json VAR1 [VAR2] [VAR3{=value}] ...
env2env  [-s] VAR1 [VAR2] [VAR3{=value}] ...
```

### Examples

`env2json -p USER HOME PATH=$HOME/bin:$PATH`

```json
{
  "HOME": "/Users/john",
  "PATH": "/Users/john/bin:/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin",
  "USER": "john"
}
```

`env2json USER HOME PATH=$HOME/bin:$PATH`

```json
{"HOME":"/Users/john","PATH":"/Users/john/bin:/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin","USER":"john"}
```

`env2yaml USER HOME PATH=$HOME/bin:$PATH`

```yaml
HOME: /Users/john
PATH: /Users/john/bin:/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin
USER: john
```

`env2env USER HOME PATH=$HOME/bin:$PATH`

```bash
HOME="/Users/john"
PATH="/Users/john/bin:/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin"
USER="john"
```

`env2env -s USER HOME PATH=$HOME/bin:$PATH`

```bash
HOME="/Users/john" PATH="/Users/john/bin:/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin" USER="john" 
```
