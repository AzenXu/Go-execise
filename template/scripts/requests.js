function _registerUser(callback) {
    let username = $("#username").val();
    let pwd = $("#pwd").val();

    if (username === '' || pwd === '') {
        callback(null, "账号密码不能为空", null);
        return
    }

    let reqBody = {
        'user_name': username,
        'pwd': pwd
    };

    let dat = {
        'url': 'http://' + window.location.hostname + ':9000/user',
        'method': 'POST',
        'req_body': JSON.stringify(reqBody)
    };

    $.ajax({
        url: 'http://' + window.location.hostname + ':8080/api',
        type: 'post',
        data: JSON.stringify(dat),
        statusCode: {
            500: function () {
                callback(null, "internal error", null);
            }
        },
        complete: function (xhr, textStatus) {
            if (xhr.status >= 400) {
                callback(null, "Error of Signin", null);
            }
        }
    }).done(function (data, statusText, xhr) {
        if (xhr.status >= 400) {
            callback(null, "Error of register", null);
            return
        }

        callback(data, null, username);
    });
}

function _signinUser(callback) {
    let username = $("#s-username").val();
    let pwd = $("#s-pwd").val();

    if (username === '' || pwd === '') {
        callback(null, "账号密码不能为空", null);
    }

    let reqBody = {
        'user_name': username,
        'pwd': pwd
    };

    let dat = {
        'url': 'http://' + window.location.hostname + ':9000/user/' + username,
        'method': 'POST',
        'req_body': JSON.stringify(reqBody)
    };

    $.ajax({
        url: 'http://' + window.location.hostname + ':8080/api',
        type: 'post',
        data: JSON.stringify(dat),
        statusCode: {
            500: function () {
                callback(null, "Internal error");
            }
        },
        complete: function (xhr, textStatus) {
            if (xhr.status >= 400) {
                callback(null, "Error of Signin", null);
            }
        }
    }).done(function (data, statusText, xhr) {
        if (xhr.status >= 400) {
            callback(null, "Error of Signin", null);
            return;
        }

        callback(data, null, username);
    });
}