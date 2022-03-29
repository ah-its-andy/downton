package hosting

type ApplicationLifetime interface {
	OnStarted()
	OnStopping()
	OnStopped()

	OnPanic(err any)

	StopApplication()
}
