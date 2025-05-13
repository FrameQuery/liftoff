package versionCmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var version = "0.1.0"

func init() {
	version = os.Getenv("LIFTOFF_VERSION")
	if version == "" {
		version = "0.0.1"
	}
}

var Cmd = &cobra.Command{
	Use:   "version",
	Short: "🌟 Show liftoff version",
	Run: func(cmd *cobra.Command, args []string) {
		art := `
 __         __     ______   ______   ______     ______   ______ 
/\ \       /\ \   /\  ___\ /\__  _\ /\  __ \   /\  ___\ /\  ___\
\ \ \____  \ \ \  \ \  __\ \/_/\ \/ \ \ \/\ \  \ \  __\ \ \  __\
 \ \_____\  \ \_\  \ \_\      \ \_\  \ \_____\  \ \_\    \ \_\  
  \/_____/   \/_/   \/_/       \/_/   \/_____/   \/_/     \/_/  

` + fmt.Sprintf("🚀  liftoff version %s", version)
		fmt.Println(art)
	},
}
