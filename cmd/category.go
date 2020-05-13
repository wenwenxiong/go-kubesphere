package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wenwenxiong/go-kubesphere/kubesphere"
)

func NewAppCategoryCommand(ctx context.Context) *cobra.Command {
	var command = &cobra.Command{
		Use:   "category [OPTIONS] [COMMANDS]",
		Short: "op for app category.",
		Long: `get create delete update app category.`,
		Run: func(cmd *cobra.Command, args []string) {
	},
	}

	command.AddCommand(NewAppCategoryGetCommand(ctx),
		NewAppCategoryCreateCommand(ctx),
		NewAppCategoryDeleteCommand(ctx),
		NewAppCategoryUpdateCommand(ctx))

	return command
}

func NewAppCategoryGetCommand(ctx context.Context) *cobra.Command {

	var categoryName string
	var command = &cobra.Command{
		Use:   "get [OPTIONS] ",
		Short: "get app category.",
		Long: `get app category.`,
		Run: func(cmd *cobra.Command, args []string) {
			accessToke := GetAccessToken(apigateway)
			client := GetClient(apigateway)
			var c *kubesphere.AppCategory
			if categoryName != "" {
				c = &kubesphere.AppCategory{
					Name: String(categoryName),
				}
			}
			r, _, err :=client.Openpitrixs.GetAppCategory(ctx, c, accessToke)
			if _, ok := err.(*kubesphere.TwoFactorAuthError); ok {
				fmt.Print("\nGitHub OTP: ")
			}

			if err != nil {
				fmt.Printf("\nerror: %v\n", err)
				return
			}
			if *r.TotalCount > 0 {
				fmt.Printf("get app category result total: \n%v\n", *r.TotalCount)
				for i, v := range r.Items {
					fmt.Printf("\tget app category at %d result:\n", i+1)
					fmt.Printf("\t\tapp category name: %s\n", *v.Name)
					fmt.Printf("\t\tapp category id: %s\n", *v.CategoryId)
				}

			}else {
				fmt.Printf("get app category result: \n%v\n", *r.TotalCount)
			}
		},
	}
	command.Flags().StringVarP(&categoryName, "categoryName", "n", "", "app category name")
	return command
}

func NewAppCategoryCreateCommand(ctx context.Context) *cobra.Command {

	var categoryName string
	var locale string
	var command = &cobra.Command{
		Use:   "create [OPTIONS] ",
		Short: "create app category.",
		Long: `create app category.`,
		Run: func(cmd *cobra.Command, args []string) {
			accessToke := GetAccessToken(apigateway)
			client := GetClient(apigateway)
			c := &kubesphere.AppCategory{
					Name: String(categoryName),
					Locale: String(locale),
			}
			r, _, err :=client.Openpitrixs.CreateAppCategory(ctx, c, accessToke)
			if _, ok := err.(*kubesphere.TwoFactorAuthError); ok {
				fmt.Print("\nkubesphere OTP: ")
			}

			if err != nil {
				fmt.Printf("\nerror: %v\n", err)
				return
			}
			if r != nil {
				fmt.Printf("create app category id: %s\n", *r.CategoryId)
			}
		},
	}
	command.Flags().StringVarP(&categoryName, "categoryName", "n", "", "app category name")
	command.MarkFlagRequired("categoryName")
	command.Flags().StringVarP(&locale, "locale", "l", "", "app category locale configuration")
	command.MarkFlagRequired("locale")
	return command
}

func NewAppCategoryDeleteCommand(ctx context.Context) *cobra.Command {

	var categoryId string
	var command = &cobra.Command{
		Use:   "delete [OPTIONS] ",
		Short: "delete app category.",
		Long: `delete app category.`,
		Run: func(cmd *cobra.Command, args []string) {
			accessToke := GetAccessToken(apigateway)
			client := GetClient(apigateway)
			c := &kubesphere.AppCategory{
					CategoryId: String(categoryId),
			}
			r, _, err :=client.Openpitrixs.DeleteAppCategory(ctx, c, accessToke)
			if _, ok := err.(*kubesphere.TwoFactorAuthError); ok {
				fmt.Print("\nkubesphere OTP: ")
			}

			if err != nil {
				fmt.Printf("\nerror: %v\n", err)
				return
			}
			if r != nil {
				fmt.Printf("delete app category id: %s, result %s\n", categoryId, *r.Message)
			}
		},
	}
	command.Flags().StringVarP(&categoryId, "categoryId", "i", "", "app category id")
	command.MarkFlagRequired("categoryId")
	return command
}

func NewAppCategoryUpdateCommand(ctx context.Context) *cobra.Command {

	var categoryId string
	var categoryNewName string
	var categoryNewDescription string
	var command = &cobra.Command{
		Use:   "update [OPTIONS] ",
		Short: "update app category.",
		Long: `update app category.`,
		Run: func(cmd *cobra.Command, args []string) {
			accessToke := GetAccessToken(apigateway)
			client := GetClient(apigateway)
			c := &kubesphere.AppCategory{
				CategoryId: String(categoryId),
			}
			if categoryNewName != "" {
				c.SetName(&categoryNewName)
			}
			if categoryNewDescription != "" {
				c.SetDescription(&categoryNewDescription)
			}
			r, _, err :=client.Openpitrixs.UpdateAppCategory(ctx, c, accessToke)
			if _, ok := err.(*kubesphere.TwoFactorAuthError); ok {
				fmt.Print("\nkubesphere OTP: ")
			}

			if err != nil {
				fmt.Printf("\nerror: %v\n", err)
				return
			}
			if r != nil {
				fmt.Printf("update app category id: %s, result %s\n", categoryId, *r.Message)
			}
		},
	}
	command.Flags().StringVarP(&categoryId, "categoryId", "i", "", "app category id")
	command.MarkFlagRequired("categoryId")
	command.Flags().StringVarP(&categoryNewName, "categoryNewName", "n", "", "app category new name")
	command.Flags().StringVarP(&categoryNewDescription, "categoryNewDescription", "d", "", "app category new description")
	return command
}