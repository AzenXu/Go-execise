let uid;

let session = _getCookie('session');
let username = _getCookie('username');

let listedVideos;
let currentVideo;

$(document).ready(function () {

    _initPage().then(() => {

        if (listedVideos !== null) {
            currentVideo = listedVideos[0];
            _selectVideo(listedVideos[0]['id']);
        }

        _asyncBindElement();
    }, (e) => {
        alert(e);
    });

    $("#upload").on('click', function () {
        $("#upload-video-modal").show();
    });

    $("#upload-form").on('submit', function (e) {
        e.preventDefault();
        var vname = $('#vname').val();

        _asyncCreateVideo(username, vname, uid, session).then((res) => {
            var obj = JSON.parse(res);
            var formData = new FormData();
            formData.append('file', $('#inputFile')[0].files[0]);

            //  上传视频
            $.ajax({
                url: 'http://' + window.location.hostname + ':8080/upload/' + obj['id'],
                type: 'POST',
                data: formData,
                crossDomain: true,
                processData: false,  // tell jQuery not to process the data
                contentType: false,  // tell jQuery not to set contentType
                complete: function (xhr, textStatus) {
                    if (xhr.status === 204) {
                        window.alert("finish");
                        return;
                    }
                    if (xhr.status === 400) {
                        $("#upload-video-modal").hide();
                        _popupErrorMsg('file is too big');
                    }
                }
            }).done((data, statusText, xhr) => {
                if (xhr.status >= 400) {
                    alert("Error of Sign in");
                    return;
                }
                alert("🎉上传成功~~~~~");
                $('#upload-video-modal').hide();
                location.reload();
            });

        }, (err) => {
            window.alert(err);
        });
    });

    $(".close").on('click', function () {
        $("#upload-video-modal").hide();
    });

    $("#logout").on('click', function () {
        _setCookie("session", "", -1);
        _setCookie("username", "", -1);
    });

    // $("#submit-comment").on('click', function() {
    //     let content = $("#comments-input").val();
    //     _postComment(currentVideo['id'], content, function(res, err) {
    //         if (err !== null) {
    //             alert("encounter and error when try to post a comment: " + content);
    //             return;
    //         }
    //
    //         if (res === "ok") {
    //             alert("New comment posted");
    //             $("#comments-input").val("");
    //
    //             // refreshComments(currentVideo['id']);
    //         }
    //     });
    // });
});

function _initPage() {

    let defer = $.Deferred();

    _asyncGetUserId(username).then(function (res) {
        let obj = JSON.parse(res);
        uid = obj['id'];
        console.log("👻 uid: ", uid);
        return _asyncListAllVideos(username);
    }, function (reason) {
        console.log("🔪", reason);
    }).then(function (res) {
        let obj = JSON.parse(res);
        console.log("👻 拿到videosInfo啦~", obj);
        listedVideos = obj['videos'];
        if (listedVideos !== null) {
            obj['videos'].forEach(function (item) {
                let ele = _htmlVideoListElement(item['id'], item['name'], item['display_ctime']);
                $("#items").append(ele);
            });
        }
        defer.resolve();
    }, function (e) {
        defer.reject(e);
    });

    return defer.promise();
}

_asyncBindElement = () => {

    $(".video-item").click(function() {
        let self = this.id;
        listedVideos.forEach(function(item, index) {
            if (item['id'] === self) {
                currentVideo = item;
            }
        });

        _selectVideo(self);
    });

    // $(".del-video-button").click(function() {
    //     var id = this.id.substring(4);
    //     deleteVideo(id, function(res, err) {
    //         if (err !== null) {
    //             //window.alert("encounter an error when try to delete video: " + id);
    //             popupErrorMsg("encounter an error when try to delete video: " + id);
    //             return;
    //         }
    //
    //         popupNotificationMsg("Successfully deleted video: " + id)
    //         location.reload();
    //     });
    // });

    $(".video-item").click(function () {
        let url = 'http://' + window.location.hostname + ':8080/videos/' + this.id;
        let video = $("#curr-video");
        video.attr('src', url);
        console.log(video);
        video.load();
    });
};

function _selectVideo(vid) {
    let url = 'http://' + window.location.hostname + ':8080/videos/' + vid;
    $("#curr-video:first-child").attr('src', url);
    $("#curr-video-name").text(currentVideo['name']);
    $("#curr-video-ctime").text('Uploaded at: ' + currentVideo['display_ctime']);
    // _refreshComments(vid);
}

function _refreshComments(vid) {
    _listAllComments(vid, function (res, err) {
        if (err !== null) {
            window.alert("encounter an error when loading comments");
            return
        }

        let obj = JSON.parse(res);
        $("#comments-history").empty();
        if (obj['comments'] === null) {
            $("#comments-total").text('0 Comments');
        } else {
            $("#comments-total").text(obj['comments'].length + ' Comments');
        }
        obj['comments'].forEach(function (item, index) {
            let ele = _htmlCommentListElement(item['id'], item['author'], item['content']);
            $("#comments-history").append(ele);
        });

    });
}

function _listAllComments(vid, callback) {
    let dat = {
        'url': 'http://' + window.location.hostname + ':9000/videos/' + vid + '/comments',
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