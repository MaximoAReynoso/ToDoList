document.addEventListener('DOMContentLoaded', () => {
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
        const contenedor = document.createElement('div');
        contenedor.classList.add('elementoDeLista');

        const titulo = document.createElement('p');
        titulo.textContent = element.title;
        titulo.style.fontWeight = 'bold';

        const descripcion = document.createElement('p');
        descripcion.textContent = element.description;

        const completada = document.createElement('p');
        completada.textContent = element.completed == true? "Completado" : "No completado";

        const botonBorrar = document.createElement('button');
        botonBorrar.classList.add('botonBorrar');
        botonBorrar.type = 'submit';
        botonBorrar.textContent = 'Borrar';
        botonBorrar.addEventListener('click', async () => {
            await deleteTask(contenedor.id);
            await load();
        });

        contenedor.appendChild(titulo);
        contenedor.appendChild(descripcion);
        contenedor.appendChild(completada);
        contenedor.appendChild(botonBorrar);
        lista.appendChild(contenedor);
    });
}

async function postTask(event) {
    event.preventDefault();

    const title = document.getElementById('title').value;
    const description = document.getElementById('description').value;
    const button = document.getElementById('checked').value;
    const valores = {title, description, button};

    try {
        const res = await fetch('/tasks', {method: 'POST', body: JSON.stringify(valores)});
        if (!res.ok) throw new Error('No se pudo crear');

        event.target.reset();
        await load();
        console.log('Se creo el task');
    } catch (error) {
        console.log(error);
    }
}

async function deleteTask(id) {
    try {
        const res = await fetch('/tasks/${id}', {method: 'DELETE'});
        if (!res.ok) throw new Error('No se pudo borrar');
        console.log('Se elimino el task de id: ${id}');
    } catch (error) {
        console.log(error);
    }
}