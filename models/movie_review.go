package models

type MovieReview struct {
	DisplayTitle    string                `json:"display_title"`
	MpaaRating      string                `json:"mpaa_rating"`
	CriticsPick     int16                 `json:"critics_pick"`
	Byline          string                `json:"byline"`
	Headline        string                `json:"headline"`
	SummaryShort    string                `json:"summary_short"`
	PublicationDate string                `json:"publication_date"`
	OpeningDate     string                `json:"opening_date"`
	Multimedia      MovieReviewMultimedia `json:"multimedia"`
	Link            MovieReviewLink       `json:"link"`
}

type MovieReviewMultimedia struct {
	Type     string `json:"type"`
	ImageUrl string `json:"src"`
	Width    int32  `json:"width"`
	Height   int32  `json:"height"`
}

type MovieReviewLink struct {
	Type              string `json:"type"`
	Url               string `json:"url"`
	SuggestedLinkText string `json:"suggested_link_text"`
}

type MovieReviewHttpAnswer struct {
	Results []MovieReview `json:"results"`
}
