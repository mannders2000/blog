function setIcon(isDark) {
    const themeIcon = document.getElementById('sunMoonIcon');
    themeIcon.textContent = isDark ? '☾' : '☼';
}

function toggleTheme() {
    const bodyEl = document.body;
    const currentTheme = bodyEl.getAttribute('data-bs-theme');

    if (currentTheme === 'dark') {
        bodyEl.setAttribute('data-bs-theme', 'light');
        localStorage.setItem('theme', 'light');
        setIcon(false);
    } else {
        bodyEl.setAttribute('data-bs-theme', 'dark');
        localStorage.setItem('theme', 'dark');
        setIcon(true);
    }
}

function loadTheme() {
    const storedTheme = localStorage.getItem('theme');
    const darkModeSwitch = document.getElementById('darkModeSwitch');

    if (storedTheme) {
        document.body.setAttribute('data-bs-theme', storedTheme);
        darkModeSwitch.checked = storedTheme === 'dark';
        setIcon(storedTheme === 'dark');
    } else {
        setIcon(false);
    }
}

document.getElementById('darkModeSwitch').addEventListener('click', toggleTheme);

// Call loadTheme when the page loads
loadTheme();