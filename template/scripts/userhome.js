$(document).ready(function () {
    //  「用户页」方法绑定
    $("#upload").on('click', function () {
        $("#uploadvideomodal").show();
    });

    $("#uploadform").on('submit', function (e) {
        e.preventDefault()
        var vname = $('#vname').val();

        _createVideo(vname, function (res, err) {
            if (err != null) {
                //window.alert('encounter an error when try to create video');
                _popupErrorMsg('encounter an error when try to create video');
                return;
            }

            var obj = JSON.parse(res);
            var formData = new FormData();
            formData.append('file', $('#inputFile')[0].files[0]);

            //  上传视频
            $.ajax({
                url: 'http://' + window.location.hostname + ':8080/upload/' + obj['id'],
                //url:'http://127.0.0.1:8080/upload/dbibi',
                type: 'POST',
                data: formData,
                //headers: {'Access-Control-Allow-Origin': 'http://127.0.0.1:9000'},
                crossDomain: true,
                processData: false,  // tell jQuery not to process the data
                contentType: false,  // tell jQuery not to set contentType
                success: function (data) {
                    console.log(data);
                    $('#uploadvideomodal').hide();
                    location.reload();
                    //window.alert("hoa");
                },
                complete: function (xhr, textStatus) {
                    if (xhr.status === 204) {
                        window.alert("finish")
                        return;
                    }
                    if (xhr.status === 400) {
                        $("#uploadvideomodal").hide();
                        _popupErrorMsg('file is too big');
                        return;
                    }
                }

            });
        });
    });

    $(".close").on('click', function () {
        $("#uploadvideomodal").hide();
    });

    $("#logout").on('click', function () {
        setCookie("session", "", -1)
        setCookie("username", "", -1)
    });


    $(".video-item").click(function () {
        var url = 'http://' + window.location.hostname + ':9000/videos/' + this.id
        var video = $("#curr-video");
        video[0].attr('src', url);
        video.load();
    });
});

function _initPage(callback) {
    _getUserId(function (res, err) {
        if (err != null) {
            window.alert("Encountered error when loading user id");
            return;
        }

        var obj = JSON.parse(res);
        uid = obj['id'];
        //window.alert(obj['id']);

        _listAllVideos(function (res, err) {
            if (err != null) {
                //window.alert('encounter an error, pls check your username or pwd');
                _popupErrorMsg('encounter an error, pls check your username or pwd');
                return;
            }

            var obj = JSON.parse(res);
            listedVideos = obj['videos'];
            obj['videos'].forEach(function (item, index) {
                var ele = _htmlVideoListElement(item['id'], item['name'], item['display_ctime']);
                $("#items").append(ele);
            });
            callback();
        });
    });
}

function _htmlVideoListElement(vid, name, ctime) {
    //  创建一个a标签，把所有元素都插到a标签里
    var ele = $('<a/>', {
        href: '#'
    });
    ele.append(
        $('<video/>', {
            width: '320',
            height: '240',
            poster: '/statics/img/preloader.jpg',
            controls: true
            //href: '#'
        })
    );
    ele.append(
        $('<div/>', {
            text: name
        })
    );
    ele.append(
        $('<div/>', {
            text: ctime
        })
    );


    var res = $('<div/>', {
        id: vid,
        class: 'video-item'
    }).append(ele);

    res.append(
        $('<button/>', {
            id: 'del-' + vid,
            type: 'button',
            class: 'del-video-button',
            text: 'Delete'
        })
    );

    res.append(
        $('<hr>', {
            size: '2'
        }).css('border-color', 'grey')
    );

    return res;
}

function _popupErrorMsg(msg) {
    var x = document.getElementById("error-bar");
    $("#error-bar").text(msg);
    x.className = "show";
    setTimeout(function () {
        x.className = x.className.replace("show", "");
    }, 2000);
}

