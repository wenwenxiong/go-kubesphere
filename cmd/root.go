package cmd

import (
	"context"
	"fmt"
	"github.com/wenwenxiong/go-kubesphere/kubesphere"

	"github.com/spf13/cobra"
)

// NewBackupRestoreCommand represents the base command when called without any subcommands
func NewKubesphereCommand(ctx context.Context) *cobra.Command {
	var RootCmd = &cobra.Command{
		Use:   "kubespherectl",
		Short: "command line utility for kubesphere app category ",
		Long: `The kubespherectl, command line utility, is built to support curd kubesphere app category
related functionality. Sub-command for this root command will support features
like get create delete update app category and get update app.`,
		Run: func(cmd *cobra.Command, args []string) {
			if version {
				printVersionInfo()
			}
		},
	}
	RootCmd.Flags().BoolVarP(&version, "version", "v", false, "print version info")
	RootCmd.PersistentFlags().StringVarP(&apigateway, "apiGateway", "a", "http://192.168.122.162:30881/", "ks-apigateway url")
	RootCmd.AddCommand(NewAppCommand(ctx),
		NewAppCategoryCommand(ctx))
	return RootCmd
}

func GetClient(apiGateway string) *kubesphere.Client {
	var client *kubesphere.Client
	if apiGateway == "" {
		client = kubesphere.NewClient(nil)
	}else {
		client = kubesphere.NewClientSpecify(nil, apiGateway)
	}
	return client
}

func GetAccessToken(apiGateway string ) *kubesphere.AccessToken{

	client := GetClient(apiGateway)
	v := &kubesphere.IamRequest{
		Username:        String("admin"),
		Password:     String("P@88w0rd"),
	}
	ctx := context.Background()

	a, _, err := client.Iams.GetAccessToken(ctx, v)
	if _, ok := err.(*kubesphere.TwoFactorAuthError); ok {
		fmt.Print("\nGitHub OTP: ")
	}

	if err != nil {
		fmt.Printf("\nerror: %v\n", err)
		return nil
	}

	fmt.Printf("Accesstoken: \n%v\n", *a.Accesstoken)
	return  a
}

func String(v string) *string { return &v }