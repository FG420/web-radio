package radio

type Order string
type StationsOrder string
type StationsBy string

const (
	OrderName         Order = "name"
	OrderStationCount Order = "stationcount"

	StationsOrderName           StationsOrder = "name"
	StationsOrderURL            StationsOrder = "url"
	StationsOrderHomepage       StationsOrder = "homepage"
	StationsOrderFavicon        StationsOrder = "favicon"
	StationsOrderTags           StationsOrder = "tags"
	StationsOrderCountry        StationsOrder = "country"
	StationsOrderState          StationsOrder = "state"
	StationsOrderLanguage       StationsOrder = "language"
	StationsOrderVotes          StationsOrder = "votes"
	StationsOrderCodec          StationsOrder = "codec"
	StationsOrderBitrate        StationsOrder = "bitrate"
	StationsOrderLastCheckOk    StationsOrder = "lastcheckok"
	StationsOrderLastCheckTime  StationsOrder = "lastchecktime"
	StationsOrderClickTimestamp StationsOrder = "clicktimestamp"
	StationsOrderClickCount     StationsOrder = "clickcount"
	StationsOrderClickTrend     StationsOrder = "clicktrend"
	StationsOrderRandom         StationsOrder = "random"

	StationsByUUID             StationsBy = "byuuid"
	StationsByName             StationsBy = "byname"
	StationsByNameExact        StationsBy = "bynameexact"
	StationsByCodec            StationsBy = "bycodec"
	StationsByCodecExact       StationsBy = "bycodecexact"
	StationsByCountry          StationsBy = "bycountry"
	StationsByCountryExact     StationsBy = "bycountryexact"
	StationsByCountryCodeExact StationsBy = "bycountrycodeexact"
	StationsByState            StationsBy = "bystate"
	StationsByStateExact       StationsBy = "bystateexact"
	StationsByLanguage         StationsBy = "bylanguage"
	StationsByLanguageExact    StationsBy = "bylanguageexact"
	StationsByTag              StationsBy = "bytag"
	StationsByTagExact         StationsBy = "bytagexact"
)
