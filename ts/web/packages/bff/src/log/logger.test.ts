import { describe, expect, it } from 'bun:test'
import { getLogger, initLogger } from './logger'

describe('initLogger', () => {
  it('init logger', async () => {
    // Only side-effects, but logger.error and logger.debug are both verified
    // to be noop.
    initLogger('info', false)
    const logger = getLogger()
    expect(logger.error).toEqual(logger.debug)
  })
})
