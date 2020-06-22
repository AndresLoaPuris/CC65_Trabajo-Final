
import React, { Component } from 'react'
export default class Home extends Component{

render(){
    return(
        <>
          <div class="jumbotron">
        <h1 style={{ color: 'black' }}>COVID-19</h1>
        <p style={{ color: 'black' }} class="lead">Trabajo Final con respecto al COVID-19 ( Machine Learning Go , Blockchain , React.js )</p>
        <p><a id="Home_btn_youtube" href="https://covid19.orcebot.com/" class="btn btn-primary btn-lg">Fuentes</a></p>
    </div>
    
    <div class="row">
        <div class="col-md-4">
            <h4 style={{ color: 'black' }}><strong>Machine Learning</strong></h4>
            <p style={{ color: 'black' }}>
            El aprendizaje automático ha desarrollado técnicas con la computación como base que permiten a los sistemas “aprender”. Es decir, identificar millones de patrones a partir de los datos obtenidos para incorporar un nuevo conocimiento que le haga reaccionar en diferentes situaciones futuras. 
            </p>
            
        </div>
        <div class="col-md-4">
            <h4 style={{ color: 'black' }}><strong>Blockchain</strong></h4>
            <p style={{ color: 'black' }}>
            La tecnología de blockchain hace uso de la firma digital del usuario a fin de llevar a cabo transacciones libres de fraude. De esta forma se hace imposible corromper o cambiar información de cualquier usuario.
            </p>
        </div>
        <div class="col-md-4">
            <h4 style={{ color: 'black' }}><strong>Golang</strong></h4>
            <p style={{ color: 'black' }}>
            La concurrencia es uno de los puntos más fuertes de Go. Todos los procesos escriben sobre el mismo espacio de memoria, es decir, comparten variables, pero debe haber algúna forma de coordinación entre ellos para que no choquen ni se bloqueen mutuamente al tomar un recurso.
            </p>
        </div>
    </div>
        </>
        );
}

}