package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wenwenxiong/go-kubesphere/kubesphere"
)

func NewAppCommand(ctx context.Context) *cobra.Command {
	var command = &cobra.Command{
		Use:   "app [OPTIONS] [COMMANDS]",
		Short: "op for app .",
		Long: `get  update app in kubesphere appstore.`,
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	command.AddCommand(NewAppGetCommand(ctx),
			NewAppUpdateCommand(ctx))

	return command
}

func NewAppGetCommand(ctx context.Context) *cobra.Command {

	var appName string
	var command = &cobra.Command{
		Use:   "get [OPTIONS] ",
		Short: "get app ",
		Long: `get app in kubesphere store`,
		Run: func(cmd *cobra.Command, args []string) {
			accessToke := GetAccessToken(apigateway)
			client := GetClient(apigateway)
			var c *kubesphere.App
			if appName != "" {
				c = &kubesphere.App{
					Name: String(appName),
				}
			}
			r, _, err :=client.Openpitrixs.GetApp(ctx, c, accessToke)
			if _, ok := err.(*kubesphere.TwoFactorAuthError); ok {
				fmt.Print("\nGitHub OTP: ")
			}

			if err != nil {
				fmt.Printf("\nerror: %v\n", err)
				return
			}
			if *r.TotalCount > 0 {
				fmt.Printf("get app result total: \n%v\n", *r.TotalCount)
				for i, v := range r.Items {
					fmt.Printf("\tget app at %d result:\n", i+1)
					fmt.Printf("\t\tapp name: %s\n", *v.Name)
					fmt.Printf("\t\tapp id: %s\n", *v.AppId)
					//fmt.Printf("\t\tapp category id: %s\n", *v.CategoryId)
				}

			}else {
				fmt.Printf("get app result: \n%v\n", *r.TotalCount)
			}
		},
	}
	command.Flags().StringVarP(&appName, "appName", "n", "", "app name")
	return command
}

func NewAppUpdateCommand(ctx context.Context) *cobra.Command {

	var appId string
	var appNewName string
	var appNewDescription string
	var appNewCategoryId string
	var command = &cobra.Command{
		Use:   "update [OPTIONS] ",
		Short: "update app ",
		Long: `update app in kubesphere store`,
		Run: func(cmd *cobra.Command, args []string) {
			accessToke := GetAccessToken(apigateway)
			client := GetClient(apigateway)
			c := &kubesphere.App{
				AppId: String(appId),
			}
			if appNewName != "" {
				c.SetName(&appNewName)
			}
			if appNewDescription != "" {
				c.SetDescription(&appNewDescription)
			}
			if appNewCategoryId != "" {
				c.SetCategoryId(&appNewCategoryId)
			}
			r, _, err :=client.Openpitrixs.UpdateApp(ctx, c, accessToke)
			if _, ok := err.(*kubesphere.TwoFactorAuthError); ok {
				fmt.Print("\nkubesphere OTP: ")
			}

			if err != nil {
				fmt.Printf("\nerror: %v\n", err)
				return
			}
			if r != nil {
				fmt.Printf("update app id: %s, result %s\n", appId, *r.Message)
			}
		},
	}
	command.Flags().StringVarP(&appId, "appId", "i", "", "app id")
	command.MarkFlagRequired("appId")
	command.Flags().StringVarP(&appNewName, "appNewName", "n", "", "app new name")
	command.Flags().StringVarP(&appNewDescription, "appNewDescription", "d", "", "app new description")
	command.Flags().StringVarP(&appNewCategoryId, "appNewCategoryId", "c", "", "app new category Id")
	return command
}