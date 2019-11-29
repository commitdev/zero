import React from 'react';
import TextField from '@material-ui/core/TextField';

export default function Service() {
  return (
    <div className="service">
        <TextField
          id="standard-basic"
          label="Name"
          margin="normal"
        />
        <TextField
          id="standard-multiline-flexible"
          label="Description"
          multiline
          rowsMax="4"
          value=""
          onChange={null}
          margin="normal"
        />
    </div>
  );
}
