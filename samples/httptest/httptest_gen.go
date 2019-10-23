package httptest

import (
	"io"
	"net/http"
	"net/http/httptest"
)

// NewServer starts a new mock http.Server using the test data.
func NewServer() *httptest.Server {
	return httptest.NewServer(
		http.HandlerFunc(router),
	)
}

func router(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.RawPath) != 0 {
		r.URL.Path = r.URL.RawPath
	}
	if len(r.URL.RawQuery) != 0 {
		r.URL.Path = r.URL.Path + "?" + r.URL.RawQuery
	}
	for _, route := range routes {
		if route.Method == r.Method && route.Path == r.URL.Path {
			for k, v := range route.Header {
				w.Header().Set(k, v)
			}
			w.WriteHeader(route.Status)
			io.WriteString(w, route.Body)
			return
		}
	}
	w.WriteHeader(404)
}

var routes = []struct {
	Method string
	Path   string
	Body   string
	Status int
	Header map[string]string
}{

	// PATCH /user
	{
		Method: "PATCH",
		Path:   "/user",
		Status: 200,
		Body:   "{\n  \"login\": \"octocat\",\n  \"id\": 1,\n  \"avatar_url\": \"https://github.com/images/error/octocat_happy.gif\",\n  \"gravatar_id\": \"\",\n  \"url\": \"https://api.github.com/users/octocat\",\n  \"html_url\": \"https://github.com/octocat\",\n  \"followers_url\": \"https://api.github.com/users/octocat/followers\",\n  \"following_url\": \"https://api.github.com/users/octocat/following{/other_user}\",\n  \"gists_url\": \"https://api.github.com/users/octocat/gists{/gist_id}\",\n  \"starred_url\": \"https://api.github.com/users/octocat/starred{/owner}{/repo}\",\n  \"subscriptions_url\": \"https://api.github.com/users/octocat/subscriptions\",\n  \"organizations_url\": \"https://api.github.com/users/octocat/orgs\",\n  \"repos_url\": \"https://api.github.com/users/octocat/repos\",\n  \"events_url\": \"https://api.github.com/users/octocat/events{/privacy}\",\n  \"received_events_url\": \"https://api.github.com/users/octocat/received_events\",\n  \"type\": \"User\",\n  \"site_admin\": false,\n  \"name\": \"monalisa octocat\",\n  \"company\": \"GitHub\",\n  \"blog\": \"https://github.com/blog\",\n  \"location\": \"San Francisco\",\n  \"email\": \"octocat@github.com\",\n  \"hireable\": false,\n  \"bio\": \"There once was...\",\n  \"public_repos\": 2,\n  \"public_gists\": 1,\n  \"followers\": 20,\n  \"following\": 0,\n  \"created_at\": \"2008-01-14T04:33:35Z\",\n  \"updated_at\": \"2008-01-14T04:33:35Z\",\n  \"total_private_repos\": 100,\n  \"owned_private_repos\": 100,\n  \"private_gists\": 81,\n  \"disk_usage\": 10000,\n  \"collaborators\": 8,\n  \"two_factor_authentication\": true,\n  \"plan\": {\n    \"name\": \"Medium\",\n    \"space\": 400,\n    \"private_repos\": 20,\n    \"collaborators\": 0\n  }\n}",
		Header: map[string]string{
			"Content-Type":          "application/json; charset=utf-8",
			"Date":                  "Mon, 01 Jul 2013 17:27:06 GMT",
			"X-Ratelimit-Limit":     "60",
			"X-Ratelimit-Remaining": "56",
			"X-Ratelimit-Reset":     "1372700873",
		},
	},

	// GET /user
	{
		Method: "GET",
		Path:   "/user",
		Status: 200,
		Body:   "{\n  \"login\": \"octocat\",\n  \"id\": 1,\n  \"avatar_url\": \"https://github.com/images/error/octocat_happy.gif\",\n  \"gravatar_id\": \"\",\n  \"url\": \"https://api.github.com/users/octocat\",\n  \"html_url\": \"https://github.com/octocat\",\n  \"followers_url\": \"https://api.github.com/users/octocat/followers\",\n  \"following_url\": \"https://api.github.com/users/octocat/following{/other_user}\",\n  \"gists_url\": \"https://api.github.com/users/octocat/gists{/gist_id}\",\n  \"starred_url\": \"https://api.github.com/users/octocat/starred{/owner}{/repo}\",\n  \"subscriptions_url\": \"https://api.github.com/users/octocat/subscriptions\",\n  \"organizations_url\": \"https://api.github.com/users/octocat/orgs\",\n  \"repos_url\": \"https://api.github.com/users/octocat/repos\",\n  \"events_url\": \"https://api.github.com/users/octocat/events{/privacy}\",\n  \"received_events_url\": \"https://api.github.com/users/octocat/received_events\",\n  \"type\": \"User\",\n  \"site_admin\": false,\n  \"name\": \"monalisa octocat\",\n  \"company\": \"GitHub\",\n  \"blog\": \"https://github.com/blog\",\n  \"location\": \"San Francisco\",\n  \"email\": \"octocat@github.com\",\n  \"hireable\": false,\n  \"bio\": \"There once was...\",\n  \"public_repos\": 2,\n  \"public_gists\": 1,\n  \"followers\": 20,\n  \"following\": 0,\n  \"created_at\": \"2008-01-14T04:33:35Z\",\n  \"updated_at\": \"2008-01-14T04:33:35Z\"\n}\n",
		Header: map[string]string{
			"Content-Type":          "application/json; charset=utf-8",
			"Date":                  "Mon, 01 Jul 2013 17:27:06 GMT",
			"X-Ratelimit-Limit":     "60",
			"X-Ratelimit-Remaining": "56",
			"X-Ratelimit-Reset":     "1372700873",
		},
	},
}
