export const typeDefs = /* GraphQL */ `
  type Session {
    access_token: String
    userGroup: String
    orgId: String
    userId: String
  }

  type Mutation {
    login(access_token: String!)
    logout()
  }

  type Query {
    session: Session
  }
`

export const defaults = {
  session: {
    __typename: 'Session',
    access_token: null,
    userGroup: null,
    userId: null,
    orgId: null,
  },
}
