package flag

const (
	FlagConfigFile           = "config"
	FlagStarAfter            = "start-after"
	FlagLimit                = "limit"
	FlagIncludeExpired       = "include-expired"
	FlagTopPrice             = "top-price"
	FlagMiddlePrice          = "middle-price"
	FlagLowPrice             = "low-price"
	FlagMinRegisterDuration  = "min-register-duration"
	FlagRegistrarAddress     = "registrar-address"
	FlagExpireAtHeight       = "expire-height"
	FlagExpireAtTime         = "expire-time"
	FlagSubdomainTextData    = "text"
	FlagSubdomainAccountData = "account"
	FlagPublicKeyData        = "public-key"
	FlagContentHashData      = "content-hash"
)

func StartAfterI(startAfter string) interface{} {
	var startAfterI interface{}
	if startAfter != "" {
		startAfterI = startAfter
	}
	return startAfterI
}

func LimitI(limit string) interface{} {
	var limitI interface{}
	if limit != "" {
		limitI = limit
	}
	return limitI
}

func TopPriceI(topPrice uint64) interface{} {
	var topPriceI interface{}
	if topPrice != 0 {
		topPriceI = topPrice
	}
	return topPriceI
}

func MiddlePriceI(middlePrice uint64) interface{} {
	var middlePriceI interface{}
	if middlePrice != 0 {
		middlePriceI = middlePrice
	}
	return middlePriceI
}

func LowPriceI(lowPrice uint64) interface{} {
	var lowPriceI interface{}
	if lowPrice != 0 {
		lowPriceI = lowPrice
	}
	return lowPriceI
}

func MinRegisterDurationI(minRegisterDuration uint64) interface{} {
	var minRegisterDurationI interface{}
	if minRegisterDuration != 0 {
		minRegisterDurationI = minRegisterDuration
	}
	return minRegisterDurationI
}

func RegistrarAddressI(registrarAddress string) interface{} {
	var registrarAddressI interface{}
	if registrarAddress != "" {
		registrarAddressI = registrarAddress
	}
	return registrarAddressI
}

func ExpireAtHeightI(expireAtHeight string) interface{} {
	var expireAtHeightI interface{}
	if expireAtHeight != "" {
		expireAtHeightI = expireAtHeight
	}
	return expireAtHeightI
}

func ExpireAtTimeI(expireAtTime uint64) interface{} {
	var expireAtTimeI interface{}
	if expireAtTime != 0 {
		expireAtTimeI = expireAtTime
	}
	return expireAtTimeI
}
