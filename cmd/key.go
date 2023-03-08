package cmd

import (
	"leizhenpeng/go-gpt3-cli/services"

	"github.com/spf13/cobra"
)

// keyCmd represents the key command
var keyCmd = &cobra.Command{
	Use:   "key",
	Short: "Manage your api key about gpt3.5",
	Long: `Manage your api key about gpt3.5.
You can set, list and clear your key.	
You can get your key from https://beta.openai.com/account/api-keys`,

	Run: func(cmd *cobra.Command, args []string) {
		keyMsg = services.GetKeyMag()
		if cmd.Flag("bowser").Value.String() == "true" {
			services.OpenLinkInBrowser("https://beta.openai.com/account/api-keys")
		}
		if cmd.Flag("set").Value.String() == "true" {
			if len(args) == 0 {
				cmd.Help()
				return
			}
			keyMsg.SetKey(keyName, args[0])
			cmd.Println("Your key is set to: ", args[0])
		} else if cmd.Flag("list").Value.String() == "true" {
			key := keyMsg.GetKey(keyName)
			if key == "" {
				cmd.Println("You don't have a key")
			} else {
				cmd.Println("Your key is: ", key)
			}
		} else if cmd.Flag("clear").Value.String() == "true" {
			keyMsg.DelKey(keyName)
			cmd.Println("Your key is cleared")
		} else {
			cmd.Help()
		}
	},
}

// sysCmd represents the key command
var sysCmd = &cobra.Command{
	Use:   "sys",
	Short: "Set system role",
	Long: `Set system role
	Run command : chat sys -s <你是我的私人智能助手>`,

	Run: func(cmd *cobra.Command, args []string) {
		keyMsg = services.GetKeyMag()
		if cmd.Flag("set").Value.String() == "true" {
			if len(args) == 0 {
				cmd.Help()
				return
			}
			keyMsg.SetKey(system, args[0])
			cmd.Printf("我将以'%s'身份和你对话\n", args[0])
		} else {
			cmd.Help()
		}
	},
}

func init() {
	rootCmd.AddCommand(keyCmd)
	keyCmd.Flags().BoolP("set", "s", false, "set api key")
	keyCmd.Flags().BoolP("list", "l", false, " list api key")
	keyCmd.Flags().BoolP("clear", "c", false, "clear api key")
	keyCmd.Flags().BoolP("bowser", "b", false, "show bowser to check  key")

	rootCmd.AddCommand(sysCmd)
	sysCmd.Flags().BoolP("set", "s", false, "Set system role")
}
