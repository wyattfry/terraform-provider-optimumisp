// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"flag"
	"log"
	"os"
	"terraform-provider-optimumisp/optimumisp"
)

var (
// these will be set by the goreleaser configuration
// to appropriate values for the compiled binary.
// version string = "dev"

// goreleaser can pass other information to the main package, such as the specific commit
// https://goreleaser.com/cookbooks/using-main.version/
)

func main() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	client := optimumisp.Client{}
	username := os.Getenv("OPTIMUM_USERNAME")
	password := os.Getenv("OPTIMUM_PASSWORD")
	if username == "" || password == "" {
		log.Fatalf("Environment variables OPTIMUM_USERNAME and OPTIMUM_PASSWORD must be set")
	}
	client.ProcessLogin(os.Getenv("OPTIMUM_USERNAME"), os.Getenv("OPTIMUM_PASSWORD"))
	// fmt.Println("Getting Port Fwd Rules")
	// rules, _ := client.GetPortForwardingRules()
	// for idx, rule := range rules {
	// 	sld, err := json.MarshalIndent(rule, "", "  ")
	// 	if err != nil {
	// 		fmt.Println("error:", err)
	// 	}
	// 	fmt.Printf("Rule %d: %v\n\n", idx, string(sld))
	// }

	// name := rules[7].PortForwardingRule.PortForwardingRuleID.Name
	// idxs := []int{
	// 	rules[7].PortForwardingRule.PortForwardings[0].Index,
	// 	rules[7].PortForwardingRule.PortForwardings[1].Index,
	// }

	// client.DeletePortForwardingRule(name, idxs)

	// opts := providerserver.ServeOpts{
	// 	// TODO: Update this string with the published name of your provider.
	// 	// Also update the tfplugindocs generate command to either remove the
	// 	// -provider-name flag or set its value to the updated provider name.
	// 	Address: "registry.terraform.io/hashicorp/scaffolding",
	// 	Debug:   debug,
	// }

	// err := providerserver.Serve(context.Background(), provider.New(version), opts)

	// if err != nil {
	// 	log.Fatal(err.Error())
	// }
}
