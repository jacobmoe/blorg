package helper

func Application() map[string]interface{} {

	helper := make(map[string]interface{})

	helper["sidebar"] = func() string {
		return "TTESThhhhhh"
	}

	return helper
}
