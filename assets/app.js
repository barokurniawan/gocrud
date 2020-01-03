var ListMessage,
    formGuestbook;

$(document).ready(function () {
    ListMessage = $("#ListMessage");
    formGuestbook = $("#formGuestbook");


    formGuestbook.submit(function () {
        var _this = $(this);
        $.post("api/create-message", _this.serialize(),
            function (response) {
                if (response == false) {
                    alert("Failed");
                }
            }).always(function () {
                loadMessage(ListMessage, true);
            });

        return false;
    });

    loadMessage(ListMessage, false);
});

function deleteHandler() {
    $("a#btnDelete").unbind("click")
        .click(function () {
            var _this = $(this);
            $.get(_this.attr("href")).always(function () {
                loadMessage(ListMessage, true);
            });
            return false;
        });
}

function loadMessage(ListMessage, reload) {
    $.get("api/list-message", function (response) {
        if (reload) {
            ListMessage.empty();
        }

        $.each(response, function (i, item) {
            ListMessage.append(function () {
                return $('<li>').addClass("list-group-item")
                    .append(function () {
                        return $('<b>').text(item.name);
                    })
                    .append(function () {
                        return $('<p>').text(item.message);
                    })
                    .append(function () {
                        return $('<a>').attr("id", "btnDelete")
                            .attr("href", "api/delete-message?id=" + item.id)
                            .text("Delete");
                    });
            });
        });
    }).always(function () {
        deleteHandler();
    }).fail(function (err) {
        alert(err.statusText)
    });
}