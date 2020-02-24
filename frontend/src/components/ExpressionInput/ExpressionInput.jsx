import React, { Component } from "react"
import "./ExpressionInput.scss"

class ExpressionInput extends Component {
    constructor(props) {
        super(props);
        this.state = {value: ''};
    
        this.handleChange = this.handleChange.bind(this);
        this.handleClick = this.handleClick.bind(this);
    }
    
    handleChange(event) {
        this.setState({value: event.target.value});
      }

    handleClick(event) {
        console.log("button clicked");
        //post to the server, handle error...
        //clear the value
        console.log(this.state);
        var req = {Value:this.state.value};
        fetch('http://localhost:8080/expression/', {
            method: 'POST',
            headers: {
                'Content-Type': `application/json`,
            },
            body: JSON.stringify(req),
        })
        .then((response) => response.json())
        .then((data) => {
            console.log('success:  ',data);
        })
        .catch((error) => {
            //not doing this very well
            console.error('error:  ', error);
            alert(error);
        })
    }

    render() {
        return (
            <div>
                <input type="text" value={this.state.value} onChange={this.handleChange}/>
                <input type="button" value="Evaluate" onClick={this.handleClick} />
            </div>
        );
    }
}

export default ExpressionInput