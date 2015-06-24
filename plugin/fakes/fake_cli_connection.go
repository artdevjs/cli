// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry/cli/plugin"
	"github.com/cloudfoundry/cli/plugin/models"
)

type FakeCliConnection struct {
	CliCommandWithoutTerminalOutputStub        func(args ...string) ([]string, error)
	cliCommandWithoutTerminalOutputMutex       sync.RWMutex
	cliCommandWithoutTerminalOutputArgsForCall []struct {
		args []string
	}
	cliCommandWithoutTerminalOutputReturns struct {
		result1 []string
		result2 error
	}
	CliCommandStub        func(args ...string) ([]string, error)
	cliCommandMutex       sync.RWMutex
	cliCommandArgsForCall []struct {
		args []string
	}
	cliCommandReturns struct {
		result1 []string
		result2 error
	}
	GetCurrentOrgStub        func() (plugin_models.GetOrgs_Model, error)
	getCurrentOrgMutex       sync.RWMutex
	getCurrentOrgArgsForCall []struct{}
	getCurrentOrgReturns     struct {
		result1 plugin_models.GetOrgs_Model
		result2 error
	}
	GetCurrentSpaceStub        func() (plugin_models.GetSpaces_Model, error)
	getCurrentSpaceMutex       sync.RWMutex
	getCurrentSpaceArgsForCall []struct{}
	getCurrentSpaceReturns     struct {
		result1 plugin_models.GetSpaces_Model
		result2 error
	}
	UsernameStub        func() (string, error)
	usernameMutex       sync.RWMutex
	usernameArgsForCall []struct{}
	usernameReturns     struct {
		result1 string
		result2 error
	}
	UserGuidStub        func() (string, error)
	userGuidMutex       sync.RWMutex
	userGuidArgsForCall []struct{}
	userGuidReturns     struct {
		result1 string
		result2 error
	}
	UserEmailStub        func() (string, error)
	userEmailMutex       sync.RWMutex
	userEmailArgsForCall []struct{}
	userEmailReturns     struct {
		result1 string
		result2 error
	}
	IsLoggedInStub        func() (bool, error)
	isLoggedInMutex       sync.RWMutex
	isLoggedInArgsForCall []struct{}
	isLoggedInReturns     struct {
		result1 bool
		result2 error
	}
	IsSSLDisabledStub        func() (bool, error)
	isSSLDisabledMutex       sync.RWMutex
	isSSLDisabledArgsForCall []struct{}
	isSSLDisabledReturns     struct {
		result1 bool
		result2 error
	}
	HasOrganizationStub        func() (bool, error)
	hasOrganizationMutex       sync.RWMutex
	hasOrganizationArgsForCall []struct{}
	hasOrganizationReturns     struct {
		result1 bool
		result2 error
	}
	HasSpaceStub        func() (bool, error)
	hasSpaceMutex       sync.RWMutex
	hasSpaceArgsForCall []struct{}
	hasSpaceReturns     struct {
		result1 bool
		result2 error
	}
	ApiEndpointStub        func() (string, error)
	apiEndpointMutex       sync.RWMutex
	apiEndpointArgsForCall []struct{}
	apiEndpointReturns     struct {
		result1 string
		result2 error
	}
	ApiVersionStub        func() (string, error)
	apiVersionMutex       sync.RWMutex
	apiVersionArgsForCall []struct{}
	apiVersionReturns     struct {
		result1 string
		result2 error
	}
	HasAPIEndpointStub        func() (bool, error)
	hasAPIEndpointMutex       sync.RWMutex
	hasAPIEndpointArgsForCall []struct{}
	hasAPIEndpointReturns     struct {
		result1 bool
		result2 error
	}
	LoggregatorEndpointStub        func() (string, error)
	loggregatorEndpointMutex       sync.RWMutex
	loggregatorEndpointArgsForCall []struct{}
	loggregatorEndpointReturns     struct {
		result1 string
		result2 error
	}
	DopplerEndpointStub        func() (string, error)
	dopplerEndpointMutex       sync.RWMutex
	dopplerEndpointArgsForCall []struct{}
	dopplerEndpointReturns     struct {
		result1 string
		result2 error
	}
	AccessTokenStub        func() (string, error)
	accessTokenMutex       sync.RWMutex
	accessTokenArgsForCall []struct{}
	accessTokenReturns     struct {
		result1 string
		result2 error
	}
	GetAppStub        func(string) (plugin_models.GetAppModel, error)
	getAppMutex       sync.RWMutex
	getAppArgsForCall []struct {
		arg1 string
	}
	getAppReturns struct {
		result1 plugin_models.GetAppModel
		result2 error
	}
	GetAppsStub        func() ([]plugin_models.GetAppsModel, error)
	getAppsMutex       sync.RWMutex
	getAppsArgsForCall []struct{}
	getAppsReturns     struct {
		result1 []plugin_models.GetAppsModel
		result2 error
	}
	GetOrgsStub        func() ([]plugin_models.GetOrgs_Model, error)
	getOrgsMutex       sync.RWMutex
	getOrgsArgsForCall []struct{}
	getOrgsReturns     struct {
		result1 []plugin_models.GetOrgs_Model
		result2 error
	}
	GetSpacesStub        func() ([]plugin_models.GetSpaces_Model, error)
	getSpacesMutex       sync.RWMutex
	getSpacesArgsForCall []struct{}
	getSpacesReturns     struct {
		result1 []plugin_models.GetSpaces_Model
		result2 error
	}
	GetOrgUsersStub        func(string, ...string) ([]plugin_models.User, error)
	getOrgUsersMutex       sync.RWMutex
	getOrgUsersArgsForCall []struct {
		arg1 string
		arg2 []string
	}
	getOrgUsersReturns struct {
		result1 []plugin_models.User
		result2 error
	}
	GetSpaceUsersStub        func(string, string) ([]plugin_models.User, error)
	getSpaceUsersMutex       sync.RWMutex
	getSpaceUsersArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getSpaceUsersReturns struct {
		result1 []plugin_models.User
		result2 error
	}
	GetServicesStub        func() ([]plugin_models.ServiceInstance, error)
	getServicesMutex       sync.RWMutex
	getServicesArgsForCall []struct{}
	getServicesReturns     struct {
		result1 []plugin_models.ServiceInstance
		result2 error
	}
	GetOrgStub        func(string) (plugin_models.Organization, error)
	getOrgMutex       sync.RWMutex
	getOrgArgsForCall []struct {
		arg1 string
	}
	getOrgReturns struct {
		result1 plugin_models.Organization
		result2 error
	}
	GetSpaceStub        func(string) (plugin_models.Space, error)
	getSpaceMutex       sync.RWMutex
	getSpaceArgsForCall []struct {
		arg1 string
	}
	getSpaceReturns struct {
		result1 plugin_models.Space
		result2 error
	}
}

