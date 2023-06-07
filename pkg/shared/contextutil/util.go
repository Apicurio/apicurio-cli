package contextutil

import (
	"context"

	"github.com/apicurio/apicurio-cli/pkg/core/localize"
	"github.com/apicurio/apicurio-cli/pkg/core/servicecontext"

	"github.com/apicurio/apicurio-cli/pkg/shared/factory"
	srsmgmtv1errors "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/registrymgmt/apiv1/error"

	registrymgmtclient "github.com/redhat-developer/app-services-sdk-core/app-services-sdk-go/registrymgmt/apiv1/client"
)

// GetContext returns the services associated with the context
func GetContext(svcContext *servicecontext.Context, localizer localize.Localizer, ctxName string) (*servicecontext.ServiceConfig, error) {

	ctx, ok := svcContext.Contexts[ctxName]
	if !ok {
		return nil, localizer.MustLocalizeError("context.common.error.context.notFound", localize.NewEntry("Name", svcContext.CurrentContext))
	}

	return &ctx, nil

}

// GetCurrentContext returns the name of the currently selected context
func GetCurrentContext(svcContext *servicecontext.Context, localizer localize.Localizer) (*servicecontext.ServiceConfig, error) {

	if svcContext.CurrentContext == "" {
		return nil, localizer.MustLocalizeError("context.common.error.notSet")
	}

	currCtx, ok := svcContext.Contexts[svcContext.CurrentContext]
	if !ok {
		return nil, localizer.MustLocalizeError("context.common.error.context.notFound", localize.NewEntry("Name", svcContext.CurrentContext))
	}

	return &currCtx, nil
}

// GetCurrentRegistryInstance returns the Service Registry instance set in the currently selected context
func GetCurrentRegistryInstance(f *factory.Factory) (*registrymgmtclient.Registry, error) {

	svcContext, err := f.ServiceContext.Load()
	if err != nil {
		return nil, err
	}

	currCtx, err := GetCurrentContext(svcContext, f.Localizer)
	if err != nil {
		return nil, err
	}

	return GetRegistryForServiceConfig(currCtx, f)

}

func GetRegistryForServiceConfig(currCtx *servicecontext.ServiceConfig, f *factory.Factory) (*registrymgmtclient.Registry, error) {
	conn, err := f.Connection()
	if err != nil {
		return nil, err
	}

	if currCtx.ServiceRegistryID == "" {
		return nil, f.Localizer.MustLocalizeError("context.common.error.noRegistryID")
	}

	registryInstance, _, err := conn.API().ServiceRegistryMgmt().GetRegistry(context.Background(), currCtx.ServiceRegistryID).Execute()
	if srsmgmtv1errors.IsAPIError(err, srsmgmtv1errors.ERROR_2) {
		return nil, f.Localizer.MustLocalizeError("context.common.error.registry.notFound")
	}

	return &registryInstance, nil
}
