import React, { Component } from 'react';
import './App.css';
import TextField from '@material-ui/core/TextField';

class App extends Component {
  
  constructor(props){
    super(props);
    this.handleTextFieldChange = this.handleTextFieldChange.bind(this);

    this.state = {textFieldValue: ''};
  }

  handleTextFieldChange(e) {
        this.setState({
            textFieldValue: e.target.value
        });
  }

  showIdentityErrMsg(){
    return (this.state.textFieldValue.length > 0 && this.state.textFieldValue.length != 11) ? "Kimlik numarası 11 karakterden oluşmalıdır." :  "" 
  }

  render(){
    return (
      <div className="app-container">
        <div className="dsc-title">
            <p>Aşağıdaki bilgileri doldurarak kredi başvurunuzu tamamlayabilirsiniz</p>  
        </div>
        <div className="app-modal">
          <div className="credit-inputs">
            <TextField
              id="outlined-number"
              label="Kimlik Numarası"
              type="number"
              InputLabelProps={{shrink: true}}
              value={this.state.textFieldValue} 
              onChange={this.handleTextFieldChange}
              variant="outlined"
              helperText={this.showIdentityErrMsg()}
              error={false}
            />
          </div>
          <div className="credit-inputs">
            <TextField 
              id="outlined-basic" 
              label="Adınız" 
              variant="outlined" 
            />
          </div>
          <div className="credit-inputs">
            <TextField 
              id="outlined-basic" 
              label="Soyadınız" 
              variant="outlined" 
            />
          </div>
        </div>
      </div>
    );
  }
}

export default App;
