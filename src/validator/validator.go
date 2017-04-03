package validator

type ValidatorInterface interface{
	// validate url and return normalized
	Validate(url string) (is_valid bool, n_url string)
}