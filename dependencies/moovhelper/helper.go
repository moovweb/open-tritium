package moovhelper

import "tritium_oss/dependencies/moovhelper/logger"
import "tritium_oss/dependencies/moovhelper/debugger"
import "tritium_oss/dependencies/moovhelper/script"

const MtkSourceAttr = "data-mtk_source"
const MtkZeroMatchAttr = "data-mtk_zero_match"

type MoovXHelper interface {
	logger.MoovXLogger
	debugger.MoovXDebugger
	script.ScriptManager
	debugger.BreakpointManager
}
