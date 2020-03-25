function showAllInterviewer(str){
    let xhr = new XMLHttpRequest();
    xhr.open("GET", "/interviewer");
    xhr.onreadystatechange = function() {
        if (xhr.status == 200 && xhr.readyState == xhr.DONE) {
            let res = JSON.parse(xhr.response)
            if (res.Result === 1) {
                webix.message({type:"error", text:res.ErrorText});
                return
            }
            $$(str).clearAll();
            $$(str).parse(xhr.response);
        }
    }       
    xhr.send();
}

function showInterviewer(){
    let selectedAssessmentId = $$("assessments").getSelectedItem().ID
    let xhr = new XMLHttpRequest();
    xhr.open("GET", "/assessment/" + selectedAssessmentId + "/interviewer");
    xhr.onreadystatechange = function() {
        if (xhr.status == 200 && xhr.readyState == xhr.DONE) {
            let res = JSON.parse(xhr.response)
            if (res.Result === 1) {
                webix.message({type:"error", text:res.ErrorText});
                return
            }
            $$("interviewerList").clearAll();
            $$("interviewerList").parse(xhr.response);
        }
    }       
    xhr.send();
}

function createInterviewer(str, surname, name, patronymic, email, phone, position){
    let xhr = new XMLHttpRequest();
    let newInterviewer = {
        Surname: surname,
	    Name: name,
	    Patronymic: patronymic,
	    Email: email,
	    PhoneNumber: phone,
	    Position: position
    }
    console.log("НОВЫЙ СОТРУДНИК:", newInterviewer)
    xhr.onreadystatechange = function() {
        console.log(xhr.readyState, xhr.status, xhr.responseText)
        if (xhr.status == 200 && xhr.readyState == xhr.DONE) {
            showAllInterviewer(str)
            showAllInterviewer("popupList")
        }
    }
    xhr.open("PUT", "/interviewer");
    xhr.setRequestHeader("Content-Type", "application/json; charset=UTF-8");       
    xhr.send(JSON.stringify(newInterviewer));
}

function addInterviewer(surname, name, patronymic, email, phoneNumber, position){
    let xhr = new XMLHttpRequest();
    let selectedAssessmentId = $$("assessments").getSelectedItem().ID
    let newInterviewer = {Surname: surname,
        Name: name,
        Patronymic: patronymic,
        Email: email, 
        PhoneNumber: phoneNumber, 
        Position: position}
    xhr.open("PUT", "/assessment/" + selectedAssessmentId + "/interviewer");
    xhr.onreadystatechange = function() {
        if (xhr.status == 200 && xhr.readyState == xhr.DONE) {/*
            let res = JSON.parse(xhr.response)
            if (res.Result === 1) {
                webix.message({type:"error", text:res.ErrorText});
                return
            }*/
            showInterviewer(selectedAssessmentId);
        }
    }       
    xhr.send(JSON.stringify(newInterviewer));
}

function removeInterviewer(){
    let selectedInterviewerId = $$("interviewerList").getSelectedItem().ID
    let selectedAssessmentId = $$("assessments").getSelectedItem().ID
    let xhr = new XMLHttpRequest();
    xhr.open("DELETE", "/assessment/" + selectedAssessmentId + "/interviewer/" + selectedInterviewerId);
    xhr.onreadystatechange = function() {
        if (xhr.status == 200 && xhr.readyState == xhr.DONE) {/*
            let res = JSON.parse(xhr.response)
            if (res.Result === 1) {
                webix.message({type:"error", text:res.ErrorText});
                return
            }*/
            showInterviewer(selectedAssessmentId)
            
        }
    }       
    xhr.send();
}

function removeInterviewerFromD(str){
    let selectedInterviewerId = $$(str).getSelectedItem().ID
    let xhr = new XMLHttpRequest();
    xhr.open("DELETE", "/interviewer/" + selectedInterviewerId);
    xhr.onreadystatechange = function() {
        if (xhr.status == 200 && xhr.readyState == xhr.DONE) {/*
            let res = JSON.parse(xhr.response)
            if (res.Result === 1) {
                webix.message({type:"error", text:res.ErrorText});
                return
            }*/
            showAllInterviewer(str)
            showAllInterviewer("popupList")
            //showInterviewer()
        }
    }       
    xhr.send();
}

function updateInterviewer(surname, name, patronymic, email, phoneNumber, position){
    let xhr = new XMLHttpRequest();
    let selectedInterviewerId = $$("InterviewerDictionary").getSelectedItem().ID
    let newInterviewer = {Surname: surname,
        Name: name,
        Patronymic: patronymic,
        Email: email, 
        PhoneNumber: phoneNumber, 
        Position: position}
    xhr.open("POST", "/interviewer/" + selectedInterviewerId);
    xhr.onreadystatechange = function() {
        if (xhr.status == 200 && xhr.readyState == xhr.DONE) {/*
            let res = JSON.parse(xhr.response)
            if (res.Result === 1) {
                webix.message({type:"error", text:res.ErrorText});
                return
            }*/
            showAllInterviewer("InterviewerDictionary")
            showAllInterviewer("popupList")
            showInterviewer()
        }
    }       
    xhr.send(JSON.stringify(newInterviewer));
}