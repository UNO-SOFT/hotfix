require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");
//require("@fortawesome/fontawesome-free/js/all.js");

window.switchLanguage = function(lang) {
	document.cookie="lang=" + lang
	window.location.reload()
}

$(() => {

});