function _listAllVideos(callback) {
    var dat = {
        'url': 'http://' + window.location.hostname + ':8000/user/' + uname + '/videos',
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
                callback(null, "Internal error");
            }
        },
        complete: function (xhr, textStatus) {
            if (xhr.status >= 400) {
                callback(null, "Error of Signin");
                return;
            }
        }
    }).done(function (data, statusText, xhr) {
        if (xhr.status >= 400) {
            callback(null, "Error of Signin");
            return;
        }
        callback(data, null);
    });
}

function _getUserId(callback) {
    var dat = {
        'url': 'http://' + window.location.hostname + ':9000/user/' + uname,
        'method': 'GET'
    };

    $.ajax({
        url: 'http://' + window.location.hostname + ':8080/api',
        type: 'post',
        data: JSON.stringify(dat),
        headers: {'X-Session-Id': session},
        statusCode: {
            500: function () {
                callback(null, "Internal Error");
            }
        },
        complete: function (xhr, textStatus) {
            if (xhr.status >= 400) {
                callback(null, "Error of getUserId");
                return;
            }
        }
    }).done(function (data, statusText, xhr) {
        callback(data, null);
    });
}

function _createVideo(vname, callback) {
    var reqBody = {
        'author_id': uid,
        'name': vname
    };

    var dat = {
        'url': 'http://' + window.location.hostname + ':8000/user/' + uname + '/videos',
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
                callback(null, "Internal error");
            }
        },
        complete: function (xhr, textStatus) {
            if (xhr.status >= 400) {
                callback(null, "Error of Signin");
                return;
            }
        }
    }).done(function (data, statusText, xhr) {
        if (xhr.status >= 400) {
            callback(null, "Error of Signin");
            return;
        }
        callback(data, null);
    });
}

function _selectVideo(vid) {
    var url = 'http://' + window.location.hostname + ':8080/videos/' + vid
    var video = $("#curr-video");
    $("#curr-video:first-child").attr('src', url);
    $("#curr-video-name").text(currentVideo['name']);
    $("#curr-video-ctime").text('Uploaded at: ' + currentVideo['display_ctime']);
    //currentVideoId = vid;
    _refreshComments(vid);
}

function _refreshComments(vid) {
    _listAllComments(vid, function (res, err) {
        if (err !== null) {
            //window.alert("encounter an error when loading comments");
            popupErrorMsg('encounter an error when loading comments');
            return
        }

        var obj = JSON.parse(res);
        $("#comments-history").empty();
        if (obj['comments'] === null) {
            $("#comments-total").text('0 Comments');
        } else {
            $("#comments-total").text(obj['comments'].length + ' Comments');
        }
        obj['comments'].forEach(function (item, index) {
            var ele = htmlCommentListElement(item['id'], item['author'], item['content']);
            $("#comments-history").append(ele);
        });

    });
}

function _listAllComments(vid, callback) {
    var dat = {
        'url': 'http://' + window.location.hostname + ':8000/videos/' + vid + '/comments',
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
                callback(null, "Internal error");
            }
        },
        complete: function (xhr, textStatus) {
            if (xhr.status >= 400) {
                callback(null, "Error of Signin");
                return;
            }
        }
    }).done(function (data, statusText, xhr) {
        if (xhr.status >= 400) {
            callback(null, "Error of Signin");
            return;
        }
        callback(data, null);
    });
}

function _deleteVideo(vid, callback) {
    var dat = {
        'url': 'http://' + window.location.hostname + ':8000/user/' + uname + '/videos/' + vid,
        'method': 'DELETE',
        'req_body': ''
    };

    $.ajax({
        url: 'http://' + window.location.hostname + ':8080/api',
        type: 'post',
        data: JSON.stringify(dat),
        headers: {'X-Session-Id': session},
        statusCode: {
            500: function () {
                callback(null, "Internal error");
            }
        },
        complete: function (xhr, textStatus) {
            if (xhr.status >= 400) {
                callback(null, "Error of Signin");
                return;
            }
        }
    }).done(function (data, statusText, xhr) {
        if (xhr.status >= 400) {
            callback(null, "Error of Signin");
            return;
        }
        callback(data, null);
    });
}