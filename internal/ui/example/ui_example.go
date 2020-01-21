package main

import (
	"fmt"

	"github.com/nori-io/norictl/internal/ui"
)

func main(){
	uiExample:=ui.NewUI()
	uiExample.GetSuccess("nori/session:0.0.0")
	uiExample.GetFailure("nori/session:0.0.1")
	uiExample.InstallSuccess("nori/session:0.0.2")
	uiExample.InstallFailure("nori/session:0.0.3")

	fmt.Println()

}
