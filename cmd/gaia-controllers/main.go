package main

import (
	goflag "flag"
	"fmt"
	"gaia.io/gaia/cmd/gaia-controllers/app"
	"gaia.io/gaia/pkg/utils"
	"github.com/spf13/pflag"
	"k8s.io/klog/v2"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	klog.InitFlags(nil)
	defer klog.Flush()

	ctx := utils.GracefulStopWithContext()
	command := app.NewGaiaControllerCmd(ctx)

	pflag.CommandLine.SetNormalizeFunc(utils.WordSepNormalizeFunc)
	pflag.CommandLine.AddGoFlagSet(goflag.CommandLine)

	if err := command.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
