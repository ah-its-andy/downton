package servicelocator

const (
	LifetimeTransient = iota
	LifetimeSingleton
	LifetimeScope
)

func GetLifetimeName(lifetime int) string {
	switch lifetime {
	case LifetimeTransient:
		return "Transient"
	case LifetimeSingleton:
		return "Singleton"
	case LifetimeScope:
		return "Scope"
	default:
		return "Unknown"
	}
}
