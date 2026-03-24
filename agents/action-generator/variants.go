package actiongenerator

func ExpandVariants(action string) []string {

	switch action {

	case "DOWNSIZE":
		return []string{"SMALL", "MEDIUM", "AGGRESSIVE"}

	case "TERMINATE":
		return []string{"SAFE", "FORCE"}

	case "SECURE":
		return []string{"PATCH", "RESTRICT"}

	default:
		return []string{"DEFAULT"}
	}
}
