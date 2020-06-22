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
                      <td style={{ color: 'black' }} ><strong>Nombes y Apellidos</strong></td>
                      <td style={{ color: 'black' }} ><strong>Etapa de Vida</strong></td>
                      <td style={{ color: 'black' }} ><strong>Departamento</strong></td>
                      <td style={{ color: 'black' }} ><strong>Sintamologia</strong></td>
                      <td style={{ color: 'black' }} ><strong>Comorbilidad</strong></td>
                  </tr>
      <tbody>{this.state.contacts.map(function(item) {
             
               return (
                  <tr>
                    <td style={{ color: 'black' }} >{item.Category}</td>
                      <td style={{ color: 'black' }} >{item.Name}</td>
                      <td style={{ color: 'black' }} >{getLifeStage(item.Values[1])}</td>
                      <td style={{ color: 'black' }} >{getDepartment(item.Values[2])}</td>
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