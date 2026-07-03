function toggleDropdown() {
    const tag = document.getElementById("dropdown-menu");
    if (tag) {
        tag.classList.toggle("show");
    }
}

window.onclick = function(event) {
    if (!event.target.matches('.dropdown-btn')) {
        const dropdowns = document.getElementsByClassName("dropdown-content");
        for (let i = 0; i < dropdowns.length; i++) {
            const openDropdown = dropdowns[i];
            if (openDropdown.classList.contains('show')) {
                openDropdown.classList.remove('show');
            }
        }
    }
};

function openFormFil() {
    console.log("Tentative d'ouverture du formulaire...");
    const formFil = document.getElementById("create-fil-form");
    
    if (formFil) {
        formFil.style.display = "flex"; 
    } else {
        console.error("Erreur : Impossible de trouver l'élément HTML avec l'ID 'create-fil-form'");
    }
}

function closeFormFil() {
    const formFil = document.getElementById("create-fil-form");
    if (formFil) {
        formFil.style.display = "none"; 
    }
}

window.addEventListener("click", function(event) {
    const formFil = document.getElementById("create-fil-form");
    if (event.target === formFil) {
        formFil.style.display = "none";
    }
});