package station

import (
	"encoding/json"
	"errors"
	"mrt-schedules-go/common/client"
	"net/http"
	"strings"
	"time"
)

type Service interface {
	GetAllStation() (response []StationResponse, err error)           //returnnya list dan error
	CheckSchedule(id string) (response []ScheduleResponse, err error) //returnnya list dan error
}

type service struct {
	client *http.Client
}

func NewService() Service {
	return &service{
		client: &http.Client{
			Timeout: 1000 * time.Second,
		},
	}
}

func (s *service) GetAllStation() (response []StationResponse, err error) {
	//layer service
	url := "https://www.jakartamrt.co.id/id/val/stasiuns"

	//hit url
	byteResponse, err := client.DoRequest(s.client, url)

	if err != nil {
		return nil, err
	}

	var stations []Station
	err = json.Unmarshal(byteResponse, &stations)

	for _, item := range stations {
		response = append(response, StationResponse{
			Id:   item.Id,
			Name: item.Name,
		})
	}
	return
}

func (s *service) CheckSchedule(id string) (response []ScheduleResponse, err error) {
	//layer service
	url := "https://www.jakartamrt.co.id/id/val/stasiuns"

	//hit url
	byteResponse, err := client.DoRequest(s.client, url)

	if err != nil {
		return nil, err
	}

	var schedules []Schedule
	err = json.Unmarshal(byteResponse, &schedules)

	if err != nil {
		return nil, err
	}
	// schedule selected by id
	var scheduleSelected Schedule
	for _, item := range schedules {
		if item.Id == id {
			scheduleSelected = item
			break
		}
	}

	if scheduleSelected.Id == "" {
		err = errors.New("Station Not Found")
		return
	}

	// fmt.Println(scheduleSelected)

	response, err = ConvertDataToResponse(scheduleSelected)

	if err != nil {
		return nil, err
	}

	return
}

func ConvertDataToResponse(schedule Schedule) (response []ScheduleResponse, err error) {
	var (
		LebakBulusTripName = "Stasiun Lebak Bulus Grab"
		BundaranHITripName = "Stasiun Bundaran HI Bank DKI"
	)

	currentStation := schedule.StationName
	scheduleLebakBulus := schedule.ScheduleLebakBulus
	scheduleBundaranHI := schedule.ScheduleBundaranHI

	scheduleLebakBulusParsed, err := ConvertScheduleToTimeFormat(scheduleLebakBulus)
	if err != nil {
		return nil, err
	}

	scheduleBundaranHIParsed, err := ConvertScheduleToTimeFormat(scheduleBundaranHI)
	if err != nil {
		return nil, err
	}

	// convert to response
	response, err = convertResponse(currentStation, scheduleLebakBulusParsed, LebakBulusTripName, response)
	
	if err != nil {
		return nil, err
	}

	response, err = convertResponse(currentStation, scheduleBundaranHIParsed, BundaranHITripName, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func ConvertScheduleToTimeFormat(schedule string) (response []time.Time, err error) {
	var parsedTime time.Time
	var schedules = strings.Split(schedule, ",")

	for _, item := range schedules {
		trimmedTime := strings.TrimSpace(item)
		if trimmedTime == "" {
			continue
		}
		parsedTime, err = time.Parse("15:04", trimmedTime)

		if err != nil {
			err = errors.New("Invalid Time Format " + trimmedTime)
			return
		}
		response = append(response, parsedTime)
	}
	return
}

func convertResponse(currentStation string, schedules []time.Time, label string, response []ScheduleResponse) ([]ScheduleResponse, error) {
	if len(schedules) == 0 {
		return response, errors.New("schedules are empty")
	}
	for _, item := range schedules {
		if item.Format("15:04") > time.Now().Format("15:04") {
			response = append(response, ScheduleResponse{
				CurrentStation: currentStation,
				StationName:    label,
				Time:           item.Format("15:04"),
			})
		}
	}
	return response, nil
}
