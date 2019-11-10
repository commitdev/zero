// TODO hook up to error monitoring service
export function trackException(err) {
  console.error(err)
  alert('Unexpected Error Occured')
}

export function handleError(err) {
  // TODO match against known error codes
}

export default {
  trackException,
  handleError,
}
