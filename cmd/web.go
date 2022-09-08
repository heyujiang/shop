package cmd

import (
	"fmt"
	"github.com/heyujiang/shop/pkg/setting"
	"os"
	"os/signal"
	"syscall"

	"github.com/heyujiang/shop/app/web/server"
	"github.com/heyujiang/shop/config"

	"github.com/spf13/cobra"
)

// webCmd represents the web command
var webCmd = &cobra.Command{
	Use:   "web",
	Short: "shop web",
	Long:  "shop web",
	Run:   runWeb,
}

func init() {
	rootCmd.AddCommand(webCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// webCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// webCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func runWeb(cmd *cobra.Command, args []string) {
	server.StartWebServer(config.GetConfig().Web.Port)

	fmt.Println("Start shop web server -", setting.AppVersion)

	sig := make(chan os.Signal, 2)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)
	<-sig
}
