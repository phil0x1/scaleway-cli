// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

package cockpit

import (
	"context"
	"reflect"

	"github.com/scaleway/scaleway-cli/v2/internal/core"
	"github.com/scaleway/scaleway-sdk-go/api/cockpit/v1beta1"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

// always import dependencies
var (
	_ = scw.RegionFrPar
)

func GetGeneratedCommands() *core.Commands {
	return core.NewCommands(
		cockpitRoot(),
		cockpitCockpit(),
		cockpitToken(),
		cockpitGrafanaUser(),
		cockpitCockpitActivate(),
		cockpitCockpitGet(),
		cockpitCockpitDeactivate(),
		cockpitCockpitResetGrafana(),
		cockpitTokenCreate(),
		cockpitTokenList(),
		cockpitTokenGet(),
		cockpitTokenDelete(),
		cockpitGrafanaUserCreate(),
		cockpitGrafanaUserList(),
		cockpitGrafanaUserDelete(),
		cockpitGrafanaUserResetPassword(),
	)
}
func cockpitRoot() *core.Command {
	return &core.Command{
		Short:     `Cockpit API`,
		Long:      `This API allows to manage Cockpits.`,
		Namespace: "cockpit",
	}
}

func cockpitCockpit() *core.Command {
	return &core.Command{
		Short:     `Cockpit management commands`,
		Long:      `Cockpit management commands.`,
		Namespace: "cockpit",
		Resource:  "cockpit",
	}
}

func cockpitToken() *core.Command {
	return &core.Command{
		Short:     `Token management commands`,
		Long:      `Token management commands.`,
		Namespace: "cockpit",
		Resource:  "token",
	}
}

func cockpitGrafanaUser() *core.Command {
	return &core.Command{
		Short:     `Grafana user management commands`,
		Long:      `Grafana user management commands.`,
		Namespace: "cockpit",
		Resource:  "grafana-user",
	}
}

func cockpitCockpitActivate() *core.Command {
	return &core.Command{
		Short:     `Activate a cockpit`,
		Long:      `Activate a cockpit associated with the given project ID.`,
		Namespace: "cockpit",
		Resource:  "cockpit",
		Verb:      "activate",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(cockpit.ActivateCockpitRequest{}),
		ArgSpecs: core.ArgSpecs{
			core.ProjectIDArgSpec(),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*cockpit.ActivateCockpitRequest)

			client := core.ExtractClient(ctx)
			api := cockpit.NewAPI(client)
			return api.ActivateCockpit(request)

		},
	}
}

func cockpitCockpitGet() *core.Command {
	return &core.Command{
		Short:     `Get cockpit`,
		Long:      `Get the cockpit associated with the given project ID.`,
		Namespace: "cockpit",
		Resource:  "cockpit",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(cockpit.GetCockpitRequest{}),
		ArgSpecs: core.ArgSpecs{
			core.ProjectIDArgSpec(),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*cockpit.GetCockpitRequest)

			client := core.ExtractClient(ctx)
			api := cockpit.NewAPI(client)
			return api.GetCockpit(request)

		},
	}
}

func cockpitCockpitDeactivate() *core.Command {
	return &core.Command{
		Short:     `Deactivate a cockpit`,
		Long:      `Deactivate a cockpit associated with the given project ID.`,
		Namespace: "cockpit",
		Resource:  "cockpit",
		Verb:      "deactivate",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(cockpit.DeactivateCockpitRequest{}),
		ArgSpecs: core.ArgSpecs{
			core.ProjectIDArgSpec(),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*cockpit.DeactivateCockpitRequest)

			client := core.ExtractClient(ctx)
			api := cockpit.NewAPI(client)
			return api.DeactivateCockpit(request)

		},
	}
}

func cockpitCockpitResetGrafana() *core.Command {
	return &core.Command{
		Short:     `Reset Grafana`,
		Long:      `Reset the Grafana of your cockpit associated with the given project ID.`,
		Namespace: "cockpit",
		Resource:  "cockpit",
		Verb:      "reset-grafana",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(cockpit.ResetCockpitGrafanaRequest{}),
		ArgSpecs: core.ArgSpecs{
			core.ProjectIDArgSpec(),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*cockpit.ResetCockpitGrafanaRequest)

			client := core.ExtractClient(ctx)
			api := cockpit.NewAPI(client)
			return api.ResetCockpitGrafana(request)

		},
	}
}

func cockpitTokenCreate() *core.Command {
	return &core.Command{
		Short:     `Create a token`,
		Long:      `Create a token associated with the given project ID.`,
		Namespace: "cockpit",
		Resource:  "token",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(cockpit.CreateTokenRequest{}),
		ArgSpecs: core.ArgSpecs{
			core.ProjectIDArgSpec(),
			{
				Name:       "name",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "scopes.query-metrics",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "scopes.write-metrics",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "scopes.setup-metrics-rules",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "scopes.query-logs",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "scopes.write-logs",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "scopes.setup-logs-rules",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "scopes.setup-alerts",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*cockpit.CreateTokenRequest)

			client := core.ExtractClient(ctx)
			api := cockpit.NewAPI(client)
			return api.CreateToken(request)

		},
	}
}

