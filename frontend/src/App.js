import React, { Component } from 'react';
import './App.css';
import { connect } from "./api";
import Header from './components/Header/Header';
import ExpressionInput from './components/ExpressionInput';
import ExpressionHistory from './components/ExpressionHistory'

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      statementHistory: []
    }
  }

  componentDidMount() {
    connect((msg) => {
      console.log("new message");
      this.setState(prevState => ({
        /*
        we need to parse here and just send the list, just wanted to
        start seeing data here which is why it is done this way.
        */
        statementHistory: [...prevState.statementHistory, msg.data]
      }))
      console.log("message", this.state);
    });
  }

  render() {
    return (
      <div className="App">
        <Header />
        <ExpressionInput />
        <ExpressionHistory statementHistory={this.state.statementHistory}/>
      </div>
    )
  }
}

export default App;
