package internal

import (
	"net/http"

	"github.com/tedsuo/rata"
)

const (
	AppInstances                          = "AppInstances"
	AppInstanceStats                      = "AppInstanceStats"
	AppRequest                            = "App"
	AppsFromRouteRequest                  = "AppsFromRoute"
	AppsRequest                           = "Apps"
	DeleteOrganizationRequest             = "DeleteOrganization"
	DeleteRouteRequest                    = "DeleteRoute"
	DeleteServiceBindingRequest           = "DeleteServiceBinding"
	InfoRequest                           = "Info"
	JobRequest                            = "Job"
	OrganizationsRequest                  = "Organizations"
	OrganizationQuotaRequest              = "OrganizationQuota"
	PrivateDomainRequest                  = "PrivateDomain"
	PrivateDomainsFromOrganizationRequest = "PrivateDomainsFromOrganization"
	RouteMappingsFromRouteRequest         = "RouteMappingsFromRoute"
	RoutesFromApplicationRequest          = "RoutesFromApplication"
	RoutesFromSpaceRequest                = "RoutesFromSpace"
	ServiceBindingsRequest                = "ServiceBindings"
	ServiceInstancesRequest               = "ServiceInstances"
	SharedDomainRequest                   = "SharedDomain"
	SharedDomainsRequest                  = "SharedDomains"
	SpaceServiceInstancesRequest          = "SpaceServiceInstances"
	SpacesRequest                         = "Spaces"
	StackRequest                          = "Stack"
	UpdateAppRequest                      = "UpdateApp"
	UsersRequest                          = "Users"
)

// APIRoutes is a list of routes used by the rata library to construct request
// URLs.
var APIRoutes = rata.Routes{
	{Path: "/v2/apps", Method: http.MethodGet, Name: AppsRequest},
	{Path: "/v2/apps/:app_guid", Method: http.MethodGet, Name: AppRequest},
	{Path: "/v2/apps/:app_guid", Method: http.MethodPut, Name: UpdateAppRequest},
	{Path: "/v2/apps/:app_guid/instances", Method: http.MethodGet, Name: AppInstances},
	{Path: "/v2/apps/:app_guid/routes", Method: http.MethodGet, Name: RoutesFromApplicationRequest},
	{Path: "/v2/apps/:app_guid/stats", Method: http.MethodGet, Name: AppInstanceStats},
	{Path: "/v2/info", Method: http.MethodGet, Name: InfoRequest},
	{Path: "/v2/jobs/:job_guid", Method: http.MethodGet, Name: JobRequest},
	{Path: "/v2/organizations", Method: http.MethodGet, Name: OrganizationsRequest},
	{Path: "/v2/organizations/:organization_guid", Method: http.MethodDelete, Name: DeleteOrganizationRequest},
	{Path: "/v2/organizations/:organization_guid/private_domains", Method: http.MethodGet, Name: PrivateDomainsFromOrganizationRequest},
	{Path: "/v2/private_domains/:private_domain_guid", Method: http.MethodGet, Name: PrivateDomainRequest},
	{Path: "/v2/quota_definitions/:organization_quota_guid", Method: http.MethodGet, Name: OrganizationQuotaRequest},
	{Path: "/v2/routes/:route_guid", Method: http.MethodDelete, Name: DeleteRouteRequest},
	{Path: "/v2/routes/:route_guid/apps", Method: http.MethodGet, Name: AppsFromRouteRequest},
	{Path: "/v2/routes/:route_guid/route_mappings", Method: http.MethodGet, Name: RouteMappingsFromRouteRequest},
	{Path: "/v2/service_bindings", Method: http.MethodGet, Name: ServiceBindingsRequest},
	{Path: "/v2/service_bindings/:service_binding_guid", Method: http.MethodDelete, Name: DeleteServiceBindingRequest},
	{Path: "/v2/service_instances", Method: http.MethodGet, Name: ServiceInstancesRequest},
	{Path: "/v2/shared_domains/:shared_domain_guid", Method: http.MethodGet, Name: SharedDomainRequest},
	{Path: "/v2/shared_domains", Method: http.MethodGet, Name: SharedDomainsRequest},
	{Path: "/v2/spaces", Method: http.MethodGet, Name: SpacesRequest},
	{Path: "/v2/spaces/:guid/service_instances", Method: http.MethodGet, Name: SpaceServiceInstancesRequest},
	{Path: "/v2/spaces/:space_guid/routes", Method: http.MethodGet, Name: RoutesFromSpaceRequest},
	{Path: "/v2/stacks/:stack_guid", Method: http.MethodGet, Name: StackRequest},
	{Path: "/v2/users", Method: http.MethodPost, Name: UsersRequest},
}
