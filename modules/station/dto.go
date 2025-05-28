package station

type Station struct{
	Id string `json:"nid"`
	Name string `json:"title"`
	Retail []StationRetail `json:"retails"`
	Facility []StationFacility `json:"fasilitas"`
}

type StationResponse struct{
	Id string `json:"id"`
	Name string `json:"name"`
}

type Schedule struct{
	Id string `json:"nid"`
	StationName string `json:"title"`
	ScheduleBundaranHI string `json:"jadwal_hi_biasa"`
	ScheduleLebakBulus string `json:"jadwal_lb_biasa"`
}

type ScheduleResponse struct{
	CurrentStation string `json:"current_station"`
	StationName string `json:"station"`
	Time string `json:"time"`
}

// cek retails di tiap stasiun
type StationRetailsResponse struct{
	Id string `json:"id"`
	StationName string `json:"title"`
	StationRetails []StationRetail `json:"retails"`
}

type StationRetail struct{
	RetailName string `json:"title"`
	Type string `json:"jenis_retail"`
	Cover string `json:"cover"`
}

type StationFacilitiesResponse struct{
	Id string `json:"id"`
	StationName string `json:"title"`
	StationFacility []StationFacility `json:"facilities"`
}

type StationFacility struct{
	RetailName string `json:"title"`
	Type string `json:"jenis_fasilitas"`
	Cover string `json:"cover"`
}