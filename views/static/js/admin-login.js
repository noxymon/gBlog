"use strict";

$(document).ready(
    $("#formLogin").submit(
        function (e) {
            e.preventDefault();
            authUser();
        }
    ),
)

function authUser() {
    var usernameInput = $("#inputUsername").val();
    var passwordInput = $("#inputPassword").val();
    var postToLoginApi = $.post("/admin/login",
        {
            username: usernameInput,
            password: passwordInput
        }
    );

    postToLoginApi.done(function (data) {
        $(location).attr('href', '/admin')
    })
}