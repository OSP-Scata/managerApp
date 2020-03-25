var interviewerTable = {
    view: "datatable", id: "interviewerList",
    columns:[
        { id:"ID", template:"#ID#", header:"#", width: 40},
        { id:"surname", template:"#Surname#", header:"Фамилия", width: 120},
        { id:"name", template:"#Name#", header:"Имя", width: 120},
        { id:"patronymic", template:"#Patronymic#", header:"Отчество", width: 120},
        { id:"email", template:"#Email#", header:"Email", width: 170},
        { id:"phoneNumber",  template:"#PhoneNumber#", header:"Телефон", width: 150},
        { id:"position", template:"#Position#", header: "Должность", width: 200},
    ],
    select: true,
    scroll: true, autoConfig: true, height: 150
}
var btnAddInterviewer = {
    view:"button", id: "addSob", label:"Cобеседователь", disabled:true, type:"icon", icon:"wxi-plus", popup: "addInterviewer"
 }

 webix.ui({
    view:"popup",
    on: {
        onFocus: function() {
            showAllInterviewer("popupList")
        }
    },
    id:"addInterviewer",
    body:{view:"form", id:"form1", scroll:false,
        elements:[
            {view:"list",
            id: "popupList",
            width:250,
            height:200,
            template:"#Surname# #Name# #Patronymic#",
            select:true,
        },
            {cols:[{ view: "button", value: "Добавить", click: function(){
                let selectedInterviewerId = $$("popupList").getSelectedItem()
                if(!$$("popupList").getSelectedId()){
                    webix.message("Сотрудник не выбран из списка");
                    return;
                }
                else{
                    addInterviewer(selectedInterviewerId.Surname, selectedInterviewerId.Name, selectedInterviewerId.Patronymic,
                    selectedInterviewerId.Email, selectedInterviewerId.PhoneNumber, selectedInterviewerId.Position);
                    this.getParentView().getParentView().hide()
                }
        }},
        { view:"button", value:"Отмена", click:function(){
            this.getTopParentView().hide(); 
          }},
        ]}
    ]}
});

var allInterviewer = {
    view: "datatable", id: "InterviewerDictionary", 
    columns:[
        { id:"ID", template:"#ID#", header:"#", width: 50},
        { id:"surname", template:"#Surname#", header:"Фамилия", width: 200},
        { id:"name", template:"#Name#", header:"Имя", width: 200},
        { id:"patronymic", template:"#Patronymic#", header:"Отчество", width: 200},
        { id:"email", template:"#Email#", header:"Email", width: 230},
        { id:"phoneNumber",  template:"#PhoneNumber#", header:"Телефон", width: 200},
        { id:"position", template:"#Position#", header: "Должность", width: 260},
    ],
    select: true,
    scroll: true, autoConfig: true
}

var btnAddIntToDictionary = {
    view:"button", id: "addSobToD", type:"icon", icon:"wxi-plus", width: 50, click: function(){
        addInterToD.show();
    }
}

var btnRemoveIntFromDictionary = {
    view:"button", id: "deleteSobFromD", type:"icon", icon:"wxi-trash", width: 50, click: function(){
        if(!$$("InterviewerDictionary").getSelectedId()){
            webix.message("Сотрудник не выбран");
            return;
        }
        removeInterviewerFromD("InterviewerDictionary")
    }
}

var btnEditIntInDictionary = {
    view:"button", id: "editSobInD", type:"icon", icon:"wxi-pencil", width: 50, click: function(){ editSob.show(); }
}
/*
var btnAddInterviewer = {
    view:"button", id: "addSob", disabled:true, type:"icon", icon:"wxi-plus", width: 60, popup: "addInterviewer", 
    click: showAllInterviewer("popupList")
}
*/
var btnRemoveInterviewer = {
    view:"button", id: "deleteSob", disabled:true, type:"icon", icon:"wxi-trash", width: 60, click: removeData 
}

function removeData(){
    if(!$$("interviewerList").getSelectedId()){
        webix.message("Сотрудник не выбран");
        return;
    }
    
    if(!$$("assessments").getSelectedId()){
        webix.message("Ассессмент не выбран");
        return;
    }

    removeInterviewer();
}

var addInterToD = webix.ui({
    view: "window",
    head: "Добавить сотрудника",
    width: 500,
    editable: true,
    position: "center",
    move: true,
    close: true,
    id:"addInterviewerToDictionary",
    body:{view:"form", id:"formAddInter", scroll:false,
        elements:[
            { view: "text", id:"surnameInt", label:"Фамилия"},
            { view: "text", id:"nameInt", label:"Имя"},
            { view: "text", id:"patronymicInt", label:"Отчество"},
            { view: "text", id: "emailInt", label: "Email"},
            { view: "text", id: "phoneInt", label: "Телефон"},
            { view: "text", id: "positionInt", label: "Должность"},
            { cols:[{ view: "button", value: "Добавить", click: function(){
                if($$("nameInt").getValue() != ""){
                    let surname = $$("surnameInt").getValue(); 
                    let name = $$("nameInt").getValue();
                    let patronymic = $$("patronymicInt").getValue();
                    let email = $$("emailInt").getValue(); 
                    let phoneNumber = $$("phoneInt").getValue();
                    let position = $$("positionInt").getValue();
                    createInterviewer("InterviewerDictionary", surname, name, patronymic, email, phoneNumber, position)
                    this.getParentView().getParentView().hide()
                }
        }},
        { view:"button", value:"Отмена", click:function(){
            $$("surnameInt").setValue("");
            $$("nameInt").setValue("");
            $$("patronymicInt").setValue("");
            $$("emailInt").setValue("");
            $$("phoneInt").setValue("");
            $$("positionInt").setValue("");
            this.getTopParentView().hide(); 
        }},
        ]}
    ]}
});

var editSob = webix.ui({
    view: "window",
    head: "Изменить сотрудника",
    id: "editInterviewer",
    width: 500,
    editable: true,
    position: "center",
    body: {
      view: "form",
      id: "editInterviewerForm",
      elements: [
        { view: "text", id: "resurnameSob", name: "Surname", label: "Фамилия"},
        { view: "text", id: "renameSob",  name: "Name",label: "Имя"},
        { view: "text", id: "repatronymicSob",  name: "Patronymic",label: "Отчество"},
        { view: "text", id: "reemailSob",  name: "Email",label: "Email"},
        { view: "text", id: "rephoneSob",  name: "PhoneNumber",label: "Телефон"},
        { view: "text", id: "repositionSob",  name: "Position", label: "Должность"},
        {cols:[{ view: "button", value: "Изменить", click: function(){
                if($$("renameSob").getValue() != ""){
                    updateInterviewer($$("resurnameSob").getValue(), 
                    $$("renameSob").getValue(), 
                    $$("repatronymicSob").getValue(), 
                    $$("reemailSob").getValue(), 
                    $$("rephoneSob").getValue(), 
                    $$("repositionSob").getValue()); 
                }
                this.getParentView().getParentView().getParentView().hide()
            }
        },
        { view:"button", value:"Отмена", click:function(){
            this.getTopParentView().hide(); 
          }}]}
      ]
    },
    move: true
});