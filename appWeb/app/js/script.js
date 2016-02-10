$(document).ready(function () {
    $(".btPlay.container-running, .btPlay.container-paused, .btPause.container-stopped, .btStop.container-paused, .btStop.container-stopped").each(function () {
        $(this).on('click', function (e) {
            e.preventDefault();
        });
        $(this).addClass("disabled");
        $(this).find(">:first-child").prop("disabled", "disabled");
    });


    $(".shorten-id").each(function () {
        var id = $(this).text();
        $(this).text(id.substr(0, 12));
    });

    var dialog = document.querySelector('dialog');
    var showDialogButton = document.querySelector('#raw');
    if (!dialog.showModal) {
        dialogPolyfill.registerDialog(dialog);
    }
    showDialogButton.addEventListener('click', function () {
        dialog.showModal();
    });
    dialog.querySelector('.close').addEventListener('click', function () {
        dialog.close();
    });
});