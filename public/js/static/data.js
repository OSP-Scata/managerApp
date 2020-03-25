idAssessCounter = 3;
idInterviewerCounter = 3;
idCandCounter = 2;

candStatus = [
    {id:1, value: "Принят"},
    {id:2, value: "Не принят"},
    {id:3, value: "Записан"},
    {id:4, value: "Пришёл"},
    {id:5, value: "Не пришёл"}
];

assessStatus = [
    {idA:1, value: "Назначен"},
    {idA:2, value: "Проведён"},
    {idA:3, value: "Не проведён"},
  ];

people = [
    {idCand:1, surname:"Иванов", name:"Иван", patronymic:"Иванович", email:"ivanov@mail.ru",
        phoneNumber:"89871234567", birthDate:"26-11-1988", address:"улица, дом", education:"высшее", status:"Записан"},
    {idCand:2, surname:"test", name:"test", patronymic:"test", email:"test@test.com",
        phoneNumber:"12345678900", birthDate:"01-01-1976", address:"test", education:"test", status:"Записан"},
]

assessments = [
    {idAssess:1, date:"2019-12-26", status:"Не проведён"},
    {idAssess:2, date:"2020-01-20", status:"Проведён"},
    {idAssess:3, date:"2020-02-14", status:"Назначен"},
]

interviewer = [
    {idInterviewer: 1, name: "Сотрудник 1"},
    {idInterviewer: 2, name: "Сотрудник 2"},
    {idInterviewer: 3, name: "Сотрудник 3"},
]