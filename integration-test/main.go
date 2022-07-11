package main

import (
	"fmt"

	"github.com/joernott/nagiosplugin"
)

func main() {
	r, err := nagiosplugin.ParseRange("@10:20")
	if err != nil {
		panic(err)
	}

	fmt.Println(r)

	check := nagiosplugin.NewCheck()
	defer check.Finish()

	check.AddResult(nagiosplugin.OK, "everything looks shiny, cap'n")

	check.AddLongPluginOutput("Lorem Ipsum\nfoo,bar\n")

	warn, err := nagiosplugin.ParseRange("0")
	if err != nil {
		panic(err)
	}
	crit, err := nagiosplugin.ParseRange("253414")
	if err != nil {
		panic(err)
	}

	value, err := nagiosplugin.NewFloatPerfDatumValue(253404)
	if err != nil {
		panic(err)
	}

	check.AddPerfDatum("/home", "MB", value, warn, crit,
		nil, nil)
}
