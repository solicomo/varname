
var isClearTabs = false;
function clearTabs() {
	var menutabs = document.getElementsByClassName("menu-tab");
	for (var i = 0; i < menutabs.length; ++i) {
		menutabs[i].style.display = "none";
	}
}
function showTabFor(menuBtn) {
	clearTabs();
	menuBtn.parentElement.getElementsByClassName("menu-tab")[0].style.display = "block";
	isClearTabs = false;
	return false;
}
document.onclick = function() {
	if (isClearTabs) {
		clearTabs();
	}
	isClearTabs = true;
};