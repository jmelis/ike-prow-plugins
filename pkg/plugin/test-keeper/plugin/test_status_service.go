package plugin

import (
	"fmt"

	"github.com/arquillian/ike-prow-plugins/pkg/github"
	"github.com/arquillian/ike-prow-plugins/pkg/log"
	"github.com/arquillian/ike-prow-plugins/pkg/scm"
)

type testStatusService struct {
	statusService scm.StatusService
}

func (gh *GitHubTestEventsHandler) newTestStatusService(log log.Logger, change scm.RepositoryChange) testStatusService {
	statusContext := github.StatusContext{BotName: "ike-plugins", PluginName: ProwPluginName}
	statusService := github.NewStatusService(gh.Client, log, change, statusContext)
	return testStatusService{statusService: statusService}
}

func (ts *testStatusService) testsExist() error {
	return ts.statusService.Success("There are some tests :)")
}

func (ts *testStatusService) noTests() error {
	return ts.statusService.Failure("No tests in this PR :(")
}

func (ts *testStatusService) okWithoutTests(approvedBy string) error {
	return ts.statusService.Success(fmt.Sprintf("PR is fine without tests says @%s", approvedBy))
}