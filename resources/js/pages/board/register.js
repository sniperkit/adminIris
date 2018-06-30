$(() => {
    let Ctx = {};

    initBoardRegister(Ctx);
    eventBoardRegister(Ctx);
});

function initBoardRegister(Ctx) {

}

function eventBoardRegister(Ctx) {

    $("#btnRegist").click(function(e) {
        e.preventDefault();
        clickBtnRegist(Ctx);
    });
}

function clickBtnRegist(Ctx) {
    let formData = {
        title: $('form > input[name=title]').val(),
        content: $('form > input[name=content]').val(),
        writer: $('form > input[name=writer]').val()
    };

    $.ajax({
        url: '/adm/board/register',
        type: 'post',
        headers: {
            'Content-Type': 'application/json',
            'X-Requested-With': 'XMLHttpRequest'
        },
        data: JSON.stringify(formData),
        success: function(result) {
            alert("등록되었습니다.");
        },
        error: function(err) {
            alert("ERROR : ", err);
        }
    });
}