func cockpitTokenList() *core.Command {
	return &core.Command{
		Short:     `List tokens`,
		Long:      `List tokens associated with the given project ID.`,
		Namespace: "cockpit",
		Resource:  "token",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(cockpit.ListTokensRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "order-by",
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"created_at_asc", "created_at_desc", "name_asc", "name_desc"},
			},
			core.ProjectIDArgSpec(),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*cockpit.ListTokensRequest)

			client := core.ExtractClient(ctx)
			api := cockpit.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			resp, err := api.ListTokens(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.Tokens, nil

		},
	}
}

func cockpitTokenGet() *core.Command {
	return &core.Command{
		Short:     `Get token`,
		Long:      `Get the token associated with the given ID.`,
		Namespace: "cockpit",
		Resource:  "token",
		Verb:      "get",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(cockpit.GetTokenRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "token-id",
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*cockpit.GetTokenRequest)

			client := core.ExtractClient(ctx)
			api := cockpit.NewAPI(client)
			return api.GetToken(request)

		},
	}
}

func cockpitTokenDelete() *core.Command {
	return &core.Command{
		Short:     `Delete token`,
		Long:      `Delete the token associated with the given ID.`,
		Namespace: "cockpit",
		Resource:  "token",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(cockpit.DeleteTokenRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "token-id",
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*cockpit.DeleteTokenRequest)

			client := core.ExtractClient(ctx)
			api := cockpit.NewAPI(client)
			e = api.DeleteToken(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "token",
				Verb:     "delete",
			}, nil
		},
	}
}

func cockpitGrafanaUserCreate() *core.Command {
	return &core.Command{
		Short:     `Create a grafana user`,
		Long:      `Create a grafana user for your grafana instance.`,
		Namespace: "cockpit",
		Resource:  "grafana-user",
		Verb:      "create",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(cockpit.CreateGrafanaUserRequest{}),
		ArgSpecs: core.ArgSpecs{
			core.ProjectIDArgSpec(),
			{
				Name:       "login",
				Required:   false,
				Deprecated: false,
				Positional: false,
			},
			{
				Name:       "role",
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"unknown_role", "editor", "viewer"},
			},
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*cockpit.CreateGrafanaUserRequest)

			client := core.ExtractClient(ctx)
			api := cockpit.NewAPI(client)
			return api.CreateGrafanaUser(request)

		},
	}
}

func cockpitGrafanaUserList() *core.Command {
	return &core.Command{
		Short:     `List grafana users`,
		Long:      `List grafana users who are able to connect to your grafana instance.`,
		Namespace: "cockpit",
		Resource:  "grafana-user",
		Verb:      "list",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(cockpit.ListGrafanaUsersRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "order-by",
				Required:   false,
				Deprecated: false,
				Positional: false,
				EnumValues: []string{"login_asc", "login_desc"},
			},
			core.ProjectIDArgSpec(),
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*cockpit.ListGrafanaUsersRequest)

			client := core.ExtractClient(ctx)
			api := cockpit.NewAPI(client)
			opts := []scw.RequestOption{scw.WithAllPages()}
			resp, err := api.ListGrafanaUsers(request, opts...)
			if err != nil {
				return nil, err
			}
			return resp.GrafanaUsers, nil

		},
	}
}

func cockpitGrafanaUserDelete() *core.Command {
	return &core.Command{
		Short:     `Delete a grafana user`,
		Long:      `Delete a grafana user from your grafana instance.`,
		Namespace: "cockpit",
		Resource:  "grafana-user",
		Verb:      "delete",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(cockpit.DeleteGrafanaUserRequest{}),
		ArgSpecs: core.ArgSpecs{
			core.ProjectIDArgSpec(),
			{
				Name:       "grafana-user-id",
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*cockpit.DeleteGrafanaUserRequest)

			client := core.ExtractClient(ctx)
			api := cockpit.NewAPI(client)
			e = api.DeleteGrafanaUser(request)
			if e != nil {
				return nil, e
			}
			return &core.SuccessResult{
				Resource: "grafana-user",
				Verb:     "delete",
			}, nil
		},
	}
}

func cockpitGrafanaUserResetPassword() *core.Command {
	return &core.Command{
		Short:     `Reset Grafana user password`,
		Long:      `Reset the Grafana user password from your grafana instance.`,
		Namespace: "cockpit",
		Resource:  "grafana-user",
		Verb:      "reset-password",
		// Deprecated:    false,
		ArgsType: reflect.TypeOf(cockpit.ResetGrafanaUserPasswordRequest{}),
		ArgSpecs: core.ArgSpecs{
			core.ProjectIDArgSpec(),
			{
				Name:       "grafana-user-id",
				Required:   true,
				Deprecated: false,
				Positional: false,
			},
		},
		Run: func(ctx context.Context, args interface{}) (i interface{}, e error) {
			request := args.(*cockpit.ResetGrafanaUserPasswordRequest)

			client := core.ExtractClient(ctx)
			api := cockpit.NewAPI(client)
			return api.ResetGrafanaUserPassword(request)

		},
	}
}
