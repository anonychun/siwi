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

$("#uploadfile").on("submit", function (e) {
  var formData = new FormData(this);

  Swal.fire({
    title: "<i>Proses Upload</i>",
    html: `
      <div class="progress">
        <div class="bar" style="width:0%">
          <p class="percent">0%</p>
        </div>
      </div>`,
    footer: "<a href>Batal</a>",
    showConfirmButton: false,
  });

  $.ajax({
    type: "POST",
    url: "upload",
    data: formData,
    cache: false,
    contentType: false,
    processData: false,

    xhr: function () {
      var myXhr = $.ajaxSettings.xhr();
      if (myXhr.upload) {
        myXhr.upload.addEventListener("progress", progress, false);
      }
      return myXhr;
    },

    success: function (data) {
      Swal.fire({
        title: "Success",
        text: "Upload file success!",
        type: "success",
        showConfirmButton: false,
      });
    },
  });

  e.preventDefault();
});

function progress(e) {
  if (e.lengthComputable) {
    var max = e.total;
    var current = e.loaded;
    var percentage = (current * 100) / max;
    $(".bar").attr("style", `width:${percentage}%`);
    $(".percent").html(`${percentage.toFixed(2)}%`);
  }
}
