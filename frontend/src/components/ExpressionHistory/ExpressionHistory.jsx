import React, { Component } from "react"
import "./ExpressionHistory.scss"

/*
this needs to be cleaned up we need to use server side and just send back an 
array just to get going we are appending it in js to see stuff
*/

class ExpressionHistory extends Component {
    render() {
        console.log("render history:  ", this.props.statementHistory);
        const history = this.props.statementHistory.map((msg, index) => <p key={index}>{msg.Body}</p>);
        return (
            <div className="ExpressionHistory">
                <h2>History</h2>
                {history}
            </div>
        );
    }
}

export default ExpressionHistory;