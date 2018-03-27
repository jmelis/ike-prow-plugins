package scm

// RepositoryIssue holds owner name, repository name and an issue number
type RepositoryIssue struct {
	Owner    string
	RepoName string
	Number   int
}

// ChangedFile is a type that contains information about created/modified/removed file within an scm repository
type ChangedFile struct {
	Name   string
	Status string
}
