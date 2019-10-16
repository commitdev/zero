import React, { Fragment } from 'react';
import Container from '@material-ui/core/Container';
import Header from 'components/layout/header';
import config from 'config';

export default function App({children}) {
  return (
    <Fragment>
      { config && config.header && config.header.enabled && <Header />}
      <Container maxWidth={false}>
        {children}
      </Container>
    </Fragment>
  );
}
