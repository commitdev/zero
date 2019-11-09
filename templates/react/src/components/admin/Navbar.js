import React from 'react'
import { withApollo, graphql, compose } from 'react-apollo'

import Tabs from '@material-ui/core/Tabs/Tabs'
import Tab from '@material-ui/core/Tab/Tab'
import PropTypes from 'prop-types'

import { Text } from '../Text'
import Navbar from '../Navbar'
import { router } from '../../utils'
import { queries } from '../../constants'

const AdminNavbar = ({ location = '', data }) => {
  const tab = (location.split('/')[2] || '').toLowerCase()

  const handleTab = (e, tab) => router.push(`/admin/${tab}`)

  return (
    <Navbar user={`Admin`}>
      <Tabs
        variant="fullWidth"
        indicatorColor="primary"
        value={tab}
        onChange={handleTab}
      >
        <Tab value="" label={<Text>dashboard</Text>} />
        <Tab value="forms" label={<Text>Forms</Text>} />
      </Tabs>
    </Navbar>
  )
}

AdminNavbar.propTypes = {
  location: PropTypes.string,
}

export default compose(
  graphql(queries.GET_SESSION),
  withApollo
)(AdminNavbar)
