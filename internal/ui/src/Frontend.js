import React from 'react';
import './Frontend.css';
import Grid from '@material-ui/core/Grid';
import ToggleButton from '@material-ui/lab/ToggleButton';
import ToggleButtonGroup from '@material-ui/lab/ToggleButtonGroup';


export default function Frontend() {
  const [provider, setProvider] = React.useState('aws');

  const handleChange = (event, newAlignment) => {
    setProvider(newAlignment);
  };

  return (
    <div className="frontend">
      <h2>Frontend</h2>
      <Grid container spacing={2} direction="column" alignItems="center">
        <Grid item>
          <ToggleButtonGroup
            value={provider}
            exclusive
            onChange={handleChange}
            aria-label="Frontend Framework"
          >
            <ToggleButton key={1} value="react" aria-label="React">
              React
            </ToggleButton>
            <ToggleButton key={2} value="angular" aria-label="Angular">
              Angular
            </ToggleButton>
            <ToggleButton key={3} value="vue" aria-label="Vue">
              Vue
            </ToggleButton>
          </ToggleButtonGroup>
        </Grid>
      </Grid>
    </div>
  );
}
