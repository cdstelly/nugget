// Generated from C:/Users/drew/GoglandProjects/nugget\Nugget2.g4 by ANTLR 4.7.

package parser // Nugget2

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseNugget2Listener is a complete listener for a parse tree produced by Nugget2Parser.
type BaseNugget2Listener struct{}

var _ Nugget2Listener = &BaseNugget2Listener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseNugget2Listener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseNugget2Listener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseNugget2Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseNugget2Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProg is called when production prog is entered.
func (s *BaseNugget2Listener) EnterProg(ctx *ProgContext) {}

// ExitProg is called when production prog is exited.
func (s *BaseNugget2Listener) ExitProg(ctx *ProgContext) {}

// EnterDefine_assign is called when production define_assign is entered.
func (s *BaseNugget2Listener) EnterDefine_assign(ctx *Define_assignContext) {}

// ExitDefine_assign is called when production define_assign is exited.
func (s *BaseNugget2Listener) ExitDefine_assign(ctx *Define_assignContext) {}

// EnterDefine is called when production define is entered.
func (s *BaseNugget2Listener) EnterDefine(ctx *DefineContext) {}

// ExitDefine is called when production define is exited.
func (s *BaseNugget2Listener) ExitDefine(ctx *DefineContext) {}

// EnterAssign is called when production assign is entered.
func (s *BaseNugget2Listener) EnterAssign(ctx *AssignContext) {}

// ExitAssign is called when production assign is exited.
func (s *BaseNugget2Listener) ExitAssign(ctx *AssignContext) {}

// EnterAsType is called when production asType is entered.
func (s *BaseNugget2Listener) EnterAsType(ctx *AsTypeContext) {}

// ExitAsType is called when production asType is exited.
func (s *BaseNugget2Listener) ExitAsType(ctx *AsTypeContext) {}

// EnterNugget_type is called when production nugget_type is entered.
func (s *BaseNugget2Listener) EnterNugget_type(ctx *Nugget_typeContext) {}

// ExitNugget_type is called when production nugget_type is exited.
func (s *BaseNugget2Listener) ExitNugget_type(ctx *Nugget_typeContext) {}

// EnterNuggetaction is called when production nuggetaction is entered.
func (s *BaseNugget2Listener) EnterNuggetaction(ctx *NuggetactionContext) {}

// ExitNuggetaction is called when production nuggetaction is exited.
func (s *BaseNugget2Listener) ExitNuggetaction(ctx *NuggetactionContext) {}
