
// Simple Example: Toggle navigation for smaller screens
const navToggle = document.getElementById('nav-toggle'); // Assuming you have an element to open/close nav
const navMenu = document.getElementById('nav-menu');   // Assuming your main navigation has an ID

navToggle.addEventListener('click', function() {
  navMenu.classList.toggle('show-nav');
});