func (fake *FakeCliConnection) CliCommandWithoutTerminalOutput(args ...string) ([]string, error) {
	fake.cliCommandWithoutTerminalOutputMutex.Lock()
	fake.cliCommandWithoutTerminalOutputArgsForCall = append(fake.cliCommandWithoutTerminalOutputArgsForCall, struct {
		args []string
	}{args})
	fake.cliCommandWithoutTerminalOutputMutex.Unlock()
	if fake.CliCommandWithoutTerminalOutputStub != nil {
		return fake.CliCommandWithoutTerminalOutputStub(args...)
	} else {
		return fake.cliCommandWithoutTerminalOutputReturns.result1, fake.cliCommandWithoutTerminalOutputReturns.result2
	}
}

func (fake *FakeCliConnection) CliCommandWithoutTerminalOutputCallCount() int {
	fake.cliCommandWithoutTerminalOutputMutex.RLock()
	defer fake.cliCommandWithoutTerminalOutputMutex.RUnlock()
	return len(fake.cliCommandWithoutTerminalOutputArgsForCall)
}

func (fake *FakeCliConnection) CliCommandWithoutTerminalOutputArgsForCall(i int) []string {
	fake.cliCommandWithoutTerminalOutputMutex.RLock()
	defer fake.cliCommandWithoutTerminalOutputMutex.RUnlock()
	return fake.cliCommandWithoutTerminalOutputArgsForCall[i].args
}

