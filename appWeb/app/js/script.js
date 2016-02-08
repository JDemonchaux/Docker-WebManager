$(document).ready(function () {
    $(".btPlay.container-running, .btPlay.container-paused, .btPause.container-stopped, .btStop.container-paused, .btStop.container-stopped").each(function () {
        $(this).on('click', function (e) {
            e.preventDefault();
        });
        $(this).addClass("disabled");
        $(this).find(">:first-child").prop("disabled", "disabled");
    })
});