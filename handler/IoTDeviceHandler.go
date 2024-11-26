package handler

import (
	"IoT-GPS-Backend/constants"
	req "IoT-GPS-Backend/dto/request"
	res "IoT-GPS-Backend/dto/response"
	"IoT-GPS-Backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IoTDeviceHandler interface {
	HandleRegisterDevice(c *gin.Context)
}

type ioTDeviceHandler struct {
	ioTDeviceService service.IoTDeviceService
}

func NewIoTDeviceHandler(ioTDeviceService service.IoTDeviceService) IoTDeviceHandler {
	return &ioTDeviceHandler{ioTDeviceService}
}

func (i *ioTDeviceHandler) HandleRegisterDevice(c *gin.Context) {
	var request req.CreateDeviceRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, res.ErrorResponse{
			StatusCode:   http.StatusBadRequest,
			ErrorMessage: constants.InvalidRequestBodyError,
		})
		return
	}

	response, err := i.ioTDeviceService.RegisterNewDevice(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, res.ErrorResponse{
			StatusCode:   http.StatusInternalServerError,
			ErrorMessage: "Unknown Error",
		})
		return
	}

	c.JSON(http.StatusCreated, res.CommonResponse{
		StatusCode: http.StatusCreated,
		Message:    constants.DeviceRegisterSuccess,
		Data:       response,
	})
	return
}
