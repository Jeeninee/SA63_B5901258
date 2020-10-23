package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oreo/app/ent"
	"github.com/oreo/app/ent/district"
)

// DistrictController defines the struct for the district controller
type DistrictController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateDistrict handles POST requests for adding district entities
// @Summary Create district
// @Description Create district
// @ID create-district
// @Accept   json
// @Produce  json
// @Param resolution body ent.District true "district entity"
// @Success 200 {object} ent.District
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /districts [post]
func (ctl *DistrictController) CreateDistrict(c *gin.Context) {
	obj := ent.District{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "district binding failed",
		})
		return
	}

	di, err := ctl.client.District.
		Create().
		SetDistrict(obj.District).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, di)
}

// GetDistrict handles GET requests to retrieve a district entity
// @Summary Get a district entity by ID
// @Description get district by ID
// @ID get-district
// @Produce  json
// @Param id path int true "District ID"
// @Success 200 {object} ent.District
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /districts/{id} [get]
func (ctl *DistrictController) GetDistrict(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	di, err := ctl.client.District.
		Query().
		Where(district.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, di)
}

// ListDistrict handles request to get a list of district entities
// @Summary List district entities
// @Description list district entities
// @ID list-district
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.District
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /districts [get]
func (ctl *DistrictController) ListDistrict(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	districts, err := ctl.client.District.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, districts)
}

// NewDistrictController creates and registers handles for the district controller
func NewDistrictController(router gin.IRouter, client *ent.Client) *DistrictController {
	dic := &DistrictController{
		client: client,
		router: router,
	}
	dic.register()
	return dic
}

// InitDistrictController registers routes to the main engine
func (ctl *DistrictController) register() {
	districts := ctl.router.Group("/districts")

	// CRUD
	districts.POST("", ctl.CreateDistrict)
	districts.GET(":id", ctl.GetDistrict)
	districts.GET("", ctl.ListDistrict)
}
