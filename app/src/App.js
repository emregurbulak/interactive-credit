import React, { useState} from 'react';
import './App.css';
import TextField from '@material-ui/core/TextField';
import MaskedInput from "react-text-mask";
import PropTypes from "prop-types";
import Button from '@material-ui/core/Button';
import SaveIcon from '@material-ui/icons/Save';

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

export function App(props) {

  const [numberIdentityValue, setNumberIdentityValue] = useState('')
  const [customerFirstName, setCustomerFirstName] = useState('')
  const [customerLastName, setCustomerLastName] = useState('')
  const [mountlySalary, setMountlySalary] = useState('')
  const [phoneNumber, setPhoneNumber] = useState('')

  function isApproveButtonDisable(){
    if(numberIdentityValue === '' || 
       customerFirstName === '' ||
       customerLastName === '' ||
       mountlySalary === '' || 
       phoneNumber === '' 
      ){
      return true
    }else{
      return false
    }
  }

  function showIdentityErrMsg(){
    return (numberIdentityValue.length > 0 && numberIdentityValue.length !== 11) ? "Kimlik numarası 11 karakterden oluşmalıdır." :  "" 
  }

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
              value={numberIdentityValue} 
              onChange={(e) => setNumberIdentityValue(e.target.value)}
              variant="outlined"
              helperText={showIdentityErrMsg()}
              error={false}
              inputComponent={TextMaskCustom}
            />
          </div>
          <div className="credit-inputs">
            <TextField 
              id="outlined-basic"
              name="customerFirstName" 
              label="Adınız" 
              value={customerFirstName} 
              onChange={(e) => setCustomerFirstName(e.target.value)}
              variant="outlined" 
            />
          </div>
          <div className="credit-inputs">
            <TextField 
              id="outlined-basic" 
              label="Soyadınız" 
              name="customerLastName" 
              value={customerLastName} 
              onChange={(e) => setCustomerLastName(e.target.value)}
              variant="outlined" 
            />
          </div>
          <div className="credit-inputs">
            <TextField 
              id="outlined-basic" 
              label="Aylık Geliriniz"
              name="mountlySalary" 
              value={mountlySalary} 
              onChange={(e) => setMountlySalary(e.target.value)}
              variant="outlined" 
            />
          </div>
          <div className="credit-inputs">
            <TextField 
              id="outlined-basic" 
              label="Telefon Bilginiz" 
              name="phoneNumber" 
              value={phoneNumber} 
              onChange={(e) => setPhoneNumber(e.target.value)}
              variant="outlined" 
            />
          </div>
          <div className="send-button">
            <Button
              variant="contained"
              color="primary"
              startIcon={<SaveIcon />}
              disabled={isApproveButtonDisable()}
            >
              Başvur
            </Button>
          </div>
        </div>
      </div>
    );
}

