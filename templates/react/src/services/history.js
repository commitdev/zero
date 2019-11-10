import { createBrowserHistory } from 'history'

const history = createBrowserHistory()
// for debugging purposes, this browser history wont' track history.pushState as there are no reliable event for it.
// note that popstate event is only for going back on user clicks
window.appHistory = history

export default history
