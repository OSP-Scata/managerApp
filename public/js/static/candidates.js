var candidateStatus = [
    { id:1, value:"Приглашён"}, 
    { id:2, value:"Не явился"}, 
    { id:3, value:"Завершил"}, 
    { id:4, value:"Не завершил"},
    { id:5, value:"Принят на обучение"},
    { id:5, value:"Принят на работу"},
    { id:7, value:"Не принят"}
]

webix.editors.setCandStatus = webix.extend({
    getValue:function(){
        let id = this.getPopup().getValue();
        let status = candidateStatus[id-1].value
        setCandidateStatus(id, status)
    }
}, webix.editors.richselect);

var peopleTable = {
    view:"datatable", id:"peopleList", autoWidth:true, editable:true, select:true, columns:[
        { id:"id", header:"#", width:40, template:"#ID#"},
        { id:"lastname", header:"Фамилия", width:100, template:"#Surname#"},
        { id:"firstname", header:"Имя", width:100, template:"#Name#"},
        { id:"midname", header:"Отчество", width:100, template:"#Patronymic#"},
        { id:"birthdate", header:"Дата рождения", width:120, template:"#BirthDate#"},
        { id:"email", header:"E-mail", width:170, template:"#Email#"},
        { id:"phone",  header:"Телефон", width:150, template:"#PhoneNumber#"},
        { id:"education", header:"Образование", width:170, template:"#Education#"},
        { id:"status", template:"#StatusName#", view:"richselect", value:3, header:"Статус", width:120},
    ], on:{onItemDblClick:editPeople}
}

var btnAddCand = {
    view:"button", id:"addCand", label:"Участник", disabled:true, type:"icon", icon:"wxi-plus", click:function () {
        addCand.show();
    }
}

var btnRemoveCand = {
    view:"button", id:"delete", label:"Удалить", disabled:true, type:"icon", icon:"wxi-trash", click:removeData
}

var btnEditCand = {
    view:"button", id:"edit",  label:"Изменить", disabled:true, type:"icon", icon:"wxi-pencil", click:function (){editPeople()}
}

var addCand = webix.ui({
    view:"window",
    head:"Добавить участника",
    modal:true,
    width:500,
    close:true,
    position:"center",
    body:{
      view:"form",
      id:"addCandidate",
      editable:true,
      elements:[
        { view:"text", id:"newLastname", name:"lastName", label:"Фамилия"},
        { view:"text", id:"newFirstname", name:"name", label:"Имя"},
        { view:"text", id:"newMidname", name:"midName", label:"Отчество"},
        { view:"datepicker", id:"birthDateCand", name:"birthDate", label:"Дата рождения", stringResult:true},
        { view:"text", id:"newEmail", name:"email", label:"E-mail"},
        { view:"text", id:"newPhone", name:"phone", label:"Телефон"},
        { view:"text", id:"newEducation", name:"education", label:"Образование"},
        { cols:[{ view:"button", value:"Добавить", click:addPeople},
        { view:"button", value:"Отмена", click:function(){
            $$("newLastname").setValue("");
            $$("newFirstname").setValue("");
            $$("newMidname").setValue("");
            $$("newEmail").setValue("");
            $$("newPhone").setValue("");
            this.getTopParentView().hide(); 
          }}]}
      ]
    },
    move:true
});

function addPeople(){
    if($$("newLastname").getValue() != ""){
        idCandCounter += 1;
        let lastName = $$("newLastname").getValue(); 
        let firstName = $$("newFirstname").getValue();
        let midName = $$("newMidname").getValue();
        let email = $$("newEmail").getValue(); 
        let phoneNumber = $$("newPhone").getValue(); 
        let birthDate = $$("birthDateCand").getValue();
        let education = $$("newEducation").getValue()
        console.log(birthDate);
        createCandidate(lastName, firstName, midName, email, phoneNumber, birthDate, education);
        $$("addCandidate").clear();
        this.getParentView().getParentView().getParentView().hide()
    }
}

function removeData(){
    if(!$$("peopleList").getSelectedId()){
        webix.message("No item is selected!");
        return;
    }
    removeCandidate()
}

function editPeople(){
    if(!$$("peopleList").getSelectedId()){
        webix.message("No item is selected!");
        return;
    }
    console.log("test");
    showCandidateById();
    editCandWindow.show()
}

var editCandWindow = webix.ui({
    view:"window",
    head:"Изменить участника",
    id:"editPeople",
    modal:true,
    width:500,
    editable:true,
    position:"center",
    body:{
      view:"form",
      id:"editForm",
      elements:[
        { view:"text", id:"resurnameCand", name:"Surname", label:"Фамилия" },
        { view:"text", id:"renameCand", name:"Name", label:"Имя"},
        { view:"text", id:"repatronymCand", name:"Patronymic", label:"Отчество"},
        { view:"text", id:"reemailCand", name:"Email", label:"E-mail"},
        { view:"text", id:"rephoneCand", name:"PhoneNumber", label:"Телефон"},
        { view:"datepicker", stringResult:true, id:"rebirthDateCand", name:"BirthDate", label:"Дата рождения"},
        { view:"text", id:"reeducationCand", name:"Education", label:"Образование"},
        { cols:[{ view:"button", value:"Изменить", click:editCand},
        { view:"button", value:"Отмена", click:function(){
            this.getTopParentView().hide(); 
          }}]}
      ]
    },
    move:true,
});

function editCand(){
    let lastName = $$("resurnameCand").getValue();
    let firstName = $$("renameCand").getValue();
    let midName = $$("repatronymCand").getValue();
    let email = $$("reemailCand").getValue();
    let phoneNumber = $$("rephoneCand").getValue();
    let birthDate = $$("rebirthDateCand").getValue();
    let education = $$("reeducationCand").getValue();
    editCandidate(lastName, firstName, midName, email, phoneNumber, birthDate, education)
    this.getParentView().getParentView().getParentView().hide()
}