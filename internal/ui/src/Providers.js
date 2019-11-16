import React from 'react';
import Grid from '@material-ui/core/Grid';
import ToggleButton from '@material-ui/lab/ToggleButton';
import ToggleButtonGroup from '@material-ui/lab/ToggleButtonGroup';
import List from '@material-ui/core/List';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';
import MenuItem from '@material-ui/core/MenuItem';
import Menu from '@material-ui/core/Menu';

export default function Providers() {
  const [provider, setProvider] = React.useState('aws');
  const [region, setRegion] = React.useState(null);
  const [profile, setProfile] = React.useState(null);
  const [selectedProfile, setSelectedProfile] = React.useState(1);
  const [selectedRegion, setSelectedRegion] = React.useState(1);

  const regions = [
    'US East (Ohio) us-east-2',
    'US East (N. Virginia) us-east-1',
    'US West (N. California) us-west-1',
    'US West (Oregon) us-west-2',
    'Asia Pacific (Hong Kong) ap-east-1',
    'Asia Pacific (Mumbai) ap-south-1',
    'Asia Pacific (Seoul) ap-northeast-2',
    'Asia Pacific (Singapore) ap-southeast-1 ',
    'Asia Pacific (Sydney) ap-southeast-2 ',
    'Asia Pacific (Tokyo) ap-northeast-1 ',
    'Canada (Central) ca-central-1 ',
    'China (Beijing) cn-north-1 ',
    'China (Ningxia) cn-northwest-1',
    'EU (Frankfurt) eu-central-1 ',
    'EU (Ireland) eu-west-1 ',
    'EU (London) eu-west-2 ',
    'EU (Paris) eu-west-3 ',
    'EU (Stockholm) eu-north-1 ',
    'Middle East (Bahrain) me-south-1 ',
    'South America (Sao Paulo) sa-east-1 '
  ];

  const profiles = [
    'default',
    'testing',
    'fake-data'
  ];

  const handleChange = (event, newAlignment) => {
    setProvider(newAlignment);
  };

  const handleClickListItemRegion = (event) => {
    setRegion(event.currentTarget);
  }

  const handleMenuItemClickRegion = (event, index) => {
    setSelectedRegion(index);
    setRegion(null);
  };

  const handleClickListItemProfile = (event) => {
    setProfile(event.currentTarget);
  }

  const handleMenuItemClickProfile = (event, index) => {
    setSelectedProfile(index);
    setProfile(null);
  };

  const handleClose = () => {
    setProfile(null);
    setRegion(null);
  };



  return (
    <div className="providers">
      <h2>Infrastructure</h2>
      <Grid container spacing={2} direction="column" alignItems="center">
        <Grid item>
          <ToggleButtonGroup
            value={provider}
            exclusive
            onChange={handleChange}
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

        {provider === "aws" &&
          <div><h3>AWS</h3>
            <Grid container spacing={4} direction="row" alignItems="center" justify="center">

              <Grid item>
                <List component="nav" aria-label="Region">
                  <ListItem
                    button
                    aria-haspopup="true"
                    aria-controls="lock-menu"
                    aria-label="Region"
                    onClick={handleClickListItemRegion}
                  >
                    <ListItemText primary="Region" secondary={regions[selectedRegion]} />
                  </ListItem>
                </List>
                <Menu
                  id="lock-menu"
                  anchorEl={region}
                  keepMounted
                  open={Boolean(region)}
                  onClose={handleClose}
                >
                  {regions.map((region, index) => (
                    <MenuItem
                      key={region}
                      selected={index === selectedRegion}
                      onClick={event => handleMenuItemClickRegion(event, index, "region")}
                    >
                      {region}
                    </MenuItem>
                  ))}
                </Menu>
              </Grid>
              <Grid item>
                <List component="nav" aria-label="Profiles">
                  <ListItem
                    button
                    aria-haspopup="true"
                    aria-controls="lock-menu"
                    aria-label="Profile"
                    onClick={handleClickListItemProfile}
                  >
                    <ListItemText primary="Profile" secondary={profiles[selectedProfile]} />
                  </ListItem>
                </List>
                <Menu
                  id="lock-menu"
                  anchorEl={profile}
                  keepMounted
                  open={Boolean(profile)}
                  onClose={handleClose}
                >
                  {profiles.map((profile, index) => (
                    <MenuItem
                      key={profile}
                      selected={index === selectedProfile}
                      onClick={event => handleMenuItemClickProfile(event, index)}
                    >
                      {profile}
                    </MenuItem>
                  ))}
                </Menu>
              </Grid>
            </Grid>
          </div>
        }
        {provider !== "aws" && <div><p>Only Amazon AWS is supported right now.</p></div>}
      </Grid>
    </div>
  );
}
