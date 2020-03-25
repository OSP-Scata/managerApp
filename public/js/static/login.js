webix.ready(function(){   
    webix.ui({  
        view:"window",
        id:"loginWindow",
        scroll:false,
        modal:true,
        head: "Авторизация",
        position: "center",
        width:400,
        body:{ view:"form", id:"loginForm", scroll:false,
            elements:form,
            rules:{
                $all:webix.rules.isNotEmpty,
            },},
        
    }).show()
});

var form = [
    { view:"text", id: "userName", labelWidth: 100, name: "userName", label:"Имя пользователя"},
    { view:"text", id: "userPassword", labelWidth: 100, name: "userPass",type:"password", label:"Пароль"},
    { view:"button", label:"Войти",  href:"/", click: function(){
        var form1 = this.getParentView();
        var username = $$("userName").getValue();
        var password = $$("userPassword").getValue();
        if (form1.validate()){
           //window.location.href = "/";
           console.log(username, password);
           login(username, password);
           //form1.hide()
        }
        else{
            webix.message({ type:"error", text: "Form data is invalid" });
        }
    }}
]

function login(username, password) {
    let xhr = new XMLHttpRequest();
    xhr.open("GET", "/login", true, username, password);
    xhr.onreadystatechange = function() {
        console.log(xhr.readyState, xhr.status, xhr.responseText)
		if (xhr.status == 200 && xhr.readyState == 4) {
			$$("loginWindow").hide();
		}
    }
    xhr.send();
}