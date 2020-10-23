package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/oreo/app/ent"
	"github.com/oreo/app/ent/subdistrict"
)

// SubdistrictController defines the struct for the subdistrict controller
type SubdistrictController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateSubdistrict handles POST requests for adding subdistrict entities
// @Summary Create subdistrict
// @Description Create subdistrict
// @ID create-subdistrict
// @Accept   json
// @Produce  json
// @Param resolution body ent.Subdistrict true "subdistrict entity"
// @Success 200 {object} ent.Subdistrict
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /subdistricts [post]
func (ctl *SubdistrictController) CreateSubdistrict(c *gin.Context) {
	obj := ent.Subdistrict{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "subdistrict binding failed",
		})
		return
	}

	sd, err := ctl.client.Subdistrict.
		Create().
		SetSubdistrict(obj.Subdistrict).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, sd)
}

// GetSubdistrict handles GET requests to retrieve a subdistrict entity
// @Summary Get a subdistrict entity by ID
// @Description get subdistrict by ID
// @ID get-subdistrict
// @Produce json
// @Param id path int true "Subdistrict ID"
// @Success 200 {object} ent.Subdistrict
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /subdistricts/{id} [get]
func (ctl *SubdistrictController) GetSubdistrict(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	sd, err := ctl.client.Subdistrict.
		Query().
		Where(subdistrict.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, sd)
}

// ListSubdistrict handles request to get a list of subdistrict entities
// @Summary List subdistrict entities
// @Description list subdistrict entities
// @ID list-subdistrict
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Subdistrict
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /subdistricts [get]
func (ctl *SubdistrictController) ListSubdistrict(c *gin.Context) {
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

	subdistricts, err := ctl.client.Subdistrict.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, subdistricts)
}

// NewSubdistrictController creates and registers handles for the subdistrict controller
func NewSubdistrictController(router gin.IRouter, client *ent.Client) *SubdistrictController {
	sdc := &SubdistrictController{
		client: client,
		router: router,
	}
	sdc.register()
	return sdc
}

// InitSubdistrictController registers routes to the main engine
func (ctl *SubdistrictController) register() {
	subdistricts := ctl.router.Group("/subdistricts")

	// CRUD
	subdistricts.POST("", ctl.CreateSubdistrict)
	subdistricts.GET(":id", ctl.GetSubdistrict)
	subdistricts.GET("", ctl.ListSubdistrict)
}
