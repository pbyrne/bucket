function Bucketeer(form, images) {
  this.form = form
  this.images = images
  this.bindEventListeners(this)
}

Bucketeer.prototype = {
  bindEventListeners: function(self) {
    self.form.find("input[data-action=filter]").bind("input", self.triggerFilter.bind(self))
  },

  // "asdf" to /a.*s.*d.*f/
  fuzzyRegExp: function(filterString) {
    var pieces = []
    for (var i = 0; i < filterString.length; i++) {
      pieces.push(filterString[i])
    }
    return new RegExp(pieces.join(".*"), "i")
  },

  triggerFilter: function(event) {
    var input = event.target || event.srcElement
    var fuzzyFinder = this.fuzzyRegExp($(input).val())
    for (var i = 0; i < this.images.length; i++) {
      var image = $(this.images[i])
      var imageName = image.attr('data-image-name')
      if (imageName.match(fuzzyFinder)) {
        image.removeClass("is-hidden")
      } else {
        image.addClass("is-hidden")
      }
    }
  }
}

(function () {
  "use strict";
  $("html.isnt-js").removeClass("isnt-js").addClass("is-js")
  var $bucketeer = new Bucketeer($(".m-actions form"), $(".m-images .image"))
})();
