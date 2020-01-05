import React, { Component } from 'react';
import './App.css';
import TextField from '@material-ui/core/TextField';
import MaskedInput from "react-text-mask";
import PropTypes from "prop-types";

const regEx = /[0-9]/;
function TextMaskCustom(props) {
  const { inputRef, ...other } = props;

  return (
    <MaskedInput
      {...other}
      ref={ref => {
        inputRef(ref ? ref.inputElement : null);
      }}
      mask={[...regEx]}
    />
  );
}

TextMaskCustom.propTypes = {
  inputRef: PropTypes.func.isRequired
};

class App extends Component {
  
  constructor(props){
    super(props);
    this.handleIdentityFieldChange = this.handleIdentityFieldChange.bind(this);

    this.state = {numberIdentityValue: ''};
  }

  handleIdentityFieldChange(e) {
        this.setState({
            numberIdentityValue: e.target.value
        });
  }

  showIdentityErrMsg(){
    return (this.state.numberIdentityValue.length > 0 && this.state.numberIdentityValue.length != 11) ? "Kimlik numarası 11 karakterden oluşmalıdır." :  "" 
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
              value={this.state.numberIdentityValue} 
              onChange={this.handleIdentityFieldChange}
              variant="outlined"
              helperText={this.showIdentityErrMsg()}
              error={false}
              inputComponent={TextMaskCustom}
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
