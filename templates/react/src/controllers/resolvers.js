import { queries } from '../constants'
import { logger, router } from '../utils'
import auth from '../services/auth'

// Ref: https://github.com/apollographql/apollo-link-state/tree/master/examples
export default {
  Mutation: {
    login: (_, { access_token, userGroup, userId, orgId }, { cache }) => {
      try {
        // const loginRedirect = storage.getItem(LOGIN_REDIRECT)
        // if (loginRedirect) {
        //   Router.push(loginRedirect)
        // }
        // storage.removeItem(LOGIN_REDIRECT)

        const { session } = cache.readQuery({ query: queries.GET_SESSION })
        cache.writeData({
          data: {
            session: {
              ...session,
              access_token,
              userGroup,
              userId,
              orgId,
            },
          },
        })
      } catch (e) {
        logger.error(e.statusText, e)
      }
      return null
    },
    logout: async () => {
      await auth.logout()
      // router.push('/')
      return null
    },
  },
}
