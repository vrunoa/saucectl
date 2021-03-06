package run

import (
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"

	"github.com/saucelabs/saucectl/cli/command"
	"github.com/saucelabs/saucectl/cli/config"
	"github.com/saucelabs/saucectl/cli/runner"
	"github.com/spf13/cobra"
)

var (
	runUse     = "run ./.sauce/config.yaml"
	runShort   = "Run a test on Sauce Labs"
	runLong    = `Some long description`
	runExample = "saucectl run ./.sauce/config.yaml"

	defaultLogFir  = "<cwd>/logs"
	defaultTimeout = 60
	defaultRegion = "us-west-1"

	cfgFilePath string
	cfgLogDir   string
	testTimeout int
	region      string
)

// Command creates the `run` command
func Command(cli *command.SauceCtlCli) *cobra.Command {
	cmd := &cobra.Command{
		Use:     runUse,
		Short:   runShort,
		Long:    runLong,
		Example: runExample,
		Run: func(cmd *cobra.Command, args []string) {
			log.Info().Msg("Start Run Command")
			exitCode, err := Run(cmd, cli, args)
			if err != nil {
				log.Err(err).Msg("failed to execute run command")
			}
			os.Exit(exitCode)
		},
	}

	defaultCfgPath := filepath.Join(".sauce", "config.yml")
	cmd.Flags().StringVarP(&cfgFilePath, "config", "c", defaultCfgPath, "config file (e.g. ./.sauce/config.yaml")
	cmd.Flags().StringVarP(&cfgLogDir, "logDir", "l", defaultLogFir, "log path")
	cmd.Flags().IntVarP(&testTimeout, "timeout", "t", 0, "test timeout in seconds (default: 60sec)")
	cmd.Flags().StringVarP(&region, "region", "r", "", "The sauce labs region. (default: us-west-1)")
	return cmd
}

// Run runs the command
func Run(cmd *cobra.Command, cli *command.SauceCtlCli, args []string) (int, error) {
	// Todo(Christian) write argument parser/validator
	if cfgLogDir == defaultLogFir {
		pwd, _ := os.Getwd()
		cfgLogDir = filepath.Join(pwd, "logs")
	}

	log.Info().Str("config", cfgFilePath).Msg("Reading config file")
	configObject, err := config.NewJobConfiguration(cfgFilePath)
	if err != nil {
		return 1, err
	}

	if testTimeout != 0 {
		configObject.Timeout = testTimeout
	}
	if configObject.Timeout == 0 {
		configObject.Timeout = defaultTimeout
	}

	if configObject.Sauce.Region == "" {
		configObject.Sauce.Region = defaultRegion
	}

	if region != "" {
		configObject.Sauce.Region = region
	}

	tr, err := runner.New(configObject, cli)
	if err != nil {
		return 1, err
	}

	log.Info().Msg("Setting up test environment")
	if err := tr.Setup(); err != nil {
		return 1, err
	}

	log.Info().Msg("Starting tests")
	exitCode, err := tr.Run()
	if err != nil {
		return 1, err
	}

	log.Info().Msg("Tearing down environment")
	if err != tr.Teardown(cfgLogDir) {
		return 1, err
	}

	log.Info().
		Int("ExitCode", exitCode).
		Msg("Command Finished")

	return exitCode, nil
}
