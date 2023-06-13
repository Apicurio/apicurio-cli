package defaultapi

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"

	// "github.com/apicurio/apicurio-cli/pkg/shared/connection"

	"github.com/apicurio/apicurio-cli/internal/build"
	"github.com/apicurio/apicurio-cli/pkg/api/generic"
	"github.com/apicurio/apicurio-cli/pkg/api/rbac"
	"github.com/apicurio/apicurio-cli/pkg/shared/connection/api"
	"github.com/apicurio/apicurio-cli/pkg/shared/svcstatus"
	ocmSdkClient "github.com/openshift-online/ocm-sdk-go"
	ocmclustersmgmtv1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"
	registryinstance "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/registryinstance/apiv1internal"
	registryinstanceclient "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/registryinstance/apiv1internal/client"
	registrymgmt "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/registrymgmt/apiv1"
	registrymgmtclient "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/registrymgmt/apiv1/client"
	"golang.org/x/oauth2"

	svcacctmgmt "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/serviceaccountmgmt/apiv1"

	svcacctmgmtclient "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/serviceaccountmgmt/apiv1/client"
)

// defaultAPI is a type which defines a number of API creator functions
type defaultAPI struct {
	api.Config
}

// New creates a new default API client wrapper
func New(cfg *api.Config) *defaultAPI {
	return &defaultAPI{
		Config: *cfg,
	}
}

func (a *defaultAPI) GetConfig() api.Config {
	return a.Config
}

// ServiceRegistryMgmt return a new Service Registry Management API client instance
func (a *defaultAPI) ServiceRegistryMgmt() registrymgmtclient.RegistriesApi {
	tc := a.CreateOAuthTransport(a.AccessToken)
	client := registrymgmt.NewAPIClient(&registrymgmt.Config{
		BaseURL:    a.ApiURL.String(),
		Debug:      a.Logger.DebugEnabled(),
		HTTPClient: tc,
		UserAgent:  build.DefaultUserAgentPrefix + build.Version,
	})

	return client.RegistriesApi
}

// ServiceAccountMgmt return a new Service Account Management API client instance
func (a *defaultAPI) ServiceAccountMgmt() svcacctmgmtclient.ServiceAccountsApi {
	tc := a.CreateOAuthTransport(a.AccessToken)
	client := svcacctmgmt.NewAPIClient(&svcacctmgmt.Config{
		BaseURL:    a.AuthURL.String(),
		Debug:      a.Logger.DebugEnabled(),
		HTTPClient: tc,
		UserAgent:  a.UserAgent,
	})

	return client.ServiceAccountsApi
}

// ServiceRegistryInstance returns a new Service Registry API client instance, with the Registry configuration object
func (a *defaultAPI) ServiceRegistryInstance(instanceID string) (*registryinstanceclient.APIClient, *registrymgmtclient.Registry, error) {
	registryAPI := a.ServiceRegistryMgmt()

	instance, resp, err := registryAPI.GetRegistry(context.Background(), instanceID).Execute()
	defer resp.Body.Close()
	if err != nil {
		return nil, nil, fmt.Errorf("%w", err)
	}

	status := svcstatus.ServiceStatus(instance.GetStatus())
	// nolint
	switch status {
	case svcstatus.StatusProvisioning, svcstatus.StatusAccepted:
		err = fmt.Errorf(`service registry instance "%v" is not ready yet`, instance.GetName())
		return nil, nil, err
	case svcstatus.StatusFailed:
		err = fmt.Errorf(`service registry instance "%v" has failed`, instance.GetName())
		return nil, nil, err
	case svcstatus.StatusDeprovision:
		err = fmt.Errorf(`service registry instance "%v" is being deprovisioned`, instance.GetName())
		return nil, nil, err
	case svcstatus.StatusDeleting:
		err = fmt.Errorf(`service registry instance "%v" is being deleted`, instance.GetName())
		return nil, nil, err
	}

	registryUrl := instance.GetRegistryUrl()
	if registryUrl == "" {
		err = fmt.Errorf(`URL is missing for Service Registry instance "%v"`, instance.GetName())

		return nil, nil, err
	}

	host, port, _ := net.SplitHostPort(registryUrl)

	var baseURL string
	if host == "localhost" {
		var apiURL = &url.URL{
			Scheme: "http",
			Host:   fmt.Sprintf("localhost:%v", port),
		}
		apiURL.Scheme = "http"
		apiURL.Path = "/data/registry"
		baseURL = apiURL.String()
		fmt.Println(baseURL)
	} else {
		baseURL = registryUrl + "/apis/registry/v2"
	}

	a.Logger.Debugf("Making request to %v", baseURL)

	token := a.AccessToken

	client := registryinstance.NewAPIClient(&registryinstance.Config{
		BaseURL:    "http://localhost:8080/apis/registry/v2",
		Debug:      a.Logger.DebugEnabled(),
		HTTPClient: a.CreateOAuthTransport(token),
		UserAgent:  build.DefaultUserAgentPrefix + build.Version,
	})

	return client, &instance, nil
}

func (a *defaultAPI) GenericAPI() generic.GenericAPI {
	tc := a.CreateOAuthTransport(a.AccessToken)
	client := generic.NewGenericAPIClient(&generic.Config{
		BaseURL:    a.ApiURL.String(),
		Debug:      a.Logger.DebugEnabled(),
		HTTPClient: tc,
	})

	return client
}

// RBAC returns a new RBAC API client instance
func (a *defaultAPI) RBAC() rbac.RbacAPI {
	rbacAPI := rbac.RbacAPI{
		PrincipalAPI: func() rbac.PrincipalAPI {
			cl := a.CreateOAuthTransport(a.AccessToken)
			cfg := rbac.Config{
				HTTPClient: cl,
				Debug:      a.Logger.DebugEnabled(),
				BaseURL:    a.ConsoleURL,
			}
			return rbac.NewPrincipalAPIClient(&cfg)
		},
	}
	return rbacAPI
}

// wraps the HTTP client with an OAuth2 Transport layer to provide automatic token refreshing
func (a *defaultAPI) CreateOAuthTransport(accessToken string) *http.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{
			AccessToken: accessToken,
		},
	)

	return &http.Client{
		Transport: &oauth2.Transport{
			Base:   a.HTTPClient.Transport,
			Source: oauth2.ReuseTokenSource(nil, ts),
		},
	}
}

func (a *defaultAPI) createOCMConnection(clusterMgmtApiUrl, accessToken string) (*ocmSdkClient.Connection, func(), error) {
	if clusterMgmtApiUrl == "" {
		clusterMgmtApiUrl = build.ProductionAPIURL
	}
	if accessToken == "" {
		accessToken = a.AccessToken
	}
	connection, err := ocmSdkClient.NewConnectionBuilder().
		URL(clusterMgmtApiUrl).
		Tokens(accessToken).
		Build()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't build connection: %v\n", err)
		return nil, nil, err
	}
	return connection, func() {
		_ = connection.Close()
	}, nil
}

// create an OCM clustermgmt client
func (a *defaultAPI) OCMClustermgmt(clusterMgmtApiUrl, accessToken string) (*ocmclustersmgmtv1.Client, func(), error) {
	connection, closeConnection, err := a.createOCMConnection(clusterMgmtApiUrl, accessToken)
	if err != nil {
		return nil, nil, err
	}
	return connection.ClustersMgmt().V1(), func() {
		closeConnection()
	}, nil
}
