package moovhelper

import "open-tritium/dependencies/moovhelper/logger"
import "open-tritium/dependencies/moovhelper/debugger"
import "open-tritium/dependencies/moovhelper/script"

const MtkSourceAttr = "data-mtk_source"
const MtkZeroMatchAttr = "data-mtk_zero_match"

type MoovXHelper interface {
	logger.MoovXLogger
	debugger.MoovXDebugger
	script.ScriptManager
	debugger.BreakpointManager
}
