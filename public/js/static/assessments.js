var statusList = [ 
    { id:1, value:"Назначен"}, 
    { id:2, value:"Проведён"}, 
    { id:3, value:"Отменён"}
]

webix.editors.setStatus = webix.extend({
    getValue:function(){
        let id = this.getPopup().getValue();
        let status = statusList[id-1].value
        setAssessmentStatus(id, status)
    }
}, webix.editors.richselect);

var assessmentTable = {
    view: "datatable", 
        id: "assessments", 
        scroll:true, 
        width: 450,
        //url:"/assessmentsDB.json",
        editable: true,
        select: true,
        columns:[
            //{ id:"countCand", header:"Участники", width: 100},
            { id:"changeStatus", template:"#StatusName#", header:"Статус",
                width: 170, editor:"setStatus", options:statusList
            },
            { id:"date", template:" #Date# <span class='webix_icon wxi-trash' title='Удалить ассессмент'></span>",
            header:["Дата ассессмента", {content:"textFilter"}], sort:"string", width: 200},
        ], 
        onClick:{
            "wxi-trash":function(ev, id){
                removeAssessment();
                $$("peopleList").clearAll();
                $$("interviewerList").clearAll();
                $$("delete").disable();
                $$("addCand").disable();
                $$("addSob").disable();
                $$("edit").disable();
            }
        },
        on: {
            onItemClick: function(){
                //showAssessmentStatus();
                //console.log("Returned:" + data)
                //$$("changeStatus").parse(data);
                //$$('peopleList').parse(people);
                //$$("status").parse();
                showCandidate();
                showInterviewer();
                //$$('interviewerList').parse(interviewer);
                $$("delete").enable();
                $$("addCand").enable();
                $$("addSob").enable();
                $$("edit").enable();
            },
        }
}

var btnAddAssessment = {
    view:"button", label:"Добавить", type:"icon", icon:"wxi-plus", click: function () {
        addAssess.show();
    }
}

var addAssess = webix.ui({
    view: "window",
    head: "Добавить ассессмент",
    id: "dateAssessment",
    width: 400,
    modal: true,
    close:true,
    position: "center",
    body:{view:"form", id:"addAssessForm", scroll:false,
        elements:[
        { view:"datepicker",
        id: "datePic",
        label: "Дата ассессмента",
        labelWidth: 150,
        name: "date",
        stringResult: true,
        timepicker: true
        },
        { view:"button", value:"Сохранить", click:addAssessment},
        ]}
});

function addAssessment(){
    if($$("datePic").getValue() != ""){
        idAssessCounter += 1;
        date = $$("datePic").getValue();
        console.log("Selected date:", date)
        sendAssessment(date);
        //$$("assessments").add({idAssess: idAssessCounter, date: $$("datePic").getValue(), status: "Назначен", countCand: 0},  0);
        $$("addAssessForm").clear();
        this.getParentView().getParentView().hide()
    }
}

var btnEditAssessment = {
    view:"button", label:"Изменить", type:"icon", icon:"wxi-pencil", click: editAssessFunc
}

function editAssessFunc(){
    if(!$$("assessments").getSelectedId()){
        webix.message("No item is selected!");
        return;
    }
    console.log($$("assessments").getSelectedItem())
    editAssess.show();
}

var editAssess = webix.ui({
    view: "window",
    head: "Изменить ассессмент",
    id: "dateAssessEdit",
    width: 400,
    modal: true,
    close:true,
    position: "center",
    body:{view:"form", id:"forma", scroll:false,
        elements:[
            { view:"datepicker",
            id: "datePica",
            label: 'Дата ассессмента',
            labelWidth: 150,
            name: "date",
            stringResult: true,
            timepicker: true
            },
            { view:"button", value:"Сохранить", click: editAssessment},
        ]}
});

function editAssessment(){
    /* $$("assessments").getSelectedItem().date = $$("datePica").getValue();
    $$("assessments").refresh(); */
    date = $$("datePica").getValue();
    console.log("Change date to:", date)
    updateAssessment(date);
    //$$("assessments").add({idAssess: idAssessCounter, date: $$("datePic").getValue(), status: "Назначен", countCand: 0},  0);
    $$("addAssessForm").clear();
    this.getParentView().getParentView().hide()
}