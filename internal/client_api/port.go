package client_api

type Port struct {
	//  {
	//  	"name": "Ajman",
	//  	"city": "Ajman",
	//  	"country": "United Arab Emirates",
	//  	"alias": [],
	//  	"regions": [],
	//  	"coordinates": [
	//  	55.5136433,
	//  	25.4052165
	//  ],
	//  	"province": "Ajman",
	//  	"timezone": "Asia/Dubai",
	//  	"unlocs": [
	//  	"AEAJM"
	//  ],
	//  	"code": "52000"
	//  }
	Name, Code, Country, Timezone, Province, City string
	Coordinates                                   [2]float64
	Unlocs                                        [1]string `json:"unlocs"`
}

func (p *Port) ID() string {
	return p.Unlocs[0]
}
