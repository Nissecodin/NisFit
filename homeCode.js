function revealSidebar() {
    var element = document.getElementById("sidebar");
    if (element.style.display === "none") {
        element.style.display = "block";
    } else {
        element.style.display = "none";
    }
}