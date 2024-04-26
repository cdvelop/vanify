# vanify
# REFACTORIZANDO !NO USAR¡

entorno de desarrollo de aplicaciones web vanilla




- [Motivación](#motivación)
- [Que hace exactamente](#que-hace-exactamente)
  - [Donde buscara vanify los archivos html, css, js del proyecto para combinarlos?](#donde-buscara-vanify-los-archivos-html-css-js-del-proyecto-para-combinarlos)
  - [Donde vanify buscara por defecto los directorios de los módulos?](#donde-vanify-buscara-por-defecto-los-directorios-de-los-módulos)



## Motivación:
este proyecto esta pensado para llevar **Go** al frontend compilado a **WebAssembly** y aprovechar su tipado estático del lado del cliente (dominio de la app), para un mejor mantenimiento (evitar **"only read code"**). usando javascript y css de sin dependencias de librerías o frameworks. 

## que hace exactamente?
levanta el navegador web chrome y escucha cambios en tus archivos, compila y recarga el navegador (hot reload) según el tipo de archivo.

tiene incluido:

browser
watch
compiler (código: html, css, js y WebAssembly)

funcionalidad parecida a Webpack solo que esta pensado específicamente para un proyecto web fullstack escrito con go.



### * donde buscara vanify los archivos html, css, js del proyecto para combinarlos?
 
para iniciar el proyecto  se toma como referencia el archivo "go.mod" como base.

```
project  
└── go.mod
```

si este no existe se inicia el modulo automáticamente y se crean los directorios y archivos web básicos. el nombre del repositorio por defecto es *github.com/currenUser/currentFolder*

vanify leerá el archivo *go.mod* del directorio actual, se extraerá el nombre  el repositorio desde el nombre del modulo (en este ejemplo es: *github.com/you*) y seleccionara todos los módulos que tengan el mismo repositorio

go.mod ej:
```mod
module github.com/you/project

go 1.20

require (
	github.com/you/htmlTheme v0.0.1
	github.com/you/module2 v0.0.1
	github.com/other/package v0.0.0
```
los módulos seleccionados en este ejemplo serán: *"**htmlTheme**","module2"*

los módulos pueden tener cualquier nombre.

la lectura de directorios para compilar es en el mismo orden que tienen los paquetes en *go.mod* por ende el paquete que contiene el tema principal de tu proyecto (css,JavaScript y pagina de inicio index.html) debe estar al principio como se muestra en el ejemplo **"htmlTheme"**

## * donde vanify buscara por defecto los directorios de los módulos?

```
 Home  
 └── User  
     └── go      
         └── src
```
para cambiar el directorio de módulos edita el archivo **vanify.yml** solo se cambia a partir de *go* por *"TuDirectorio"*. home y user siempre se mantendrán ej:
```
 Home  
 └── User  
     └── TuDirectorio
```


## en que ruta inicia el navegador?

si el puerto configurado contiene el numero "44" el protocolo es **https**  
ej:  
port "4433" = **https**://localhost:4433/..... (default: **http** port: **8080**)

## JavaScript
el orden de compilación es primero la raíz solo los archivos que comienzan con mayúsculas.
luego la carpeta **js** y al final los archivos jsTest (en las carpetas no se distinguen mayúsculas y minúsculas) se carga todo por orden alfabético.
```
Module  
└── js  
     1xFun.js #2
     func.js #3
     Help.js #4
     main.js #5

└── jsTest
     test.js #6

Load.js #1
noLoad.js
```
## Css
se compilara de la mima forma que el javascript solo cambia el nombre de la carpeta a css
```
Module  
└── css  
     1xStyle.css #2
     Help.css #3
     main.css #4
     other.css #5

Load.css #1
noLoad.css
```


## Este proyecto no hubiese sido posible sin:

### github.com/fsnotify
### github.com/chromedp
### github.com/tdewolff/minify