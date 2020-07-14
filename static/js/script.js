var $fileInput = $(".file-input");
var $droparea = $(".file-drop-area");

// highlight drag area
$fileInput.on("dragenter focus click", function () {
  $droparea.addClass("is-active");
});

// back to normal state
$fileInput.on("dragleave blur drop", function () {
  $droparea.removeClass("is-active");
});

// change inner text
$fileInput.on("change", function () {
  var filesCount = $(this)[0].files.length;
  var $textContainer = $(this).prev();

  if (filesCount === 1) {
    // if single file is selected, show file name
    var fileName = $(this).val().split("\\").pop();
    $textContainer.text(fileName);
  } else {
    // otherwise show number of files
    $textContainer.text(filesCount + " files selected");
  }
});

// submit click
$("#file_upload").click(function () {
  var formData = new FormData();
  var files = $("input[type=file]")[0].files;
  $.each(files, function(key, file){
    formData.append("upload_file", file);
  });

  Swal.fire({
    title: '<i>Proses Upload</i>',
    html:
    `<div class="progress">
        <div class="bar" style="width:0%">
          <p class="percent">0%</p>
        </div>
      </div>`,
    footer: '<a href>Batal</a>',
    showConfirmButton: false,
  })

  $.ajax({
    xhr: function () {
      var xhr = new window.XMLHttpRequest();
      xhr.upload.addEventListener(
        "progress",
        function (evt) {
          if (evt.lengthComputable) {
            var percentComplete = evt.loaded / evt.total;
            percentComplete = parseInt(percentComplete * 100);
            $(".bar").attr('style', `width:${percentComplete}%`); 
            $(".percent").html(`${percentComplete}%`); 
            if (percentComplete === 100) {
            }
          }
        },
        false
      );
      return xhr;
    },
    url: "upload",
    type: "POST",
    data: formData,
    contentType: false,
    processData: false,
    success: function (result) {
      Swal.fire({
        title: 'Success',
        text: "Upload file success!",
        type: 'success',
        showConfirmButton: false,
      })
      
    },
  });
});
