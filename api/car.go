package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

type VesApiResponse struct {
	RegistrationNumber       string `json:"registrationNumber"`
	Co2Emissions             int    `json:"co2Emissions"`
	EngineCapacity           int    `json:"engineCapacity"`
	MarkedForExport          bool   `json:"markedForExport"`
	FuelType                 string `json:"fuelType"`
	MotStatus                string `json:"motStatus"`
	Colour                   string `json:"colour"`
	Make                     string `json:"make"`
	TypeApproval             string `json:"typeApproval"`
	YearOfManufacture        int    `json:"yearOfManufacture"`
	TaxDueDate               string `json:"taxDueDate"`
	TaxStatus                string `json:"taxStatus"`
	DateOfLastV5CIssued      string `json:"dateOfLastV5CIssued"`
	MotExpiryDate            string `json:"motExpiryDate"`
	Wheelplan                string `json:"wheelplan"`
	MonthOfFirstRegistration string `json:"monthOfFirstRegistration"`
	EuroStatus               string `json:"euroStatus"`
}

func (server *Server) GetCarRegistration(ctx echo.Context) error {
	carRegistration := ctx.Param("reg")

	if carRegistration == "" {
		err := errors.New("reg param not provided")
		return echo.NewHTTPError(http.StatusBadRequest, errorResponse(err))
	}

	carInfo, err := getCarInfoRequest(server.config.VesApiUrl, server.config.VesApiKey, carRegistration)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, carInfo)
}

func getCarInfoRequest(address, apiKey, carRegistration string) (*VesApiResponse, error) {
	body, err := json.Marshal(map[string]string{"registrationNumber": carRegistration})
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, errorResponse(err))
	}

	request, err := http.NewRequest("POST", address+"/vehicle-enquiry/v1/vehicles", bytes.NewReader(body))
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, errorResponse(err))
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("x-api-key", apiKey)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, errorResponse(err))
	}

	defer response.Body.Close()

	if response.StatusCode > 399 {
		if response.StatusCode == 404 {
			err := errors.New("failed find car with license plate: " + carRegistration)
			return nil, echo.NewHTTPError(response.StatusCode, errorResponse(err))
		}
		err := errors.New("failed to get car info")
		return nil, echo.NewHTTPError(response.StatusCode, errorResponse(err))
	}

	responseBody := &VesApiResponse{}
	if err = json.NewDecoder(response.Body).Decode(responseBody); err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, errorResponse(err))
	}

	return responseBody, nil

}
