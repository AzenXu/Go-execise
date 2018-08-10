const DEFAULT_COOKIE_EXPIRE_TIME = 30;

$(document).ready(function () {

    let uname = '';
    let session = '';

    session = _getCookie('session');
    uname = _getCookie('username');

    $("#reg-btn").on('click', function (e) {
        $("#reg-btn").text('Loading...');
        e.preventDefault();

        _registerUser(function (res, err, username) {
            if (err != null) {
                $('#reg-btn').text("Register");
                _popupErrorMsg(err);
                return;
            }

            let obj = JSON.parse(res);
            console.log("result: %s", obj);

            uname = username;
            session = obj["session_id"];

            _setCookie("session", session, DEFAULT_COOKIE_EXPIRE_TIME);
            _setCookie("username", uname, DEFAULT_COOKIE_EXPIRE_TIME);

            $("#reg-submit").submit();
        });
    });

    $("#sign-in-btn").on('click', function (e) {

        $("#sign-in-btn").text('Loading...')
        e.preventDefault();

        _signinUser(function (res, err, username) {
            if (err != null) {
                $('#sign-in-btn').text("Sign In");
                _popupErrorMsg(err);
                return;
            }

            let obj = JSON.parse(res);
            console.log("result: %s", obj);

            uname = username;
            session = obj["session_id"];

            _setCookie("session", session, DEFAULT_COOKIE_EXPIRE_TIME);
            _setCookie("username", uname, DEFAULT_COOKIE_EXPIRE_TIME);

            $("#sign-in-submit").submit();
        });
    });

    $("#sign-in-href").on('click', function () {
        $("#reg-submit").hide();
        $("#sign-in-submit").show();
    });

    $("#register-href").on('click', function () {
        $("#reg-submit").show();
        $("#sign-in-submit").hide();
    });
});