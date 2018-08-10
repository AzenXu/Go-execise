function _setCookie(cname, cvalue, exmin) {
    let d = new Date();
    d.setTime(d.getTime() + (exmin * 60 * 1000));

    let expires = "expires=" + d.toUTCString();
    document.cookie = cname + "=" + cvalue + ";" + expires + ";path=/";
}

function _getCookie(cname) {
    var name = cname + "=";
    var ca = document.cookie.split(';');
    for (var i = 0; i < ca.length; i++) {
        var c = ca[i];
        while (c.charAt(0) == ' ') {
            c = c.substring(1);
        }
        if (c.indexOf(name) == 0) {
            return c.substring(name.length, c.length);
        }
    }
    return "";
}

function _popupErrorMsg(msg) {
    var x = document.getElementById("error-bar");
    $("#error-bar").text(msg);
    x.className = "show";
    setTimeout(function () {
        x.className = x.className.replace("show", "");
    }, 2000);
}

function _htmlVideoListElement(vid, name, ctime) {
    //  创建一个a标签，把所有元素都插到a标签里
    let ele = $('<a/>', {
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


    let res = $('<div/>', {
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