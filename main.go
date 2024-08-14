package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"slices"
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

var name string

func init() {
	flag.BoolVar(&pretty, "p", false, "pretty print json")
	flag.Parse()
}

func main() {
	defer recovery()

	name = path.Base(os.Args[0])
	args := flag.Args()

	if len(args) == 0 {
		printUsage()
	}

	env := make(map[string]string)

	for _, arg := range args {
		if strings.Contains(arg, "=") {
			v := strings.SplitN(arg, "=", 2)
			env[v[0]] = v[1]
		} else {
			env[arg] = os.Getenv(arg)
		}
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
		m := make(map[string]string)
		keys := make([]string, 0, len(env))
		for k, v := range env {
			m[k] = v
			keys = append(keys, k)
		}
		slices.Sort(keys)

		e := make([]string, 0, len(keys))
		for _, k := range keys {
			e = append(e, fmt.Sprintf("%s=%q", k, m[k]))
		}
		o = []byte(strings.Join(e, "\n") + "\n")
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
