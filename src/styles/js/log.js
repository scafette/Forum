// Autor: Gio , Elias
function changePassword() {
    var password = document.getElementById('password');
    var password2 = document.getElementById('password2');
    if (password.type === 'password') {
        password.type = 'text';
        password2.type = 'text';
    } else {
        password.type = 'password';
        password2.type = 'password';
    }
}
function changeUsername () {
    var username = document.getElementById('username');
    if (username.type === 'text') {
        username.type = 'password';
    } else {
        username.type = 'text';
    }
}
function enableAll() {
    var password = document.getElementById('password');
    var password2 = document.getElementById('password2');
    var username = document.getElementById('username');
    password.disabled = false;
    password2.disabled = false;
    username.disabled = false;
}

document.addEventListener("DOMContentLoaded", () => {
    const foodSelect = document.getElementById('food-select');
    foodSelect.addEventListener('change', () => {
        if (foodSelect.value === 'Create') {
            window.location.href = '/create-categorie';
        }
    });
});