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

function _asyncGetUserId(username) {
    let defer = $.Deferred();

    let dat = {
        'url': 'http://' + window.location.hostname + ':9000/user/' + username,
        'method': 'GET'
    };

    $.ajax({
        url: 'http://' + window.location.hostname + ':8080/api',
        type: 'post',
        data: JSON.stringify(dat),
        headers: {'X-Session-Id': session},
        statusCode: {
            500: function () {
                console.log("Internal Error");
                defer.reject("Internal Error");
            }
        },
        complete: function (xhr, textStatus) {
            if (xhr.status >= 400) {
                console.log("Error of getUserId");
                defer.reject("Error of getUserId");
            }
        }
    }).done(function (data, statusText, xhr) {
        console.log(data);
        defer.resolve(data);
    });

    return defer.promise();
}

function _asyncListAllVideos(uname) {
    let defer = $.Deferred();

    let dat = {
        'url': 'http://' + window.location.hostname + ':9000/user/' + uname + '/videos',
        'method': 'GET',
        'req_body': ''
    };

    $.ajax({
        url: 'http://' + window.location.hostname + ':8080/api',
        type: 'post',
        data: JSON.stringify(dat),
        headers: {'X-Session-Id': session},
        statusCode: {
            500: function () {
                defer.reject("Internal error");
            }
        },
        complete: function (xhr, textStatus) {
            if (xhr.status >= 400) {
                defer.reject("Error of Sign in");
            }
        }
    }).done(function (data, statusText, xhr) {
        if (xhr.status >= 400) {
            defer.reject("Error of Sign in");
            return;
        }
        defer.resolve(data);
    });

    return defer.promise();
}

let _asyncCreateVideo = (uname, vname, uid, session) => {

    let defer = $.Deferred();

    let reqBody = {
        'author_id': uid,
        'name': vname
    };

    let dat = {
        'url': 'http://' + window.location.hostname + ':9000/user/' + uname + '/videos',
        'method': 'POST',
        'req_body': JSON.stringify(reqBody)
    };

    $.ajax({
        url: 'http://' + window.location.hostname + ':8080/api',
        type: 'post',
        data: JSON.stringify(dat),
        headers: {'X-Session-Id': session},
        statusCode: {
            500: function () {
                defer.reject("Internal error");
            }
        },
        complete: function (xhr, textStatus) {
            if (xhr.status >= 400) {
                defer.reject("Error of Signin");
            }
        }
    }).done(function (data, statusText, xhr) {
        if (xhr.status >= 400) {
            defer.reject("Error of Signin");
            return;
        }
        defer.resolve(data);
    });

    return defer;
};

function _postComment(vid, content, callback) {
    var reqBody = {
        'author_id': uid,
        'content': content
    };


    var dat = {
        'url': 'http://' + window.location.hostname + ':9000/videos/' + vid + '/comments',
        'method': 'POST',
        'req_body': JSON.stringify(reqBody)
    };

    $.ajax({
        url  : 'http://' + window.location.hostname + ':8080/api',
        type : 'post',
        data : JSON.stringify(dat),
        headers: {'X-Session-Id': session},
        statusCode: {
            500: function() {
                callback(null, "Internal error");
            }
        },
        complete: function(xhr, textStatus) {
            if (xhr.status >= 400) {
                callback(null, "Error of Signin");
                return;
            }
        }
    }).done(function(data, statusText, xhr){
        if (xhr.status >= 400) {
            callback(null, "Error of Signin");
            return;
        }
        callback(data, null);
    });
}