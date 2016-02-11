$(document).ready(function () {

    $("#bt-safe-delete").prop("disabled", true);

    $(".modal-trigger").leanModal({
        dismissible: true,
        ready: function () {
            var data = $('.json-data').text();
            $("#jjson").jJsonViewer(data);
        },
        complete: function () {
            $('.lean-overlay').each(function () {
                $(this).remove();
            })
        }
    });

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

    $(".button-collapse").sideNav();

    $("#input-safe-delete").on('keyup', function () {
        var deleteOk = false;
        console.log($(this).val());
        deleteOk = $(this).val() == "DELETE";


        if (deleteOk) {
            var url = $(".btDelete").data("href");
            $("#form-safe-delete").attr("action", url);
            $("#bt-safe-delete").prop("disabled", false);
        }
    });

    $("#bt-safe-delete").on('click', function () {
        $("#form-safe-delete").submit();
    });
});