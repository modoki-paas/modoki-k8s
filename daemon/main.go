package main

import (
	"flag"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/modoki-paas/modoki-k8s/daemon/store"
)

type commandArg struct {
	DSN  string
	Help bool
}

func (arg *commandArg) Init() {
	flag.StringVar(&arg.DSN, "db", "", "database source name")
	flag.BoolVar(&arg.Help, "help", false, "show usage")

	flag.Parse()

	if arg.Help {
		flag.Usage()

		os.Exit(1)
	}
}

func main() {
	carg := &commandArg{}
	carg.Init()

	d, err := sqlx.Open("mysql", carg.DSN)

	if err != nil {
		panic(err)
	}

	store.NewDB(d)
}
