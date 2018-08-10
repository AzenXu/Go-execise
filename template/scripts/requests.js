function _registerUser(callback) {
    var username = $("#username").val();
    var pwd = $("#pwd").val();
    var apiUrl = window.location.hostname + ':8080/api';

    if (username === '' || pwd === '') {
        callback(null, "账号密码不能为空", null);
        return
    }

    var reqBody = {
        'user_name': username,
        'pwd': pwd
    };

    var dat = {
        'url': 'http://'+ window.location.hostname + ':9000/user',
        'method': 'POST',
        'req_body': JSON.stringify(reqBody)
    };

    $.ajax({
        url  : 'http://' + window.location.hostname + ':8080/api',
        type : 'post',
        data : JSON.stringify(dat),
        statusCode: {
            500: function() {
                callback(null, "internal error", null);
            }
        },
        complete: function(xhr, textStatus) {
            if (xhr.status >= 400) {
                callback(null, "Error of Signin", null);
            }
        }
    }).done(function(data, statusText, xhr){
        if (xhr.status >= 400) {
            callback(null, "Error of register", null);
            return
        }

        callback(data, null, username);
    });
}