import React, { Component } from "react";


function SelectOptions(props){
    const items = props.items;
    const optionList = items.map(x => 
        <option key={x.value.toString() + x.name} value={x.value}>{x.name}</option>
    );
    return optionList;
};

class PostForm extends Component {
	constructor(props) {
        super(props);
		this.state = {
			Name: '',
			Department: 0.0000000,
            LifeStage: 0.0000000,
            Gender:0.0000000,
            Comorbidity:0.0000000,
            Symptomatology:0.0000000
        };
        this.handleChange = this.handleChange.bind(this);
        this.submitHandler = this.submitHandler.bind(this);
	}

	handleChange = event => {
        console.log("Name: ", event.target.name);
        console.log("Value: ", event.target.value);
        this.setState({ 
            [event.target.name]: event.target.value 
        });
	}

	submitHandler = e => {
		e.preventDefault()
        console.log(this.state);
        const payload = {
            name: this.state.Name,
            department: this.state.Department,
            lifestage: this.state.LifeStage,
            gender: this.state.Gender,
            comorbidity: this.state.Comorbidity,
            symptomatology: this.state.Symptomatology
        };
		fetch('http://localhost:3000/knn', {
            method: 'post',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({payload})
        })
        .then(res=>res.json())
        .then(res => {
            //Aca devuelve las respuesta / res: bool
            console.log(res); // Imprime true o false
        });
	}

	render() {
        const departments = [
            {
                value: 0.0620,
                name: "Ica"
            },
            {
                value: 0.0620,
                name: "Lambayeque"
            },
            {
                value: 0.0580,
                name: "Ancash"
            },
            {
                value: 0.0550,
                name: "Tumbes"
            },
            {
                value: 0.0490,
                name: "Piura"
            },
            {
                value: 0.440,
                name: "Loreto"
            },
            {
                value: 0.0430,
                name: "La Libertad"
            },
            {
                value: 0.0270,
                name: "Amazonas"
            },
            {
                value: 0.0220,
                name: "Apurimac"
            },
            {
                value: 0.0210,
                name: "Lima"
            },
            {
                value: 0.0200,
                name: "Ucayali"
            },
            {
                value: 0.0180,
                name: "Arequipa"
            },
            {
                value: 0.0180,
                name: "San Martin"
            },
            {
                value: 0.0170,
                name: "Huanuco"
            },
            {
                value: 0.0170,
                name: "Madre de Dios"
            },
            {
                value: 0.0170,
                name: "Puno"
            },
            {
                value: 0.140,
                name: "Huanuco"
            },
            {
                value: 0.170,
                name: "Junin"
            },
            {
                value: 0.120,
                name: "Cajamarca"
            },
            {
                value: 0.110,
                name: "Ayacucho"
            },
            {
                value: 0.0090,
                name: "Tacna"
            },
            {
                value: 0.0080,
                name: "Huancavelica"
            },
            {
                value: 0.0070,
                name: "Moquegua"
            },
            {
                value: 0.0040,
                name: "Cusco"
            }
        ];
        const lifeStages = [
            {
                value:0.6042,
                name: "Adulto"
            },
            {
                value:0.1867,
                name:"Joven"
            },
            {
                value:0.1675,
                name: "Adulto Mayor"
    
            },
            {
                value:0.0256,
                name:"Niño"
                
            },
            {
                value:0.0161,
                name:"Adolescente"
            }
        ];
        const genders = [
            {
                value: 0.6020,
                name:"Masculino"
            },
            {
                value: 0.3980,
                name:"Femenino"
            }
        ];
        const comorbidities = [
            {
                value: 0.3010,
                name:"Presion Arterial Alta"
    
            },
            {
                value: 0.2768,
                name:"Diabetes"
            },
            {
                value: 0.1488,
                name:"Obesidad"
            },
            {
                value: 0.0969,
                name:"Enfermedad Pulmonar"
            },
            {
                value: 0.0519,
                name:"Cardiopatia"     
            },
            {
                value: 0.0519,
                name:"Enfermedad Cronica Renal"
            },
            {
                value: 0.0381,
                name:"Enfermedad Neurologica"
            },
            {
                value: 0.0346,
                name:"Cancer"
            }
        ];
        const symptomatologies = [
            {
                value: 0.3010,
                name: "Tos"
            },
            {
                value: 0.2768,
                name: "Fiebre"
            },
            {
                value: 0.1488,
                name: "Malestar General"
            },
            {
                value: 0.1476,
                name: "Dolor de Garganta"
            },
            {
                value: 0.0809,
                name: "Cefalea"
            },
            {
                value: 0.0756,
                name: "Congestion Nasal"
            },
            {
                value: 0.0747,
                name: "Dificultad Respiratoria"
            },
            {
                value: 0.0747,
                name: "Dolor Muscular"
            },
            {
                value: 0.0320,
                name:"Diarrea"
            },
            {
                value: 0.0228,
                name:"Nauseas"
            },
            {
                value: 0.0129,
                name:"Dolor de Articulaciones"
            },
            {
                value: 0.0201,
                name:"Dolor Abdominal"
            },
            {
                value: 0.0050,
                name:"Confusion"
            }
        ];
        const styles = {
            marginLeft: '50px'
        };
		return (
			<div style={styles}>
				<form onSubmit={this.submitHandler}>
                    <label>Nombre:</label>
					<input
							type="text"
							name="Name"
							value={this.state.Name}
							onChange={this.handleChange}
						/>
                    <br/>
                    <label>
                        Departamento:
                        <select name="Department" value={this.state.Department} onChange={this.handleChange}>
                            <SelectOptions items={departments}/>
                        </select>
                    </label>
                    <br/>
                    <label>
                        Etapa de Vida:
                        <select name="LifeStage" value={this.state.LifeStage} onChange={this.handleChange}>
                           <SelectOptions items ={lifeStages}/>
                        </select>
                    </label>
                    <br/>
                    <label>
                        Genero:
                        <select name="Gender" value={this.state.Gender} onChange={this.handleChange}>
                            <SelectOptions items={genders}/>
                        </select>
                    </label>
                    <br/>
                    <label>
                        Comorbilidad:
                        <select name="Comorbidity" value={this.state.Comorbidity} onChange={this.handleChange}>
                           <SelectOptions items={comorbidities}/>
                        </select>
                    </label>
                    <br/>
                    <label>
                        Sintomatologia:
                        <select name="Symptomatology" value={this.state.Symptomatology} onChange={this.handleChange}>
                            <SelectOptions items={symptomatologies}/>
                        </select>
                    </label>
                    <br/>
					<button type="submit" onClick={this.submitHandler}>Enviar</button>
				</form>
			</div>
		)
	}
}


export default PostForm