package orchestrator

type Config struct {
	Scenario string
	Seed     int
}

func DefaultConfig() Config {
	return Config{
		Scenario: "FULL_CHAOS",
		Seed:     42,
	}
}
