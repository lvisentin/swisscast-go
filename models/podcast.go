package models

type PodcastFeed struct {
	Title           string         `json:"title"`
	Description     string         `json:"description"`
	Link            string         `json:"link"`
	FeedLink        string         `json:"feedLink"`
	Links           []string       `json:"links"`
	Updated         string         `json:"updated"`
	UpdatedParsed   string         `json:"updatedParsed"`
	Published       string         `json:"published"`
	PublishedParsed string         `json:"publishedParsed"`
	Author          struct {
		Name string `json:"name"`
	} `json:"author"`
	Authors []struct {
		Name string `json:"name"`
	} `json:"authors"`
	Image struct {
		URL   string `json:"url"`
		Title string `json:"title"`
	} `json:"image"`
	Copyright string `json:"copyright"`
	Generator string `json:"generator"`
	Extensions struct {
		Atom struct {
			Link []struct {
				Name     string `json:"name"`
				Value    string `json:"value"`
				Attrs    struct {
					Href string `json:"href"`
					Rel  string `json:"rel"`
					Type string `json:"type"`
				} `json:"attrs"`
				Children interface{} `json:"children"`
			} `json:"link"`
		} `json:"atom"`
	} `json:"extensions"`
	Items      []PodcastEpisode `json:"items"`
	FeedType   string        `json:"feedType"`
	FeedVersion string        `json:"feedVersion"`
}

type PodcastEpisode struct {
	Title           string `json:"title"`
	Description     string `json:"description"`
	Link            string `json:"link"`
	Links           []string `json:"links"`
	Published       string `json:"published"`
	PublishedParsed string `json:"publishedParsed"`
	Author struct {
		Name string `json:"name"`
	} `json:"author"`
	Authors []struct {
		Name string `json:"name"`
	} `json:"authors"`
	Guid       string `json:"guid"`
	Enclosures []struct {
		URL    string `json:"url"`
		Length string `json:"length"`
		Type   string `json:"type"`
	} `json:"enclosures"`
	DcExt struct {
		Creator []string `json:"creator"`
	} `json:"dcExt"`
	Extensions struct {
		Dc struct {
			Creator []struct {
				Name     string `json:"name"`
				Value    string `json:"value"`
				Attrs    interface{} `json:"attrs"`
				Children interface{} `json:"children"`
			} `json:"creator"`
		} `json:"dc"`
	} `json:"extensions"`
}

