/*
Copyright © 2026 Alex Duzi <duzihd@gmail.com>
*/
package cmd

import (
	"github.com/alexduzi/labstresstest/dto"
	"github.com/alexduzi/labstresstest/service"
	"github.com/spf13/cobra"
)

func newRequestCmd(requestsService *service.RequestsService) *cobra.Command {
	return &cobra.Command{
		Use:     "request",
		Aliases: []string{"req"},
		Short:   "Request command makes a HTTP GET request to a given URL",
		Long: `Request command makes a HTTP GET request to a given URL with the parameters:
				--url -> URL to be called
				--requests -> Number of requests
				--concurrency -> Number of concurrent calls`,
		RunE: func(cmd *cobra.Command, args []string) error {
			url, _ := cmd.Flags().GetString("url")
			requests, _ := cmd.Flags().GetInt("requests")
			concurrency, _ := cmd.Flags().GetInt("concurrency")

			req, err := dto.NewRequestsDto(url, requests, concurrency)
			if err != nil {
				return err
			}

			requestsService.ProcessRequests(req)

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
