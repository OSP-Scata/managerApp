var candidateStatus = [
    { id:1, value:"Приглашен"}, 
    { id:2, value:"Пришел"}, 
    { id:3, value:"Не пришел"},
    { id:4, value:"Завершил"}, 
    { id:5, value:"Не завершил"},
    { id:6, value:"Принят"}, 
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
        { id:"ID", header:"#", width:40},
        { id:"Surname", header:"Фамилия", width:100},
        { id:"Name", header:"Имя", width:100},
        { id:"Patronymic", header:"Отчество", width:100},
        { id:"Email", header:"E-mail", width:170},
        { id:"BirthDate", header:"Дата рождения", width:120},
        { id:"PhoneNumber",  header:"Телефон", width:150},
        { id:"Status", template:"#StatusName#", editor:"setCandStatus", options:candidateStatus, value:3, header:"Статус", width:120},
        { id:"Address", header:"Адрес", width:170},
        { id:"Education", header:"Образование", width:170},
        { id:"Resume", header:"Резюме", width:170},
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
        { view:"text", id:"surnameCand", name:"surname", label:"Фамилия"},
        { view:"text", id:"nameCand", name:"name", label:"Имя"},
        { view:"text", id:"patronymicCand", name:"patronymic", label:"Отчество"},
        { view:"text", id:"emailCand", name:"email", label:"E-mail"},
        { view:"text", id:"phoneCand", name:"phone", label:"Телефон"},
        //{ view:"richselect", suggest:{ data:candStatus}, id:"statusCand", value:3, name:"status", label:"Статус"},
        { view:"text", id:"addressCand", name:"address", label:"Адрес"},
        { view:"datepicker", id:"birthDateCand", name:"birthDate", label:"Дата рождения", stringResult:true},
        { view:"text", id:"educationCand", name:"education", label:"Образование"},
        { view:"text", id:"resumeCand", name:"resume", label:"Резюме"},
        { cols:[{ view:"button", value:"Добавить", click:addPeople},
        { view:"button", value:"Отмена", click:function(){
            $$("surnameCand").setValue("");
            $$("nameCand").setValue("");
            $$("patronymicCand").setValue("");
            $$("emailCand").setValue("");
            $$("phoneCand").setValue("");
            $$("statusCand").setValue("");
            this.getTopParentView().hide(); 
          }}]}
      ]
    },
    move:true
});

function addPeople(){
    if($$("nameCand").getValue() != ""){
        idCandCounter += 1;
        /*
        $$("peopleList").add({idCand:idCandCounter,
            surname:$$("surnameCand").getValue(), 
            name:$$("nameCand").getValue(),
            patronymic:$$("patronymicCand").getValue(),
            email:$$("emailCand").getValue(), 
            phoneNumber:$$("phoneCand").getValue(), 
            status:$$("statusCand").getValue(),
            //resume:$$("resumeCand").getValue(),
            address:$$("addressCand").getValue(),
            birthDate:$$("birthDateCand").getValue(),
            education:$$("educationCand").getValue()},  0);
        */
        //let idCand = idCandCounter;
        let surname = $$("surnameCand").getValue(); 
        let name = $$("nameCand").getValue();
        let patronymic = $$("patronymicCand").getValue();
        let email = $$("emailCand").getValue(); 
        let phoneNumber = $$("phoneCand").getValue(); 
        //let status = $$("statusCand").getValue();
        let resume = $$("resumeCand").getValue();
        let address = $$("addressCand").getValue();
        let birthDate = $$("birthDateCand").getValue();
        let education = $$("educationCand").getValue()
        console.log(birthDate);
        createCandidate(surname, name, patronymic, email, phoneNumber, address, birthDate, education, resume);
        $$("addCandidate").clear();
        this.getParentView().getParentView().getParentView().hide()
    }
}

function removeData(){
    if(!$$("peopleList").getSelectedId()){
        webix.message("No item is selected!");
        return;
    }
    //$$("peopleList").remove($$("peopleList").getSelectedId());
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
    //$$("editPeople").show();
    //$$("editPeople").attachEvent("onFocus", function() {console.log("test");})
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
        //{ view:"richselect", suggest:{ data:candStatus}, id:"restatusCand", name:"status", label:"Статус"},
        { view:"text", id:"reresumeCand", name:"Resume", label:"Резюме"},
        { view:"text", id:"readdressCand", name:"Address", label:"Адрес"},
        { view:"datepicker", stringResult:true, id:"rebirthDateCand", name:"BirthDate", label:"Дата рождения"},
        { view:"text", id:"reeducationCand", name:"Education", label:"Образование"},
        { cols:[{ view:"button", value:"Изменить", click:editCand},
        { view:"button", value:"Отмена", click:function(){
            this.getTopParentView().hide(); 
          }}]}
      ]
    },
    move:true,
    //on: {onFocus: function() {showCandidateById(); console.log("test");}}
});

function editCand(){
    let surname = $$("resurnameCand").getValue();
    let name = $$("renameCand").getValue();
    let patronymic = $$("repatronymCand").getValue();
    let email = $$("reemailCand").getValue();
    let phoneNumber = $$("rephoneCand").getValue();
    //let status = $$("restatusCand").getValue();
    let resume = $$("reresumeCand").getValue();
    let birthDate = $$("rebirthDateCand").getValue();
    let address = $$("readdressCand").getValue();
    let education = $$("reeducationCand").getValue();
    editCandidate(surname, name, patronymic, email, phoneNumber, address, birthDate, education, resume)
    this.getParentView().getParentView().getParentView().hide()
}
