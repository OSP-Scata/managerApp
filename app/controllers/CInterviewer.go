package controllers

import (
	"managerApp/app/helpers"
	"managerApp/app/models/entities"
	"managerApp/app/models/providers"
	//"database/sql"
	//"encoding/base64"
	"encoding/json"
	"fmt"
	//"strings"

	"io/ioutil"

	//"log"
	"strconv"

	"github.com/revel/revel"
)

type CInterviewer struct {
	*revel.Controller
	provider *providers.InterviewerProvider
}

func (c *CInterviewer) Init() {
	c.provider = new(providers.InterviewerProvider)
	c.provider.Init()
}

//получить выбранного сотрудника
func (c *CInterviewer) GetInterviewerByID() revel.Result {
	c.Init()

	// получаем ID сотрудника
	sInterviewerId := c.Params.Get("interviewerID")
	// конвертируем в int
	interviewerId, err := strconv.ParseInt(sInterviewerId, 10, 64)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}

	// вызваем метод GetInterviewerById провайдера
	interviewer, err := c.provider.GetInterviewerById(interviewerId)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	return c.RenderJSON(helpers.Success(interviewer))
}

//получить всех сотрудников
func (c *CInterviewer) GetAllInterviewers() revel.Result {
	c.Init()
	// вызваем метод GetAllInterviewers провайдера
	interviewers, err := c.provider.GetAllInterviewers()
	if err != nil {
		return c.RenderJSON(err)
	}
	return c.RenderJSON(interviewers)
}

//получить сотрудников, относящихся с выбранному ассессменту
func (c *CInterviewer) GetInterviewers() revel.Result {
	c.Init()

	//получаем ID ассессмента и конвертируем его в int
	sAssessmentId := c.Params.Get("assessmentID")
	assessmentId, err := strconv.ParseInt(sAssessmentId, 10, 64)
	if err != nil {
		fmt.Println(err)
		return c.RenderJSON(err)
	}

	//вызываем метод GetInterviewers
	interviewers, err := c.provider.GetInterviewers(assessmentId)
	if err != nil {
		return c.RenderJSON(err)
	}
	return c.RenderJSON(interviewers)
}

//добавление сотрудника в выбранный ассессмент
func (c *CInterviewer) PutInterviewer() revel.Result {
	c.Init()
	sAssessmentId := c.Params.Get("assessmentID")
	assessmentId, err := strconv.ParseInt(sAssessmentId, 10, 64)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}

	var newInterviewer entities.Interviewer

	//считываем данные с фронта
	b, err := ioutil.ReadAll(c.Request.GetBody())
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	//анмаршалим
	err = json.Unmarshal(b, &newInterviewer)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}

	// вызываем метод PutInterviewer в провайдере
	createdInterviewer, err := c.provider.PutInterviewer(&newInterviewer, assessmentId)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	return c.RenderJSON(helpers.Success(createdInterviewer))
}

//добавляем сотрудника (не в ассессмент, а в список сотрудников)
func (c *CInterviewer) SetInterviewer(newInterviewer entities.Interviewer) revel.Result {
	c.Init()
	createdInterviewer, err := c.provider.SetInterviewer(&newInterviewer)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	return c.RenderJSON(helpers.Success(createdInterviewer))
}

//изменить сотрудника
func (c *CInterviewer) PostInterviewer() revel.Result {
	c.Init()

	//получаем ID сотрудника
	sInterviewerId := c.Params.Get("interviewerID")
	interviewerId, err := strconv.ParseInt(sInterviewerId, 10, 64)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}

	var newInterviewer entities.Interviewer
	b, err := ioutil.ReadAll(c.Request.GetBody())
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	err = json.Unmarshal(b, &newInterviewer)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	updatedInterviewer, err := c.provider.PostInterviewer(&newInterviewer, interviewerId)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}
	return c.RenderJSON(helpers.Success(updatedInterviewer))
}

//удаление сотрудника из ассессмента
func (c *CInterviewer) DeleteInterviewerByID() revel.Result {
	c.Init()
	sInterviewerId := c.Params.Get("interviewerID")
	interviewerId, err := strconv.ParseInt(sInterviewerId, 10, 64)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}

	sAssessmentId := c.Params.Get("assessmentID")
	assessmentId, err := strconv.ParseInt(sAssessmentId, 10, 64)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}

	erro := c.provider.DeleteInterviewer(interviewerId, assessmentId)
	if erro != nil {
		return c.RenderJSON(erro)
	}
	return nil
}

//удаление сотрудника из списка сотрудников
func (c *CInterviewer) DeleteInterviewer() revel.Result {
	c.Init()
	sInterviewerId := c.Params.Get("interviewerID")
	interviewerId, err := strconv.ParseInt(sInterviewerId, 10, 64)
	if err != nil {
		return c.RenderJSON(helpers.Failed(err))
	}

	erro := c.provider.DeleteInterviewerFromD(interviewerId)
	if erro != nil {
		return c.RenderJSON(erro)
	}
	return nil
}
