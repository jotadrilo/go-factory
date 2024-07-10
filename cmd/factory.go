package cmd

import (
	"fmt"
	"github.com/jotadrilo/go-factory/pkg/config"
	"github.com/jotadrilo/go-factory/pkg/generator"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"strings"
)

var (
	version string
	cfgFile string
	cfg     config.Config
	vip     *viper.Viper
)

func init() {
	cobra.OnInitialize(func() {
		if err := loadConfig(); err != nil {
			panic(err)
		}
	})
}

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "go-factory",
		Version: version,
		Short:   "Generate factory helpers for your Golang structs",
		RunE: func(cmd *cobra.Command, args []string) error {
			return generator.NewFactories(&cfg).Generate()
		},
	}

	pfs := cmd.PersistentFlags()
	pfs.StringVar(&cfgFile, "config", "", "Configuration file to use")
	pfs.StringP("name", "n", "", "Struct type name")

	if err := initViper(pfs); err != nil {
		panic(err)
	}

	return cmd
}

func initViper(pfs *pflag.FlagSet) error {
	vip = viper.New()

	vip.SetConfigName(".gofactory")
	vip.SetConfigType("yaml")
	vip.SetEnvPrefix("GOFACTORY")
	vip.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	vip.AutomaticEnv()

	if err := vip.BindEnv("file", "GOFILE"); err != nil {
		return err
	}

	if err := vip.BindPFlags(pfs); err != nil {
		return err
	}

	if vip.IsSet("config") {
		vip.SetConfigFile(vip.GetString("config"))
		return nil
	}

	vip.AddConfigPath(".")

	homeDir, err := homedir.Dir()
	if err != nil {
		return fmt.Errorf("cannot find home directory: %w", err)
	}

	vip.AddConfigPath(homeDir)

	projectDir, err := config.LocateProjectRootDir()
	if err != nil {
		return fmt.Errorf("cannot find project directory: %w", err)
	}

	vip.AddConfigPath(projectDir)

	return nil
}

func loadConfig() error {
	if vip == nil {
		return fmt.Errorf("viper is not configured")
	}
	if err := vip.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	if err := vip.Unmarshal(&cfg); err != nil {
		return fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return nil
}
