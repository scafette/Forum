document.addEventListener("DOMContentLoaded", () => {
    const filterSelectDessert = document.getElementById('filter-select-dessert')
    const filterSelectPlat = document.getElementById ('filter-select-plat')
    const filterSelectEntrer = document.getElementById ('filter-select-entrer')

    if (filterSelectDessert) {
    filterSelectDessert.addEventListener('change', () => {
        window.location.href = '/dessert?'+filterSelectDessert.value;
    });}

    if (filterSelectPlat){
    filterSelectPlat.addEventListener('change', () => {
        window.location.href = '/plat?'+filterSelectPlat.value;
    });}

    if(filterSelectEntrer){
    filterSelectEntrer.addEventListener('change', () => {
        window.location.href = '/entrer?'+filterSelectEntrer.value;
    });}
});