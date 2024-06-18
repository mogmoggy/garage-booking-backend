package api

import (
	"bytes"
	"encoding/json"
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
		return echo.NewHTTPError(http.StatusBadRequest, "no car registration was provided")
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
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "failed to marshal json data")
	}

	request, err := http.NewRequest("POST", address+"/vehicle-enquiry/v1/vehicles", bytes.NewReader(body))
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "failed to make new http request")
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("x-api-key", apiKey)

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "failed to get car info")
	}

	defer response.Body.Close()

	if response.StatusCode > 399 {
		if response.StatusCode == 404 {
			return nil, echo.NewHTTPError(response.StatusCode, errorResponse("failed find car with license plate: "+carRegistration))
		}
		return nil, echo.NewHTTPError(response.StatusCode, errorResponse("failed to get car info"))
	}

	responseBody := &VesApiResponse{}
	if err = json.NewDecoder(response.Body).Decode(responseBody); err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, errorResponse("failed to decode json data from car info response"))
	}

	return responseBody, nil

}
