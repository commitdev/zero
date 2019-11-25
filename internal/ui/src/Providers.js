import React from 'react';
import Grid from '@material-ui/core/Grid';
import ToggleButton from '@material-ui/lab/ToggleButton';
import ToggleButtonGroup from '@material-ui/lab/ToggleButtonGroup';
import AWSProvider from './AWSProvider';

export default class Providers extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      provider: props.provider
    }
  }

  handleChange = (_, newAlignment) => {
    this.setState({
      provider: newAlignment
    })
    this.props.setProvider(newAlignment);
  };

  render() {
    return (
      <div className="providers">
        <h2>Infrastructure</h2>
        <Grid container spacing={2} direction="column" alignItems="center">
          <Grid item>
            <ToggleButtonGroup
              value={this.state.provider}
              exclusive
              onChange={this.handleChange}
              aria-label="Cloud Provider"
            >
              <ToggleButton key={1} value="aws" aria-label="Amazon AWS">
                Amazon AWS
              </ToggleButton>
              <ToggleButton key={2} value="gcp" aria-label="Google Cloud">
                Google Cloud
              </ToggleButton>
              <ToggleButton key={3} value="azure" aria-label="Microsoft Azure">
                Microsoft Azure
              </ToggleButton>
            </ToggleButtonGroup>
          </Grid>

          {this.state.provider === "aws" && <AWSProvider setProvider={this.props.setProvider} />}
          {this.props.provider !== "aws" && <div><p>Only Amazon AWS is supported right now.</p></div>}
        </Grid>
      </div>
    );
  }
}
