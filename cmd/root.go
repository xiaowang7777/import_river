package cmd

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"import_river/config"
	"import_river/pkg/client"
)

var cfgFile string

var ROOT_CMD = cobra.Command{
	Use:   "river",
	Short: "导入河流点位信息",
	Long: `
导入河流点位信息。
eg:...`,
	//Run: func(cmd *cobra.Command, args []string) {
	//	c := &config.Config{}
	//	if err := c.Load(); err != nil {
	//		logrus.Fatal(err)
	//	}
	//	client.Run(c)
	//},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		c := &config.Config{}
		if err := c.Load(); err != nil {
			logrus.Fatal(err)
		}
		client.Run(c)
	},
}

func Execute() {
	if err := ROOT_CMD.Execute(); err != nil {
		logrus.Fatal(err)
	}
}

func init() {
	initConfig()
	ROOT_CMD.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}
	viper.SetConfigName(config.Config_File_Name)
	viper.AddConfigPath("$HOME")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Using config path", viper.ConfigFileUsed())
	}
}
