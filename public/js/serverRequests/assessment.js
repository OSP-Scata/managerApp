function sendAssessment(datePick){
    let xhr = new XMLHttpRequest();
    let newAssessment = {Date: datePick}
    //console.log(JSON.stringify(newAssessment))
    /*xhr.onreadystatechange = function() {
        console.log(xhr.readyState, xhr.status, xhr.responseText)
    }*/
    xhr.open("PUT", "/assessment", true);
    xhr.setRequestHeader("Content-Type", "application/json; charset=UTF-8");       
    xhr.send(JSON.stringify(newAssessment));
    xhr.onreadystatechange = function() {
        if (xhr.status == 200 && xhr.readyState == 4) {
            showAssessment();
        }
    }
}

function showAssessment(){
    let xhr = new XMLHttpRequest();
    xhr.open("GET", "/assessment");
    xhr.setRequestHeader("Content-Type", "application/json; charset=UTF-8");
    xhr.onreadystatechange = function() {
        console.log(xhr.readyState, xhr.status, xhr.responseText)
        if (xhr.status == 200 && xhr.readyState == 4) {
            console.log(xhr.readyState, xhr.status, xhr.responseText)
            let res = JSON.parse(xhr.response)
            if (res.Result == 1) {
                webix.message({type:"error", text:res.ErrorText});
                return
            }
            $$("assessments").clearAll();
            $$("assessments").parse(res)
        }
    }
    xhr.send();
}

function removeAssessment(){
    let selectedAssessmentId = $$("assessments").getSelectedItem().ID
    console.log("RemoveAssessment called. ID to remove:", selectedAssessmentId)
    let xhr = new XMLHttpRequest();
    xhr.open("DELETE", "/assessment/" +selectedAssessmentId);
    xhr.onreadystatechange = function() {
        console.log(xhr.readyState, xhr.status, xhr.responseText)
        if (xhr.status == 200 && xhr.readyState == 4) {
            showAssessment()
        }
    }       
    xhr.send();
}

function updateAssessment(datePick){
    let selectedAssessmentId = $$("assessments").getSelectedItem().ID
    let newAssessment = { Date: datePick }
    console.log("New date:", JSON.stringify(newAssessment))
    let xhr = new XMLHttpRequest();
    xhr.open("POST", "/assessment/" + selectedAssessmentId, true);
    xhr.onreadystatechange = function() {
        console.log(xhr.readyState, xhr.status, xhr.responseText)
        if (xhr.status == 200 && xhr.readyState == 4) {
            showAssessment();
        }
    }
    xhr.send(JSON.stringify(newAssessment));
}

function showAssessmentStatus(){
    let xhr = new XMLHttpRequest();
    let selectedAssessmentId = $$("assessments").getSelectedItem().ID
    xhr.open("GET", "/assessment/" + selectedAssessmentId + "/status");
    xhr.onreadystatechange = function() {
        console.log(xhr.readyState, xhr.status, xhr.responseText)
    }
    xhr.send();
}

function setAssessmentStatus(selectedStatusId, status){
    let selectedAssessmentId = $$("assessments").getSelectedItem().ID
    let xhr = new XMLHttpRequest();
    let newStatus = {ID: selectedStatusId,
        Status: status}
    //console.log(JSON.stringify(newStatus));
    xhr.open("POST", "/assessment/" + selectedAssessmentId + "/status/" + selectedStatusId, true);
    xhr.onreadystatechange = function() {
        console.log(xhr.readyState, xhr.status, xhr.responseText)
        if (xhr.status == 200 && xhr.readyState == 4) {
            showAssessment();
        }
    }
    setStatusesInAssessment(selectedAssessmentId, selectedStatusId)
    xhr.send(JSON.stringify(newStatus));
}