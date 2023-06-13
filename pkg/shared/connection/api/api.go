package api

import (
	"net/http"
	"net/url"

	ocmclustersmgmtv1 "github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1"

	"github.com/apicurio/apicurio-cli/pkg/api/generic"
	"github.com/apicurio/apicurio-cli/pkg/api/rbac"
	"github.com/apicurio/apicurio-cli/pkg/core/logging"

	registryinstanceclient "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/registryinstance/apiv1internal/client"
	registrymgmtclient "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/registrymgmt/apiv1/client"
	svcacctmgmtclient "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/serviceaccountmgmt/apiv1/client"
)

type API interface {
	ServiceRegistryMgmt() registrymgmtclient.RegistriesApi
	ServiceAccountMgmt() svcacctmgmtclient.ServiceAccountsApi
	ServiceRegistryInstance(instanceID string) (*registryinstanceclient.APIClient, *registrymgmtclient.Registry, error)
	RBAC() rbac.RbacAPI
	GenericAPI() generic.GenericAPI
	GetConfig() Config
	OCMClustermgmt(apiGateway, accessToken string) (*ocmclustersmgmtv1.Client, func(), error)
}

type Config struct {
	AccessToken string
	ApiURL      *url.URL
	AuthURL     *url.URL
	ConsoleURL  *url.URL
	UserAgent   string
	HTTPClient  *http.Client
	Logger      logging.Logger
}
