
/*
import React, { Component } from 'react'
import { CSVReader } from 'react-papaparse'

const buttonRef = React.createRef()

export default class CSVReader1 extends Component {


  constructor(props) {
    super(props);
    this.state = {value: ''};

    this.handleChange = this.handleChange.bind(this);
  }

  handleChange(event) {
    this.setState({value: event.target.value});
    console.log(event.target.value)
  }

  handleOpenDialog = (e) => {
    if (buttonRef.current) {
      buttonRef.current.open(e)
    }
  }

   handleOnFileLoad = (data) => {

    for (var i = 0; i < data.length; i++) {

      fetch("http://localhost:8000/people", {
        method: 'post',
        mode:'no-cors',
        body: JSON.stringify({
          Name:      data[i].data[0],
          Values:    [parseFloat(data[i].data[1]), parseFloat(data[i].data[2]), parseFloat(data[i].data[3]), parseFloat(data[i].data[4]), parseFloat(data[i].data[5])],
          //Category:  parseFloat(this.state.value),
          Category: Math.floor((Math.random() * parseFloat(this.state.value)) + 1),
          Algorithm: 'kmeans'
        })
      
      });

    }

  }

  handleOnError = (err, file, inputElem, reason) => {
    console.log(err)
  }


  render() {
    return (
      <div className="card shadow mb-4">
        <div class="card-header py-3">
          <h6 class="m-0 font-weight-bold text-primary">Subir Conjunto de Datos</h6>
        </div>

        <div class="card-body">
        <div class="col-md-3 mb-3">
        <input className = "form-control" type="number" value={this.state.value} onChange={this.handleChange} min="0" max="100" placeholder="N° de Grupos"/>
        
        </div>
        
        
        <CSVReader
          ref={buttonRef}
          onFileLoad={this.handleOnFileLoad}
          onError={this.handleOnError}
          noClick
          noDrag
        >
          {({ file }) => (
            <aside
              style={{
                display: 'flex',
                flexDirection: 'row',
                marginBottom: 30
              }}
            >


                <button style={{ marginLeft: 10 }} type='button' className="btn btn-secondary" onClick={this.handleOpenDialog} >
                  Adjuntar Archivo
                </button>
              
                
                <div style={{ marginLeft: 10, width: '50%' }} >
                  {file && file.name}
                </div>
                
                

             
              
            </aside>
          )}
        </CSVReader>
        <div class="col-md-3 mb-3"></div>
        <div class="col-md-3 mb-3">
        <button type='button' className="btn btn-primary">
               Enviar Datos
        </button>
        </div>
        </div>
        


      </div>

    )
  }
}

*/



import React, { Component } from 'react';



class App extends Component {

  state = {
    contacts: []
  }

  async componentDidMount() {
    await fetch('http://localhost:8000/people')
    .then(res => res.json())
    .then((data) => {
      this.setState({ contacts: data })
    })
    .catch(console.log)
  }
  


  render() {

    function getDepartment (number) {

      switch (number) {
        case 0.065: return "Ica";
        case 0.062: return "Lambayeque";
        case 0.058: return "Ancash";
        case 0.055: return "Tumbes";
        case 0.049: return "Piura";
        case 0.044: return "Loreto";
        case 0.043: return "La Libertad";
        case 0.029: return "Callao";
        case 0.027: return "Amazonas";
        case 0.022: return "Apurímac";
        case 0.021: return "Lima";
        case 0.020: return "Ucayali";
        case 0.019: return "Pasco";
        case 0.018: return "Arequipa";
        case 0.017: return "Madre de Dios";
        case 0.014: return "Junín";
        case 0.012: return "Cajamarca";
        case 0.011: return "Ayacucho";
        case 0.009: return "Tacna";
        case 0.008: return "Huancavelica";
        case 0.007: return "Moquegua";
        case 0.004: return "Cusco";
      }
      
    }

    function getLifeStage (number) {

      switch (number) {
        case 0.6042: return "Adulto";
        case 0.1867: return "Joven";
        case 0.1675: return "Adulto Mayor";
        case 0.0256: return "Niño";
        case 0.0161: return "Adolescente";
      }
      
    }

    function getComorbidity (number) {

      switch (number) {
        case 0.3010: return "Presion Arterial Alta";
        case 0.2768: return "Diabetes";
        case 0.1488: return "Obesidad";
        case 0.0969: return "Enfermedad Pulmonar";
        case 0.0519: return "Cardiopatia";
        case 0.0381: return "Enfermedad Neurologica";
        case 0.0346: return "Cancer";
      }
      
    }

    function getSymptomatology (number) {

      switch (number) {
        case 0.1806: return "Tos";
        case 0.1458: return "Fiebre";
        case 0.1431: return "Malestar General";
        case 0.1426: return "Dolor Garganta";
        case 0.0809: return "Cefalea";
        case 0.0756: return "Congestion Nasal";
        case 0.0747: return "Dificultad Respiratoria";
        case 0.0541: return "Dolor Muscular";
        case 0.032: return "Diarrea";
        case 0.0228: return "Nauseas";
        case 0.0201: return "Dolor de Pecho";
        case 0.0129: return "Dolor de Articulaciones";
        case 0.0099: return "Dolor Abdominal";
        case 0.005: return "Confusion";
      }
      
    }

  {this.state.contacts.sort(function (a, b) { return (a.Category - b.Category) })}

    return (

      <div className="card shadow mb-4">
            <div class="card-header py-3">
        <h5 class="m-0 font-weight-bold text-primary">Detalle Grupo de Riesgo</h5>
    </div>
 <div className="card-body">
            <div className="table-responsive">

            <table className="table table-bordered">
                  <tr>
                      <td style={{ color: 'black' }} ><strong>Grupo N°</strong></td>
                      <td style={{ color: 'black' }} ><strong>Name and LastName</strong></td>
                      <td style={{ color: 'black' }} ><strong>Department</strong></td>
                      <td style={{ color: 'black' }} ><strong>LifeStage</strong></td>
                      <td style={{ color: 'black' }} ><strong>Symptomatology</strong></td>
                      <td style={{ color: 'black' }} ><strong>Comorbidity</strong></td>
                  </tr>
      <tbody>{this.state.contacts.map(function(item) {
             
               return (
                  <tr>
                    <td style={{ color: 'black' }} >{item.Category}</td>
                      <td style={{ color: 'black' }} >{item.Name}</td>
                      <td style={{ color: 'black' }} >{getDepartment(item.Values[2])}</td>
                      <td style={{ color: 'black' }} >{getLifeStage(item.Values[1])}</td>
                      <td style={{ color: 'black' }} >{getSymptomatology(item.Values[3])}</td>
                      <td style={{ color: 'black' }} >{getComorbidity(item.Values[4])}</td>
                  </tr>
                )
             
             })}</tbody>
       </table>
            </div>
           </div>
      </div>
          
      
    )
  }
}

export default App;

