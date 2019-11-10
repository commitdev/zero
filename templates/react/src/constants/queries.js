import gql from 'graphql-tag'

const GET_ALL_FORM_TEMPLATES = gql`
  query {
    formTemplates {
      id
      title
      description
    }
  }
`

const GET_FORM_TEMPLATE = gql`
  query FormTemplate($id: ID!) {
    formTemplate(id: $id) {
      title
      description
      fieldTypes {
        name
        label
        helpText
        component
        required
        type
        width
        fieldset
        options {
          value
          label
        }
      }
    }
  }
`

const GET_FORM = gql`
  query Form($templateId: ID!, $userId: ID!) {
    formTemplate(id: $templateId) {
      title
      description
      fieldTypes {
        name
        label
        helpText
        component
        required
        type
        width
        fieldset
        options {
          value
          label
        }
      }
    }
    formResponse(userId: $userId, formTemplateId: $templateId) {
      id
      formFields {
        name
        value
      }
    }
  }
`

const DESTROY_TEMPLATE_MUTATION = gql`
  mutation DestroyTemplateMutation($id: ID!) {
    destroyTemplate(id: $id)
  }
`

const SUBMIT_FORM_MUTATION = gql`
  mutation SubmitForm(
    $responseId: ID
    $templateId: ID!
    $formFields: [FormFieldInput]!
  ) {
    submitForm(
      formResponseId: $responseId
      formTemplateId: $templateId
      formFields: $formFields
    ) {
      id
      userId
      formTemplateId
      formFields {
        name
        value
      }
      createdAt
      updatedAt
    }
  }
`

const UPSERT_TEMPLATE_MUTATION = gql`
  mutation UpsertTemplateMutation(
    $id: ID
    $title: String!
    $fieldTypes: [FieldTypeInput]
  ) {
    upsertTemplate(id: $id, title: $title, fieldTypes: $fieldTypes) {
      id
      title
    }
  }
`

export default {
  GET_ALL_FORM_TEMPLATES,
  GET_FORM,
  GET_FORM_TEMPLATE,
  DESTROY_TEMPLATE_MUTATION,
  UPSERT_TEMPLATE_MUTATION,
  SUBMIT_FORM_MUTATION,

  GET_SESSION: gql`
    query {
      session @client {
        access_token
        userGroup
        orgId
        userId
      }
    }
  `,

  LOGIN: gql`
    mutation Login(
      $accessToken: String!
      $userId: String
      $orgId: String
      $userGroup: String
    ) {
      login(
        access_token: $accessToken
        userId: $userId
        orgId: $orgId
        userGroup: $userGroup
      ) @client
    }
  `,

  LOGOUT: gql`
    mutation Logout {
      logout @client
    }
  `,
}
