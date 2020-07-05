let zipAPI = "/PDF2ZIP",
    jpgAPI = "/PDF2JPG",
    api = zipAPI;

$(document).ready(function(){

    $("#jpg").click(function(){
        $(this).removeClass("btn-secondary").addClass("btn-primary");
        $("#zip").removeClass("btn-primary").addClass("btn-secondary");
        api = jpgAPI;
    })

    $("#zip").click(function(){
        $(this).removeClass("btn-secondary").addClass("btn-primary");
        $("#jpg").removeClass("btn-primary").addClass("btn-secondary");
        api = zipAPI;
    })

    let $inputFile = $("#file");

    $("#submit").click(function(){
        let files = $inputFile.prop("files");
        if (!files.length) {
            $inputFile.addClass("is-invalid")
        } else {
            $inputFile.removeClass("is-invalid")

            let fileData = new FormData();
            fileData.append("file", files[0]);
            let req = {
                url: new URL(api, window.location.href),
                type: "POST",
                data: fileData,
                success: function(result) {
                    let dataUrl = new URL("/data/" + result.filename, window.location.href);
                    $("#url").val(decodeURI(dataUrl));
                },
                cache: false,
                processData: false,
                contentType: false,
            }
            $.ajax(req);
        }
    })

    $(document).ajaxSend(function(){
        $("#submit").hide();
        $(".spinner").show();
    })
    $(document).ajaxComplete(function(){
        $(".spinner").hide();
        $("#submit").show();
    })
})