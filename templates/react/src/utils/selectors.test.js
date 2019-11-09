import selectors from './selectors'

describe('selectors', () => {
  it('getFormProgress form 3', () => {
    const data = {
      myProfile: {
        formResponses: [
          { formTemplateId: 'PROFILE' },
          { formTemplateId: 'ADDRESS' },
        ],
      },
    }
    const { formProgress } = selectors.getFormProgress({ data })

    expect(formProgress.getNextIndex()).toBe(2)
    expect(formProgress.getActiveTab()).toBe(0)
    expect(formProgress.isComplete()).toBe(false)
    expect(formProgress.getTabProgress()).toBe(40)
  })
})
