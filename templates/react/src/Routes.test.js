import React from 'react'
import { Routes } from './Routes'
import { render, waitForElement, fireEvent } from 'react-testing-library'
import { setupPolly } from 'setup-polly-jest'
import FSPersister from '@pollyjs/persister-fs'
import { Polly } from '@pollyjs/core'
import 'jest-dom/extend-expect'
import HTTPAdapter from '@pollyjs/adapter-node-http'
import router from './utils/router'

Polly.register(HTTPAdapter)
Polly.register(FSPersister)
// doc: https://github.com/kentcdodds/react-testing-library#readme
// doc: https://reactjs.org/docs/test-utils.html

const nextLoop = new Promise(resolve => setTimeout(resolve, 0))

describe('Renders pages:', () => {
  let context = setupPolly({
    adapters: ['node-http'],
    persister: 'fs',
  })

  it('/ - should render landing page', () => {
    const { getByTestId } = render(<Routes />)
    expect(getByTestId('start-btn')).toHaveTextContent('start')
  })

  it('/login', () => {
    const { getByTestId } = render(<Routes location="/login" />)
    expect(getByTestId('login-btn')).toBeVisible()
  })

  it('/signup', () => {
    const { getByTestId } = render(<Routes location="/signup" />)
    expect(getByTestId('signup-btn')).toBeVisible()
  })

  it('/password-reset', () => {
    const { getByTestId } = render(<Routes location="/password-reset" />)
    expect(getByTestId('reset-password-btn')).toBeVisible()
  })

  it('/app - should not render when unauthenticated', async () => {
    const errorLogger = global.console.error
    global.console.error = jest.fn()
    const { queryByTestId } = render(<Routes location="/app" />)

    // waiting for next event loop to render out components, for some reason, not working with waitForElement
    await nextLoop
    expect(
      queryByTestId(document.documentElement, 'dashboard-nav')
    ).not.toBeInTheDocument()
    expect(console.error).toBeCalled()

    global.console.error = errorLogger
  })

  it('/login - should be able to login and trigger redirect', async () => {
    // NOTE: because aws creates unique Ids we need to create custom recorders for this
    context.polly.stop()
    const assignMock = jest.fn()
    global.location.assign = assignMock

    const { queryByTestId, getByLabelText } = render(
      <Routes location="/login" />
    )
    setInputValue(getByLabelText('email'), 'thomas@meta.re')
    setInputValue(getByLabelText('password'), '@Testing1')
    fireEvent.click(queryByTestId('login-btn'))
    await waitForElement(() => {
      expect(window.location.assign).toBeCalled()
      return queryByTestId('login-btn')
    })
  })

  describe('Authenticated Routes:', () => {
    it('/app/ - should render investor dashboard', async () => {
      const { getByTestId } = render(<Routes location="/app" />)

      await waitForElement(() => {
        return getByTestId('app-navbar')
      })
    })
    it('/admin - should redirect away from admin routes', async () => {
      render(<Routes location="/admin" />)
      const spy = jest.spyOn(router, 'push')

      await waitForElement(() => {
        expect(spy).toHaveBeenCalledWith('/app')
        return true
      })
    })
  })
})

function setInputValue(inputEl, value) {
  fireEvent.change(inputEl, {
    target: { value },
  })
}
