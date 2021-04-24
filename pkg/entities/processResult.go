package entities

type ProcessedResult struct {
	CurrentGeo                     Geo   `json:"currentGeo"`
	TravelToCurrentGeoSuspicious   bool  `json:"travelToCurrentGeoSuspicious"`
	TravelFromCurrentGeoSuspicious bool  `json:"travelFromCurrentGeoSuspicious"`
	PrecedingIpAccess              IpLog `json:"precedingIpAccess"`
	SubsequentIpAccess             IpLog `json:"subsequentIpAccess"`
}
