
package proxmoxcli

import (
	"fmt"
	"net/http"
	"net/url"
	"os"

	"golang.org/x/net/http/httpproxy"

	"yunion.io/x/pkg/util/shellutils"
	"yunion.io/x/structarg"

	"yunion.io/x/cloudmux/pkg/cloudprovider"
	"yunion.io/x/cloudmux/pkg/multicloud/proxmox"
	_ "yunion.io/x/cloudmux/pkg/multicloud/proxmox/shell"
)

type BaseOptions struct {
	Debug      bool   `help:"debug mode"`
	Username   string `help:"Username" default:"$PROXMOX_USERNAME" metavar:"PROXMOX_USERNAME"`
	Password   string `help:"Password" default:"$PROXMOX_PASSWORD" metavar:"PROXMOX_PASSWORD"`
	Host       string `help:"Host" default:"$PROXMOX_HOST" metavar:"PROXMOX_HOST"`
	Port       int    `help:"Port" default:"$PROXMOX_PORT|8006" metavar:"PROXMOX_PORT"`
	SUBCOMMAND string `help:"proxmoxcli subcommand" subcommand:"true"`
}

func getSubcommandParser() (*structarg.ArgumentParser, error) {
	parse, e := structarg.NewArgumentParserWithHelp(&BaseOptions{},
		"proxmoxcli",
		"Command-line interface to proxmoxc API.",
		`See "proxmoxcli COMMAND --help" for help on a specific command.`)

	if e != nil {
		return nil, e
	}

	subcmd := parse.GetSubcommand()
	if subcmd == nil {
		return nil, fmt.Errorf("No subcommand argument.")
	}
	for _, v := range shellutils.CommandTable {
		_, e := subcmd.AddSubParserWithHelp(v.Options, v.Command, v.Desc, v.Callback)
		if e != nil {
			return nil, e
		}
	}
	return parse, nil
}

func showErrorAndExit(e error) {
	fmt.Fprintf(os.Stderr, "%s", e)
	fmt.Fprintln(os.Stderr)
	os.Exit(1)
}

func newClient(options *BaseOptions) (*proxmox.SRegion, error) {
	if len(options.Host) == 0 {
		return nil, fmt.Errorf("Missing host")
	}

	if len(options.Username) == 0 {
		return nil, fmt.Errorf("Missing username")
	}

	if len(options.Password) == 0 {
		return nil, fmt.Errorf("Missing password")
	}

	cfg := &httpproxy.Config{
		HTTPProxy:  os.Getenv("HTTP_PROXY"),
		HTTPSProxy: os.Getenv("HTTPS_PROXY"),
		NoProxy:    os.Getenv("NO_PROXY"),
	}
	cfgProxyFunc := cfg.ProxyFunc()
	proxyFunc := func(req *http.Request) (*url.URL, error) {
		return cfgProxyFunc(req.URL)
	}

	cli, err := proxmox.NewProxmoxClient(
		proxmox.NewProxmoxClientConfig(
			options.Username,
			options.Password,
			options.Host,
			options.Port,
		).Debug(options.Debug).
			CloudproviderConfig(
				cloudprovider.ProviderConfig{
					ProxyFunc: proxyFunc,
				},
			),
	)
	if err != nil {
		return nil, err
	}

	return cli.GetRegion(), nil
}

func Main() {
	parser, e := getSubcommandParser()
	if e != nil {
		showErrorAndExit(e)
	}

	e = parser.ParseArgs(os.Args[1:], false)
	options := parser.Options().(*BaseOptions)

	if parser.IsHelpSet() {
		fmt.Print(parser.HelpString())
		return
	}
	subcmd := parser.GetSubcommand()
	subparser := subcmd.GetSubParser()
	if e != nil || subparser == nil {
		if subparser != nil {
			fmt.Print(subparser.Usage())
		} else {
			fmt.Print(parser.Usage())
		}
		showErrorAndExit(e)
		return
	}
	suboptions := subparser.Options()
	if subparser.IsHelpSet() {
		fmt.Print(subparser.HelpString())
		return
	}
	var region *proxmox.SRegion
	region, e = newClient(options)
	if e != nil {
		showErrorAndExit(e)
	}
	e = subcmd.Invoke(region, suboptions)
	if e != nil {
		showErrorAndExit(e)
	}
}
