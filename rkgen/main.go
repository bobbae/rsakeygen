package main

import (
	"flag"
	"os/user"
	"path/filepath"
	"rsakeygen"
	"strings"

	"github.com/qeof/q"
)

func init() {
	q.O = "stderr"
	q.P = ".*"
}

func main() {
	usr, _ := user.Current()
	dir := usr.HomeDir
	ppath := flag.String("p", "~/.ssh/example_id_rsa", "key files path")
	flag.Parse()
	path := *ppath
	if strings.HasPrefix(path, "~/") {
		path = filepath.Join(dir, path[2:])
	}
	q.Q(path)
	err := rsakeygen.RsaGenerateKeys(path)
	if err != nil {
		q.Q(err)
	}
}
