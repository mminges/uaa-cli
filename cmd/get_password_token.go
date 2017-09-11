package cmd

import (
	"code.cloudfoundry.org/uaa-cli/config"
	"code.cloudfoundry.org/uaa-cli/help"
	"code.cloudfoundry.org/uaa-cli/uaa"
	"github.com/spf13/cobra"
	"os"
)

var getPasswordToken = &cobra.Command{
	Use:   "get-password-token CLIENT_ID -s CLIENT_SECRET -u USERNAME -p PASSWORD",
	Short: "Obtain an access token using the password grant type",
	Long:  help.PasswordGrant(),
	PreRun: func(cmd *cobra.Command, args []string) {
		EnsureTarget()
	},
	Run: func(cmd *cobra.Command, args []string) {
		clientId := args[0]
		requestedType := uaa.TokenFormat(tokenFormat)

		ccClient := uaa.ResourceOwnerPasswordClient{
			ClientId:     clientId,
			ClientSecret: clientSecret,
			Username:     username,
			Password:     password,
		}
		c := GetSavedConfig()
		tokenResponse, err := ccClient.RequestToken(GetHttpClient(), c, requestedType)
		if err != nil {
			log.Error("An error occurred while fetching token.")
			TraceRetryMsg(c)
			os.Exit(1)
		}

		activeContext := c.GetActiveContext()
		activeContext.ClientId = clientId
		activeContext.GrantType = uaa.PASSWORD
		activeContext.Username = username
		activeContext.TokenResponse = tokenResponse
		c.AddContext(activeContext)
		config.WriteConfig(c)
		log.Info("Access token successfully fetched and added to context.")
	},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			MissingArgument("client_id", cmd)
		}
		if clientSecret == "" {
			MissingArgument("client_secret", cmd)
		}
		if password == "" {
			MissingArgument("password", cmd)
		}
		if username == "" {
			MissingArgument("username", cmd)
		}
		validateTokenFormat(cmd, tokenFormat)
		return nil
	},
}

func init() {
	RootCmd.AddCommand(getPasswordToken)
	getPasswordToken.Annotations = make(map[string]string)
	getPasswordToken.Annotations[TOKEN_CATEGORY] = "true"
	getPasswordToken.Flags().StringVarP(&clientSecret, "client_secret", "s", "", "client secret")
	getPasswordToken.Flags().StringVarP(&username, "username", "u", "", "username")
	getPasswordToken.Flags().StringVarP(&password, "password", "p", "", "user password")
	getPasswordToken.Flags().StringVarP(&tokenFormat, "format", "", "jwt", "available formats include "+availableFormatsStr())
}
