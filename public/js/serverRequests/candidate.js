function showAllCandidates(){
    let xhr = new XMLHttpRequest();
    xhr.open("GET", "/candidate");
    xhr.onreadystatechange = function() {
        //console.log(xhr.readyState, xhr.status, xhr.responseText)
        if (xhr.status == 200 && xhr.readyState == xhr.DONE) {
            let res = JSON.parse(xhr.response)
            if (res.Result === 1) {
                webix.message({type:"error", text:res.ErrorText});
                return
            }
            //$$("interviewerList").clearAll();
            $$("CandidateList").clearAll();
            $$("CandidateList").parse(xhr.response);
        }
    }       
    xhr.send();
}

function showCandidate(){
    let selectedAssessmentId = $$("assessments").getSelectedItem().ID
    let xhr = new XMLHttpRequest();
    xhr.open("GET", "/assessment/" + selectedAssessmentId + "/candidate");
    xhr.onreadystatechange = function() {
        if (xhr.status == 200 && xhr.readyState == 4) {
            let res = JSON.parse(xhr.response)
            if (res.Result === 1) {
                webix.message({type:"error", text:res.ErrorText});
                return
            }
            $$("peopleList").clearAll();
            $$("peopleList").parse(xhr.response);
            console.log("showCandidate called")
        }
    } 
    xhr.send();
}

function showCandidateById(){
    let selectedAssessmentId = $$("assessments").getSelectedItem().ID
    let selectedCandidateId = $$("peopleList").getSelectedItem().ID
    let xhr = new XMLHttpRequest();
    xhr.open("GET", "/assessment/" + selectedAssessmentId + "/candidate/" + selectedCandidateId);
    xhr.onreadystatechange = function() {
        if (xhr.status == 200 && xhr.readyState == xhr.DONE) {
            let res = JSON.parse(xhr.response)
            if (res.Result === 1) {
                webix.message({type:"error", text:res.ErrorText});
                return
            }
            //console.log("КАНДИДАТ: ", res.Data); 
            $$("editForm").parse(res.Data);
            $$("rebirthDateCand").setValue(new Date(res.Data.BirthDate));
        }
    }
    xhr.send();
}

function createCandidate(lastName, firstName, midName, email, phone, birthDate, education){
    let xhr = new XMLHttpRequest();
    let selectedAssessmentId = $$("assessments").getSelectedItem().ID
    //console.log(selectedAssessmentId)
    let newCandidate = {
        Surname: lastName,
	    Name: firstName,
	    Patronymic: midName,
	    Email: email,
	    PhoneNumber: phone,
	    Education: education,
        BirthDate: birthDate,
    }
    //console.log("НОВЫЙ КАНДИДАТ:", newCandidate)
    //console.log(JSON.stringify(newCandidate))
    xhr.onreadystatechange = function() {
        //console.log(xhr.readyState, xhr.status, xhr.responseText)
        if (xhr.status == 200 && xhr.readyState == 4) {
            showCandidate(selectedAssessmentId);
        }
    }
    xhr.open("PUT", "/assessment/" + selectedAssessmentId + "/candidate");
    xhr.setRequestHeader("Content-Type", "application/json; charset=UTF-8");       
    xhr.send(JSON.stringify(newCandidate));
}

function editCandidate(lastName, firstName, midName, email, phone, birthDate, education){
    let xhr = new XMLHttpRequest();
    let selectedAssessmentId = $$("assessments").getSelectedItem().ID
    let selectedCandidateId = $$("peopleList").getSelectedItem().ID
    console.log(selectedAssessmentId, selectedCandidateId)
    let editedCandidate = {
        Surname: lastName,
	    Name: firstName,
	    Patronymic: midName,
	    Email: email,
	    PhoneNumber: phone,
	    Education: education,
        BirthDate: birthDate,
    }
    //console.log("ИЗМЕНЁННЫЙ КАНДИДАТ:", editedCandidate)
    //console.log(JSON.stringify(newCandidate))
    xhr.onreadystatechange = function() {
        //console.log(xhr.readyState, xhr.status, xhr.responseText)
        if (xhr.status == 200 && xhr.readyState == xhr.DONE) {
            showCandidate(selectedAssessmentId);
        }
    }
    xhr.open("POST", "/assessment/" + selectedAssessmentId + "/candidate/" + selectedCandidateId);
    xhr.setRequestHeader("Content-Type", "application/json; charset=UTF-8");       
    xhr.send(JSON.stringify(editedCandidate));
}

function removeCandidate() {
    let selectedCandidateId = $$("peopleList").getSelectedItem().ID
    let selectedAssessmentId = $$("assessments").getSelectedItem().ID
    let xhr = new XMLHttpRequest();
    xhr.open("DELETE", "/assessment/" + selectedAssessmentId + "/candidate/" + selectedCandidateId);
    xhr.onreadystatechange = function() {
        if (xhr.status == 200 && xhr.readyState == xhr.DONE) { /*
            let res = JSON.parse(xhr.response)
            if (res.Result === 1) {
                webix.message({type:"error", text:res.ErrorText});
                return
            } */
            $$("peopleList").remove(selectedCandidateId);
            showCandidate(selectedAssessmentId)
        }
    }       
    xhr.send();
}

function setCandidateStatus(selectedStatusId, status){
    let xhr = new XMLHttpRequest();
    let selectedAssessmentId = $$("assessments").getSelectedItem().ID
    let selectedCandidateId = $$("peopleList").getSelectedItem().ID
    let newStatus = {ID: selectedStatusId,
        Status: status}
    xhr.open("POST", "/assessment/" + selectedAssessmentId + "/candidate/" + selectedCandidateId + "/status/" + selectedStatusId);
    xhr.onreadystatechange = function() {
        console.log(xhr.readyState, xhr.status, xhr.responseText)
        if (xhr.status == 200 && xhr.readyState == xhr.DONE) {
            showCandidate(selectedAssessmentId);
        }
    }       
    xhr.send(JSON.stringify(newStatus));
}

function showCandidateStatus(){
    let xhr = new XMLHttpRequest();
    let selectedAssessmentId = $$("assessments").getSelectedItem().ID
    let selectedCandidateId = $$("peopleList").getSelectedItem().ID
    xhr.open("GET", "/assessment/" + selectedAssessmentId + "/candidate/" + selectedCandidateId + "/status/");
    xhr.onreadystatechange = function() {
        if (xhr.status == 200 && xhr.readyState == xhr.DONE) {
            let res = JSON.parse(xhr.response)
            if (res.Result === 1) {
                webix.message({type:"error", text:res.ErrorText});
                return
            }
            //$$("peopleStatusList").clearAll();
            $$("peopleList").parse(res.Data)
            console.log(res.Data)
        }
        }
    xhr.send();
}