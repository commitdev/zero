// A simple in-house i18n management tool
// based on API from i18next https://www.i18next.com/overview/api
// i18next is Webtranslateit compatible
import englishResources from '../assets/locales/en'
import { getProp } from './helpers'

class Translator {
  constructor({ locale = 'en', resources = {} }) {
    this.locale = locale
    this.resources = resources
  }

  setLocale(locale) {
    this.locale = locale
  }

  setResources(resources) {
    this.resources = resources
  }

  t(_keyPath = '', fallback) {
    if (!_keyPath.trim) {
      console.error('invalid keyPath', _keyPath)
      return fallback
    }
    const keyPath = _keyPath.trim()
    const text = getProp(this.resources || {}, `${this.locale}.${keyPath}`)
    if (text === undefined) {
      console.warn(`Missing translation for: [${this.locale}]`, keyPath)
      return fallback === undefined ? keyPath : fallback
    } else {
      return text
    }
  }
}

const translator = new Translator({
  resources: {
    en: englishResources,
  },
})

export default translator
