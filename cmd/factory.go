package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/jotadrilo/go-factory/pkg/config"
	"github.com/jotadrilo/go-factory/pkg/generator"
	"github.com/jotadrilo/go-factory/pkg/log"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	version = "dev"
	cfgFile string
	cfg     config.Config
	vip     *viper.Viper
)

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "go-factory",
		Version: version,
		Short:   "Generate factory helpers for your Golang structs",
		RunE: func(_ *cobra.Command, _ []string) error {
			log.Logger.Infof("Running go-factory version %s", version)

			if err := loadConfig(); err != nil {
				return err
			}

			return generator.NewFactories(&cfg).Generate()
		},
	}

	pfs := cmd.PersistentFlags()
	pfs.StringVarP(&cfgFile, "config", "c", "", "Configuration .gofactory.yaml file to use. It is auto-detected by default")
	pfs.StringP("name", "n", "", "Struct type name to generate the factory for in go:generate inline mode")
	pfs.BoolP("version", "v", false, "Shows the go-factory version")
	pfs.BoolP("help", "h", false, "Shows the go-factory help")

	if err := initViper(pfs); err != nil {
		panic(err)
	}

	return cmd
}

func newConfig() config.Config {
	projectDir, err := config.LocateProjectRootDir()
	if err != nil {
		log.Logger.Warnf("Cannot locate project root directory: %s", err.Error())
	} else {
		return config.NewConfig(version, projectDir)
	}

	projectDir, err = os.Getwd()
	if err != nil {
		log.Logger.Warnf("Cannot locate current directory: %s", err.Error())
	} else {
		return config.NewConfig(version, projectDir)
	}

	return config.NewConfig(version, projectDir)
}

func initViper(pfs *pflag.FlagSet) error {
	vip = viper.New()

	vip.SetConfigName(".gofactory")
	vip.SetConfigType("yaml")
	vip.SetEnvPrefix("GOFACTORY")
	vip.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	vip.AutomaticEnv()

	if err := vip.BindPFlags(pfs); err != nil {
		return err
	}

	return nil
}

const ConfigFile = ".gofactory.yaml"

func locateConfigFile() (string, error) {
	var configFile string

	if vip.IsSet("config") {
		configFile = vip.GetString("config")
	} else if cfg.ProjectDir != "" {
		configFile = filepath.Join(cfg.ProjectDir, ConfigFile)
	} else {
		configFile = ConfigFile
	}

	if _, err := os.Stat(configFile); err != nil {
		log.Logger.Errorf("Cannot locate the configuration (%s) file. You can use the --config flag, or place it in your project root directory.", ConfigFile)
		return "", fmt.Errorf("%s file not found", ConfigFile)
	}

	return configFile, nil
}

func loadConfig() error {
	if vip == nil {
		return fmt.Errorf("viper is not configured")
	}

	if vip.IsSet("name") {
		// Inline mode
		//
		// We can infer the file to modify from the GOFILE env set by
		// the go:generate binding

		log.Logger.Infof("Running in inline mode")

		if err := vip.BindEnv("file", "GOFILE"); err != nil {
			return err
		}
	} else {
		// Configuration mode
		//
		// We must locate the configuration file (I prefer do it myself)
		// and read it in.

		log.Logger.Infof("Running in configuration mode")

		configFile, err := locateConfigFile()
		if err != nil {
			return err
		}

		log.Logger.Infof("Using %s configuration file", configFile)

		vip.SetConfigFile(configFile)

		if err := vip.ReadInConfig(); err != nil {
			return err
		}
	}

	cfg = newConfig()

	if err := vip.Unmarshal(&cfg); err != nil {
		return fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return nil
}
