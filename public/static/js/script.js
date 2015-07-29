
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

function toggle_comments(id) {
	var blk = document.getElementById(id);
	
	if (blk === null) {
		return;
	}

	if (blk.style.display != 'none') {
		blk.setAttribute('data-display', blk.style.display);
		blk.style.display = 'none';
	} else {
		var display = blk.getAttribute('data-display');

		if (display === null || display.length < 1) {
			display = 'block';
		}
		blk.style.display = display;
	}

	var btn = document.getElementById('entry-comments-tgicon');
	if (btn !== null) {
		btn.classList.toggle('fa-rotate-180');
	}
}

document.onclick = function() {
	if (isClearTabs) {
		clearTabs();
	}
	isClearTabs = true;
};