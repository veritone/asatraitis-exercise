package sosearch

import (
	"time"
	"translate/types"
)

type SearchResponse struct {
	Data struct {
		SearchMedia struct {
			Jsondata struct {
				Results []struct {
					Recording struct {
						AbsoluteStartTimeMs time.Time `json:"absoluteStartTimeMs"`
						AbsoluteStopTimeMs  time.Time `json:"absoluteStopTimeMs"`
						RecordingID         string    `json:"recordingId"`
						ProgramID           string    `json:"programId"`
						ProgramName         string    `json:"programName"`
						MediaSourceID       string    `json:"mediaSourceId"`
						MediaSourceTypeID   string    `json:"mediaSourceTypeId"`
						ParentTreeObjectIds []string  `json:"parentTreeObjectIds"`
						RelativeStopTimeMs  int       `json:"relativeStopTimeMs"`
						SliceTime           int       `json:"sliceTime"`
						MediaStartTime      int       `json:"mediaStartTime"`
						AibDuration         int       `json:"aibDuration"`
						IsOwn               bool      `json:"isOwn"`
						HitStartTime        int       `json:"hitStartTime"`
						HitEndTime          int       `json:"hitEndTime"`
					} `json:"recording"`
					StartDateTime int `json:"startDateTime"`
					StopDateTime  int `json:"stopDateTime"`
					Hits          []struct {
						Sdo struct {
							Series []struct {
								Email           string `json:"email"`
								JobTitle        string `json:"jobTitle"`
								OriginalFileURL string `json:"originalFileUrl"`
								TranslatedTo    string `json:"translatedTo"`
								Start           int    `json:"start"`
								End             int    `json:"end"`
							} `json:"series"`
						} `json:"sdo"`
					} `json:"hits"`
				} `json:"results"`
				TotalResults int    `json:"totalResults"`
				Limit        int    `json:"limit"`
				From         int    `json:"from"`
				To           int    `json:"to"`
				SearchToken  string `json:"searchToken"`
				Timestamp    int    `json:"timestamp"`
			} `json:"jsondata"`
		} `json:"searchMedia"`
	} `json:"data"`
	Errors types.Errors `json:"errors"`
}
type SearchInput struct {
	Value string `json:"value"`
}