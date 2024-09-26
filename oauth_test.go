package gitlab

import "os"

var testTokenCredential = &TokenCredential{
	AccessToken: os.Getenv("GITLAB_TOKEN"),
}
