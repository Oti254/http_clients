package main


func fetchTasks(baseURL, availability string) []Issue {
	var issues string
	switch {
	case availability == "High":
		issues = "5"
	case availability == "Medium":
		issues = "3"
	default:
		issues = "1"
	}
	fullURL := baseURL + "?sort=estimate&limit=" + issues
	return getIssues(fullURL)
}
