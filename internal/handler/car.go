package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nekvil/cars-go-api/internal/model"
	"github.com/nekvil/cars-go-api/internal/utils"
)

// @Summary Retrieve all cars
// @Description Retrieves all cars optionally filtered and paginated.
// @ID get-all-cars
// @Accept json
// @Produce json
// @Param filter query string false "Optional filter to apply"
// @Param page query integer false "Optional page number for pagination"
// @Success 200 {object} successResponse "Successful response containing the list of cars"
// @Failure 400 {object} errorResponse "Invalid page number format"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /cars [get]
func (h *Handler) GetAllCars(c *gin.Context) {
	filter := c.Query("filter")
	page := c.Query("page")

	var pageNum int
	var err error
	if page == "" {
		pageNum = 1
	} else {
		pageNum, err = strconv.Atoi(page)
		if err != nil {
			utils.Logger.Debugf("Invalid page number: %v", err)
			errorResponseJSON(c, http.StatusBadRequest, fmt.Errorf("invalid page number format: %v", err))
			return
		}
	}

	cars, err := h.services.Car.GetAll(filter, pageNum)
	if err != nil {
		utils.Logger.Errorf("Failed to get cars: %v", err)
		errorResponseJSON(c, http.StatusInternalServerError, err)
		return
	}

	utils.Logger.Infof("Retrieved %d cars", len(cars))

	successResponseJSON(c, http.StatusOK, "Cars obtained successfully", gin.H{
		"cars": cars,
	})
}

// @Summary Delete a car
// @Description Deletes a car by its ID.
// @ID delete-car
// @Accept json
// @Produce json
// @Param id path integer true "Car ID to delete"
// @Success 200 {object} successResponse "Successful response indicating the deletion"
// @Failure 400 {object} errorResponse "Invalid car ID format"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /cars/{id} [delete]
func (h *Handler) DeleteCar(c *gin.Context) {
	id := c.Param("id")

	idNum, err := strconv.Atoi(id)
	if err != nil {
		utils.Logger.Errorf("Failed to convert carID to int: %v", err)
		errorResponseJSON(c, http.StatusBadRequest, fmt.Errorf("invalid car ID"))
		return
	}

	err = h.services.Car.Delete(idNum)
	if err != nil {
		utils.Logger.Errorf("Failed to delete car: %v", err)
		errorResponseJSON(c, http.StatusInternalServerError, err)
		return
	}

	utils.Logger.Infof("Car deleted successfully: %s", id)

	successResponseJSON(c, http.StatusOK, "Cars deleted successfully", gin.H{
		"deleted_car_id": id,
	})
}

// @Summary Update a car
// @Description Updates a car by its ID.
// @ID update-car
// @Accept json
// @Produce json
// @Param id path integer true "Car ID to update"
// @Param requestBody body model.Car true "New car data"
// @Success 200 {object} successResponse "Successful response indicating the update"
// @Failure 400 {object} errorResponse "Invalid car ID format or JSON payload"
// @Failure 500 {object} errorResponse "Internal server error"
// @Router /cars/{id} [put]
func (h *Handler) UpdateCar(c *gin.Context) {
	id := c.Param("id")

	idNum, err := strconv.Atoi(id)
	if err != nil {
		utils.Logger.Errorf("Failed to convert carID to int: %v", err)
		errorResponseJSON(c, http.StatusBadRequest, fmt.Errorf("invalid car ID"))
		return
	}

	var requestBody model.Car
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		utils.Logger.Debugf("Invalid JSON payload: %v", err)
		errorResponseJSON(c, http.StatusBadRequest, err)
		return
	}

	err = h.services.Car.Update(idNum, &requestBody)
	if err != nil {
		utils.Logger.Errorf("Failed to update car: %v", err)
		errorResponseJSON(c, http.StatusInternalServerError, err)
		return
	}

	utils.Logger.Infof("Car updated successfully: %s", id)

	successResponseJSON(c, http.StatusOK, "Car updated successfully", gin.H{
		"updated_car_id": idNum,
	})
}

type addCarRequestBody struct {
	RegNums []string `json:"regNums"`
}

// @Summary Add multiple cars
// @Description Adds multiple cars using their registration numbers.
// @ID add-cars
// @Accept json
// @Produce json
// @Param requestBody body addCarRequestBody true "Array of car registration numbers"
// @Success 200 {object} successResponse "Successful response indicating the addition of cars"
// @Failure 400 {object} errorResponse "Invalid JSON payload"
// @Failure 500 {object} errorResponse "Failed to add cars or internal server error"
// @Router /cars [post]
func (h *Handler) AddCars(c *gin.Context) {
	var requestBody addCarRequestBody

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		utils.Logger.Debugf("Invalid JSON payload: %v", err)
		errorResponseJSON(c, http.StatusBadRequest, err)
		return
	}

	cars, err := h.services.ClientApi.GetByRegNum(requestBody.RegNums)
	if err != nil {
		utils.Logger.Errorf("Failed to add cars: %v", err)
		errorResponseJSON(c, http.StatusInternalServerError, err)
		return
	}

	utils.Logger.Infof("All cars added successfully")

	successResponseJSON(c, http.StatusOK, "Cars added successfully", gin.H{
		"added_car_numbers": requestBody.RegNums,
		"cars":              cars,
	})
}
