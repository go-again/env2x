package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"slices"
	"strconv"
	"strings"
	"time"

	"encoding/json"
	"github.com/goccy/go-yaml"
)

var names = []string{
	"env2json",
	"env2yaml",
	"env2env",
}

//go:embed usage.txt
var usage string

var pretty bool
var single bool

var name string

func init() {
	flag.BoolVar(&pretty, "p", false, "pretty print json")
	flag.BoolVar(&single, "s", false, "single line env")
	flag.Parse()
}

func main() {
	defer recovery()

	name = path.Base(os.Args[0])
	args := flag.Args()

	if len(args) == 0 {
		printUsage()
	}

	env := make(map[string]any)

	for _, arg := range args {
		var k string
		var v any
		if strings.Contains(arg, "=") {
			s := strings.SplitN(arg, "=", 2)
			k = s[0]
			v = s[1]
		} else {
			k = arg
			v = os.Getenv(k)
		}

		switch strings.ToLower(v.(string)) {
		case "true", "false":
			v, _ = strconv.ParseBool(v.(string))
		default:
			x, err := strconv.ParseFloat(v.(string), 64)
			if err == nil {
				if strings.Contains(v.(string), ".") &&
					fmt.Sprintf("%f", v) != fmt.Sprintf("%v", x) {
					v = x
				} else {
					v = int64(x)
				}
			}
		}
		env[k] = v
	}

	var (
		o   []byte
		err error
	)
	switch name {
	case "env2json":
		if pretty {
			o, err = json.MarshalIndent(env, "", "  ")
		} else {
			o, err = json.Marshal(env)
		}
		o = append(o, '\n')
	case "env2yaml":
		o, err = yaml.Marshal(env)
	case "env2env":
		m := make(map[string]any)
		keys := make([]string, 0, len(env))
		for k, v := range env {
			m[k] = v
			keys = append(keys, k)
		}
		slices.Sort(keys)

		e := make([]string, 0, len(keys))
		for _, k := range keys {
			v := env[k]
			switch v.(type) {
			case float64:
				s := fmt.Sprintf("%s=%f", k, v)
				s = strings.TrimRight(s, "0")
				if s[len(s)-1] == '.' {
					s += "0"
				}
				e = append(e, s)
			case int64:
				e = append(e, fmt.Sprintf("%s=%d", k, v))
			case bool:
				e = append(e, fmt.Sprintf("%s=%v", k, v))
			default:
				e = append(e, fmt.Sprintf("%s=%q", k, v))
			}
		}
		sep := "\n"
		if single {
			sep = " "
		}
		o = []byte(strings.Join(e, sep) + sep)
	default:
		printUsage()
	}
	if err != nil {
		log.Fatal(err)
	}

	_, err = os.Stdout.Write(o)
	if err != nil {
		log.Fatal(err)
	}
}

func printUsage() {
	fmt.Printf(usage, toAnySlice(names)...)
	os.Exit(0)
}

func toAnySlice[T any](collection []T) []any {
	result := make([]any, len(collection))
	for i := range collection {
		result[i] = collection[i]
	}
	return result
}

func recovery() {
	if r := recover(); r != nil {
		pwd, _ := os.Getwd()
		file := path.Join(pwd, name+"-"+time.Now().Format("2006-01-02_15-04-05")+".log")
		os.WriteFile(file, []byte(fmt.Sprintf("%v", r)), 0644)
		fmt.Println("Unhandled error saved to", file)
		os.Exit(1)
	}
}
