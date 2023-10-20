package searchengines

type SearchEngine struct {
	Url string `json:"url"`
}

var (
	ahmia = "http://juhanurmihxlp77nkq76byazcldy2hlmovfu2epvl5ankdibsot4csyd.onion/search/?q=%s"
)
