const menuButton = document.querySelector('[data-collapse-toggle="navbar-cta"]');
const navbarMenu = document.getElementById('navbar-cta');

// Add a click event listener to the menu button
menuButton.addEventListener('click', () => {
    // Toggle the hidden class to show/hide the menu
    navbarMenu.classList.toggle('hidden');
});
