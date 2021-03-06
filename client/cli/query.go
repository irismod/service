package cli

import (
	"context"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"

	"github.com/irismod/service/client/utils"
	"github.com/irismod/service/types"
)

// GetQueryCmd returns the cli query commands for the module.
func GetQueryCmd() *cobra.Command {
	serviceQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the service module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	serviceQueryCmd.AddCommand(
		GetCmdQueryServiceDefinition(),
		GetCmdQueryServiceBinding(),
		GetCmdQueryServiceBindings(),
		GetCmdQueryWithdrawAddr(),
		GetCmdQueryServiceRequest(),
		GetCmdQueryServiceRequests(),
		GetCmdQueryServiceResponse(),
		GetCmdQueryRequestContext(),
		GetCmdQueryServiceResponses(),
		GetCmdQueryEarnedFees(),
		GetCmdQuerySchema(),
		GetCmdQueryParams(),
	)

	return serviceQueryCmd
}

// GetCmdQueryServiceDefinition implements the query service definition command.
func GetCmdQueryServiceDefinition() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "definition [service-name]",
		Short: "Query a service definition",
		Long: strings.TrimSpace(
			fmt.Sprintf(
				"Query details of a service definition.\n\n"+
					"Example:\n"+
					"$ %s query service definition <service-name>\n",
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			if err := types.ValidateServiceName(args[0]); err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Definition(
				context.Background(),
				&types.QueryDefinitionRequest{
					ServiceName: args[0],
				},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res.ServiceDefinition)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryServiceBinding implements the query service binding command
func GetCmdQueryServiceBinding() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "binding [service-name] [provider-address]",
		Short: "Query a service binding",
		Long: strings.TrimSpace(
			fmt.Sprintf(
				"Query details of a service binding.\n\n"+
					"Example:\n"+
					"$ %s query service binding <service-name> <provider-address>\n",
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			if err := types.ValidateServiceName(args[0]); err != nil {
				return err
			}

			provider, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Binding(
				context.Background(),
				&types.QueryBindingRequest{
					ServiceName: args[0],
					Provider:    provider,
				},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res.ServiceBinding)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryServiceBindings implements the query service bindings command
func GetCmdQueryServiceBindings() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "bindings [service-name]",
		Short: "Query all bindings of a service definition with an optional owner",
		Long: strings.TrimSpace(
			fmt.Sprintf(
				"Query all bindings of a service definition with an optional owner.\n\n"+
					"Example:\n"+
					"$ %s query service bindings <service-name> --owner=<address>\n",
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			if err := types.ValidateServiceName(args[0]); err != nil {
				return err
			}

			var owner sdk.AccAddress
			ownerStr := viper.GetString(FlagOwner)
			if len(ownerStr) > 0 {
				owner, err = sdk.AccAddressFromBech32(ownerStr)
				if err != nil {
					return err
				}
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Bindings(
				context.Background(),
				&types.QueryBindingsRequest{
					ServiceName: args[0],
					Owner:       owner,
				},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res)
		},
	}

	cmd.Flags().AddFlagSet(FsQueryServiceBindings)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryWithdrawAddr implements the query withdraw address command
func GetCmdQueryWithdrawAddr() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdraw-addr [address]",
		Short: "Query the withdrawal address of an owner",
		Long: strings.TrimSpace(
			fmt.Sprintf(
				"Query the withdrawal address of an owner.\n\n"+
					"Example:\n"+
					"$ %s query service withdraw-addr <address>\n",
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			owner, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.WithdrawAddress(
				context.Background(),
				&types.QueryWithdrawAddressRequest{
					Owner: owner,
				},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryServiceRequest implements the query service request command
func GetCmdQueryServiceRequest() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "request [request-id]",
		Short: "Query a request by the request ID",
		Long: strings.TrimSpace(
			fmt.Sprintf(
				"Query details of a service request.\n\n"+
					"Example:\n"+
					"$ %s query service request <request-id>\n",
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			requestID, err := types.ConvertRequestID(args[0])
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Request(
				context.Background(),
				&types.QueryRequestRequest{
					RequestId: requestID,
				},
			)
			if err != nil {
				return err
			}

			if res.Request == nil {
				request, err := utils.QueryRequestByTxQuery(clientCtx, types.QuerierRoute, requestID)
				if err != nil {
					return err
				}
				res.Request = &request
			}

			if res.Request.Empty() {
				return fmt.Errorf("unknown request: %s", requestID)
			}

			return clientCtx.PrintOutput(res.Request)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryServiceRequests implements the query service requests command
func GetCmdQueryServiceRequests() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "requests [service-name] [provider] | [request-context-id] [batch-counter]",
		Short: "Query active requests by the service binding or request context ID",
		Long: strings.TrimSpace(
			fmt.Sprintf(
				"Query active requests by the service binding or request context ID.\n\n"+
					"Example:\n"+
					"$ %s query service requests <service-name> <provider> | <request-context-id> <batch-counter>\n",
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			queryByBinding := true
			provider, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				queryByBinding = false
			}

			queryClient := types.NewQueryClient(clientCtx)

			if queryByBinding {
				res, err := queryClient.Requests(context.Background(), &types.QueryRequestsRequest{
					ServiceName: args[0],
					Provider:    provider,
				})
				if err != nil {
					return err
				}
				return clientCtx.PrintOutput(res)
			}

			requestContextID, err := hex.DecodeString(args[0])
			if err != nil {
				return err
			}

			batchCounter, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			res, err := queryClient.RequestsByReqCtx(
				context.Background(),
				&types.QueryRequestsByReqCtxRequest{
					RequestContextId: requestContextID,
					BatchCounter:     batchCounter,
				},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryServiceResponse implements the query service response command
func GetCmdQueryServiceResponse() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "response [request-id]",
		Short: "Query a response by the request ID",
		Long: strings.TrimSpace(
			fmt.Sprintf(
				"Query details of a service response.\n\n"+
					"Example:\n"+
					"$ %s query service response <request-id>\n",
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			requestID, err := types.ConvertRequestID(args[0])
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Response(
				context.Background(),
				&types.QueryResponseRequest{
					RequestId: requestID,
				},
			)
			if err != nil {
				return err
			}

			if res.Response == nil {
				response, err := utils.QueryResponseByTxQuery(clientCtx, types.QuerierRoute, requestID)
				if err != nil {
					return err
				}
				res.Response = &response
			}

			if res.Response.Empty() {
				return fmt.Errorf("unknown response: %s", requestID)
			}

			return clientCtx.PrintOutput(res.Response)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryServiceResponses implements the query service responses command
func GetCmdQueryServiceResponses() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "responses [request-context-id] [batch-counter]",
		Short: "Query active responses by the request context ID and batch counter",
		Long: strings.TrimSpace(
			fmt.Sprintf(
				"Query active responses by the request context ID and batch counter.\n\n"+
					"Example:\n"+
					"$ %s query service responses <request-context-id> <batch-counter>\n",
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())

			if err != nil {
				return err
			}

			requestContextID, err := hex.DecodeString(args[0])
			if err != nil {
				return err
			}

			batchCounter, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Responses(
				context.Background(),
				&types.QueryResponsesRequest{
					RequestContextId: requestContextID,
					BatchCounter:     batchCounter,
				},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryRequestContext implements the query request context command
func GetCmdQueryRequestContext() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "request-context [request-context-id]",
		Short: "Query a request context",
		Long: strings.TrimSpace(
			fmt.Sprintf(
				"Query a request context.\n\n"+
					"Example:\n"+
					"$ %s query service request-context <request-context-id>\n",
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			requestContextID, err := hex.DecodeString(args[0])
			if err != nil {
				return err
			}

			res, err := utils.QueryRequestContext(
				clientCtx,
				types.QuerierRoute,
				types.QueryRequestContextRequest{
					RequestContextId: requestContextID,
				},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(&res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryEarnedFees implements the query earned fees command
func GetCmdQueryEarnedFees() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fees [provider-address]",
		Short: "Query the earned fees of a provider",
		Long: strings.TrimSpace(
			fmt.Sprintf(
				"Query the earned fees of a provider.\n\n"+
					"Example:\n"+
					"$ %s query service fees <provider-address>\n",
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			provider, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.EarnedFees(
				context.Background(),
				&types.QueryEarnedFeesRequest{
					Provider: provider,
				},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQuerySchema implements the query schema command
func GetCmdQuerySchema() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "schema [schema-name]",
		Short: "Query the system schema by the schema name",
		Long: strings.TrimSpace(
			fmt.Sprintf(
				"Query the system schema by the schema name, only pricing and result allowed.\n\n"+
					"Example:\n"+
					"$ %s query service schema <schema-name>\n",
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)
			res, err := queryClient.Schema(
				context.Background(),
				&types.QuerySchemaRequest{
					SchemaName: args[0],
				},
			)
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryParams implements the query params command.
func GetCmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Short: "Query the current service parameter values",
		Long: strings.TrimSpace(
			fmt.Sprintf(
				"Query values set as service parameters.\n\n"+
					"Example:\n"+
					"$ %s query service params\n",
				version.AppName,
			),
		),
		Args: cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			clientCtx, err := client.ReadQueryCommandFlags(clientCtx, cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Params(context.Background(), &types.QueryParamsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintOutput(&res.Params)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
