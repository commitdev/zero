import React from 'react';
import Button from '@material-ui/core/Button';

export default class GenerateButton extends React.Component {
  render() {
    return (
      <p>
        <Button color="primary" variant="contained" onClick={this.props.generate}>Generate</Button>
      </p>
    )
  }
}
