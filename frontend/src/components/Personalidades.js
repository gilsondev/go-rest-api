import React, { Component } from 'react';
import axios from 'axios';
import '../components/Personalidades.css';

const { REACT_APP_API_URL_DOMAIN } = process.env
const ENDPOINT = `http://${REACT_APP_API_URL_DOMAIN}/api/personalidades`

console.log(ENDPOINT)

export default class Personalidades extends Component {

    state = {
        personalidades: []
    }

    componentDidMount() {
      axios.get(ENDPOINT)
          .then(res => {
              const personalidades = res.data;
              this.setState({ personalidades })
          })
    }

    render() {
        return (
            <div>
                {this.state.personalidades.map((p, id )=>
                    <div className="CardPersonalidades" key={id}>
                        <h3>{p.nome}</h3>
                        <p>{p.historia}</p>
                    </div>)}
            </div>
        );
    }
}
