package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	logger "github.com/Sirupsen/logrus"
	config "github.com/mgreau/licenses-report/config"
	report "github.com/mgreau/licenses-report/report"
	types "github.com/mgreau/licenses-report/types"
	cli "gopkg.in/urfave/cli.v1"
)

var (
	// Version is the version of the software
	Version string
	// BuildStmp is the build date
	BuildStmp string
	// GitHash is the git build hash
	GitHash string

	logLevel = "warning"
	// path to config file
	cfgFile string

	// Params to configure the report
	params *types.Params = &types.Params{
		Format: "json",
		Name:   "dependencies-licenses-report",
	}

	pathFlag    string
	projectFlag string
	outputFlag  string = "."
)

// preload initializes any global options and configuration
// before the main or sub commands are run.
func preload(c *cli.Context) (err error) {
	if c.GlobalBool("debug") {
		logger.SetLevel(logger.DebugLevel)
	}

	return nil
}

func main() {

	// set timezone as UTC for bson/json time marshalling
	time.Local = time.UTC

	// new app
	app := cli.NewApp()
	app.Name = "licences-report"
	app.Usage = "Generate Licenses Report of 3rd party dependencies"

	timeStmp, err := strconv.Atoi(BuildStmp)
	if err != nil {
		timeStmp = 0
	}
	app.Version = Version + ", build on " + time.Unix(int64(timeStmp), 0).String() + ", git hash " + GitHash
	app.Author = "@mgreau"
	app.Email = "maxime.greau@elastic.co"
	app.Before = preload

	// command line flags
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Value: logLevel,
			Name:  "logl",
			Usage: "Set the output log level (debug, info, warning, error)",
		},
		cli.StringFlag{
			Name:        "config",
			Usage:       "Set the path to the config file (/my-path/conf.yaml)",
			Destination: &cfgFile,
		},
		cli.StringFlag{
			Name:        "project",
			Usage:       "Set the project name",
			Destination: &projectFlag,
		},
		cli.StringFlag{
			Name:        "path",
			Destination: &pathFlag,
			Usage:       "Set the path to look for dependencies",
		},
		cli.StringFlag{
			Name:        "output",
			Destination: &outputFlag,
			Usage:       "Set the output path to generate the report",
		},
	}

	app.Commands = []cli.Command{

		{
			Name:    "generate",
			Aliases: []string{"g"},
			Usage:   "Generate the report ",
			Action: func(c *cli.Context) error {

				if cfgFile != "" {
					// loading YAML file
					byt, err := ioutil.ReadFile(cfgFile)
					if err != nil {
						return fmt.Errorf("error reading configuration: %s", err)
					}
					params, err = config.Parse(byt)
					if err != nil {
						return fmt.Errorf("error reading configuration: %s", err)
					}

				} else {

					// override project name
					if projectFlag != "" {
						params.Project = projectFlag
					}
					// override path
					if pathFlag != "" {
						params.Path = pathFlag
					}
					params.Output = outputFlag
				}

				report.GenerateReport(params)
				return nil
			},
		},
	}

	// run the appcd
	err = app.Run(os.Args)
	if err != nil {
		logger.Fatalf("Run error %q\n", err)
	}
}
