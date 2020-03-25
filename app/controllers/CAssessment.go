package controllers

import (
	"managerApp/app/helpers"
	"managerApp/app/models/providers"
	"managerApp/app/models/entities"
	"database/sql"
	_"encoding/base64"
	"encoding/json"
	"fmt"
	//"strings"

	"io/ioutil"

	_"log"
	"strconv"

	"github.com/revel/revel"
)

type CAssessment struct {
	*revel.Controller
	provider *providers.AssessmentProvider
	//assessments *mappers.AssessmentMapper
	db       *sql.DB
}

func (c *CAssessment) Init() {
	c.provider = new(providers.AssessmentProvider)
	c.provider.Init()
}

func (c *CAssessment) DeleteAssessmentByID() revel.Result {
	c.Init()
	//получаем ID удаляемого ассессментаи и конвертируем в int
	sAssessmentId := c.Params.Get("assessmentID")
	assessmentId, err := strconv.ParseInt(sAssessmentId, 10, 64)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	fmt.Println(assessmentId)
	//вызываем метод удаления ассессмента из провайдера
	erro := c.provider.DeleteAssessment(assessmentId)
	if erro != nil {
		return c.RenderJSON(erro)
	}
	return nil
}

func (c *CAssessment) PostAssessmentByID() revel.Result {
	c.Init()
	sAssessmentId := c.Params.Get("assessmentID")
	assessmentId, err := strconv.ParseInt(sAssessmentId, 10, 64)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	var newAssessment entities.Assessment
	b, err := ioutil.ReadAll(c.Request.GetBody())
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	err = json.Unmarshal(b, &newAssessment)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	updatedAssessment, err := c.provider.PostAssessment(&newAssessment, assessmentId)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	return c.RenderJSON(helpers.Success(updatedAssessment))
}

func (c *CAssessment) GetAssessments() revel.Result {
	c.Init()
	assessments, err := c.provider.GetAssessments()
	if err != nil {
		panic(err)
	}
	return c.RenderJSON(assessments)
}

func (c *CAssessment) PutAssessment(newAssessment entities.Assessment) revel.Result {
	c.Init()
	createdAssessment, err := c.provider.PutAssessment(newAssessment)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	return c.RenderJSON(helpers.Success(createdAssessment))
}
//получить возможные статусы ассессмента
func (c *CAssessment) GetStatus() revel.Result {
	c.Init()

	// достаём ID нужного ассессмента и конвертируем его в int
	sAssessmentId := c.Params.Get("assessmentID")
	assessmentId, err := strconv.ParseInt(sAssessmentId, 10, 64)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}

	assessment, err := c.provider.GetAssessmentStatus(assessmentId)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	return c.RenderJSON(helpers.Success(assessment))
}
func (c *CAssessment) SetStatus(newStatus entities.AssessmentStatus) revel.Result {
	c.Init()
	sAssessmentId := c.Params.Get("assessmentID")
	assessmentId, err := strconv.ParseInt(sAssessmentId, 10, 64)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	// достаём ID статуса и конвертируем его в int
	sStatusId := c.Params.Get("statusID")
	statusId, err := strconv.ParseInt(sStatusId, 10, 64)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	updatedStatus, err := c.provider.SetAssessmentStatus(&newStatus, statusId, assessmentId)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	return c.RenderJSON(helpers.Success(updatedStatus))
}

func (c *CAssessment) GetAssessmentByID() revel.Result {
	c.Init()
	// достаём ID ассессмента и конвертируем его в int
	sAssessmentId := c.Params.Get("assessmentID")
	assessmentId, err := strconv.ParseInt(sAssessmentId, 10, 64)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	fmt.Println("Assessment ID:", sAssessmentId)
	//вызываем метод GetAssessmentById из провайдера
	assessment, err := c.provider.GetAssessmentById(assessmentId)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	//fmt.Println(helpers.Success(assessment))
	return c.RenderJSON(helpers.Success(assessment))
}