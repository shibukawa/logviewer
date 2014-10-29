$('.input-group.date').datepicker({
    format: "yyyy-mm-dd",
    todayBtn: true,
    calendarWeeks: true,
    autoclose: true,
    todayHighlight: true
});

$('.input-group.date').datepicker('update', new Date());

hljs.initHighlightingOnLoad();

$('#query').submit(function (event) {
    console.log($("#user-input").val());
    console.log($("#date-input").val());
    $.get('/search', {user: $("#user-input").val(), date: $("#date-input").val()}, function (data) {
        console.log("result: ", data);
    })
    return false;
});
