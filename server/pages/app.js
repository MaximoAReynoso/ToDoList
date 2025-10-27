/*document.addEventListener('DOMContentLoaded', () => {
    load();

    const formulario = document.getElementById('formulario');
    formulario.addEventListener('submit', Enviar());
});

async function load() {
    try {
        const contenido = await fetch('/tasks');
        if (!contenido.ok) throw new Error('No se pudo cargar.');

        const elementos = await contenido.json();
        cargarContenido(contenido);
    } catch (error) {
        console.log(error);
    }
}

async function cargarContenido(contenido) {
    const lista = document.getElementById("listaElementos");
    lista.innerHTML = "";

    contenido.forEach(element => {
        const el = document.createElement('div');

    });
}*/

const lista = document.getElementById("listaElementos");
for (let index = 0; index < 4; index++) {
    const contenedor = document.createElement('div');
    contenedor.classList.add('elementoDeLista');

    const titulo = document.createElement('p');
    titulo.textContent = 'Titulo';
    titulo.style.fontWeight = 'bold';

    const descripcion = document.createElement('p');
    descripcion.textContent = 'Descripcion';

    const completada = document.createElement('p');
    completada.textContent = 'Completada';

    const botonBorrar = document.createElement('button');
    botonBorrar.classList.add('botonBorrar');
    botonBorrar.type = 'submit';
    botonBorrar.textContent = 'Borrar';

    contenedor.appendChild(titulo);
    contenedor.appendChild(descripcion);
    contenedor.appendChild(completada);
    contenedor.appendChild(botonBorrar);
    lista.appendChild(contenedor);
}