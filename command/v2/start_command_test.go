package v2_test

import (
	"errors"
	"time"

	"code.cloudfoundry.org/cli/actor/sharedaction"
	"code.cloudfoundry.org/cli/actor/v2action"
	"code.cloudfoundry.org/cli/api/cloudcontroller/ccv2"
	"code.cloudfoundry.org/cli/command"
	"code.cloudfoundry.org/cli/command/commandfakes"
	"code.cloudfoundry.org/cli/command/v2"
	"code.cloudfoundry.org/cli/command/v2/shared"
	"code.cloudfoundry.org/cli/command/v2/v2fakes"
	"code.cloudfoundry.org/cli/util/configv3"
	"code.cloudfoundry.org/cli/util/ui"
	"github.com/cloudfoundry/bytefmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
)

var _ = Describe("Start Command", func() {
	var (
		cmd             v2.StartCommand
		testUI          *ui.UI
		fakeConfig      *commandfakes.FakeConfig
		fakeSharedActor *commandfakes.FakeSharedActor
		fakeActor       *v2fakes.FakeStartActor
		binaryName      string
		executeErr      error
	)

	BeforeEach(func() {
		testUI = ui.NewTestUI(nil, NewBuffer(), NewBuffer())
		fakeConfig = new(commandfakes.FakeConfig)
		fakeSharedActor = new(commandfakes.FakeSharedActor)
		fakeActor = new(v2fakes.FakeStartActor)

		cmd = v2.StartCommand{
			UI:          testUI,
			Config:      fakeConfig,
			SharedActor: fakeSharedActor,
			Actor:       fakeActor,
		}

		cmd.RequiredArgs.AppName = "some-app"

		binaryName = "faceman"
		fakeConfig.BinaryNameReturns(binaryName)

		// TODO: remove when experimental flag is removed
		fakeConfig.ExperimentalReturns(true)

		var err error
		testUI.TimezoneLocation, err = time.LoadLocation("America/Los_Angeles")
		Expect(err).NotTo(HaveOccurred())

		fakeActor.StartApplicationStub = func(app v2action.Application, client v2action.NOAAClient, config v2action.Config) (<-chan *v2action.LogMessage, <-chan error, <-chan string, <-chan error) {
			messages := make(chan *v2action.LogMessage)
			logErrs := make(chan error)
			warnings := make(chan string)
			errs := make(chan error)

			go func() {
				close(messages)
				close(logErrs)
				close(warnings)
				close(errs)
			}()

			return messages, logErrs, warnings, errs
		}
	})

	JustBeforeEach(func() {
		executeErr = cmd.Execute(nil)
	})

	// TODO: remove when experimental flag is removed
	It("Displays the experimental warning message", func() {
		Expect(testUI.Out).To(Say(command.ExperimentalWarning))
	})

	Context("when checking target fails", func() {
		BeforeEach(func() {
			fakeSharedActor.CheckTargetReturns(sharedaction.NotLoggedInError{BinaryName: binaryName})
		})

		It("returns an error if the check fails", func() {
			Expect(executeErr).To(MatchError(command.NotLoggedInError{BinaryName: "faceman"}))

			Expect(fakeSharedActor.CheckTargetCallCount()).To(Equal(1))
			_, checkTargetedOrg, checkTargetedSpace := fakeSharedActor.CheckTargetArgsForCall(0)
			Expect(checkTargetedOrg).To(BeTrue())
			Expect(checkTargetedSpace).To(BeTrue())
		})
	})

	Context("when the user is logged in, and org and space are targeted", func() {
		BeforeEach(func() {
			fakeConfig.HasTargetedOrganizationReturns(true)
			fakeConfig.TargetedOrganizationReturns(configv3.Organization{Name: "some-org"})
			fakeConfig.HasTargetedSpaceReturns(true)
			fakeConfig.TargetedSpaceReturns(configv3.Space{
				GUID: "some-space-guid",
				Name: "some-space"})
			fakeConfig.CurrentUserReturns(
				configv3.User{Name: "some-user"},
				nil)
		})

		Context("when getting the current user returns an error", func() {
			var expectedErr error

			BeforeEach(func() {
				expectedErr = errors.New("getting current user error")
				fakeConfig.CurrentUserReturns(
					configv3.User{},
					expectedErr)
			})

			It("returns the error", func() {
				Expect(executeErr).To(MatchError(expectedErr))
			})
		})

		It("displays flavor text", func() {
			Expect(testUI.Out).To(Say("Starting app some-app in org some-org / space some-space as some-user..."))
		})

		Context("when the app exists", func() {
			Context("when the app is already started", func() {
				BeforeEach(func() {
					fakeActor.GetApplicationByNameAndSpaceReturns(
						v2action.Application{State: ccv2.ApplicationStarted},
						v2action.Warnings{"warning-1", "warning-2"},
						nil,
					)
				})

				It("short circuits and displays message", func() {
					Expect(executeErr).ToNot(HaveOccurred())

					Expect(testUI.Out).To(Say("App some-app is already started"))

					Expect(testUI.Err).To(Say("warning-1"))
					Expect(testUI.Err).To(Say("warning-2"))

					Expect(fakeActor.StartApplicationCallCount()).To(Equal(0))
				})
			})

			Context("when the app is not already started", func() {
				BeforeEach(func() {
					fakeActor.GetApplicationByNameAndSpaceReturns(
						v2action.Application{GUID: "app-guid", State: ccv2.ApplicationStopped},
						v2action.Warnings{"warning-1", "warning-2"},
						nil,
					)
				})

				It("starts the app", func() {
					Expect(executeErr).ToNot(HaveOccurred())

					Expect(testUI.Err).To(Say("warning-1"))
					Expect(testUI.Err).To(Say("warning-2"))

					Expect(fakeActor.StartApplicationCallCount()).To(Equal(1))
					app, _, config := fakeActor.StartApplicationArgsForCall(0)
					Expect(app.GUID).To(Equal("app-guid"))
					Expect(config).To(Equal(fakeConfig))
				})

				Context("when passed a log message", func() {
					BeforeEach(func() {
						fakeActor.StartApplicationStub = func(app v2action.Application, client v2action.NOAAClient, config v2action.Config) (<-chan *v2action.LogMessage, <-chan error, <-chan string, <-chan error) {
							messages := make(chan *v2action.LogMessage)
							logErrs := make(chan error)
							warnings := make(chan string)
							errs := make(chan error)

							go func() {
								messages <- v2action.NewLogMessage("log message 1", 1, time.Unix(0, 0), "APP", "1")
								messages <- v2action.NewLogMessage("log message 2", 1, time.Unix(0, 0), "APP", "1")
								close(messages)
								close(logErrs)
								close(warnings)
								close(errs)
							}()

							return messages, logErrs, warnings, errs
						}
					})

					It("displays the log", func() {
						Expect(executeErr).ToNot(HaveOccurred())
						Expect(testUI.Out).To(Say("log message 1"))
						Expect(testUI.Out).To(Say("log message 2"))
					})
				})

				Context("when passed an log err", func() {
					Context("NOAA connection times out/closes", func() {
						BeforeEach(func() {
							fakeActor.StartApplicationStub = func(app v2action.Application, client v2action.NOAAClient, config v2action.Config) (<-chan *v2action.LogMessage, <-chan error, <-chan string, <-chan error) {
								messages := make(chan *v2action.LogMessage)
								logErrs := make(chan error)
								warnings := make(chan string)
								errs := make(chan error)

								go func() {
									messages <- v2action.NewLogMessage("log message 1", 1, time.Unix(0, 0), "APP", "1")
									messages <- v2action.NewLogMessage("log message 2", 1, time.Unix(0, 0), "APP", "1")
									messages <- v2action.NewLogMessage("log message 3", 1, time.Unix(0, 0), "APP", "1")
									logErrs <- v2action.NOAATimeoutError{}
									close(messages)
									close(logErrs)
									close(warnings)
									close(errs)
								}()

								return messages, logErrs, warnings, errs
							}

							applicationSummary := v2action.ApplicationSummary{
								Application: v2action.Application{
									Name:                 "some-app",
									GUID:                 "some-app-guid",
									Instances:            3,
									Memory:               128,
									PackageUpdatedAt:     time.Unix(0, 0),
									DetectedBuildpack:    "some-buildpack",
									State:                "STARTED",
									DetectedStartCommand: "some start command",
								},
								Stack: v2action.Stack{
									Name: "potatos",
								},
								Routes: []v2action.Route{
									{
										Host:   "banana",
										Domain: "fruit.com",
										Path:   "/hi",
									},
									{
										Domain: "foobar.com",
										Port:   13,
									},
								},
							}
							warnings := []string{"app-summary-warning"}

							applicationSummary.RunningInstances = []v2action.ApplicationInstanceWithStats{}

							fakeActor.GetApplicationSummaryByNameAndSpaceReturns(applicationSummary, warnings, nil)
						})

						It("displays a warning and continues until app has started", func() {
							Expect(executeErr).To(BeNil())
							Expect(testUI.Out).To(Say("message 1"))
							Expect(testUI.Out).To(Say("message 2"))
							Expect(testUI.Out).To(Say("message 3"))
							Expect(testUI.Err).To(Say("timeout connecting to log server, no log will be shown"))
							Expect(testUI.Out).To(Say("Name:\\s+some-app"))
						})
					})

					Context("an unexpected error occurs", func() {
						var expectedErr error

						BeforeEach(func() {
							expectedErr = errors.New("err log message")
							fakeActor.StartApplicationStub = func(app v2action.Application, client v2action.NOAAClient, config v2action.Config) (<-chan *v2action.LogMessage, <-chan error, <-chan string, <-chan error) {
								messages := make(chan *v2action.LogMessage)
								logErrs := make(chan error)
								warnings := make(chan string)
								errs := make(chan error)

								go func() {
									logErrs <- expectedErr
									close(messages)
									close(logErrs)
									close(warnings)
									close(errs)
								}()

								return messages, logErrs, warnings, errs
							}
						})

						It("stops logging and returns the error", func() {
							Expect(executeErr).To(MatchError(expectedErr))
						})
					})
				})

				Context("when passed a warning", func() {
					Context("while NOAA is still logging", func() {
						BeforeEach(func() {
							fakeActor.StartApplicationStub = func(app v2action.Application, client v2action.NOAAClient, config v2action.Config) (<-chan *v2action.LogMessage, <-chan error, <-chan string, <-chan error) {
								messages := make(chan *v2action.LogMessage)
								logErrs := make(chan error)
								warnings := make(chan string)
								errs := make(chan error)

								go func() {
									warnings <- "warning 1"
									warnings <- "warning 2"
									close(messages)
									close(logErrs)
									close(warnings)
									close(errs)
								}()

								return messages, logErrs, warnings, errs
							}
						})

						It("displays the warnings to STDERR", func() {
							Expect(executeErr).ToNot(HaveOccurred())
							Expect(testUI.Err).To(Say("warning 1"))
							Expect(testUI.Err).To(Say("warning 2"))
						})
					})

					Context("while NOAA is no longer logging", func() {
						BeforeEach(func() {
							fakeActor.StartApplicationStub = func(app v2action.Application, client v2action.NOAAClient, config v2action.Config) (<-chan *v2action.LogMessage, <-chan error, <-chan string, <-chan error) {
								messages := make(chan *v2action.LogMessage)
								logErrs := make(chan error)
								warnings := make(chan string)
								errs := make(chan error)

								go func() {
									warnings <- "warning 1"
									warnings <- "warning 2"
									logErrs <- v2action.NOAATimeoutError{}
									close(messages)
									close(logErrs)
									warnings <- "warning 3"
									warnings <- "warning 4"
									close(warnings)
									close(errs)
								}()

								return messages, logErrs, warnings, errs
							}
						})

						It("displays the warnings to STDERR", func() {
							Expect(executeErr).ToNot(HaveOccurred())
							Expect(testUI.Err).To(Say("warning 1"))
							Expect(testUI.Err).To(Say("warning 2"))
							Expect(testUI.Err).To(Say("timeout connecting to log server, no log will be shown"))
							Expect(testUI.Err).To(Say("warning 3"))
							Expect(testUI.Err).To(Say("warning 4"))
						})
					})
				})

				Context("when passed an API err", func() {
					Context("while NOAA is still logging", func() {
						Describe("an unexpected error", func() {
							var expectedErr error

							BeforeEach(func() {
								expectedErr = errors.New("err log message")
								fakeActor.StartApplicationStub = func(app v2action.Application, client v2action.NOAAClient, config v2action.Config) (<-chan *v2action.LogMessage, <-chan error, <-chan string, <-chan error) {
									messages := make(chan *v2action.LogMessage)
									logErrs := make(chan error)
									warnings := make(chan string)
									errs := make(chan error)

									go func() {
										errs <- expectedErr
										close(messages)
										close(logErrs)
										close(warnings)
										close(errs)
									}()

									return messages, logErrs, warnings, errs
								}
							})

							It("stops logging and returns the error", func() {
								Expect(executeErr).To(MatchError(expectedErr))
							})
						})

						Describe("staging error", func() {
							BeforeEach(func() {
								fakeActor.StartApplicationStub = func(app v2action.Application, client v2action.NOAAClient, config v2action.Config) (<-chan *v2action.LogMessage, <-chan error, <-chan string, <-chan error) {
									messages := make(chan *v2action.LogMessage)
									logErrs := make(chan error)
									warnings := make(chan string)
									errs := make(chan error)

									go func() {
										errs <- v2action.StagingFailedError{Reason: "Something, but not nothing"}
										close(messages)
										close(logErrs)
										close(warnings)
										close(errs)
									}()

									return messages, logErrs, warnings, errs
								}
							})

							It("stops logging and returns StagingFailedError", func() {
								Expect(executeErr).To(MatchError(shared.StagingFailedError{BinaryName: "faceman", Message: "Something, but not nothing"}))
							})
						})
					})
				})

				Context("when the app finishes starting", func() {
					var (
						applicationSummary v2action.ApplicationSummary
						warnings           []string
					)

					BeforeEach(func() {
						applicationSummary = v2action.ApplicationSummary{
							Application: v2action.Application{
								Name:                 "some-app",
								GUID:                 "some-app-guid",
								Instances:            3,
								Memory:               128,
								PackageUpdatedAt:     time.Unix(0, 0),
								DetectedBuildpack:    "some-buildpack",
								State:                "STARTED",
								DetectedStartCommand: "some start command",
							},
							Stack: v2action.Stack{
								Name: "potatos",
							},
							Routes: []v2action.Route{
								{
									Host:   "banana",
									Domain: "fruit.com",
									Path:   "/hi",
								},
								{
									Domain: "foobar.com",
									Port:   13,
								},
							},
						}
						warnings = []string{"app-summary-warning"}

						applicationSummary.RunningInstances = []v2action.ApplicationInstanceWithStats{
							{
								ID:          0,
								State:       v2action.ApplicationInstanceState(ccv2.ApplicationInstanceRunning),
								Since:       1403140717.984577,
								CPU:         0.73,
								Disk:        50 * bytefmt.MEGABYTE,
								DiskQuota:   2048 * bytefmt.MEGABYTE,
								Memory:      100 * bytefmt.MEGABYTE,
								MemoryQuota: 128 * bytefmt.MEGABYTE,
								Details:     "info from the backend",
							},
						}

						fakeActor.GetApplicationSummaryByNameAndSpaceReturns(applicationSummary, warnings, nil)
					})

					It("displays the app summary and it's warnings", func() {
						Expect(executeErr).ToNot(HaveOccurred())
						Expect(testUI.Out).To(Say("Name:\\s+some-app"))
						Expect(testUI.Out).To(Say("Requested state:\\s+started"))
						Expect(testUI.Out).To(Say("Instances:\\s+1\\/3"))
						Expect(testUI.Out).To(Say("Usage:\\s+128M x 3 instances"))
						Expect(testUI.Out).To(Say("Routes:\\s+banana.fruit.com/hi, foobar.com:13"))
						Expect(testUI.Out).To(Say("Last uploaded:\\s+1970-01-01T00:00:00Z"))
						Expect(testUI.Out).To(Say("Stack:\\s+potatos"))
						Expect(testUI.Out).To(Say("Buildpack:\\s+some-buildpack"))
						Expect(testUI.Out).To(Say("Start command:\\s+some start command"))

						Expect(testUI.Err).To(Say("app-summary-warning"))
					})

					It("should display the instance table", func() {
						Expect(executeErr).ToNot(HaveOccurred())
						Expect(testUI.Out).To(Say("State\\s+Since\\s+CPU\\s+Memory\\s+Disk"))
						Expect(testUI.Out).To(Say(`#0\s+running\s+2014-06-19T01:18:37Z\s+73.0%\s+100M of 128M\s+50M of 2G\s+info from the backend`))
					})
				})
			})
		})

		Context("when the app does *not* exists", func() {
			BeforeEach(func() {
				fakeActor.GetApplicationByNameAndSpaceReturns(
					v2action.Application{},
					v2action.Warnings{"warning-1", "warning-2"},
					v2action.ApplicationNotFoundError{Name: "some-app"},
				)
			})

			It("returns back an error", func() {
				Expect(executeErr).To(MatchError(command.ApplicationNotFoundError{Name: "some-app"}))

				Expect(testUI.Err).To(Say("warning-1"))
				Expect(testUI.Err).To(Say("warning-2"))
			})
		})
	})
})
