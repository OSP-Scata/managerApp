package controllers

import (
	"managerApp/app/helpers"
	"managerApp/app/models/providers"
	"managerApp/app/models/entities"
	"database/sql"
	_"encoding/base64"
	_"encoding/json"
	"fmt"
	"strconv"
	//"strings"

	_"io/ioutil"
	"github.com/revel/revel"
)

type CCandidate struct {
	*revel.Controller
	provider *providers.CandidateProvider
	db       *sql.DB
}

func (c *CCandidate) Init() {
	c.provider = new(providers.CandidateProvider)
	c.provider.Init()
}

func (c *CCandidate) GetAllCandidates() revel.Result {
	c.Init()
	candidates, err := c.provider.GetAllCandidates()
	if err != nil {
		return c.RenderJSON(err)
	}

	return c.RenderJSON(candidates)
}

//получить всех кандидатов, состоящих в выбранном ассессменте
func (c *CCandidate) GetCandidates() revel.Result {
	c.Init()

	//получаем ID выбранного ассессмента и конвертируем в int
	sAssessmentId := c.Params.Get("assessmentID")
	assessmentId, err := strconv.ParseInt(sAssessmentId, 10, 64)
	if err != nil {
		fmt.Println(err)
		return c.RenderJSON(err)
	}

	//вызываем метод GetCandidates провайдера
	candidates, err := c.provider.GetCandidates(assessmentId)
	if err != nil {
		return c.RenderJSON(err)
	}
	return c.RenderJSON(candidates)
}

func (c *CCandidate) GetCandidateByID() revel.Result {
	c.Init()

	//получаем ID выбранного кандидата и конвертируем в int
	sCandidateId := c.Params.Get("candidateID")
	candidateId, err := strconv.ParseInt(sCandidateId, 10, 64)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}

	//получаем ID выбранного ассессмента и конвертируем в int
	sAssessmentId := c.Params.Get("assessmentID")
	assessmentId, err := strconv.ParseInt(sAssessmentId, 10, 64)
	if err != nil {
		fmt.Println(err)
		return c.RenderJSON(err)
	}

	//вызываем метод GetCandidateById провайдера
	candidate, err := c.provider.GetCandidateById(candidateId, assessmentId)
	if err != nil {
		return c.RenderJSON(err)
	}
	return c.RenderJSON(helpers.Success(candidate))
}

//получить возможные статусы кандидата
func (c *CCandidate) GetCandidateStatus() revel.Result {
	c.Init()

	sAssessmentId := c.Params.Get("assessmentID")
	assessmentId, err := strconv.ParseInt(sAssessmentId, 10, 64)
	if err != nil {
		fmt.Println(err)
		return c.RenderJSON(err)
	}

	sCandidatetId := c.Params.Get("candidateID")
	candidateId, err := strconv.ParseInt(sCandidatetId, 10, 64)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	candidate, err := c.provider.GetCandidateStatus(candidateId, assessmentId)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	fmt.Println(c.RenderJSON(helpers.Success(candidate)))
	return c.RenderJSON(helpers.Success(candidate))
}

func (c *CCandidate) SetStatus(newStatus entities.CandidateStatus) revel.Result {
	c.Init()
	sCandidateId := c.Params.Get("candidateID")
	candidateId, err := strconv.ParseInt(sCandidateId, 10, 64)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	sStatusId := c.Params.Get("statusID")
	statusId, err := strconv.ParseInt(sStatusId, 10, 64)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	updatedStatus, err := c.provider.SetCandidateStatus(&newStatus, statusId, candidateId)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	return c.RenderJSON(helpers.Success(updatedStatus))
}

//добавляем кандидата в выбранный ассессмент
func (c *CCandidate) PutCandidate(newCandidate entities.Candidate) revel.Result {
	c.Init()
	sAssessmentId := c.Params.Get("assessmentID")
	assessmentId, err := strconv.ParseInt(sAssessmentId, 10, 64)
	fmt.Println(sAssessmentId, assessmentId)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	createdCandidate, err := c.provider.PutCandidate(&newCandidate, assessmentId)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	return c.RenderJSON(helpers.Success(createdCandidate))
}

//изменение кандидата
func (c *CCandidate) PostCandidateByID(newCandidate entities.Candidate) revel.Result {
	c.Init()
	sCandidateId := c.Params.Get("candidateID")
	candidateId, err := strconv.ParseInt(sCandidateId, 10, 64)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}

	sAssessmentId := c.Params.Get("assessmentID")
	assessmentId, err := strconv.ParseInt(sAssessmentId, 10, 64)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	/*var newCandidate entities.Candidate
	b, err := ioutil.ReadAll(c.Request.GetBody())
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	fmt.Println("Change candidate request:", b)
	err = json.Unmarshal(b, &newCandidate)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}*/
	updatedCandidate, err := c.provider.PostCandidate(&newCandidate, candidateId, assessmentId)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	return c.RenderJSON(helpers.Success(updatedCandidate))
}

//удаление кандидата
func (c *CCandidate) DeleteCandidateByID() revel.Result {
	c.Init()
	sCandidateId := c.Params.Get("candidateID")
	candidateId, err := strconv.ParseInt(sCandidateId, 10, 64)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}

	sAssessmentId := c.Params.Get("assessmentID")
	assessmentId, err := strconv.ParseInt(sAssessmentId, 10, 64)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	fmt.Printf("ID Assessment from CCandidate:", assessmentId, ", ", candidateId)
	erro := c.provider.DeleteCandidate(candidateId, assessmentId)
	if err != nil {
		return c.RenderJSON(erro)
	}
	return nil
}
