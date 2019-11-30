import React from 'react';
import './Frontend.css';
import Grid from '@material-ui/core/Grid';
import ToggleButton from '@material-ui/lab/ToggleButton';
import ToggleButtonGroup from '@material-ui/lab/ToggleButtonGroup';


export default class Frontend extends React.Component {
  handleChange = (e, v) => {
    this.props.setFrontend(v);
  }
  render() {
    return (
      <div className="frontend">
        <h2>Frontend</h2>
        <Grid container spacing={2} direction="column" alignItems="center">
          <Grid item>
            <ToggleButtonGroup
              value={this.props.frontend}
              exclusive
              children="true"
              onChange={this.handleChange}
              aria-label="Frontend Framework"
            >
              <ToggleButton key={1} value="react" aria-label="React">
                React
            </ToggleButton>
              <ToggleButton key={2} value="angular" aria-label="Angular" disabled={true}>
                Angular
            </ToggleButton>
              <ToggleButton key={3} value="vue" aria-label="Vue" disabled={true}>
                Vue
            </ToggleButton>
            </ToggleButtonGroup>
          </Grid>
        </Grid>
      </div>
    );
  }
}
