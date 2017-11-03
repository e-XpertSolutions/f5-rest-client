package ltm

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"

	"e-xpert_solutions/go-jsonfilter/jsonfilter"
	"e-xpert_solutions/go-pretty/pretty"

	"github.com/e-XpertSolutions/f5-rest-client/f5"
	"github.com/e-XpertSolutions/f5-rest-client/f5/ltm"

	"github.com/urfave/cli"
)

func SliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

func CreateNode(c *cli.Context, f5Client *f5.Client) error {
	if c.NArg() < 3 {
		// TODO return a true error message
		fmt.Println("invalid number of args")
		return nil
	}

	name := c.Args()[0]

	if !arrayContains(c.Args(), "address") {
		// TODO return a true error message
		fmt.Println("missing arg address")
		return nil
	}

	address := c.Args()[SliceIndex(c.NArg(), func(i int) bool { return c.Args()[i] == "address" })+1]

	nodeConfig := ltm.Node{
		Name:    name,
		Address: address,
	}

	idx := SliceIndex(c.NArg(), func(i int) bool { return c.Args()[i] == "monitor" })
	if idx > 0 {
		nodeConfig.Monitor = c.Args()[idx+1]
	}

	connectionLimit, err := strconv.Atoi((c.Args()[SliceIndex(c.NArg(), func(i int) bool { return c.Args()[i] == "connection-limit" })+1]))
	if err == nil {
		nodeConfig.ConnectionLimit = connectionLimit
	}

	dynamicRatio, err := strconv.Atoi((c.Args()[SliceIndex(c.NArg(), func(i int) bool { return c.Args()[i] == "dynamic-ratio" })+1]))
	if err == nil {
		nodeConfig.DynamicRatio = dynamicRatio
	}

	ratio, err := strconv.Atoi((c.Args()[SliceIndex(c.NArg(), func(i int) bool { return c.Args()[i] == "ratio" })+1]))
	if err == nil {
		nodeConfig.Ratio = ratio
	}

	idx = SliceIndex(c.NArg(), func(i int) bool { return c.Args()[i] == "logging" })
	if idx > 0 {
		nodeConfig.Logging = c.Args()[idx+1]
	}

	idx = SliceIndex(c.NArg(), func(i int) bool { return c.Args()[i] == "rate-limit" })
	if idx > 0 {
		nodeConfig.RateLimit = c.Args()[idx+1]
	}

	idx = SliceIndex(c.NArg(), func(i int) bool { return c.Args()[i] == "session" })
	if idx > 0 {
		nodeConfig.Session = c.Args()[idx+1]
	}

	idx = SliceIndex(c.NArg(), func(i int) bool { return c.Args()[i] == "state" })
	if idx > 0 {
		nodeConfig.State = c.Args()[idx+1]
	}

	ltmClient := ltm.New(f5Client)

	if err := ltmClient.Node().Create(nodeConfig); err != nil {
		log.Fatal(err)
	}

	return nil
}

func DeleteNode(c *cli.Context, f5Client *f5.Client) error {
	if c.NArg() != 1 {
		// TODO return a true error message
		return nil
	}

	ltmClient := ltm.New(f5Client)

	if err := ltmClient.Node().Delete(c.Args().First()); err != nil {
		log.Fatal(err)
	}

	return nil
}

func EnableNode(c *cli.Context, f5Client *f5.Client) error {
	if c.NArg() == 1 {
		ltmClient := ltm.New(f5Client)

		err := ltmClient.Node().Enable(c.Args().First())
		if err != nil {
			fmt.Println("err :", err)
			return err
		}
		return nil
	}
	return nil
}

func DisableNode(c *cli.Context, f5Client *f5.Client) error {
	if c.NArg() == 1 {
		ltmClient := ltm.New(f5Client)

		err := ltmClient.Node().Disable(c.Args().First())
		if err != nil {
			fmt.Println("err :", err)
			return err
		}
		return nil
	}
	return nil
}

func OfflineNode(c *cli.Context, f5Client *f5.Client) error {
	if c.NArg() == 1 {
		ltmClient := ltm.New(f5Client)

		err := ltmClient.Node().ForceOffline(c.Args().First())
		if err != nil {
			fmt.Println("err :", err)
			return err
		}
		return nil
	}
	return nil
}

func StatsNode(c *cli.Context, f5Client *f5.Client) error {

	ltmClient := ltm.New(f5Client)

	if c.NArg() == 1 {
		nodeStats, err := ltmClient.Node().ShowStats(c.Args().First())
		if err != nil {
			fmt.Println("err :", err)
			return err
		}

		jsonFlag := c.GlobalBool("json")

		if jsonFlag {
			bs, err := json.MarshalIndent(nodeStats, "", "   ")
			if err != nil {
				return err
			}
			fmt.Println(string(bs))
			return nil
		}
		for _, ns := range nodeStats.Entries {
			_, err = pretty.Print(ns.NestedNodeStats)
			if err != nil {
				log.Print(err)
			}
		}
		return nil
	} else {
		nodeStats, err := ltmClient.NodeStats().All()
		if err != nil {
			fmt.Println("err :", err)
			return err
		}

		jsonFlag := c.GlobalBool("json")

		if jsonFlag {
			bs, err := json.MarshalIndent(nodeStats, "", "   ")
			if err != nil {
				return err
			}
			fmt.Println(string(bs))
			return nil
		}
		for _, ns := range nodeStats.Entries {
			_, err = pretty.Print(ns.NestedNodeStats)
			if err != nil {
				log.Print(err)
			}
		}
	}
	return nil
}

func ListNode(c *cli.Context, f5Client *f5.Client) error {

	ltmClient := ltm.New(f5Client)

	var nodes *ltm.NodeList

	if c.NArg() >= 1 && c.Args()[0] != "all-properties" && !arrayContains(getAllFields(ltm.Node{}), c.Args()[0]) {

		node, err := ltmClient.Node().Get(c.Args().First())

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
			bs, err := jsonfilter.MarshalJSON(node, opts)
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

		_, err = printer.Print(node)
		if err != nil {
			log.Print(err)
		}
		return nil
	}

	nodes, err := ltmClient.Node().ListAll()
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
		bs, err := jsonfilter.MarshalJSON(nodes, opts)
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

	_, err = printer.Print(nodes.Items)
	if err != nil {
		log.Print(err)
	}
	//	fmt.Printf("%+v\n", nodes)
	return nil
}

func getAllFields(v interface{}) []string {
	t := reflect.TypeOf(v)
	if t.Kind() != reflect.Struct {
		return []string{}
	}
	var fields []string
	for i := 0; i < t.NumField(); i++ {
		name := t.Field(i).Name
		if name[:1] == strings.ToUpper(name[:1]) {
			fields = append(fields, string(pretty.FormatFieldName(name)))
		}
	}
	return fields
}

func arrayContains(ary []string, s string) bool {
	for _, elmt := range ary {
		if elmt == s {
			return true
		}
	}
	return false
}

func formatName(fieldName, jsonTagName string) string {
	if jsonTagName != "" {
		return string(pretty.FormatFieldName(jsonTagName))
	}
	return string(pretty.FormatFieldName(fieldName))
}
