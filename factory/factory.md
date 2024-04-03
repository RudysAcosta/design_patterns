# Patrón Factory (Factoría)

El patrón Factory es un patrón de diseño creacional que se utiliza cuando necesitamos crear objetos sin especificar la clase exacta del objeto que se creará. En lugar de instanciar objetos directamente utilizando el constructor de su clase, el patrón Factory proporciona un método para crear objetos basado en un conjunto de condiciones.

## Estructura del patrón Factory:

- **Creator (Creador):** Define un método de fábrica que devuelve un objeto de tipo Product. Puede ser una clase abstracta o una interfaz.

- **Concrete Creator (Creador Concreto):** Implementa el método de fábrica definido en el Creador y devuelve instancias concretas de Product.

- **Product (Producto):** Define la interfaz de los objetos que la fábrica crea.

- **Concrete Product (Producto Concreto):** Implementa la interfaz Product.

## Ejemplo del patrón Factory:

Supongamos que estamos desarrollando una aplicación de gestión de contenido y necesitamos crear diferentes tipos de publicaciones, como artículos y videos. Podemos aplicar el patrón Factory para manejar la creación de estas publicaciones de manera flexible.
