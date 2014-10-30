$('.input-group.date').datepicker({
    format: "yyyy-mm-dd",
    todayBtn: true,
    calendarWeeks: true,
    autoclose: true,
    todayHighlight: true
});

$('.input-group.date').datepicker('update', new Date());

var allLogs = {};

var myViewModel = {
  sessions: ko.observableArray(),
  logs: ko.observableArray()
};

function SessionModel(session) {
    this.session = session;
    this.active = ko.observable(false);
    this.click = function () {
        sessions = myViewModel.sessions();
        console.log("click", sessions.length);
        for (var i = 0; i < sessions.length; i++) {
            sessions[i].active(this == sessions[i]);
        }
        this.active(true);
        myViewModel.logs.removeAll();
        var sessions = allLogs[this.session];
        for (var i = 0; i < sessions.length; i++) {
            myViewModel.logs.push(sessions[i]);
        }
        hljs.initHighlightingOnLoad();
    };
}

ko.applyBindings(myViewModel);

$('#query').submit(function (event) {
    $.get('/search', {user: $("#user-input").val(), date: $("#date-input").val()}, function (data) {
        var lines = data.split("\n");
        myViewModel.sessions.removeAll();
        allLogs = {};
        for (var i = 0; i < lines.length; i++) {
            var line = lines[i];
            var cleanedLine = line.split($("#date-input").val() + "T")[1];
            if (!cleanedLine) {
                continue;
            }
            var tokens = cleanedLine.split("\t");
            var session = tokens[1].slice(tokens[1].lastIndexOf(".") + 1);
            if (!allLogs.hasOwnProperty(session)) {
                allLogs[session] = [];
                myViewModel.sessions.push(new SessionModel(session));
            }
            allLogs[session].push({date: tokens[0], log: tokens[2]});
        }
    })
    return false;
});
