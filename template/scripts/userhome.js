$(document).ready(function () {
    //  「用户页」方法绑定
    $("#upload").on('click', function() {
        $("#uploadvideomodal").show();
    });

    $("#uploadform").on('submit', function(e) {
        e.preventDefault()
        var vname = $('#vname').val();

        _createVideo(vname, function(res, err) {
            if (err != null ) {
                //window.alert('encounter an error when try to create video');
                _popupErrorMsg('encounter an error when try to create video');
                return;
            }

            var obj = JSON.parse(res);
            var formData = new FormData();
            formData.append('file', $('#inputFile')[0].files[0]);

            //  上传视频
            $.ajax({
                url : 'http://' + window.location.hostname + ':8080/upload/' + obj['id'],
                //url:'http://127.0.0.1:8080/upload/dbibi',
                type : 'POST',
                data : formData,
                //headers: {'Access-Control-Allow-Origin': 'http://127.0.0.1:9000'},
                crossDomain: true,
                processData: false,  // tell jQuery not to process the data
                contentType: false,  // tell jQuery not to set contentType
                success : function(data) {
                    console.log(data);
                    $('#uploadvideomodal').hide();
                    location.reload();
                    //window.alert("hoa");
                },
                complete: function(xhr, textStatus) {
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

    $(".close").on('click', function() {
        $("#uploadvideomodal").hide();
    });

    $("#logout").on('click', function() {
        setCookie("session", "", -1)
        setCookie("username", "", -1)
    });


    $(".video-item").click(function () {
        var url = 'http://' + window.location.hostname + ':9000/videos/'+ this.id
        var video = $("#curr-video");
        video[0].attr('src', url);
        video.load();
    });
})