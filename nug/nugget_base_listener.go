// Generated from Nugget.g4 by ANTLR 4.7.

package parser // Nugget

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseNuggetListener is a complete listener for a parse tree produced by NuggetParser.
type BaseNuggetListener struct{}

var _ NuggetListener = &BaseNuggetListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseNuggetListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseNuggetListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseNuggetListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseNuggetListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterNugget is called when production nugget is entered.
func (s *BaseNuggetListener) EnterNugget(ctx *NuggetContext) {}

// ExitNugget is called when production nugget is exited.
func (s *BaseNuggetListener) ExitNugget(ctx *NuggetContext) {}

// EnterInitextract is called when production initextract is entered.
func (s *BaseNuggetListener) EnterInitextract(ctx *InitextractContext) {}

// ExitInitextract is called when production initextract is exited.
func (s *BaseNuggetListener) ExitInitextract(ctx *InitextractContext) {}

// EnterAssign is called when production assign is entered.
func (s *BaseNuggetListener) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production assign is exited.
func (s *BaseNuggetListener) ExitAssign(ctx *AssignContext) {}

// EnterExecute is called when production execute is entered.
func (s *BaseNuggetListener) EnterExecute(ctx *ExecuteContext) {}

// ExitExecute is called when production execute is exited.
func (s *BaseNuggetListener) ExitExecute(ctx *ExecuteContext) {}

// EnterStreamTask is called when production streamTask is entered.
func (s *BaseNuggetListener) EnterStreamTask(ctx *StreamTaskContext) {}

// ExitStreamTask is called when production streamTask is exited.
func (s *BaseNuggetListener) ExitStreamTask(ctx *StreamTaskContext) {}

// EnterSave is called when production save is entered.
func (s *BaseNuggetListener) EnterSave(ctx *SaveContext) {}

// ExitSave is called when production save is exited.
func (s *BaseNuggetListener) ExitSave(ctx *SaveContext) {}

// EnterFilter is called when production filter is entered.
func (s *BaseNuggetListener) EnterFilter(ctx *FilterContext) {}

// ExitFilter is called when production filter is exited.
func (s *BaseNuggetListener) ExitFilter(ctx *FilterContext) {}

// EnterFilename is called when production filename is entered.
func (s *BaseNuggetListener) EnterFilename(ctx *FilenameContext) {}

// ExitFilename is called when production filename is exited.
func (s *BaseNuggetListener) ExitFilename(ctx *FilenameContext) {}

// EnterTimefilter is called when production timefilter is entered.
func (s *BaseNuggetListener) EnterTimefilter(ctx *TimefilterContext) {}

// ExitTimefilter is called when production timefilter is exited.
func (s *BaseNuggetListener) ExitTimefilter(ctx *TimefilterContext) {}

// EnterReference is called when production reference is entered.
func (s *BaseNuggetListener) EnterReference(ctx *ReferenceContext) {}

// ExitReference is called when production reference is exited.
func (s *BaseNuggetListener) ExitReference(ctx *ReferenceContext) {}

// EnterDate is called when production date is entered.
func (s *BaseNuggetListener) EnterDate(ctx *DateContext) {}

// ExitDate is called when production date is exited.
func (s *BaseNuggetListener) ExitDate(ctx *DateContext) {}

// EnterSubtype is called when production subtype is entered.
func (s *BaseNuggetListener) EnterSubtype(ctx *SubtypeContext) {}

// ExitSubtype is called when production subtype is exited.
func (s *BaseNuggetListener) ExitSubtype(ctx *SubtypeContext) {}

// EnterTask is called when production task is entered.
func (s *BaseNuggetListener) EnterTask(ctx *TaskContext) {}

// ExitTask is called when production task is exited.
func (s *BaseNuggetListener) ExitTask(ctx *TaskContext) {}

// EnterTarget is called when production target is entered.
func (s *BaseNuggetListener) EnterTarget(ctx *TargetContext) {}

// ExitTarget is called when production target is exited.
func (s *BaseNuggetListener) ExitTarget(ctx *TargetContext) {}

// EnterField is called when production field is entered.
func (s *BaseNuggetListener) EnterField(ctx *FieldContext) {}

// ExitField is called when production field is exited.
func (s *BaseNuggetListener) ExitField(ctx *FieldContext) {}

// EnterSourceidentifier is called when production sourceidentifier is entered.
func (s *BaseNuggetListener) EnterSourceidentifier(ctx *SourceidentifierContext) {}

// ExitSourceidentifier is called when production sourceidentifier is exited.
func (s *BaseNuggetListener) ExitSourceidentifier(ctx *SourceidentifierContext) {}

// EnterPrintId is called when production printId is entered.
func (s *BaseNuggetListener) EnterPrintId(ctx *PrintIdContext) {}

// ExitPrintId is called when production printId is exited.
func (s *BaseNuggetListener) ExitPrintId(ctx *PrintIdContext) {}

// EnterSin is called when production sin is entered.
func (s *BaseNuggetListener) EnterSin(ctx *SinContext) {}

// ExitSin is called when production sin is exited.
func (s *BaseNuggetListener) ExitSin(ctx *SinContext) {}
