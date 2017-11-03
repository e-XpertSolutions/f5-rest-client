package main

import (
	"e-xpert_solutions/f5-rest-client/cmd/tmsh-client/ltm"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/e-XpertSolutions/f5-rest-client/f5"
	"github.com/howeyc/gopass"

	"github.com/urfave/cli"
)

type config struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
}

func main() {
	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "json,j",
			Usage: "pretty print output using JSON format",
		},
		cli.BoolFlag{
			Name:  "insecure,k",
			Usage: "Allow connections to unsecure F5 iControl APIs",
		},
		cli.StringFlag{
			Name:  "proxy,x",
			Value: "",
			Usage: "use a proxy",
		},
		cli.StringFlag{
			Name:  "host,H",
			Value: "",
			Usage: "hostname or IP address of the F5 device",
		},
		cli.StringFlag{
			Name:  "port,P",
			Value: "443",
			Usage: "hostname or IP address of the F5 device",
		},
		cli.StringFlag{
			Name:  "user,u",
			Value: "admin",
			Usage: "Administrative user",
		},
		cli.StringFlag{
			Name:  "password,p",
			Value: "",
			Usage: "Administrative user",
		},
		cli.StringFlag{
			Name:  "config,c",
			Value: "",
			Usage: "CONFIGURATION_FILE Path to configuration file",
		},
	}

	var f5Client *f5.Client
	app.Before = func(c *cli.Context) error {
		var cfg config

		if configPath := c.GlobalString("config"); configPath != "" {
			tomlData, err := ioutil.ReadFile(configPath)
			if err != nil {
				fmt.Fprintln(os.Stderr, "cannot read configuration file: ", err)
				os.Exit(1)
			}
			if _, err := toml.Decode(string(tomlData), &cfg); err != nil {
				fmt.Fprintln(os.Stderr, "cannot decode configuration file: ", err)
				os.Exit(1)
			}
		}

		if host := c.String("host"); host != "" {
			cfg.Host = host
		}

		if user := c.String("user"); user != "" {
			cfg.User = user
		}

		if password := c.String("password"); password != "" {
			cfg.Password = password
		}
		if cfg.Password == "" {
			fmt.Fprint(os.Stdout, "password: ")
			pass, err := gopass.GetPasswd()
			if err != nil {
				// Handle gopass.ErrInterrupted or getch() read error
				fmt.Fprintln(os.Stderr, "cannot read Password input: ", err)
				os.Exit(1)
			}
			cfg.Password = string(pass)
		}
		var err error
		f5Client, err = f5.NewBasicClient("https://"+cfg.Host, cfg.User, cfg.Password)
		if err != nil {
			fmt.Println("err ?", err)
			return err
		}

		insecure := c.GlobalBool("insecure")
		if insecure {
			f5Client.DisableCertCheck()
		}
		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:     "node",
			Category: "Local Traffic Manager",
			Usage:    "Query or modify node objects",
			Aliases:  []string{"n"},
			Subcommands: []cli.Command{
				{
					Name: "create",
					Action: func(c *cli.Context) error {
						return ltm.CreateNode(c, f5Client)
					},
					Flags: []cli.Flag{
						cli.StringFlag{
							Name: "name,n",
						},
						cli.StringFlag{
							Name: "address,a",
						},
					},
				},
				{
					Name: "delete",
					Action: func(c *cli.Context) error {
						return ltm.DeleteNode(c, f5Client)
					},
				},
				{
					Name: "stats",
					Action: func(c *cli.Context) error {
						return ltm.StatsNode(c, f5Client)
					},
				},
				{
					Name: "enable",
					Action: func(c *cli.Context) error {
						return ltm.EnableNode(c, f5Client)
					},
				},
				{
					Name: "disable",
					Action: func(c *cli.Context) error {
						return ltm.DisableNode(c, f5Client)
					},
				},
				{
					Name: "offline",
					Action: func(c *cli.Context) error {
						return ltm.OfflineNode(c, f5Client)
					},
				},
				{
					Name: "list",
					Action: func(c *cli.Context) error {
						return ltm.ListNode(c, f5Client)
					},
				},
			},
		},
		{
			Name:     "virtual",
			Category: "Local Traffic Manager",
			Usage:    "Query or modify Virtual Server objects",
			Aliases:  []string{"n"},
			Subcommands: []cli.Command{
				{
					Name: "create",
					Action: func(c *cli.Context) error {
						return ltm.CreateVS(c, f5Client)
					},
					Flags: []cli.Flag{
						cli.StringFlag{
							Name: "name,n",
						},
						cli.StringFlag{
							Name: "address,a",
						},
					},
				},
				{
					Name: "delete",
					Action: func(c *cli.Context) error {
						return ltm.DeleteVS(c, f5Client)
					},
				},
				{
					Name: "stats",
					Action: func(c *cli.Context) error {
						return ltm.StatsVS(c, f5Client)
					},
				},
				{
					Name: "enable",
					Action: func(c *cli.Context) error {
						return ltm.EnableVS(c, f5Client)
					},
				},
				{
					Name: "disable",
					Action: func(c *cli.Context) error {
						return ltm.DisableVS(c, f5Client)
					},
				},
				{
					Name: "list",
					Action: func(c *cli.Context) error {
						return ltm.ListVS(c, f5Client)
					},
				},
			},
		},
	}

	app.Run(os.Args)
}
