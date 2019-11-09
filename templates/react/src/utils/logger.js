class Logger {
  constructor({ FS, ga }) {
    this.FS = FS
    this.ga = ga
  }

  isLoaded(app) {
    return process.env.NODE_ENV === 'production' && this[app] !== undefined
  }

  setUser(userId) {
    this.isLoaded('FS') &&
      this.FS.identify(userId, {
        displayName: userId,
        email: userId,
        // Add your own custom user variables here, details at
        // http://help.fullstory.com/develop-js/setuservars
        reviewsWritten_int: 14,
      })

    this.isLoaded('ga') && userId && this.ga('set', '&uid', userId)
  }

  pageView(href) {
    if (this.isLoaded('ga') && href) {
      // NOTE: setting the page variable then calling a general track allows other events to work properly
      // https://developers.google.com/analytics/devguides/collection/analyticsjs/single-page-applications
      this.ga('set', 'page', href)
      this.ga('send', 'pageview')
    }
  }

  track() {
    if (this.isLoaded('ga')) {
      // ref: https://developers.google.com/analytics/devguides/collection/analyticsjs/events
      this.ga.call(this, 'send', 'event', ...arguments)
    }
  }

  error(description, fatal = false) {
    // https://developers.google.com/analytics/devguides/collection/analyticsjs/exceptions
    if (this.isLoaded('ga')) {
      this.ga('send', 'exception', {
        exDescription: description,
        exFatal: fatal,
      })
    }
  }
}

export default new Logger(window)
