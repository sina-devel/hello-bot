package translategooglefree

import (
	"encoding/json"
	"net/http"
)

type Result struct {
	Sentences []struct {
		Trans   string `json:"trans"`
		Orig    string `json:"orig"`
		Backend int    `json:"backend"`
	} `json:"sentences"`
	Src        string  `json:"src"`
	Confidence float64 `json:"confidence"`
	Spell      struct {
	} `json:"spell"`
	LdResult struct {
		Srclangs            []string  `json:"srclangs"`
		SrclangsConfidences []float64 `json:"srclangs_confidences"`
		ExtendedSrclangs    []string  `json:"extended_srclangs"`
	} `json:"ld_result"`
}

func Translate(source, sourceLang, targetLang string) (*Result, error) {
	var result Result

	req, err := http.NewRequest("GET", "https://translate.google.com/translate_a/single?client=at&dt=t&dt=ld&dt=qca&dt=rm&dt=bd&dj=1&hl=es-ES&ie=UTF-8&oe=UTF-8&inputm=2&otf=2", nil)
	if err != nil {
		return nil, err
	}

	fields := req.URL.Query()
	fields.Set("sl", sourceLang)
	fields.Set("tl", targetLang)
	fields.Set("q", source)

	req.URL.RawQuery = fields.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}