func (fake *FakeCliConnection) CliCommandWithoutTerminalOutputReturns(result1 []string, result2 error) {
	fake.CliCommandWithoutTerminalOutputStub = nil
	fake.cliCommandWithoutTerminalOutputReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) CliCommand(args ...string) ([]string, error) {
	fake.cliCommandMutex.Lock()
	fake.cliCommandArgsForCall = append(fake.cliCommandArgsForCall, struct {
		args []string
	}{args})
	fake.cliCommandMutex.Unlock()
	if fake.CliCommandStub != nil {
		return fake.CliCommandStub(args...)
	} else {
		return fake.cliCommandReturns.result1, fake.cliCommandReturns.result2
	}
}

func (fake *FakeCliConnection) CliCommandCallCount() int {
	fake.cliCommandMutex.RLock()
	defer fake.cliCommandMutex.RUnlock()
	return len(fake.cliCommandArgsForCall)
}

func (fake *FakeCliConnection) CliCommandArgsForCall(i int) []string {
	fake.cliCommandMutex.RLock()
	defer fake.cliCommandMutex.RUnlock()
	return fake.cliCommandArgsForCall[i].args
}

func (fake *FakeCliConnection) CliCommandReturns(result1 []string, result2 error) {
	fake.CliCommandStub = nil
	fake.cliCommandReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) GetCurrentOrg() (plugin_models.GetOrgs_Model, error) {
	fake.getCurrentOrgMutex.Lock()
	fake.getCurrentOrgArgsForCall = append(fake.getCurrentOrgArgsForCall, struct{}{})
	fake.getCurrentOrgMutex.Unlock()
	if fake.GetCurrentOrgStub != nil {
		return fake.GetCurrentOrgStub()
	} else {
		return fake.getCurrentOrgReturns.result1, fake.getCurrentOrgReturns.result2
	}
}

func (fake *FakeCliConnection) GetCurrentOrgCallCount() int {
	fake.getCurrentOrgMutex.RLock()
	defer fake.getCurrentOrgMutex.RUnlock()
	return len(fake.getCurrentOrgArgsForCall)
}

