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
