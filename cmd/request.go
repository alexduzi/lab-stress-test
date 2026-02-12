/*
Copyright © 2026 Alex Duzi <duzihd@gmail.com>
*/
package cmd

import (
	"strconv"

	"github.com/alexduzi/labstresstest/service"
	"github.com/spf13/cobra"
)

func newRequestCmd(requestsService *service.RequestsService) *cobra.Command {
	return &cobra.Command{
		Use:     "request",
		Aliases: []string{"req"},
		Short:   "Request command makes a HTTP GET request to a given URL",
		Long: `Request command makes a HTTP GET request to a given URL with the parameters:
				--url -> url to be called
				--requests -> amount of the requests
				--concurrency -> amount of the concurrent calls`,
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.ValidateRequiredFlags(); err != nil {
				return err
			}

			url := args[0]
			requests, _ := strconv.Atoi(args[1])
			concurrent, _ := strconv.Atoi(args[2])

			req := service.NewRequestsDto(url, requests, concurrent)

			if err := requestsService.Call(req); err != nil {
				return err
			}

			return nil
		},
	}
}

func init() {
	requestsService := service.NewRequestsService()
	requestCmd := newRequestCmd(requestsService)
	requestCmd.Flags().StringP("url", "u", "", "Help message for toggle")
	requestCmd.Flags().IntP("requests", "r", 0, "Help message for toggle")
	requestCmd.Flags().IntP("concurrency", "c", 0, "Help message for toggle")
	requestCmd.MarkFlagRequired("url")
	requestCmd.MarkFlagRequired("requests")
	requestCmd.MarkFlagRequired("concurrency")
	rootCmd.AddCommand(requestCmd)
}
