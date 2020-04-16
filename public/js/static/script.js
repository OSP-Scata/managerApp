webix.ready(function(){
    webix.ui({
        rows:[
        {
            view:"toolbar", cols:[
                {view: "label", label: "Assessment Manager"},
                {view: "button", label: "Выйти", href:"/logout",  width: 100, click: function(){
                    window.location.href = "/logout";
                    $$("loginWindow").show();
                    }
                }
            ]
        },
        {view:"tabview",
        animate:true,
        cells:[ 
            {header:"Ассессменты",
            body:{
                rows:[
                    {cols:[
                        {rows:[
                            {cols:[
                            { view:"label", label: "<span class='webix_icon wxi-calendar' title='Calendar'></span>Ассессменты"},
                            btnAddAssessment,
                            btnEditAssessment
                        ]},
                            assessmentTable
                        ]},
                    {view: "resizer"},
                    {rows:[
                        {cols:[
                        {view:"label", label: "<span class='webix_icon wxi-user' title='User'></span>Кандидаты"},
                        btnAddCand,
                        btnEditCand,
                        btnRemoveCand,
                    ]},
                        peopleTable,
                        {cols: [
                        { view:"label", label: "<span class='webix_icon wxi-user' title='User'></span>Сотрудники"},
                        btnAddInterviewer,
                    ]},
                        interviewerTable,
                    ]}
                    ]}  
                ]
            }
        },
        {
            header:"Сотрудники",
            body:{
                rows:[
                    {cols:[
                        {view:"label", label: "<span class='webix_icon wxi-user' title='User'></span>Сотрудники"},
                        btnAddIntToDictionary,
                        btnEditIntInDictionary,
                        btnRemoveIntFromDictionary
                    ]},
                    allInterviewer
                ]
            }
          }
    ]}]
    }),
    showAssessment();
    showAllInterviewer("InterviewerDictionary");
});