func (fake *FakeCliConnection) GetCurrentOrgReturns(result1 plugin_models.GetOrgs_Model, result2 error) {
	fake.GetCurrentOrgStub = nil
	fake.getCurrentOrgReturns = struct {
		result1 plugin_models.GetOrgs_Model
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) GetCurrentSpace() (plugin_models.GetSpaces_Model, error) {
	fake.getCurrentSpaceMutex.Lock()
	fake.getCurrentSpaceArgsForCall = append(fake.getCurrentSpaceArgsForCall, struct{}{})
	fake.getCurrentSpaceMutex.Unlock()
	if fake.GetCurrentSpaceStub != nil {
		return fake.GetCurrentSpaceStub()
	} else {
		return fake.getCurrentSpaceReturns.result1, fake.getCurrentSpaceReturns.result2
	}
}

func (fake *FakeCliConnection) GetCurrentSpaceCallCount() int {
	fake.getCurrentSpaceMutex.RLock()
	defer fake.getCurrentSpaceMutex.RUnlock()
	return len(fake.getCurrentSpaceArgsForCall)
}

func (fake *FakeCliConnection) GetCurrentSpaceReturns(result1 plugin_models.GetSpaces_Model, result2 error) {
	fake.GetCurrentSpaceStub = nil
	fake.getCurrentSpaceReturns = struct {
		result1 plugin_models.GetSpaces_Model
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) Username() (string, error) {
	fake.usernameMutex.Lock()
	fake.usernameArgsForCall = append(fake.usernameArgsForCall, struct{}{})
	fake.usernameMutex.Unlock()
	if fake.UsernameStub != nil {
		return fake.UsernameStub()
	} else {
		return fake.usernameReturns.result1, fake.usernameReturns.result2
	}
}

func (fake *FakeCliConnection) UsernameCallCount() int {
	fake.usernameMutex.RLock()
	defer fake.usernameMutex.RUnlock()
	return len(fake.usernameArgsForCall)
}

func (fake *FakeCliConnection) UsernameReturns(result1 string, result2 error) {
	fake.UsernameStub = nil
	fake.usernameReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) UserGuid() (string, error) {
	fake.userGuidMutex.Lock()
	fake.userGuidArgsForCall = append(fake.userGuidArgsForCall, struct{}{})
	fake.userGuidMutex.Unlock()
	if fake.UserGuidStub != nil {
		return fake.UserGuidStub()
	} else {
		return fake.userGuidReturns.result1, fake.userGuidReturns.result2
	}
}

func (fake *FakeCliConnection) UserGuidCallCount() int {
	fake.userGuidMutex.RLock()
	defer fake.userGuidMutex.RUnlock()
	return len(fake.userGuidArgsForCall)
}

func (fake *FakeCliConnection) UserGuidReturns(result1 string, result2 error) {
	fake.UserGuidStub = nil
	fake.userGuidReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) UserEmail() (string, error) {
	fake.userEmailMutex.Lock()
	fake.userEmailArgsForCall = append(fake.userEmailArgsForCall, struct{}{})
	fake.userEmailMutex.Unlock()
	if fake.UserEmailStub != nil {
		return fake.UserEmailStub()
	} else {
		return fake.userEmailReturns.result1, fake.userEmailReturns.result2
	}
}

func (fake *FakeCliConnection) UserEmailCallCount() int {
	fake.userEmailMutex.RLock()
	defer fake.userEmailMutex.RUnlock()
	return len(fake.userEmailArgsForCall)
}

func (fake *FakeCliConnection) UserEmailReturns(result1 string, result2 error) {
	fake.UserEmailStub = nil
	fake.userEmailReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) IsLoggedIn() (bool, error) {
	fake.isLoggedInMutex.Lock()
	fake.isLoggedInArgsForCall = append(fake.isLoggedInArgsForCall, struct{}{})
	fake.isLoggedInMutex.Unlock()
	if fake.IsLoggedInStub != nil {
		return fake.IsLoggedInStub()
	} else {
		return fake.isLoggedInReturns.result1, fake.isLoggedInReturns.result2
	}
}

func (fake *FakeCliConnection) IsLoggedInCallCount() int {
	fake.isLoggedInMutex.RLock()
	defer fake.isLoggedInMutex.RUnlock()
	return len(fake.isLoggedInArgsForCall)
}

func (fake *FakeCliConnection) IsLoggedInReturns(result1 bool, result2 error) {
	fake.IsLoggedInStub = nil
	fake.isLoggedInReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) IsSSLDisabled() (bool, error) {
	fake.isSSLDisabledMutex.Lock()
	fake.isSSLDisabledArgsForCall = append(fake.isSSLDisabledArgsForCall, struct{}{})
	fake.isSSLDisabledMutex.Unlock()
	if fake.IsSSLDisabledStub != nil {
		return fake.IsSSLDisabledStub()
	} else {
		return fake.isSSLDisabledReturns.result1, fake.isSSLDisabledReturns.result2
	}
}

func (fake *FakeCliConnection) IsSSLDisabledCallCount() int {
	fake.isSSLDisabledMutex.RLock()
	defer fake.isSSLDisabledMutex.RUnlock()
	return len(fake.isSSLDisabledArgsForCall)
}

func (fake *FakeCliConnection) IsSSLDisabledReturns(result1 bool, result2 error) {
	fake.IsSSLDisabledStub = nil
	fake.isSSLDisabledReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) HasOrganization() (bool, error) {
	fake.hasOrganizationMutex.Lock()
	fake.hasOrganizationArgsForCall = append(fake.hasOrganizationArgsForCall, struct{}{})
	fake.hasOrganizationMutex.Unlock()
	if fake.HasOrganizationStub != nil {
		return fake.HasOrganizationStub()
	} else {
		return fake.hasOrganizationReturns.result1, fake.hasOrganizationReturns.result2
	}
}

func (fake *FakeCliConnection) HasOrganizationCallCount() int {
	fake.hasOrganizationMutex.RLock()
	defer fake.hasOrganizationMutex.RUnlock()
	return len(fake.hasOrganizationArgsForCall)
}

func (fake *FakeCliConnection) HasOrganizationReturns(result1 bool, result2 error) {
	fake.HasOrganizationStub = nil
	fake.hasOrganizationReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) HasSpace() (bool, error) {
	fake.hasSpaceMutex.Lock()
	fake.hasSpaceArgsForCall = append(fake.hasSpaceArgsForCall, struct{}{})
	fake.hasSpaceMutex.Unlock()
	if fake.HasSpaceStub != nil {
		return fake.HasSpaceStub()
	} else {
		return fake.hasSpaceReturns.result1, fake.hasSpaceReturns.result2
	}
}

func (fake *FakeCliConnection) HasSpaceCallCount() int {
	fake.hasSpaceMutex.RLock()
	defer fake.hasSpaceMutex.RUnlock()
	return len(fake.hasSpaceArgsForCall)
}

func (fake *FakeCliConnection) HasSpaceReturns(result1 bool, result2 error) {
	fake.HasSpaceStub = nil
	fake.hasSpaceReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) ApiEndpoint() (string, error) {
	fake.apiEndpointMutex.Lock()
	fake.apiEndpointArgsForCall = append(fake.apiEndpointArgsForCall, struct{}{})
	fake.apiEndpointMutex.Unlock()
	if fake.ApiEndpointStub != nil {
		return fake.ApiEndpointStub()
	} else {
		return fake.apiEndpointReturns.result1, fake.apiEndpointReturns.result2
	}
}

func (fake *FakeCliConnection) ApiEndpointCallCount() int {
	fake.apiEndpointMutex.RLock()
	defer fake.apiEndpointMutex.RUnlock()
	return len(fake.apiEndpointArgsForCall)
}

func (fake *FakeCliConnection) ApiEndpointReturns(result1 string, result2 error) {
	fake.ApiEndpointStub = nil
	fake.apiEndpointReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) ApiVersion() (string, error) {
	fake.apiVersionMutex.Lock()
	fake.apiVersionArgsForCall = append(fake.apiVersionArgsForCall, struct{}{})
	fake.apiVersionMutex.Unlock()
	if fake.ApiVersionStub != nil {
		return fake.ApiVersionStub()
	} else {
		return fake.apiVersionReturns.result1, fake.apiVersionReturns.result2
	}
}

func (fake *FakeCliConnection) ApiVersionCallCount() int {
	fake.apiVersionMutex.RLock()
	defer fake.apiVersionMutex.RUnlock()
	return len(fake.apiVersionArgsForCall)
}

func (fake *FakeCliConnection) ApiVersionReturns(result1 string, result2 error) {
	fake.ApiVersionStub = nil
	fake.apiVersionReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) HasAPIEndpoint() (bool, error) {
	fake.hasAPIEndpointMutex.Lock()
	fake.hasAPIEndpointArgsForCall = append(fake.hasAPIEndpointArgsForCall, struct{}{})
	fake.hasAPIEndpointMutex.Unlock()
	if fake.HasAPIEndpointStub != nil {
		return fake.HasAPIEndpointStub()
	} else {
		return fake.hasAPIEndpointReturns.result1, fake.hasAPIEndpointReturns.result2
	}
}

func (fake *FakeCliConnection) HasAPIEndpointCallCount() int {
	fake.hasAPIEndpointMutex.RLock()
	defer fake.hasAPIEndpointMutex.RUnlock()
	return len(fake.hasAPIEndpointArgsForCall)
}

func (fake *FakeCliConnection) HasAPIEndpointReturns(result1 bool, result2 error) {
	fake.HasAPIEndpointStub = nil
	fake.hasAPIEndpointReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) LoggregatorEndpoint() (string, error) {
	fake.loggregatorEndpointMutex.Lock()
	fake.loggregatorEndpointArgsForCall = append(fake.loggregatorEndpointArgsForCall, struct{}{})
	fake.loggregatorEndpointMutex.Unlock()
	if fake.LoggregatorEndpointStub != nil {
		return fake.LoggregatorEndpointStub()
	} else {
		return fake.loggregatorEndpointReturns.result1, fake.loggregatorEndpointReturns.result2
	}
}

func (fake *FakeCliConnection) LoggregatorEndpointCallCount() int {
	fake.loggregatorEndpointMutex.RLock()
	defer fake.loggregatorEndpointMutex.RUnlock()
	return len(fake.loggregatorEndpointArgsForCall)
}

func (fake *FakeCliConnection) LoggregatorEndpointReturns(result1 string, result2 error) {
	fake.LoggregatorEndpointStub = nil
	fake.loggregatorEndpointReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) DopplerEndpoint() (string, error) {
	fake.dopplerEndpointMutex.Lock()
	fake.dopplerEndpointArgsForCall = append(fake.dopplerEndpointArgsForCall, struct{}{})
	fake.dopplerEndpointMutex.Unlock()
	if fake.DopplerEndpointStub != nil {
		return fake.DopplerEndpointStub()
	} else {
		return fake.dopplerEndpointReturns.result1, fake.dopplerEndpointReturns.result2
	}
}

func (fake *FakeCliConnection) DopplerEndpointCallCount() int {
	fake.dopplerEndpointMutex.RLock()
	defer fake.dopplerEndpointMutex.RUnlock()
	return len(fake.dopplerEndpointArgsForCall)
}

func (fake *FakeCliConnection) DopplerEndpointReturns(result1 string, result2 error) {
	fake.DopplerEndpointStub = nil
	fake.dopplerEndpointReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) AccessToken() (string, error) {
	fake.accessTokenMutex.Lock()
	fake.accessTokenArgsForCall = append(fake.accessTokenArgsForCall, struct{}{})
	fake.accessTokenMutex.Unlock()
	if fake.AccessTokenStub != nil {
		return fake.AccessTokenStub()
	} else {
		return fake.accessTokenReturns.result1, fake.accessTokenReturns.result2
	}
}

func (fake *FakeCliConnection) AccessTokenCallCount() int {
	fake.accessTokenMutex.RLock()
	defer fake.accessTokenMutex.RUnlock()
	return len(fake.accessTokenArgsForCall)
}

func (fake *FakeCliConnection) AccessTokenReturns(result1 string, result2 error) {
	fake.AccessTokenStub = nil
	fake.accessTokenReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) GetApp(arg1 string) (plugin_models.GetAppModel, error) {
	fake.getAppMutex.Lock()
	fake.getAppArgsForCall = append(fake.getAppArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.getAppMutex.Unlock()
	if fake.GetAppStub != nil {
		return fake.GetAppStub(arg1)
	} else {
		return fake.getAppReturns.result1, fake.getAppReturns.result2
	}
}

func (fake *FakeCliConnection) GetAppCallCount() int {
	fake.getAppMutex.RLock()
	defer fake.getAppMutex.RUnlock()
	return len(fake.getAppArgsForCall)
}

func (fake *FakeCliConnection) GetAppArgsForCall(i int) string {
	fake.getAppMutex.RLock()
	defer fake.getAppMutex.RUnlock()
	return fake.getAppArgsForCall[i].arg1
}

func (fake *FakeCliConnection) GetAppReturns(result1 plugin_models.GetAppModel, result2 error) {
	fake.GetAppStub = nil
	fake.getAppReturns = struct {
		result1 plugin_models.GetAppModel
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) GetApps() ([]plugin_models.GetAppsModel, error) {
	fake.getAppsMutex.Lock()
	fake.getAppsArgsForCall = append(fake.getAppsArgsForCall, struct{}{})
	fake.getAppsMutex.Unlock()
	if fake.GetAppsStub != nil {
		return fake.GetAppsStub()
	} else {
		return fake.getAppsReturns.result1, fake.getAppsReturns.result2
	}
}

func (fake *FakeCliConnection) GetAppsCallCount() int {
	fake.getAppsMutex.RLock()
	defer fake.getAppsMutex.RUnlock()
	return len(fake.getAppsArgsForCall)
}

func (fake *FakeCliConnection) GetAppsReturns(result1 []plugin_models.GetAppsModel, result2 error) {
	fake.GetAppsStub = nil
	fake.getAppsReturns = struct {
		result1 []plugin_models.GetAppsModel
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) GetOrgs() ([]plugin_models.GetOrgs_Model, error) {
	fake.getOrgsMutex.Lock()
	fake.getOrgsArgsForCall = append(fake.getOrgsArgsForCall, struct{}{})
	fake.getOrgsMutex.Unlock()
	if fake.GetOrgsStub != nil {
		return fake.GetOrgsStub()
	} else {
		return fake.getOrgsReturns.result1, fake.getOrgsReturns.result2
	}
}

func (fake *FakeCliConnection) GetOrgsCallCount() int {
	fake.getOrgsMutex.RLock()
	defer fake.getOrgsMutex.RUnlock()
	return len(fake.getOrgsArgsForCall)
}

func (fake *FakeCliConnection) GetOrgsReturns(result1 []plugin_models.GetOrgs_Model, result2 error) {
	fake.GetOrgsStub = nil
	fake.getOrgsReturns = struct {
		result1 []plugin_models.GetOrgs_Model
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) GetSpaces() ([]plugin_models.GetSpaces_Model, error) {
	fake.getSpacesMutex.Lock()
	fake.getSpacesArgsForCall = append(fake.getSpacesArgsForCall, struct{}{})
	fake.getSpacesMutex.Unlock()
	if fake.GetSpacesStub != nil {
		return fake.GetSpacesStub()
	} else {
		return fake.getSpacesReturns.result1, fake.getSpacesReturns.result2
	}
}

func (fake *FakeCliConnection) GetSpacesCallCount() int {
	fake.getSpacesMutex.RLock()
	defer fake.getSpacesMutex.RUnlock()
	return len(fake.getSpacesArgsForCall)
}

func (fake *FakeCliConnection) GetSpacesReturns(result1 []plugin_models.GetSpaces_Model, result2 error) {
	fake.GetSpacesStub = nil
	fake.getSpacesReturns = struct {
		result1 []plugin_models.GetSpaces_Model
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) GetOrgUsers(arg1 string, arg2 ...string) ([]plugin_models.User, error) {
	fake.getOrgUsersMutex.Lock()
	fake.getOrgUsersArgsForCall = append(fake.getOrgUsersArgsForCall, struct {
		arg1 string
		arg2 []string
	}{arg1, arg2})
	fake.getOrgUsersMutex.Unlock()
	if fake.GetOrgUsersStub != nil {
		return fake.GetOrgUsersStub(arg1, arg2...)
	} else {
		return fake.getOrgUsersReturns.result1, fake.getOrgUsersReturns.result2
	}
}

func (fake *FakeCliConnection) GetOrgUsersCallCount() int {
	fake.getOrgUsersMutex.RLock()
	defer fake.getOrgUsersMutex.RUnlock()
	return len(fake.getOrgUsersArgsForCall)
}

func (fake *FakeCliConnection) GetOrgUsersArgsForCall(i int) (string, []string) {
	fake.getOrgUsersMutex.RLock()
	defer fake.getOrgUsersMutex.RUnlock()
	return fake.getOrgUsersArgsForCall[i].arg1, fake.getOrgUsersArgsForCall[i].arg2
}

func (fake *FakeCliConnection) GetOrgUsersReturns(result1 []plugin_models.User, result2 error) {
	fake.GetOrgUsersStub = nil
	fake.getOrgUsersReturns = struct {
		result1 []plugin_models.User
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) GetSpaceUsers(arg1 string, arg2 string) ([]plugin_models.User, error) {
	fake.getSpaceUsersMutex.Lock()
	fake.getSpaceUsersArgsForCall = append(fake.getSpaceUsersArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.getSpaceUsersMutex.Unlock()
	if fake.GetSpaceUsersStub != nil {
		return fake.GetSpaceUsersStub(arg1, arg2)
	} else {
		return fake.getSpaceUsersReturns.result1, fake.getSpaceUsersReturns.result2
	}
}

func (fake *FakeCliConnection) GetSpaceUsersCallCount() int {
	fake.getSpaceUsersMutex.RLock()
	defer fake.getSpaceUsersMutex.RUnlock()
	return len(fake.getSpaceUsersArgsForCall)
}

func (fake *FakeCliConnection) GetSpaceUsersArgsForCall(i int) (string, string) {
	fake.getSpaceUsersMutex.RLock()
	defer fake.getSpaceUsersMutex.RUnlock()
	return fake.getSpaceUsersArgsForCall[i].arg1, fake.getSpaceUsersArgsForCall[i].arg2
}

func (fake *FakeCliConnection) GetSpaceUsersReturns(result1 []plugin_models.User, result2 error) {
	fake.GetSpaceUsersStub = nil
	fake.getSpaceUsersReturns = struct {
		result1 []plugin_models.User
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) GetServices() ([]plugin_models.ServiceInstance, error) {
	fake.getServicesMutex.Lock()
	fake.getServicesArgsForCall = append(fake.getServicesArgsForCall, struct{}{})
	fake.getServicesMutex.Unlock()
	if fake.GetServicesStub != nil {
		return fake.GetServicesStub()
	} else {
		return fake.getServicesReturns.result1, fake.getServicesReturns.result2
	}
}

func (fake *FakeCliConnection) GetServicesCallCount() int {
	fake.getServicesMutex.RLock()
	defer fake.getServicesMutex.RUnlock()
	return len(fake.getServicesArgsForCall)
}

func (fake *FakeCliConnection) GetServicesReturns(result1 []plugin_models.ServiceInstance, result2 error) {
	fake.GetServicesStub = nil
	fake.getServicesReturns = struct {
		result1 []plugin_models.ServiceInstance
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) GetOrg(arg1 string) (plugin_models.Organization, error) {
	fake.getOrgMutex.Lock()
	fake.getOrgArgsForCall = append(fake.getOrgArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.getOrgMutex.Unlock()
	if fake.GetOrgStub != nil {
		return fake.GetOrgStub(arg1)
	} else {
		return fake.getOrgReturns.result1, fake.getOrgReturns.result2
	}
}

func (fake *FakeCliConnection) GetOrgCallCount() int {
	fake.getOrgMutex.RLock()
	defer fake.getOrgMutex.RUnlock()
	return len(fake.getOrgArgsForCall)
}

func (fake *FakeCliConnection) GetOrgArgsForCall(i int) string {
	fake.getOrgMutex.RLock()
	defer fake.getOrgMutex.RUnlock()
	return fake.getOrgArgsForCall[i].arg1
}

func (fake *FakeCliConnection) GetOrgReturns(result1 plugin_models.Organization, result2 error) {
	fake.GetOrgStub = nil
	fake.getOrgReturns = struct {
		result1 plugin_models.Organization
		result2 error
	}{result1, result2}
}

func (fake *FakeCliConnection) GetSpace(arg1 string) (plugin_models.Space, error) {
	fake.getSpaceMutex.Lock()
	fake.getSpaceArgsForCall = append(fake.getSpaceArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.getSpaceMutex.Unlock()
	if fake.GetSpaceStub != nil {
		return fake.GetSpaceStub(arg1)
	} else {
		return fake.getSpaceReturns.result1, fake.getSpaceReturns.result2
	}
}

func (fake *FakeCliConnection) GetSpaceCallCount() int {
	fake.getSpaceMutex.RLock()
	defer fake.getSpaceMutex.RUnlock()
	return len(fake.getSpaceArgsForCall)
}

func (fake *FakeCliConnection) GetSpaceArgsForCall(i int) string {
	fake.getSpaceMutex.RLock()
	defer fake.getSpaceMutex.RUnlock()
	return fake.getSpaceArgsForCall[i].arg1
}

func (fake *FakeCliConnection) GetSpaceReturns(result1 plugin_models.Space, result2 error) {
	fake.GetSpaceStub = nil
	fake.getSpaceReturns = struct {
		result1 plugin_models.Space
		result2 error
	}{result1, result2}
}

var _ plugin.CliConnection = new(FakeCliConnection)
