# env2x

`env2x` outputs env variables in `json`, `yaml`, `env` or string formats.

### Usage

```
env2json [-p] VAR1 [VAR2] [VAR3{=value}] ...
env2json VAR1 [VAR2] [VAR3{=value}] ...
env2env  [-s] [-e] VAR1 [VAR2] [VAR3{=value}] ...
env2file VAR [file] [mode]
```

### Examples

`env2json -p USER HOME PATH=$HOME/bin:$PATH`

Output:
```json
{
  "HOME": "/Users/john",
  "PATH": "/Users/john/bin:/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin",
  "USER": "john"
}
```
---
`env2json USER HOME PATH=$HOME/bin:$PATH`

Output:
```json
{"HOME":"/Users/john","PATH":"/Users/john/bin:/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin","USER":"john"}
```
---
`env2yaml USER HOME PATH=$HOME/bin:$PATH`

Output:
```yaml
HOME: /Users/john
PATH: /Users/john/bin:/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin
USER: john
```
---
`env2env USER HOME PATH=$HOME/bin:$PATH`

Output:
```bash
HOME="/Users/john"
PATH="/Users/john/bin:/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin"
USER="john"
```
---
`env2env -e USER HOME PATH=$HOME/bin:$PATH`

Output:
```bash
export HOME="/Users/john"
export PATH="/Users/john/bin:/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin"
export USER="john"
```
---
`env2env -s USER HOME PATH=$HOME/bin:$PATH`

Output:
```bash
HOME="/Users/john" PATH="/Users/john/bin:/usr/local/bin:/usr/bin:/bin:/usr/sbin:/sbin" USER="john" 
```
---
`env2file TOKEN token.txt 600`

Writes TOKEN value to `token.txt` and sets file mode `0600`