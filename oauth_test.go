package gitlab

import "os"

var testTokenCredential = &TokenCredential{
	Endpoint:    os.Getenv("GITLAB_HOST"),
	AccessToken: os.Getenv("GITLAB_TOKEN"),
}
