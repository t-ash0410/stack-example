import pino, { type LevelWithSilentOrString, type Logger } from 'pino'

let logger: Logger

export const initLogger = (
  level: LevelWithSilentOrString,
  enabled?: boolean,
) => {
  logger = pino({
    level: level,
    enabled: enabled,
  })
}

export const getLogger = () => {
  return (
    logger ??
    pino({
      enabled: false,
    })
  )
}
