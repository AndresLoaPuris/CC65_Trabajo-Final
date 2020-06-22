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

    var sendjson=[];


  for (var i = 0; i < data.length; i++) { 
    sendjson.push({
      Name:      data[i].data[0],
      Values:    [parseFloat(data[i].data[1]), parseFloat(data[i].data[2]), parseFloat(data[i].data[3]), parseFloat(data[i].data[4]), parseFloat(data[i].data[5])],
      Category:  parseFloat(this.state.value),
      Algorithm: 'kmeans'
    });
  }


      fetch("http://localhost:8000/people", {
        method: 'post',
        mode:'no-cors',
        body: JSON.stringify(sendjson)
      });

  }

  handleOnError = (err, file, inputElem, reason) => {
    console.log(err)
  }


  render() {
    return (
      <div className="card shadow mb-4">
        <div class="card-header py-3">
          <h5 class="m-0 font-weight-bold text-primary">Subir Conjunto de Datos</h5>
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


                <button style={{ marginLeft: 10 }} type='button' className="btn btn-primary" onClick={this.handleOpenDialog} >
                  Adjuntar Archivo
                </button>
              
                
                <div style={{ marginLeft: 10, width: '50%' }} >
                  {file && file.name}
                </div>
                
                

             
              
            </aside>
          )}
        </CSVReader>
  
        </div>
        


      </div>

    )
  }
}