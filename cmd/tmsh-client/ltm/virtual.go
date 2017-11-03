package ltm

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"

	"e-xpert_solutions/go-jsonfilter/jsonfilter"
	"e-xpert_solutions/go-pretty/pretty"

	"github.com/e-XpertSolutions/f5-rest-client/f5"
	"github.com/e-XpertSolutions/f5-rest-client/f5/ltm"

	"github.com/urfave/cli"
)

func CreateVS(c *cli.Context, f5Client *f5.Client) error {
	if c.NumFlags() < 2 {
		// TODO return a true error message
		return nil
	}

	if c.String("address") == "" {
		// TODO return a true error message
		return nil
	}

	ipAddr := net.ParseIP(c.String("address"))

	if ipAddr.String() == "<nil>" {
		// TODO return a true error message
		return nil
	}

	ltmClient := ltm.New(f5Client)

	nodeConfig := ltm.Node{
		Name:    c.String("name"),
		Address: c.String("address"),
	}
	if err := ltmClient.Node().Create(nodeConfig); err != nil {
		log.Fatal(err)
	}

	return nil
}

func DeleteVS(c *cli.Context, f5Client *f5.Client) error {
	if c.NArg() != 1 {
		// TODO return a true error message
		return nil
	}

	ltmClient := ltm.New(f5Client)

	if err := ltmClient.Virtual().Delete(c.Args().First()); err != nil {
		log.Fatal(err)
	}

	return nil
}

func EnableVS(c *cli.Context, f5Client *f5.Client) error {
	if c.NArg() == 1 {
		ltmClient := ltm.New(f5Client)

		err := ltmClient.Virtual().Enable(c.Args().First())
		if err != nil {
			fmt.Println("err :", err)
			return err
		}
		return nil
	}
	return nil
}

func DisableVS(c *cli.Context, f5Client *f5.Client) error {
	if c.NArg() == 1 {
		ltmClient := ltm.New(f5Client)

		err := ltmClient.Virtual().Disable(c.Args().First())
		if err != nil {
			fmt.Println("err :", err)
			return err
		}
		return nil
	}
	return nil
}

func StatsVS(c *cli.Context, f5Client *f5.Client) error {

	ltmClient := ltm.New(f5Client)

	if c.NArg() == 1 {
		virtualStats, err := ltmClient.Virtual().ShowStats(c.Args().First())
		if err != nil {
			fmt.Println("err :", err)
			return err
		}

		jsonFlag := c.GlobalBool("json")

		if jsonFlag {
			bs, err := json.MarshalIndent(virtualStats, "", "   ")
			if err != nil {
				return err
			}
			fmt.Println(string(bs))
			return nil
		}
		for _, ns := range virtualStats.Entries {
			_, err = pretty.Print(ns.NestedVirtualStats)
			if err != nil {
				log.Print(err)
			}
		}
		return nil
	} else {
		virtualStats, err := ltmClient.VirtualStats().All()
		if err != nil {
			fmt.Println("err :", err)
			return err
		}

		jsonFlag := c.GlobalBool("json")

		if jsonFlag {
			bs, err := json.MarshalIndent(virtualStats, "", "   ")
			if err != nil {
				return err
			}
			fmt.Println(string(bs))
			return nil
		}
		for _, ns := range virtualStats.Entries {
			_, err = pretty.Print(ns.NestedVirtualStats)
			if err != nil {
				log.Print(err)
			}
		}
	}
	return nil
}

func ListVS(c *cli.Context, f5Client *f5.Client) error {

	ltmClient := ltm.New(f5Client)

	var virtuals *ltm.VirtualServerList

	if c.NArg() >= 1 && c.Args()[0] != "all-properties" && !arrayContains(getAllFields(ltm.VirtualServer{}), c.Args()[0]) {

		virtual, err := ltmClient.Virtual().Get(c.Args().First())

		if err != nil {
			fmt.Println("err :", err)
			return err
		}

		jsonFlag := c.GlobalBool("json")
		if jsonFlag {
			opts := jsonfilter.Options{PrettyPrint: true, TagName: "pretty", FormatFunc: formatName}
			if c.NArg() > 1 {
				if c.Args()[1] == "all-properties" {
					opts.Expanded = true
				} else {
					opts.Whitelist = c.Args()[1:]
				}
			}
			bs, err := jsonfilter.MarshalJSON(virtual, opts)
			if err != nil {
				return err
			}
			fmt.Println(string(bs))
			return nil
		}

		opts := pretty.DefaultPrinterOptions
		if c.NArg() > 1 {
			if c.Args()[1] == "all-properties" {
				opts.Expanded = true
			} else {
				opts.Whitelist = c.Args()[1:]
			}
		}

		printer := pretty.NewPrinter(os.Stdout, opts)

		_, err = printer.Print(virtual)
		if err != nil {
			log.Print(err)
		}
		return nil
	}

	virtuals, err := ltmClient.Virtual().ListAll()
	if err != nil {
		fmt.Println("err :", err)
		return err
	}

	jsonFlag := c.GlobalBool("json")
	if jsonFlag {
		opts := jsonfilter.Options{PrettyPrint: true, TagName: "pretty", FormatFunc: formatName}
		if c.NArg() >= 1 {
			if c.Args()[0] == "all-properties" {
				opts.Expanded = true
			} else {
				opts.Whitelist = c.Args()[0:]
			}
		}
		bs, err := jsonfilter.MarshalJSON(virtuals, opts)
		if err != nil {
			return err
		}
		fmt.Println(string(bs))
		return nil
	}

	opts := pretty.DefaultPrinterOptions
	if c.NArg() >= 1 {
		if c.Args()[0] == "all-properties" {
			opts.Expanded = true
		} else {
			opts.Whitelist = c.Args()[0:]
		}
	}

	printer := pretty.NewPrinter(os.Stdout, opts)

	_, err = printer.Print(virtuals.Items)
	if err != nil {
		log.Print(err)
	}
	//	fmt.Printf("%+v\n", nodes)
	return nil
}
