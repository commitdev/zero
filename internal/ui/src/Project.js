import React from 'react';
import './Frontend.css';
import Grid from '@material-ui/core/Grid';
import TextField from '@material-ui/core/TextField';


export default class Project extends React.Component {
  constructor(props) {
    super(props);
  }
  updateName = (event) => {
    this.props.setProjectName(event.target.value);
  }
  updateDesc = (event) => {
    this.props.setProjectDescription(event.target.value);
  }
  render() {
    return (
      <div className="project">
        <Grid container spacing={2} direction="column" alignItems="center">
          <Grid item>
            <TextField
            id="standard-basic"
            label="Name"
            margin="normal"
            onChange={this.updateName}
            />
            <TextField
            id="standard-multiline-flexible"
            label="Description"
            multiline
            rowsMax="4"
            onChange={this.updateDesc}
            margin="normal"
            />
          </Grid>
        </Grid>
      </div>
    );
  }
}
