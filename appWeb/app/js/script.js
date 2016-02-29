$(document).ready(function () {

    $('a.disabled').on('click', function (e) {
        e.preventDefault();
    });
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

    // Show edit button close to container name
    $('td.name').hover(function () {
        // mouseIn
        $(this).find(".edit-button").css("visibility", "visible");
    }, function () {
        // mouseOut
        $(this).find(".edit-button").css("visibility", "hidden");
    });

    // pop an input for name editing
    $(".edit-button").on('click', function () {
        $(this).parent().find('.input-value').show();
        var value = $(this).parent().find('.edit-input').val();
        $(this).parent().find('.edit-input').val($.trim(value));
        $(this).parent().find('.show-value').hide();
        $(this).hide();
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
    $("#bt-safe-delete").prop("disabled", true);
    $("#bt-safe-delete").on('click', function () {
        var url = $("#form-safe-delete").attr("action", url);
        windows.location.href = url;
    });


    // Authentification
    $(".authentificate").on('click', function () {
        var username = $("#username").val();
        var pass = $("#password").val();
        var url = $("#formLogin").attr("action");
        console.log(url);
        var crypted = sha256(username + pass);
        var data = { auth: crypted };
        $.ajax({
            type: "POST",
            url: url,
            data: data
        });
        return false;
    });
});