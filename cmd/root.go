package cmd

import (
	"github.com/heyujiang/shop/config"
	"os"

	"github.com/spf13/cobra"
)

var cfgConfig string

// rootCmd represents the root command
var rootCmd = &cobra.Command{
	Use:   "shop",
	Short: "heyujiang`s shop demo",
	Long:  "heyujiang`s shop demo",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rootCmd.PersistentFlags().String("foo", "", "A help for foo")
	rootCmd.PersistentFlags().StringVar(&cfgConfig, "config", "./config/shop.toml", "config file default ./config/shop.toml")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	config.InitConfig(cfgConfig)
}
