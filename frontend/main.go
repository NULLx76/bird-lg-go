package main

import (
	"flag"
	"github.com/caarlos0/env/v6"
	"log"
	"strings"
)

type settingType struct {
	servers         []string `env:"BIRDLG_SERVERS" envSeparator:","`
	proxyPort       int      `env:"BIRDLG_PROXY_PORT"`
	whoisServer     string   `env:"BIRDLG_WHOIS"`
	listen          string   `env:"BIRDLG_LISTEN"`
	dnsInterface    string   `env:"BIRDLG_DNS_INTERFACE"`
	netSpecificMode string   `env:"BIRDLG_NET_SPECIFIC_MODE"`
	titleBrand      string   `env:"BIRDLG_TITLE_BRAND"`
	navBarBrand     string   `env:"BIRDLG_NAVBAR_BRAND"`
}

var setting settingType

func main() {
	var settingDefault = settingType{
		servers:      []string{""},
		proxyPort:    8000,
		whoisServer:  "whois.verisign-grs.com",
		listen:       ":5000",
		dnsInterface: "asn.cymru.com",
		titleBrand:   "Bird-lg Go",
		navBarBrand:  "Bird-lg Go",
	}

	if err := env.Parse(&settingDefault); err != nil {
		log.Fatalf("%+v\n", err)
	}

	serversPtr := flag.String("servers", strings.Join(settingDefault.servers, ","), "server name prefixes, separated by comma")
	proxyPortPtr := flag.Int("proxy-port", settingDefault.proxyPort, "port bird-lgproxy is running on")
	whoisPtr := flag.String("whois", settingDefault.whoisServer, "whois server for queries")
	listenPtr := flag.String("listen", settingDefault.listen, "address bird-lg is listening on")
	dnsInterfacePtr := flag.String("dns-interface", settingDefault.dnsInterface, "dns zone to query ASN information")
	netSpecificModePtr := flag.String("net-specific-mode", settingDefault.netSpecificMode, "network specific operation mode, [(none)|dn42]")
	titleBrandPtr := flag.String("title-brand", settingDefault.titleBrand, "prefix of page titles in browser tabs")
	navBarBrandPtr := flag.String("navbar-brand", settingDefault.navBarBrand, "brand to show in the navigation bar")
	flag.Parse()

	if *serversPtr == "" {
		panic("no server set")
	}

	setting = settingType{
		strings.Split(*serversPtr, ","),
		*proxyPortPtr,
		*whoisPtr,
		*listenPtr,
		*dnsInterfacePtr,
		strings.ToLower(*netSpecificModePtr),
		*titleBrandPtr,
		*navBarBrandPtr,
	}

	webServerStart()
}
